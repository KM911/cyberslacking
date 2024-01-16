package network

// 这个是 client 发送的数据包格式

type Payload struct {
	index uint64
	data  []byte
}
