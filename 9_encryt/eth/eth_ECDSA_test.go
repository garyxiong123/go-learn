package encrypto

import (
	"crypto/ecdsa"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	"log"
	"testing"
)

func Test_create_eth_ecdsa_By_random_private_key(t *testing.T) {
	privateKey, err := crypto.GenerateKey()
	privateKeyBytes := crypto.FromECDSA(privateKey)
	privateKeyHex := hexutil.Encode(privateKeyBytes)

	key, err := crypto.HexToECDSA(privateKeyHex)

	log.Println("Generated private key address :", key)

	if err != nil {

	}
	publicKey := key.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("error casting public key to ECDSA")
	}
	address := crypto.PubkeyToAddress(*publicKeyECDSA).Hex()

	log.Println("Generated private key address :", address)

}

func Test_create_eth_ecdsa_By__private_string(t *testing.T) {
	//0xfd22f209df3f718b509bda00eb07a5c3f971c56a77c82b693dc691e3756b360c
	privateKey, err := crypto.HexToECDSA("fd22f209df3f718b509bda00eb07a5c3f971c56a77c82b693dc691e3756b360c")
	privateKeyBytes := crypto.FromECDSA(privateKey)
	privateKeyHex := hexutil.Encode(privateKeyBytes)
	key, err := crypto.HexToECDSA(privateKeyHex)

	log.Println("Generated private key address :", key)

	if err != nil {

	}
	publicKey := key.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("error casting public key to ECDSA")
	}
	address := crypto.PubkeyToAddress(*publicKeyECDSA).Hex()

	log.Println("Generated private key address :", address)

}
