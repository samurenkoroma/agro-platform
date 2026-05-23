package auth

import (
	"encoding/json"
	"net/http"

	account "github.com/samurenkoroma/agro-platform/internal/domain/account/aggregate/user"
	domain "github.com/samurenkoroma/agro-platform/internal/domain/account/repository"
	"github.com/samurenkoroma/agro-platform/internal/infrastructure/repository/providers"
	"github.com/samurenkoroma/agro-platform/internal/interfaces/http/response"
	"github.com/samurenkoroma/agro-platform/internal/shared/repository"
)

// RegisterRequest запрос на регистрацию
type RegisterRequest struct {
	Email     string `json:"email"`
	Username  string `json:"username"`
	Password  string `json:"password"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Phone     string `json:"phone"`
	Role      string `json:"role"`
}

// RegisterResponse ответ на регистрацию
type RegisterResponse struct {
	UserID   string `json:"user_id"`
	Email    string `json:"email"`
	Username string `json:"username"`
	Role     string `json:"role"`
	Message  string `json:"message"`
}

// POST /auth/register
func (h *AuthHandler) Register(w http.ResponseWriter, r *http.Request) {
	var req RegisterRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		response.WriteValidationError(w, "invalid request body")
		return
	}

	// Валидация
	if req.Email == "" {
		response.WriteValidationError(w, "email is required")
		return
	}
	if req.Username == "" {
		response.WriteValidationError(w, "username is required")
		return
	}
	if req.Password == "" || len(req.Password) < 6 {
		response.WriteValidationError(w, "password must be at least 6 characters")
		return
	}

	ctx := r.Context()
	h.uow.Execute(ctx, providers.NewAccountProvider, func(provider repository.RepositoryProvider) (any, error) {
		// Приводим провайдер к нужному типу
		authProvider, ok := provider.(domain.AccountProvider)
		if !ok {
			if !ok {
			}
			response.WriteInternalError(w, repository.ErrInvalidProviderType.Error())
			return nil, repository.ErrInvalidProviderType
		}
		userRepo := authProvider.Users()

		// Проверяем, не существует ли пользователь
		existing, _ := userRepo.FindByEmail(ctx, req.Email)
		if existing != nil {
			return nil, account.ErrUserAlreadyExists
		}

		existing, _ = userRepo.FindByUsername(ctx, req.Username)
		if existing != nil {
			return nil, account.ErrUserAlreadyExists
		}

		// Определяем роль
		role := account.Role(req.Role)
		if role == "" {
			role = account.RoleClient // роль по умолчанию
		}

		// Создаем пользователя
		user, err := account.NewUser(
			req.Email, req.Username, req.Password,
			req.FirstName, req.LastName, req.Phone,
		)
		if err != nil {
			response.WriteValidationError(w, err.Error())
			return nil, err
		}

		// Сохраняем
		if err := userRepo.Save(ctx, user); err != nil {
			return nil, err
		}

		h.uow.RegisterAggregate(user)
		response.Success(RegisterResponse{
			UserID:   user.ID,
			Email:    user.Email,
			Username: user.Username,
			Role:     string(user.Role),
			Message:  "User registered successfully",
		}).WriteJSON(w, http.StatusCreated)
		return nil, nil
	})
	return

}
