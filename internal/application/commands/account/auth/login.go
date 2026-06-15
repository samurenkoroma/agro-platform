package auth

import (
	"encoding/json"
	"net/http"

	"github.com/samurenkoroma/agro-platform/internal/application/uow"
	account "github.com/samurenkoroma/agro-platform/internal/domain/account/aggregate/user"
	domain "github.com/samurenkoroma/agro-platform/internal/domain/account/repository"
	"github.com/samurenkoroma/agro-platform/internal/infrastructure/jwt"
	"github.com/samurenkoroma/agro-platform/internal/infrastructure/repository/providers"
	"github.com/samurenkoroma/agro-platform/internal/interfaces/http/response"
	"github.com/samurenkoroma/agro-platform/internal/shared/repository"
)

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
type User struct {
	Id    string `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
	Role  string `json:"role"`
}
type LoginResult struct {
	TokenPair    *jwt.TokenPair `json:"tokenPair"`
	User         User           `json:"user"`
	CurrentOrgId string         `json:"currentOrgId,omitempty"`
}

func (h *AuthHandler) Login(w http.ResponseWriter, r *http.Request) {
	var req LoginRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		response.WriteValidationError(w, "invalid request body")
		return
	}

	if req.Email == "" {
		response.WriteValidationError(w, "email is required")
		return
	}
	if req.Password == "" || len(req.Password) < 6 {
		response.WriteValidationError(w, "password must be at least 6 characters")
		return
	}

	ctx := r.Context()
	h.uow.Execute(ctx, providers.NewAccountProvider, func(provider repository.RepositoryProvider, exec uow.Execution) (any, error) {
		// Приводим провайдер к нужному типу
		authProvider, ok := provider.(domain.AccountProvider)
		if !ok {
			if !ok {
			}
			response.WriteInternalError(w, repository.ErrInvalidProviderType.Error())
			return nil, repository.ErrInvalidProviderType
		}

		userRepo := authProvider.Users()
		membershipRepo := authProvider.Memberships()

		// Ищем пользователя
		user, err := userRepo.FindByEmail(ctx, req.Email)
		if err != nil {
			response.WriteNotFound(w, err.Error())
			return nil, account.ErrInvalidCredentials
		}

		// Проверяем пароль
		if !user.CheckPassword(req.Password) {
			response.WriteValidationError(w, account.ErrInvalidCredentials.Error())
			return nil, account.ErrInvalidCredentials
		}

		// Проверяем статус
		if !user.IsActive() {
			return nil, account.ErrUserInactive
		}

		currentOrgID := user.GetCurrentOrganizationID()
		var orgRole string
		if currentOrgID != "" {
			// Получаем все членства пользователя
			membership, err := membershipRepo.FindByUserAndOrganization(ctx, user.ID, currentOrgID)
			if err != nil {
				return nil, err
			}
			orgRole = membership.GetRoleName()
		}

		// Генерируем токены с текущей организацией

		tokenPair, err := h.jwtService.GenerateTokenPair(
			user.ID,
			user.Username,
			user.Email,
			string(user.Role),
			currentOrgID,
			orgRole,
		)
		if err != nil {
			return nil, err
		}

		// Обновляем время последнего входа
		user.UpdateLastLogin()
		userRepo.Update(ctx, user)

		exec.RegisterAggregate(user)
		response.Success(
			LoginResult{
				TokenPair: tokenPair,
				User: User{
					Id:    user.ID,
					Name:  user.Username,
					Email: user.Email,
					Role:  user.Role.String(),
				},
				CurrentOrgId: currentOrgID,
			}).WriteJSON(w, http.StatusOK)

		return nil, nil
	})
	return
}
