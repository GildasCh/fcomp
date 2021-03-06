package fcomp

import (
	"fmt"
	"os"
	"testing"
)

const (
	smallFile  = "testFiles/smallFile"  // 52K
	mediumFile = "testFiles/mediumFile" // 70M
	bigFile    = "testFiles/bigFile"    // 350M
)

func TestHashSmall(t *testing.T) {
	f, _ := os.Open(smallFile)

	h := Hash(f)

	fmt.Println(h)
}

func TestHashBig(t *testing.T) {
	f, _ := os.Open(bigFile)

	h := Hash(f)

	fmt.Println(h)
}

func BenchmarkHashSmall(b *testing.B) {
	f, _ := os.Open(smallFile)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Hash(f)
	}

}

func BenchmarkHashMedium(b *testing.B) {
	f, _ := os.Open(mediumFile)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Hash(f)
	}

}

func BenchmarkHashBig(b *testing.B) {
	f, _ := os.Open(bigFile)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Hash(f)
	}
}

func TestHashesSmall(t *testing.T) {
	f, _ := os.Open(smallFile)

	h := Hashes(f)

	fmt.Println(h)
}

func TestHashesBig(t *testing.T) {
	f, _ := os.Open(bigFile)

	Hashes(f)

	// fmt.Println(h)
}

func BenchmarkHashesSmall(b *testing.B) {
	f, _ := os.Open(smallFile)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Hashes(f)
	}

}

func BenchmarkHashesMedium(b *testing.B) {
	f, _ := os.Open(mediumFile)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Hashes(f)
	}

}

func BenchmarkHashesBig(b *testing.B) {
	f, _ := os.Open(bigFile)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Hashes(f)
	}
}
