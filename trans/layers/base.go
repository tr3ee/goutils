package layers

import (
	"bytes"
	"encoding/binary"
	"errors"
	"tr3e/utils/cipher"
)

const SizeOfDataHeader = 36

type DataHeader struct {
	Length    uint32 //32byte
	Signature []byte //256byte
}

func NewDataHeader(data []byte) ([]byte, error) {
	dh := new(DataHeader)
	dh.Length, dh.Signature = uint32(len(data)), DHash(data)
	if dh.Length <= 0 || len(dh.Signature) != SizeOfDataHeader-4 {
		return nil, errors.New("data size is 0 or cannot get the signature from data")
	}
	buf := make([]byte, 4)
	binary.BigEndian.PutUint32(buf, dh.Length)
	return append(buf, dh.Signature...), nil
}

func GetDataHeader(data []byte) (*DataHeader, error) {
	length := SizeOfDataHeader
	if len(data) < length {
		return nil, errors.New("data is too short to parse")
	}

	//parse header
	header := data[:length]
	dlen := binary.BigEndian.Uint32([]byte{
		uint8(header[0]),
		uint8(header[1]),
		uint8(header[2]),
		uint8(header[3])})
	dsig := header[4:]

	return &DataHeader{
		Length:    dlen,
		Signature: dsig,
	}, nil
}

func DHash(data []byte) []byte {
	return cipher.Sha256(data)
}

func ZeroPadding(data []byte, blocksize int) []byte {
	padlen := blocksize - len(data)%blocksize
	padtext := bytes.Repeat([]byte{0}, padlen)
	return append(data, padtext...)
}
