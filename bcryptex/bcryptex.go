package bcryptex

// go get golang.org/x/crypto/bcrypt

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

func BcryptExample() {
	password := "mysecretpassword"
	// Hash the password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		fmt.Println("Error hashing password:", err)
		return
	}

	fmt.Println("Hashed Password:", string(hashedPassword))

	// Check the password
	err = bcrypt.CompareHashAndPassword(hashedPassword, []byte(password))
	if err != nil {
		fmt.Println("Password does not match:", err)
	} else {
		fmt.Println("Password matches!")
	}
}
