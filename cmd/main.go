package main

import (
	"fmt"
	"github.com/tonkeeper/tongo"
	"github.com/tonkeeper/tongo/liteapi"
	"github.com/tonkeeper/tongo/wallet"
	"os"
	"strings"
)

var SEED = "" // TODO: set invalid seed here or as ENV variable

func main() {
	if len(SEED) == 0 {
		SEED = os.Getenv("SEED")
	}
	if len(SEED) == 0 {
		fmt.Printf("Need to set the seed in the code or as ENV variable")
		os.Exit(1)
	}
	client, err := liteapi.NewClientWithDefaultMainnet()
	if err != nil {
		panic(err)
	}
	seed, addr, err := recoverSeed(SEED, client)
	if err != nil {
		fmt.Printf("%v", err)
		os.Exit(1)
	}
	fmt.Printf("Valid seed: %s\nAddress: %s\n", seed, addr.ToHuman(false, false))
	os.Exit(0)
}

func recoverSeed(seedString string, client *liteapi.Client) (string, *tongo.AccountID, error) {
	seed := strings.Split(seedString, " ")
	if len(seed) < 23 || len(seed) > 24 {
		return "", nil, fmt.Errorf("len of seed must be 23 or 24 words")
	}

	for _, word := range seed {
		if !contains(word, wallet.WORDLIST) {
			return "", nil, fmt.Errorf("invalid word: %s", word)
		}
	}

	if len(seed) == 24 {
		valid, addr := checkSeed(seed, client)
		if valid {
			fmt.Printf("You seed is valid.\n")
			return seedString, addr, nil
		}
		for i := range seed {
			seed2 := copySeed(seed)
			if bruteforce(seed2, i, client) {
				_, addr := checkSeed(seed2, client)
				return strings.Join(seed2, " "), addr, nil
			}
		}
		return "", nil, fmt.Errorf("can not find valid seed")
	}

	for i := 0; i < 24; i++ {
		seed2 := insertEmpty(seed, i)
		if bruteforce(seed2, i, client) {
			_, addr := checkSeed(seed2, client)
			return strings.Join(seed2, " "), addr, nil
		}
	}
	return "", nil, fmt.Errorf("can not find valid seed")
}
