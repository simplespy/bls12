package bls12

import (
	"testing"
)

func TestHalfPoint(t *testing.T) {
	for i := 1; i < 10000; i+=2 {
		t1 := new(Fq).SetInt64(int64(i)+124124)
		fouqueHalfPoint(t1)
	}
}

