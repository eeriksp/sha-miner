package main

import (
	"crypto/sha256"
	"fmt"
	"math/rand"
	"strconv"
)

var smallestHash = []byte("ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff")
var smallestNonss string
var content string = "Foo"
var numTries = 0

func main() {
	for true {
		nons := generateRandomNonss()
		sum := sha256.Sum256([]byte(content + nons))
		if isFirstHashSmaller(sum[:], smallestHash) {
			smallestHash = sum[:]
			smallestNonss = nons
			fmt.Println("New smallest:")
			fmt.Printf(nons)
			fmt.Println()
			fmt.Printf("%x", sum)
			fmt.Println()
		}
		numTries++
		if numTries%10000000 == 0 {
			fmt.Print("Num of tries: ")
			fmt.Println(numTries)
		}

	}
}

func intToHex(n uint32) string {
	return strconv.FormatInt(int64(n), 16)
}

func generateRandomNonss() string {
	return intToHex(rand.Uint32()) + intToHex(rand.Uint32()) + intToHex(rand.Uint32()) + intToHex(rand.Uint32()) + "0"
}

func isFirstHashSmaller(hash1 []byte, hash2 []byte) bool {
	return hashToInt(hash1) < hashToInt(hash2)
}

func hashToInt(hex []byte) uint64 {
	r, err := strconv.ParseInt(fmt.Sprintf("%x", hex[:5]), 16, 64)
	if err != nil {
		panic(err)
	}
	return uint64(r)
}
