package params

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common"
)

// CHANGE(hashkey): hashkey contracts
var (
	// The sandbox policy contracts.
	SandboxPolicyAddress = common.HexToAddress("0x4200000000000000000000000000000000000425")

	// The slot number of trustContractList in the policy contract.
	TrustContractSlot = big.NewInt(0)
)
