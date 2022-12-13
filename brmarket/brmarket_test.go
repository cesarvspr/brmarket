package brmarket

import (
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
