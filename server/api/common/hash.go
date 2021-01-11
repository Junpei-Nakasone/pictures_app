package common

import "golang.org/x/crypto/bcrypt"

// HashPassword パスワードを暗号化する
func HashPassword(rawPassword string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(rawPassword), bcrypt.DefaultCost)

	return string(hashedPassword), err
}

// VerifyPassword 暗号化されたパスワードとユーザーが入力したパスワードを比較する
func VerifyPassword(hashedPassword string, entryPassword string) error {
	// password := []byte(hashedPassword)
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(entryPassword))
	return err
}
