package main

import (
	"code.google.com/p/go.crypto/nacl/box"
	"crypto/rand"
	"fmt"
	"io"
	"os"
)

func main() {
	keySet("local")
	keySet("remote")
	createNonce()

	fmt.Println("Created your keyset in local")
	fmt.Println("Created distributable keyset in remote")
	fmt.Println("Created nonce")
}

func createNonce() {
	nonce := new([24]byte)
	_, err := io.ReadFull(rand.Reader, nonce[:])
	f, err := os.Create("nonce.bin")
	if err != nil {
		panic(err)
	}
	defer f.Close()
	f.Write(nonce[:])
}

func keySet(name string) (pub, prv *[32]byte) {
	pub, priv, err := box.GenerateKey(rand.Reader)
	if err != nil {
		panic(err.Error())
	}

	fpub, err := os.Create(name + ".pub")
	if err != nil {
		panic(err.Error())
	}
	defer fpub.Close()
	fpub.Write(pub[0:])

	fprv, err := os.Create(name + ".prv")
	if err != nil {
		panic(err.Error())
	}
	defer fprv.Close()
	fprv.Write(priv[0:])

	return pub, priv
}
