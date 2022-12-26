package brmarket

import (
	"errors"
	"testing"
)

func TestIbovIndex(t *testing.T) {
	data, err := GetSharesForIndex("IBOV")
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}
	if data.Index.Symbol != "IBOV" {
		t.Fatalf("Expected IBOV, got %s", data.Index.Symbol)
	}
}

func TestIBXXIndex(t *testing.T) {
	data, err := GetSharesForIndex(IBXX.String())
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}
	if data.Index.Symbol != "IBXX" {
		t.Fatalf("Expected IBXX, got %s", data.Index.Symbol)
	}
}

func TestErrorOnUnknownIndex(t *testing.T) {
	_, err := GetSharesForIndex("unknown")
	if !errors.Is(err, ErrIndexNotFound) {
		t.Fatalf("Unexpected error: %v", err)
	}
}
