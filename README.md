# Stupid Crypto Tricks with Go's NaCl API

Go's crypto APIs are **way** too low-level.  This is high-level
example code showing you how to:

* Create a pair of public and private keys.
* Sign/encrypt ("seal") a document with your private key.
* Decrypt/verify ("open") a document with your public key.

It uses NaCl, [read the documentation](http://godoc.org/code.google.com/p/go.crypto/nacl/box).


# Usage

```go
go run setup.go
go run encrypt.go
go run decrypt.go
```

# License

Public Domain

# Author

Mike Perham, [Contributed Systems](http://contribsys.com)

Note: I'm not a crypto expert, please don't contact me with
crypto questions.  I'm a poor slob who spent hours figuring this
out and want to save other poor slobs from the same fate.
