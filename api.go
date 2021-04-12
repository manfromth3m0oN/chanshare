package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"time"
)

type Catalog []struct {
	Page    int `json:"page"`
	Threads []struct {
		No          uint32 `json:"no"`
		Sticky      int    `json:"sticky,omitempty"`
		Closed      int    `json:"closed,omitempty"`
		Now         string `json:"now"`
		Name        string `json:"name"`
		Sub         string `json:"sub,omitempty"`
		Com         string `json:"com,omitempty"`
		Filename    string `json:"filename"`
		Ext         string `json:"ext"`
		W           int    `json:"w"`
		H           int    `json:"h"`
		TnW         int    `json:"tn_w"`
		TnH         int    `json:"tn_h"`
		Tim         int64  `json:"tim"`
		Time        int    `json:"time"`
		Md5         string `json:"md5"`
		Fsize       int    `json:"fsize"`
		Resto       int    `json:"resto"`
		Capcode     string `json:"capcode,omitempty"`
		SemanticURL string `json:"semantic_url"`
		Replies     int    `json:"replies"`
		Images      int    `json:"images"`
		LastReplies []struct {
			No      int    `json:"no"`
			Now     string `json:"now"`
			Name    string `json:"name"`
			Com     string `json:"com"`
			Time    int    `json:"time"`
			Resto   int    `json:"resto"`
			Capcode string `json:"capcode"`
		} `json:"last_replies"`
		LastModified  int `json:"last_modified"`
		Bumplimit     int `json:"bumplimit,omitempty"`
		Imagelimit    int `json:"imagelimit,omitempty"`
		OmittedPosts  int `json:"omitted_posts,omitempty"`
		OmittedImages int `json:"omitted_images,omitempty"`
	} `json:"threads"`
}

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

	log.Printf("Requested from url: %s", url)

	httpClient := http.Client{
		Timeout: time.Second * 5,
	}

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		log.Fatalf("Error creating new request: %v", err)
	}

	res, getErr := httpClient.Do(req)
	if getErr != nil {
		log.Fatalf("Error executing request: %v", getErr)
	}

	if res.Body != nil {
		defer res.Body.Close()
	}

	body, readErr := ioutil.ReadAll(res.Body)
	if readErr != nil {
		log.Fatalf("Error reading response: %v", readErr)
	}

	thread := Thread{}
	jsonErr := json.Unmarshal(body, &thread)
	if jsonErr != nil {
		log.Fatalf("Error marshaling request into json: %v", jsonErr)
	}

	return thread
}

func getThreads(board string) Catalog {
	url := fmt.Sprintf("https://a.4cdn.org/%s/catalog.json", board)

	log.Printf("Requested from url: %s", url)

	httpClient := http.Client{
		Timeout: time.Second * 5,
	}

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		log.Fatalf("Error creating new request: %v", err)
	}

	res, getErr := httpClient.Do(req)
	if getErr != nil {
		log.Fatalf("Error executing request: %v", getErr)
	}

	if res.Body != nil {
		defer res.Body.Close()
	}

	body, readErr := ioutil.ReadAll(res.Body)
	if readErr != nil {
		log.Fatalf("Error reading response: %v", readErr)
	}

	catalog := Catalog{}
	jsonErr := json.Unmarshal(body, &catalog)
	if jsonErr != nil {
		log.Fatalf("Error unmarshaling Board data: %v", err)
	}

	return catalog
}

func getRandomThread() {
	catalog := getThreads("gif")
	rn := rand.Intn(len(catalog[0].Threads)-1) + 1
	id := catalog[0].Threads[rn].No
	thread := getThread(id, "gif")
	media = extractMedia(thread, "gif")
}

func extractMedia(t Thread, board string) []string {
	var media []string
	for _, post := range t.Posts {
		if post.Tim == 0 {
			media = append(media, "")
		} else {
			url := fmt.Sprintf("https://i.4cdn.org/%s/%d%s", board, post.Tim, post.Ext)
			media = append(media, url)
		}
	}
	return media
}
