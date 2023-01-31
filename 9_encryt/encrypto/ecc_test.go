package encrypto

import (
	"testing"
)

func Test_ECDSA_Encrypto(t *testing.T) {
	// Generate a new private key
	//privateKey, x, y, err := elliptic.GenerateKey(elliptic.P256(), rand.Reader)
	//if err != nil {
	//	fmt.Println(err)
	//	return
	//}
	//
	//// Encrypt the data
	//message := []byte("Hello, World!")
	//elliptic.P256()
	//ciphertext, err := elliptic.Encrypt(rand.Reader, &privateKey.PublicKey, message)
	//if err != nil {
	//	fmt.Println(err)
	//	return
	//}
	//
	//// Decrypt the data
	//plaintext, err := elliptic.Decrypt(privateKey, x, y, ciphertext)
	//if err != nil {
	//	fmt.Println(err)
	//	return
	//}
	//
	//fmt.Println("Plaintext:", string(plaintext))
}
