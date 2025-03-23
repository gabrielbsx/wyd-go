package identity

type UserRepository interface {
	Save(user *User) error
	FindByUsername(username string) (*User, error)
}
