package main

import (
	"log"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMainAppSuccess(t *testing.T) {
	oldArgs := os.Args
	defer func() { os.Args = oldArgs }()
	os.Args = []string{"cmd", "help"}
	main()
}

func TestMainAppFail(t *testing.T) {
	oldArgs := os.Args
	defer func() { os.Args = oldArgs }()
	os.Args = []string{"cmd", "nonExistantAction"}
	err := startApp()
	if err != nil {
		log.Fatalf("Error found: %+v", err)
	}
	assert.Error(t, err)
}
