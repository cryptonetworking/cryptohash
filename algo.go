package cryptohash

import "github.com/multiformats/go-multihash"

type Algo uint64

const (
	SHA256 Algo = multihash.SHA2_256
	SHA512 Algo = multihash.SHA2_512
)
