package encryption

import (
	"fmt"
	"crypto/rand"
	"crypto/elliptic"
	"math/big"
	"github.com/ThePiachu/Split-Vanity-Miner-Golang/src/pkg/ripemd160"
)

func CreateKey(log chan string) ([]byte, *big.Int, *big.Int) {
	priv, x, y, err := elliptic.GenerateKey(elliptic.P256(), rand.Reader)
	if err != nil {
		log <- "Key Generation Error"
	}
	return priv, x, y
}

func GetAddress(log chan string, x, y *big.Int) ([]byte, string) {
	pubKey := elliptic.Marshal(elliptic.P256(), x, y)
	ripemd := ripemd160.New()

	fmt.Println(pubKey)
	fmt.Println(ripemd)
	
	return nil, ""
}