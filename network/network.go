package network

import (
	"bytes"
	"encoding/gob"
)

type NetProtocol interface {
	Listener(_address string)
	EstablishConnection(_address string)
	SendData(_data []byte) int
	ReceiveData() []byte
}

var (
	Buffer = bytes.Buffer{}
)

func ToGob(s any) (b []byte) {
	encoder := gob.NewEncoder(&Buffer)
	err := encoder.Encode(s)
	if err != nil {
		panic(err.Error())
	}
	b = Buffer.Bytes()
	Buffer.Reset()
	return
}

func ParseGob(b []byte, s any) {
	reader := bytes.NewReader(b)
	decoder := gob.NewDecoder(reader)
	err := decoder.Decode(s)
	if err != nil {
		panic(err.Error())
	}
}
