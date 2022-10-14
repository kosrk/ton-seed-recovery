package main

import (
	"fmt"
	"github.com/startfellows/tongo/liteclient"
	"github.com/startfellows/tongo/wallet"
	"os"
	"strings"
)

const SEED = ""

func main() {
	client, err := liteclient.NewClientWithDefaultMainnet()
	if err != nil {
		panic(err)
	}
	seed, err := recoverSeed(SEED, client)
	if err != nil {
		fmt.Printf("%v", err)
		os.Exit(1)
	}
	fmt.Printf("Valid seed:\n%s", seed)
	os.Exit(0)
}

func recoverSeed(seedString string, client *liteclient.Client) (string, error) {
	seed := strings.Split(seedString, " ")
	if len(seed) < 23 || len(seed) > 24 {
		return "", fmt.Errorf("can not recover")
	}

	for _, word := range seed {
		if !contains(word, wallet.WORDLIST) {
			return "", fmt.Errorf("invalid word: %s", word)
		}
	}

	if len(seed) == 24 {
		valid := checkSeed(seed, client)
		if valid {
			fmt.Println("you seed is valid")
			return seedString, nil
		}
		for i := range seed {
			seed2 := copySeed(seed)
			if bruteforce(seed2, i, client) {
				return strings.Join(seed2, " "), nil
			}
		}
		return "", fmt.Errorf("can not find valid seed")
	}

	for i := 0; i < 24; i++ {
		seed2 := insertEmpty(seed, i)
		if bruteforce(seed2, i, client) {
			return strings.Join(seed2, " "), nil
		}
	}
	return "", fmt.Errorf("can not find valid seed")
}
