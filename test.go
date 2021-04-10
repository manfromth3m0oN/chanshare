package main

import (
	"log"
	"testing"
)

func TestThread(t *testing.T) {
	thread := getThread(19543747, "gif")
	log.Println(thread)
}
