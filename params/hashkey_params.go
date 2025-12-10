package params

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common"
)

// CHANGE(hashkey): hashkey relate params.
var (
	// The sandbox policy contracts.
	SandboxPolicyAddress = common.HexToAddress("0xB189CAF4D787c991E767e02f054da5208428f218")

	// The slot number of trustContractList in the policy contract.
	TrustContractSlot = big.NewInt(0)

	// The slot number of blackList in the policy contract.
	BlackListSlotNumber = big.NewInt(2)
)
