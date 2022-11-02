package main

import (
	"log"
)

func FailIfError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
