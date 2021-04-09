package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

type Thread struct {
	Posts []struct {
		No          int    `json:"no"`
		Now         string `json:"now"`
		Name        string `json:"name"`
		Sub         string `json:"sub,omitempty"`
		Com         string `json:"com"`
		Filename    string `json:"filename,omitempty"`
		Ext         string `json:"ext,omitempty"`
		W           int    `json:"w,omitempty"`
		H           int    `json:"h,omitempty"`
		TnW         int    `json:"tn_w,omitempty"`
		TnH         int    `json:"tn_h,omitempty"`
		Tim         int64  `json:"tim,omitempty"`
		Time        int    `json:"time"`
		Md5         string `json:"md5,omitempty"`
		Fsize       int    `json:"fsize,omitempty"`
		Resto       int    `json:"resto"`
		Bumplimit   int    `json:"bumplimit,omitempty"`
		Imagelimit  int    `json:"imagelimit,omitempty"`
		SemanticURL string `json:"semantic_url,omitempty"`
		Replies     int    `json:"replies,omitempty"`
		Images      int    `json:"images,omitempty"`
		UniqueIps   int    `json:"unique_ips,omitempty"`
	} `json:"posts"`
}

func getThread(no uint32, board string) Thread {
	url := fmt.Sprintf("https://a.4cdn.org/%s/thread/%d.json", board, no)

	httpClient := http.Client{
		Timeout: time.Second * 10,
	}

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		log.Fatal(err)
	}

	res, getErr := httpClient.Do(req)
	if getErr != nil {
		log.Fatal(getErr)
	}

	if res.Body != nil {
		defer res.Body.Close()
	}

	body, readErr := ioutil.ReadAll(res.Body)
	if readErr != nil {
		log.Fatal(readErr)
	}

	thread := Thread{}
	jsonErr := json.Unmarshal(body, &thread)
	if jsonErr != nil {
		log.Fatal(jsonErr)
	}

	return thread
}

func extractMedia(t Thread, board string) []string {
	var media []string
	for _, post := range t.Posts {
		url := fmt.Sprintf("https://i.4cdn.org/%s/%d.webm", board, post.Tim)
		media = append(media, url)
	}
	return media
}
