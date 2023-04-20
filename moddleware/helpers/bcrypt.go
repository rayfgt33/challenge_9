package helpers

import (
	"log"

	"golang.org/x/crypto/bcrypt"
)

// Untuk menjaga keamanan password/enkripsi
func HashPass(p string) string {
	salt := 8 //jumlah hasil encrypt
	password := []byte(p)
	hash, err := bcrypt.GenerateFromPassword(password, salt)
	if err != nil {
		log.Fatal(err)
	}

	return string(hash)
}

// untuk membandingkan hasil encrypt password
func ComparePass(h, p []byte) bool {
	err := bcrypt.CompareHashAndPassword(h, p)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	return err == nil
}
