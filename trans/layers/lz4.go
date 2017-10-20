package layers

import (
	"github.com/bkaradzic/go-lz4"
)

func LZ4Encode(data []byte) ([]byte, error) {
	return lz4.Encode(nil, data)
}

func LZ4Decode(data []byte) ([]byte, error) {
	return lz4.Decode(nil, data)
}
