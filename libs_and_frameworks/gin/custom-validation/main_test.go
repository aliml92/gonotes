package main

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/stretchr/testify/assert"
	"golang.org/x/crypto/sha3"

	"github.com/ethereum/go-ethereum/common/hexutil"
)

func dummyEthAddr() string {
	privateKey, _ := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)

    // Convert the private key to bytes
    privateKeyBytes := privateKey.D.Bytes()

    // Hash the private key using SHA3-256
    hash := sha3.NewLegacyKeccak256()
    hash.Write(privateKeyBytes)
    hashed := hash.Sum(nil)

    // Take the last 20 bytes of the hashed private key to generate the Ethereum address
    address := hexutil.Encode(hashed[12:])
	return address
}


func TestSaveWallet(t *testing.T) {
	router := setupRouter()
	// wa := gofakeit.Regex(`^0x[a-fA-F0-9]{40}$`)
	// wa := "0xdbF03B407c01E7cD3CBea99509d93f8DDDC8C6FB"
	wa := dummyEthAddr()
	expected := `{"wallet":{"user_id":"` + gofakeit.UUID() + `","wallet_address":"` + wa + `"}}`
	// expected := `{"wallet":{"user_id":"123","wallet_address":"0x02F9AE5f22EA3fA88F05780B30385bECFacbf130"}}`
	req, err := http.NewRequest("POST", "/api/v1/wallets", strings.NewReader(expected))
	if err != nil {
		t.Fatal(err)
	}

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, 201, w.Code)
	assert.JSONEq(t, expected, w.Body.String())
}

