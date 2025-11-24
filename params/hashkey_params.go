package params

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common"
)

// CHANGE(hashkey): hashkey relate params.
var (
	// The sandbox router contract.
	SandboxRouterAddress = common.HexToAddress("0x4200000000000000000000000000000000000425")

	// The sandbox policy contract.
	SandboxPolicyAddress = common.HexToAddress("0x4200000000000000000000000000000000000426")

	// The slot number of blackList in the policy contract.
	BlackListSlotNumber = big.NewInt(3)
)
