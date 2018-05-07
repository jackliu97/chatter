package data

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"path/filepath"
)

var (
	videos = map[string]bool{
		".mpeg": true,
		".mpg":  true,
		".avi":  true,
		".mp4":  true,
	}

	images = map[string]bool{
		".jpeg": true,
		".jpg":  true,
		".gif":  true,
		".tiff": true,
		".bmp":  true,
	}
)

// Define our message object
type Message struct {
	Username string `json:"username"`
	Message  string `json:"message"`
}

type ImageMessage struct {
	Width  int    `json:"width"`
	Height int    `json:"height"`
	Url    string `json:"url"`
}

type VideoMessage struct {
	Length int    `json:"length"` // in seconds
	Source string `json:"source"`
	Url    string `json:"url"`
}

// ParseMessage checks if the message is an image or video
func ParseMessage(msg *Message) {
	// check if message is a url, if not we exit early.
	url, err := url.ParseRequestURI(msg.Message)
	if err != nil {
		fmt.Printf("ParseRequestURI : %s", err)
		return
	}

	_, err = http.Get(url.String())
	if err != nil {
		// not a proper url
		fmt.Printf("Get : %s", err)
		return
	}

	if isVideo(url) {
		newMsg, err := json.Marshal(VideoMessage{
			Length: 120,
			Source: "YouTube",
			Url:    url.String(),
		})

		if err != nil {
			fmt.Printf("Marshall video : %s", err)
			return
		}

		msg.Message = string(newMsg)
	} else if isImage(url) {
		newMsg, err := json.Marshal(ImageMessage{
			Height: 100,
			Width:  200,
			Url:    url.String(),
		})

		if err != nil {
			fmt.Printf("Marshall video : %s", err)
			return
		}

		msg.Message = string(newMsg)
	}
}

func isVideo(url *url.URL) bool {
	ext := filepath.Ext(url.Path)

	_, ok := videos[ext]
	return ok
}

func isImage(url *url.URL) bool {
	ext := filepath.Ext(url.Path)

	_, ok := images[ext]
	return ok
}
