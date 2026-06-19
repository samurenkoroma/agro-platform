//go:build integration

package e2e

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/samurenkoroma/agro-platform/internal/bootstrap"
	configs "github.com/samurenkoroma/agro-platform/internal/shared/config"
	dbtest "github.com/samurenkoroma/agro-platform/internal/testutil/postgres"
	"github.com/samurenkoroma/agro-platform/pkg/logger"
)

// Envelope — стандартный конверт ответов.
type Envelope struct {
	Success bool            `json:"success"`
	Data    json.RawMessage `json:"data,omitempty"`
	Error   *struct {
		Code    string `json:"code"`
		Message string `json:"message"`
	} `json:"error,omitempty"`
}

// Client — HTTP-клиент поверх httptest.Server.
type Client struct {
	t           *testing.T
	server      *httptest.Server
	httpClient  *http.Client
	accessToken string
}

func NewClient(t *testing.T, modules ...string) *Client {
	t.Helper()

	db := dbtest.NewTestDB(t, modules...)
	handler := buildApp(t, db.Pool, db.DSN)

	srv := httptest.NewServer(handler)
	t.Cleanup(srv.Close)

	return &Client{
		t:          t,
		server:     srv,
		httpClient: &http.Client{Timeout: 10 * time.Second},
	}
}

func buildApp(t *testing.T, pool *pgxpool.Pool, dsn string) http.Handler {
	t.Helper()

	conf := &configs.Config{
		Db: configs.DbConfig{Dsn: dsn},
		Auth: configs.AuthConfig{
			SecretKey:     "test-secret-key-for-e2e",
			AccessExpiry:  time.Hour,
			RefreshExpiry: 24 * time.Hour,
			Issuer:        "agro-platform-test",
		},
		Server: configs.ServerConfig{ApiPort: ":0"},
		Logger: configs.LoggerConfig{Level: "error"},
	}

	ctx := logger.WithContext(context.Background(), logger.New("error", nil))

	app, err := bootstrap.Build(ctx, pool, conf)
	if err != nil {
		t.Fatalf("bootstrap.Build: %v", err)
	}
	return app.HTTPHandler
}

// SetupOrg регистрирует пользователя, создаёт организацию, переключается на неё.
// После вызова клиент авторизован с валидным organization_id в токене.
func (c *Client) SetupOrg(t *testing.T, email, username, password, orgName string) string {
	t.Helper()

	if env := c.postPublic("/auth/register", map[string]any{
		"email": email, "username": username, "password": password,
		"first_name": "Test", "last_name": "User",
	}); !env.Success {
		t.Fatalf("register: %+v", env.Error)
	}

	loginEnv := c.postPublic("/auth/login", map[string]any{
		"email": email, "password": password,
	})
	if !loginEnv.Success {
		t.Fatalf("login: %+v", loginEnv.Error)
	}
	c.accessToken = extractToken(loginEnv.Data)

	createEnv := c.Command("account.create_organization", map[string]any{"name": orgName})
	if !createEnv.Success {
		t.Fatalf("create_organization: %+v", createEnv.Error)
	}
	// DTO отдаёт {id, name}
	var orgData struct {
		ID string `json:"id"`
	}
	json.Unmarshal(createEnv.Data, &orgData)
	if orgData.ID == "" {
		t.Fatalf("create_organization returned empty id, data=%s", string(createEnv.Data))
	}

	switchEnv := c.Command("account.switch_organization", map[string]any{
		"organization_id": orgData.ID,
	})
	if !switchEnv.Success {
		t.Fatalf("switch_organization: %+v", switchEnv.Error)
	}
	c.accessToken = extractToken(switchEnv.Data)

	return orgData.ID
}

func extractToken(data json.RawMessage) string {
	var d struct {
		TokenPair struct {
			AccessToken string `json:"accessToken"`
		} `json:"tokenPair"`
	}
	json.Unmarshal(data, &d)
	return d.TokenPair.AccessToken
}

// Command бьёт POST /api/commands {"command": name, "data": data}.
func (c *Client) Command(name string, data any) Envelope {
	return c.postAuthed("/api/commands", map[string]any{"command": name, "data": data})
}

// Query бьёт POST /api/queries {"query": name, "data": data}.
func (c *Client) Query(name string, data any) Envelope {
	return c.postAuthed("/api/queries", map[string]any{"query": name, "data": data})
}

// PostWithoutAuth отправляет запрос без токена авторизации.
// Нужен для проверки, что middleware возвращает 401.
func (c *Client) PostWithoutAuth(path string, body any) (int, Envelope) {
	return c.doRequest(path, body, "")
}

func (c *Client) postPublic(path string, body any) Envelope {
	_, env := c.doRequest(path, body, "")
	return env
}

func (c *Client) postAuthed(path string, body any) Envelope {
	if c.accessToken == "" {
		c.t.Fatal("postAuthed: no access token (call SetupOrg first)")
	}
	_, env := c.doRequest(path, body, c.accessToken)
	return env
}

func (c *Client) doRequest(path string, body any, token string) (int, Envelope) {
	c.t.Helper()

	raw, err := json.Marshal(body)
	if err != nil {
		c.t.Fatalf("marshal: %v", err)
	}

	req, err := http.NewRequest(http.MethodPost, c.server.URL+path, bytes.NewReader(raw))
	if err != nil {
		c.t.Fatalf("new request: %v", err)
	}
	req.Header.Set("Content-Type", "application/json")
	if token != "" {
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", token))
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		c.t.Fatalf("do request to %s: %v", path, err)
	}
	defer resp.Body.Close()

	var env Envelope
	json.NewDecoder(resp.Body).Decode(&env)
	return resp.StatusCode, env
}
