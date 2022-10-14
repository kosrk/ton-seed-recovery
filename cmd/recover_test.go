package main

import (
	"github.com/startfellows/tongo/liteclient"
	"strings"
	"testing"
)

const (
	_VALIDSEED   = "purpose pistol subject pact panic citizen door treat citizen pact doctor panic elder sport doctor elder panic doctor universe citizen citizen embark chunk door"
	_SEED23      = "purpose subject pact panic citizen door treat citizen pact doctor panic elder sport doctor elder panic doctor universe citizen citizen embark chunk door"
	_SEED24      = "purpose test subject pact panic citizen door treat citizen pact doctor panic elder sport doctor elder panic doctor universe citizen citizen embark chunk door"
	_BADWORDSEED = "purpose blah subject pact account citizen door treat citizen doctor panic elder sport doctor elder panic doctor universe citizen citizen embark chunk door"
)

func TestRecoverSeed23Words(t *testing.T) {
	checksCounter = 0
	client, err := liteclient.NewClientWithDefaultTestnet()
	if err != nil {
		panic(err)
	}
	seed, err := recoverSeed(_SEED23, client)
	if err != nil {
		panic(err)
	}
	if seed != _VALIDSEED {
		panic("recovered seed not valid")
	}
}

func TestRecoverSeed24Words(t *testing.T) {
	checksCounter = 0
	client, err := liteclient.NewClientWithDefaultTestnet()
	if err != nil {
		panic(err)
	}
	seed, err := recoverSeed(_SEED24, client)
	if err != nil {
		panic(err)
	}
	if seed != _VALIDSEED {
		panic("recovered seed not valid")
	}
}

func TestRecoverBadWordSeed(t *testing.T) {
	_, err := recoverSeed(_BADWORDSEED, nil)
	println(err.Error())
	if !strings.Contains(err.Error(), "invalid word") {
		panic(err)
	}
}
