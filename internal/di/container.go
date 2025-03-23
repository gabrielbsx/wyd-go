package container

import (
	"github.com/gabrielbsx/wyd-go/app/handlers/identity"
	hasher "github.com/gabrielbsx/wyd-go/infra/bcrypt"
	"github.com/gabrielbsx/wyd-go/infra/logger"
	"github.com/gabrielbsx/wyd-go/infra/memory"
)

type AppContainer struct {
	Logger *logger.Logger
	SignUp *identity.SignUpHandler
}

func NewAppContainer() *AppContainer {
	userRepo := memory.NewInMemoryUserRepo()
	hasher := hasher.NewBCryptPasswordHasher()
	logger := logger.NewLogger()
	signUp := identity.NewSignUpHandler(userRepo, hasher, logger)

	return &AppContainer{
		Logger: logger,
		SignUp: signUp,
	}
}
