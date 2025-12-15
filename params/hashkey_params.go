package params

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common"
)

// CHANGE(hashkey): hashkey relate params.
var (
	// The sandbox policy contracts.
	SandboxPolicyAddress = common.HexToAddress("0xdbD572bB7E3C92A67Acaddc68452381E99183164")

	// The slot number of trustContractList in the policy contract.
	TrustContractSlot = big.NewInt(0)

	// The slot number of blackList in the policy contract.
	BlackListSlotNumber = big.NewInt(2)
)
