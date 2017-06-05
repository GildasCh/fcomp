package fcomp

import (
	"encoding/base64"
	"io"

	blake2b "github.com/minio/blake2b-simd"
)

const block = 1 * 1024 * 1024 // 10MB

func Hash(r io.Reader) string {
	buf := make([]byte, block)

	h := blake2b.New512()

	var err error
	for err == nil {
		var n int
		n, err = r.Read(buf)

		if n == 0 {
			break
		}

		h.Write(buf)
	}

	a := h.Sum(nil)
	return base64.StdEncoding.EncodeToString(a[:])
}

func Hashes(r io.Reader) []string {
	var ret []string

	buf := make([]byte, block)

	var err error
	for err == nil {
		var n int
		n, err = r.Read(buf)

		if n == 0 {
			break
		}

		a := blake2b.Sum512(buf)
		ret = append(ret, base64.StdEncoding.EncodeToString(a[:]))
	}

	return ret
}
