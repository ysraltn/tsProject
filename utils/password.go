package utils

import "golang.org/x/crypto/bcrypt"

func CheckPasswordHash(password, hashedPassword string) bool {
	// bcrypt.CompareHashAndPassword hash'lenmiş şifre ile düz metin şifreyi karşılaştırır
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	return err == nil
}

func HashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashedPassword), nil
}
