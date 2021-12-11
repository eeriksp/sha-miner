package main

import (
	"crypto/sha256"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
	"sync"
	"sync/atomic"
)
var stdout = log.New(os.Stdout, "", log.Ldate | log.Ltime)
var stderr = log.New(os.Stdout, "", log.Ldate | log.Ltime)

var smallestHashMutex sync.RWMutex
var smallestHash = []byte("ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff")
var content string = "Foo"
var triesCount uint64

func main() {
	var wg sync.WaitGroup
	for _, v := range "0123456789abcdef" {
		wg.Add(1)
		go mine(string(v))
	}
	wg.Wait()
}

func mine(nonssFirstLetter string) {
	for true {
		nons := generateRandomNonss(nonssFirstLetter)
		hash := sha256.Sum256([]byte(content + nons))
		updateTriesCount()
		if isSmallestHash(hash[:]) {
			updateSmallestHash(nons, hash[:])
		}
	}
}

func updateTriesCount() {
	atomic.AddUint64(&triesCount, 1)
	currentCount := atomic.LoadUint64(&triesCount)
	if currentCount%10000000 == 0 {
		stderr.Println("Num of tries: " + strconv.FormatUint(currentCount, 10))
	}
}

func intToHex(n uint32) string {
	return strconv.FormatInt(int64(n), 16)
}

func generateRandomNonss(nonssFirstLetter string) string {
	return nonssFirstLetter + intToHex(rand.Uint32()) + intToHex(rand.Uint32()) + intToHex(rand.Uint32()) + intToHex(rand.Uint32())
}

func isSmallestHash(hash1 []byte) bool {
	smallestHashMutex.RLock()
	defer smallestHashMutex.RUnlock()
	return hashToInt(hash1) < hashToInt(smallestHash)
}

func hashToInt(hex []byte) uint64 {
	r, err := strconv.ParseInt(fmt.Sprintf("%x", hex[:5]), 16, 64)
	if err != nil {
		panic(err)
	}
	return uint64(r)
}

func updateSmallestHash(nonss string, hash []byte) {
	smallestHashMutex.Lock()
	smallestHash = hash
	smallestHashMutex.Unlock()
	printNewNonss(nonss, hash)
}

func printNewNonss(nonss string, hash []byte) {
	stdout.Println(nonss + " " + string(hash))
	println(hash)
}
