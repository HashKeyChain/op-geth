package params

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common"
)

// CHANGE(hashkey): hashkey relate params.
var (
	HashKeyBlacklist = common.HexToAddress("0x4200000000000000000000000000000000000100")

	// The slot number of blackList in the policy contract.
	BlackListSlotNumber = big.NewInt(0)

	// The sandbox policy contracts.
	SandboxPolicyAddress = common.HexToAddress("0x4200000000000000000000000000000000000101")

	// The slot number of trustContractList in the policy contract.
	TrustContractSlot = big.NewInt(0)
)
