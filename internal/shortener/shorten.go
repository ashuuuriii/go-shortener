package shortener

import (
	"github.com/btcsuite/btcutil/base58"
	"hash/fnv"
)

// Takes a long url and encodes it as an 8 character long string
func ShortenURL(url string) string {
	// FNV is a quick non-cryptographic hashing function with good dispersion
	h := fnv.New64a()
	h.Write([]byte(url))

	// base58 is url safe
	s := base58.Encode(h.Sum(nil))

	return s[:8]
}
