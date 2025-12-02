package params

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common"
)

// CHANGE(hashkey): hashkey relate params.
var (
	// The sandbox black contract.
	BlackAddress = common.HexToAddress("0x4200000000000000000000000000000000000425")

	// The slot number of blackList in the policy contract.
	BlackListSlotNumber = big.NewInt(0)
)
