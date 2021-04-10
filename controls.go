package main

import (
	"log"
	"strconv"

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

func loadThread(thread string, board string) {
	threadId, err := strconv.ParseInt(thread, 10, 32)
	if err != nil {
		log.Fatalf("Couldnt parse threadId to int: %v", err)
	}
	requesting = true
	var threadStruct Thread
	threadStruct = getThread(uint32(threadId), board)
	media = []string{}
	media = extractMedia(threadStruct, board)
	requesting = false
	mediaPos = 0
	loadFile(media[mediaPos])

}

func next() {
	log.Printf("Current mediaPos: %d", mediaPos)
	if mediaPos > len(media) {
		mediaPos = 0
	} else {
		mediaPos = mediaPos + 1
	}
	mediaPos += 1
	log.Printf("New mediaPos: %d", mediaPos)
	log.Println(media)
	loadFile(media[mediaPos])
}

func prev() {
	log.Printf("Current mediaPos: %d", mediaPos)
	if mediaPos < 0 {
		mediaPos = 0
	} else {
		mediaPos -= 1 
	}
	log.Printf("New mediaPos: %d", mediaPos)
	log.Println(media)
	loadFile(media[mediaPos])
}
