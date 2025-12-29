package params

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common"
)

// CHANGE(hashkey): hashkey relate params.
var (
	// The BlackList contract.
	HashKeyBlacklist = common.HexToAddress("0x4200000000000000000000000000000000000100")

	// The slot number of blackList in the BlackList contract.
	BlackListSlotNumber = func() *big.Int {
		slot, ok := new(big.Int).SetString("0x4ce4469ea4e93d3a899cd1b7c07c17b40c36ac473b8622212e32ec1ea12c3000", 0)
		if !ok {
			panic("initialize BlackListSlotNumber failed")
		}
		return slot
	}()

	// The sandbox policy contracts.
	SandboxPolicyAddress = common.HexToAddress("0x4200000000000000000000000000000000000101")

	// The slot number of trustContractList in the policy contract.
	TrustContractSlot = func() *big.Int {
		slot, ok := new(big.Int).SetString("0xa05566b51c5978a75b6e841561de44fd654454533c2aa3dac6035f0ba50b6d00", 0)
		if !ok {
			panic("initialize TrustContractSlot failed")
		}
		return slot
	}()

	// The slot number of sandboxBoundaryExceptions in the policy contract.
	SandboxBoundaryExceptionsSlot = func() *big.Int {
		slot, ok := new(big.Int).SetString("0xa05566b51c5978a75b6e841561de44fd654454533c2aa3dac6035f0ba50b6d01", 0)
		if !ok {
			panic("initialize SandboxBoundaryExceptionsSlot failed")
		}
		return slot
	}()
)
