package hasher

import "golang.org/x/crypto/bcrypt"

type BCryptPasswordHasher struct{}

func NewBCryptPasswordHasher() *BCryptPasswordHasher {
	return &BCryptPasswordHasher{}
}

func (h *BCryptPasswordHasher) Hash(password string) (string, error) {
	hashed, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

	if err != nil {
		return "", err
	}

	return string(hashed), nil
}

func (h *BCryptPasswordHasher) Check(password, hash string) error {
	return bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
}
