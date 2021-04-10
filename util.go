package main

import "log"

func loadFile(file string) {
	err := m.Command([]string{"loadfile", file})
	if err != nil {
		log.Fatalf("Failed to load file: %v",err)
	}
}
