package main

import (
	"C"
	"crypto/rand"
	"crypto/sha512"
	"encoding/base64"

	_ "github.com/GehirnInc/crypt/sha512_crypt"
)
import (
	"fmt"

	"github.com/GehirnInc/crypt"
)

// Define salt size
const saltSize = 16

func main() {

	//https://www.cyberciti.biz/faq/understanding-etcshadow-file/
	//https://www.computernetworkingnotes.com/linux-tutorials/etc-shadow-file-in-linux-explained-with-examples.html

	// salt := generateRandomSalt(saltSize)
	// fmt.Println(salt, string(salt))

	// char := C.CString("$6$Vi.DuMQS")
	// fmt.Println(char)

	// bss := []byte("$6$Vi.DuMQS")
	// fmt.Println(bss)

	// hashPass := hashPassword("test", salt)
	// fmt.Println(hashPass)

	//o3U0A8v/2i7x23pwM//on8IQhak2gaFvEZBZ0zlOdqHSBhD1hOMxxApqm6HyEWgOd/LeFkm7mDYKcdYsqWSr90

	// hashedPassword := "lp4fOhKXEIz5lbVdbSgIKmsHVFxIkogY25tttRFt2RqXkJdCZ88N/WF/NiUqi0CnBEkrlZb4i.FD3SWfRYkOl."
	// currPassword := "test123"
	// salt := []byte("$6$C.pW7Z3S5gn1ftqG")
	// isMatch := passCompare(hashedPassword, currPassword, salt)
	// fmt.Println(isMatch)

	crypt := crypt.SHA512.New()
	generatedHashPass, _ := crypt.Generate([]byte("test123"), []byte("$6$SApW7Z3S5gn1ftLT")) //$6$salt
	fmt.Println(generatedHashPass)

	//hashpass := "$6$cFMM8PU5cD0fSO4D$o3U0A8v/2i7x23pwM//on8IQhak2gaFvEZBZ0zlOdqHSBhD1hOMxxApqm6HyEWgOd/LeFkm7mDYKcdYsqWSr90"
	// hashpass := "$6$C.pW7Z3S5gn1ftqG$lp4fOhKXEIz5lbVdbSgIKmsHVFxIkogY25tttRFt2RqXkJdCZ88N/WF/NiUqi0CnBEkrlZb4i.FD3SWfRYkOl."
	err := crypt.Verify(generatedHashPass, []byte("test123"))
	fmt.Println(err)

	entr, err := Lookup("mostain2")
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("pass:", entr.IsAccountValid(), entr.Pass)
	err = entr.VerifyPassword("test")
	fmt.Println(err)

	//show all user details
	// readAllUser, err := Read()
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }
	// for i, user := range readAllUser {
	// 	fmt.Println(i, user.Name, user.LastChange, user.MinPassAge, user.MaxPassAge, user.IsAccountValid(), user.IsPasswordValid(), user.AcctExpiry)
	// }

}

func generateRandomSalt(saltSize int) []byte {
	var salt = make([]byte, saltSize)

	_, err := rand.Read(salt[:])

	if err != nil {
		panic(err)
	}

	return salt
}

func hashPassword(password string, salt []byte) string {
	// Convert password string to byte slice
	var passwordBytes = []byte(password)

	// Create sha-512 hasher

	var sha512Hasher = sha512.New()

	// Append salt to password
	passwordBytes = append(passwordBytes, salt...)

	// Write password bytes to the hasher
	sha512Hasher.Write(passwordBytes)

	// Get the SHA-512 hashed password
	var hashedPasswordBytes = sha512Hasher.Sum(nil)

	// Convert the hashed password to a base64 encoded string
	var base64EncodedPasswordHash = base64.URLEncoding.EncodeToString(hashedPasswordBytes)

	return base64EncodedPasswordHash
}

// Check if two passwords match
func passCompare(hashedPassword, currPassword string, salt []byte) bool {

	var currPasswordHash = hashPassword(currPassword, salt)
	return hashedPassword == currPasswordHash
}
