package params

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common"
)

// CHANGE(hashkey): hashkey contracts
var (
	// The sandbox policy contracts.
	SandboxPolicyAddress = common.HexToAddress("0x4200000000000000000000000000000000000426")

	// The slot number of trustContractList in the policy contract.
	TrustContractSlot = big.NewInt(0)

	// The slot number of kycList in the policy contract.
	KycSlot = big.NewInt(1)
)
