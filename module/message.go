package module

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"time"
)

type Message struct {
	Action  string
	Content []byte
}

// use gob to
func (_msg *Message) ToBuffer() []byte {
	var buf bytes.Buffer
	encoder := gob.NewEncoder(&buf)
	err := encoder.Encode(_msg)
	if err != nil {
		panic(err.Error())
	}
	bf := buf.Bytes()
	return bf
}

func ResolveMessage(_data []byte) *Message {
	_msg := &Message{}
	reader := bytes.NewReader(_data)

	decoder := gob.NewDecoder(reader)
	err := decoder.Decode(_msg)
	if err != nil {
		fmt.Println("ResolveMessage  err ", err.Error())
	}
	return _msg
}

func Chat(_content []byte) *Message {
	// timestamp := []byte("12:12:12")
	timestamp := []byte(time.Now().Format("  15:04:05  "))
	return &Message{
		Action:  "chat",
		Content: append(timestamp, _content...),
	}
}

func Quit() *Message {
	return &Message{
		Action:  "quit",
		Content: []byte(""),
	}
}

func Rename(_content []byte) *Message {
	return &Message{
		Action:  "rename",
		Content: _content,
	}
}

func List() *Message {
	return &Message{
		Action:  "list",
		Content: []byte(""),
	}
}
