package cryptohash

import (
	"bytes"
	"github.com/multiformats/go-multihash"
)

type Hash []byte

func (h Hash) String() string {
	return multihash.Multihash(h).B58String()
}

func FromString(str string) (Hash, error) {
	h, err := multihash.FromB58String(str)
	return Hash(h), err
}

func (h Hash) Equal(data []byte) bool {
	decoded, err := multihash.Decode(h)
	if err != nil {
		return false
	}
	h2, err := multihash.Sum(data, decoded.Code, -1)
	if err != nil {
		return false
	}
	return bytes.Equal(h2, h)

}

func Sum(data []byte, algo Algo) []byte {
	hash, err := multihash.Sum(data, uint64(algo), -1)
	if err != nil {
		panic(err)
	}
	return hash
}
