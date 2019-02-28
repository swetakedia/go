package ethereum

import (
	"math/big"

	"github.com/hcnet/go/services/bifrost/common"
)

func (t Transaction) ValueToHcnet() string {
	valueEth := new(big.Rat)
	valueEth.Quo(new(big.Rat).SetInt(t.ValueWei), weiInEth)
	return valueEth.FloatString(common.HcnetAmountPrecision)
}
