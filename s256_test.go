package ecc

import "testing"

func TestGetS256(t *testing.T) {
	curve := S256()
	t.Log(curve)
}
