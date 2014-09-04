package main

import (
	"code.google.com/p/go.crypto/nacl/box"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

type Document struct {
	Name   string
	Email  string
	Color  string
	FavNum int
	Nonce  int
}

func main() {
	mypub, _ := readKeys("local")
	_, prv := readKeys("remote")

	enc, err := ioutil.ReadFile("document.bin")
	if err != nil {
		panic(err)
	}
	nonce := loadNonce()

	dec, ok := box.Open(nil, enc, nonce, mypub, prv)
	if !ok {
		panic("invalid")
	}

	newdoc := &Document{}
	json.Unmarshal(dec, newdoc)
	fmt.Println("Decrypted this document")
	fmt.Println(newdoc)
}

func loadNonce() *[24]byte {
	nonce := new([24]byte)
	f, err := os.Open("nonce.bin")
	if err != nil {
		panic(err)
	}
	defer f.Close()
	_, err = io.ReadFull(f, nonce[:])
	return nonce
}

func readKeys(name string) (pub, prv *[32]byte) {
	pub = new([32]byte)
	f, err := os.Open(name + ".pub")
	if err != nil {
		panic(err)
	}
	_, err = io.ReadFull(f, pub[:])
	if err != nil {
		panic(err)
	}

	prv = new([32]byte)
	f, err = os.Open(name + ".prv")
	if err != nil {
		panic(err)
	}
	_, err = io.ReadFull(f, prv[:])
	if err != nil {
		panic(err)
	}
	return
}
