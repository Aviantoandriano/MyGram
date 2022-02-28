package helpers

import "golang.org/x/crypto/bcrypt"

func HashPass(pass string) (hashPassword string) {
	hashPass, err := bcrypt.GenerateFromPassword([]byte(pass), bcrypt.DefaultCost)
	if err != nil {
		panic(err)
	}
	hashPassword = string(hashPass)
	return
}

func PassCheck(hashPass, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashPass), []byte(password))
	if err != nil {
		return false
	}
	return true
}
