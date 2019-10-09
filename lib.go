package gglib

import (
	"time"
	"github.com/pkg/errors"
)

func GenerateAnnoyingErrors(e chan error) {
	ticker := time.NewTicker(5 * time.Second)
	for {
		select {
		case <-ticker.C:
			e <- errors.New("Annoying Error message, something went wrong")
		}
	}
}
