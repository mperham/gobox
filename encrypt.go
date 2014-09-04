package main

import (
	"code.google.com/p/go.crypto/nacl/box"
	"encoding/json"
	"fmt"
	"io"
	"math/rand"
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
	_, myprv := readKeys("local")
	pub, _ := readKeys("remote")

	doc := Document{
		"mike",
		"mike@contribsys.com",
		"blue",
		47,
		int(rand.Int31()),
	}
	fmt.Println("Encrypting this document")
	fmt.Println(doc)
	data, err := json.Marshal(doc)
	if err != nil {
		panic(err)
	}

	nonce := loadNonce()

	enc := box.Seal(nil, data, nonce, pub, myprv)

	f, err := os.Create("document.bin")
	if err != nil {
		panic(err)
	}
	defer f.Close()
	f.Write(enc)
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
