package __eth

import (
	"context"
	"log"
	"testing"
)

func TestGetNonce(t *testing.T) {
	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		log.Fatal(err)
	}
	println(nonce)
}
