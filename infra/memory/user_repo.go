package memory

import (
	"sync"

	"github.com/gabrielbsx/wyd-go/domain/identity"
)

type InMemoryUserRepository struct {
	mu    sync.RWMutex
	users map[string]*identity.User
}

func NewInMemoryUserRepo() *InMemoryUserRepository {
	return &InMemoryUserRepository{
		users: make(map[string]*identity.User),
	}
}

func (r *InMemoryUserRepository) FindByUsername(username string) (*identity.User, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	user, ok := r.users[username]

	if !ok {
		return nil, identity.ErrUserNotFound
	}

	return user, nil
}

func (r *InMemoryUserRepository) Save(user *identity.User) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	r.users[user.Username] = user

	return nil
}
