package helpers

import "golang.org/x/crypto/bcrypt"


func HashPass(pass string) (string, error) {
	p := []byte(pass)
	salt := bcrypt.DefaultCost
	hashed, err := bcrypt.GenerateFromPassword(p, salt)
	if err != nil {
		return "", err
	}
	return string(hashed), nil
}


func ComparePass(hash, pass []byte) (bool, error) {
	err := bcrypt.CompareHashAndPassword(hash, pass)
	if err != nil {
		return false, err
	}
	return true, nil
}

