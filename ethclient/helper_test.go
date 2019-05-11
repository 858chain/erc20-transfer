package ethclient

import (
	"math/big"
	"testing"
)

func TestfloatToBigInt(t *testing.T) {
	if floatToBigInt(0.1, 10) == big.NewInt(1) {
		t.Errorf("floatToBigInt(0.1, 10) != big.NewInt(1) ")
	}
}
