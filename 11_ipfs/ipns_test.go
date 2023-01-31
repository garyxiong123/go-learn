package _1_ipfs

import (
	"testing"
	//"time"
	//ipns "github.com/ipfs/go-ipns"
	//crypto "github.com/libp2p/go-libp2p-crypto"
)

func Test_1(t *testing.T) {

	// Generate a private key to sign the IPNS record with. Most of the time,
	// however, you'll want to retrieve an already-existing key from IPFS using the
	// go-ipfs/core/coreapi CoreAPI.KeyAPI() interface.
	//privateKey, publicKey, err := crypto.GenerateKeyPair(crypto.RSA, 2048)
	//println(publicKey)
	//if err != nil {
	//	panic(err)
	//}
	//
	//// Create an IPNS record that expires in one hour and points to the IPFS address
	//// /ipfs/Qme1knMqwt1hKZbc1BmQFmnm9f36nyQGwXxPGVpVJ9rMK5
	//ipnsRecord, err := ipns.Create(privateKey, []byte("/ipfs/Qme1knMqwt1hKZbc1BmQFmnm9f36nyQGwXxPGVpVJ9rMK5"), 0, time.Now().Add(1*time.Hour), 0)
	//if err != nil {
	//	panic(err)
	//}
	//println(ipnsRecord)
	//here are several other major operations you can do with go-ipns. Check out the API docs or look at the tests in this repo for examples.
}
