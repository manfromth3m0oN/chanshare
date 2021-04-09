package main

import (
	"log"

	"github.com/yourok/go-mpv/mpv"
)

func pause() {
	switch pauseTextState{
	case 0: pauseTextState = 1
	case 1: pauseTextState = 0
	}
	state, err := m.GetProperty("pause", mpv.FORMAT_FLAG)
	if err != nil {
		log.Fatalf("Unable to get pause state: %v", err)
	}
	err = m.SetProperty("pause", mpv.FORMAT_FLAG, !(state.(bool)))
	if err != nil {
		log.Fatalf("Unable to pause playback: %v", err)
	}
}
