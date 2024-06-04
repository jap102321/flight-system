package utils

import "golang.org/x/crypto/bcrypt"

func HashPassword(password string) (string, error){

	hashedPass, err := bcrypt.GenerateFromPassword([]byte(password), 14)


	return string(hashedPass), err
}

func DecryptPassword(hashedPassword, password string) bool{
	
	err := bcrypt.CompareHashAndPassword([]byte(password), []byte(hashedPassword))


	return err == nil
}