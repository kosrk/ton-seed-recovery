package main

import (
	"context"
	"fmt"
	"github.com/tonkeeper/tongo"
	"github.com/tonkeeper/tongo/liteapi"
	"github.com/tonkeeper/tongo/tlb"
	"github.com/tonkeeper/tongo/wallet"
	"strings"
)

func contains(s string, sl []string) bool {
	for i := range sl {
		if sl[i] == s {
			return true
		}
	}
	return false
}

func copySeed(old []string) []string {
	n := make([]string, len(old))
	copy(n, old)
	return n
}

func insertEmpty(seed []string, n int) []string {
	seed2 := copySeed(seed)
	if len(seed2) == n { // nil or empty slice or after last element
		return append(seed2, "")
	}
	seed2 = append(seed2[:n+1], seed2[n:]...) // index < len(a)
	seed2[n] = ""
	return seed2
}

func bruteforce(seed []string, wordNumber int, client *liteapi.Client) bool {
	for _, w := range wallet.WORDLIST {
		seed[wordNumber] = w
		if ok, _ := checkSeed(seed, client); ok {
			return true
		}
	}
	return false
}

var checksCounter = 0

func checkSeed(seed []string, client *liteapi.Client) (bool, *tongo.AccountID) {
	checksCounter++
	if checksCounter%1000 == 0 {
		fmt.Printf("%v interations\n", checksCounter)
	}
	addresses, err := toAddresses(seed)
	if err != nil {
		return false, nil
	}

	for _, a := range addresses {
		state, err := client.GetAccountState(context.TODO(), a)
		if err != nil {
			continue
		}
		if state.Account.Status() == tlb.AccountActive || state.Account.Status() == tlb.AccountUninit {
			return true, &a
		}
	}
	return false, nil
}

func toAddresses(seed []string) ([]tongo.AccountID, error) {
	key, err := wallet.SeedToPrivateKey(strings.Join(seed, " "))
	if err != nil {
		return nil, err
	}
	w4, err := wallet.New(key, wallet.V4R2, nil)
	if err != nil {
		return nil, err
	}
	w3, err := wallet.New(key, wallet.V3R2, nil)
	if err != nil {
		return nil, err
	}
	w5, err := wallet.New(key, wallet.V5R1, nil)
	if err != nil {
		return nil, err
	}
	return []tongo.AccountID{w5.GetAddress(), w4.GetAddress(), w3.GetAddress()}, nil
}
