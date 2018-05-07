package data

import (
	"encoding/json"
	"github.com/magiconair/properties/assert"
	"testing"
)
// To test everything: go test ./data
// go test ./data -run Test_ParseMessage_IsImage -v
func Test_ParseMessage_IsImage(t *testing.T) {
	testUrl := "https://images.pexels.com/photos/67636/rose-blue-flower-rose-blooms-67636.jpeg"
	m := &Message{
		Message:  testUrl,
		Username: "jackliu97",
	}

	ParseMessage(m)

	expected, _ := json.Marshal(ImageMessage{
		Width:  200,
		Height: 100,
		Url:    testUrl,
	})

	assert.Equal(t, m.Message, string(expected))
}

// go test ./data -run Test_ParseMessage_IsVideo -v
func Test_ParseMessage_IsVideo(t *testing.T) {
	testUrl := "https://images.pexels.com/photos/67636/rose-blue-flower-rose-blooms-67636.mp4"
	m := &Message{
		Message:  testUrl,
		Username: "jackliu97",
	}

	ParseMessage(m)

	expected, _ := json.Marshal(VideoMessage{
		Length:  120,
		Source: "YouTube",
		Url:    testUrl,
	})

	assert.Equal(t, m.Message, string(expected))
}

// go test ./data -run Test_ParseMessage_OtherURL -v
func Test_ParseMessage_OtherURL(t *testing.T) {
	testUrl := "https://google.com"
	m := &Message{
		Message:  testUrl,
		Username: "jackliu97",
	}

	ParseMessage(m)

	expected := testUrl

	assert.Equal(t, m.Message, expected)
}

// go test ./data -run Test_ParseMessage_NormalMessage -v
func Test_ParseMessage_NormalMessage(t *testing.T) {
	testUrl := "A normal message"
	m := &Message{
		Message:  testUrl,
		Username: "jackliu97",
	}

	ParseMessage(m)

	expected := testUrl

	assert.Equal(t, m.Message, expected)
}
