package identity

import (
	"github.com/gabrielbsx/wyd-go/domain"
	"github.com/gabrielbsx/wyd-go/domain/identity"
)

type SignUpUserCommand struct {
	Name     string
	Username string
	Email    string
	Password string
}

type SignUpHandler struct {
	Repo   identity.UserRepository
	Hasher identity.PasswordHasher
	Logger domain.Logger
}

func NewSignUpHandler(
	repo identity.UserRepository,
	hasher identity.PasswordHasher,
	logger domain.Logger,
) *SignUpHandler {
	return &SignUpHandler{
		Repo:   repo,
		Hasher: hasher,
		Logger: logger,
	}
}

func (h *SignUpHandler) Handle(cmd SignUpUserCommand) error {
	userExists, err := h.Repo.FindByUsername(cmd.Username)

	if userExists != nil || err == nil {
		h.Logger.Errorf("User already exists: %s", cmd.Username)
		return identity.ErrUserAlreadyExists
	}

	passwordHashed, err := h.Hasher.Hash(cmd.Password)

	if err != nil {
		h.Logger.Errorf("Error hashing password: %s", err)
		return identity.ErrInHashPassword
	}

	h.Logger.Infof("Password hashed: %s", passwordHashed)

	user, err := identity.NewUser(
		cmd.Name,
		cmd.Username,
		cmd.Email,
		passwordHashed,
	)

	if err != nil {
		h.Logger.Errorf("Error creating user: %s", err)
		return err
	}

	h.Logger.Infof("User created: %s", user.Username)

	return h.Repo.Save(user)
}
