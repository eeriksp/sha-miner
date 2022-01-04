package main

import (
	"crypto/sha256"
	"flag"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
	"sync"
	"sync/atomic"
)

var stdout = log.New(os.Stdout, "", log.Ldate|log.Ltime)
var stderr = log.New(os.Stdout, "", log.Ldate|log.Ltime)

var smallestHashMutex sync.RWMutex
var smallestHash = []byte("ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff")
var triesCount uint64

func main() {
	content, numOfGoroutines := parseArgs()
	var wg sync.WaitGroup
	for i := 0; i < int(numOfGoroutines); i++ {
		wg.Add(1)
		go mine(content)
	}
	wg.Wait()
}

func parseArgs() (content string, threadsCount uint) {
	h := flag.Bool("h", false, "Prints this help message.")
	help := flag.Bool("help", false, "Prints this help message.")
	numOfGoroutines := flag.Uint("threads", 16, "The number of goroutines the the task will divided between.")
	flag.Parse()
	if flag.Arg(0) == "" || *h || *help {
		println("Usage: miner <content> [--threads <threads_count>]")
		flag.PrintDefaults()
		os.Exit(1)
	}
	return flag.Arg(0), *numOfGoroutines
}

func mine(content string) {
	for true {
		nons := generateRandomNonss()
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

func generateRandomNonss() string {
	return fmt.Sprintf("%032s", intToHex(rand.Uint32())+intToHex(rand.Uint32())+intToHex(rand.Uint32())+intToHex(rand.Uint32()))
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
	stdout.Println(nonss + " " + fmt.Sprintf("%x", hash))
}
