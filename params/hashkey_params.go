package params

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common"
)

// CHANGE(hashkey): hashkey relate params.
var (
	// The BlackList contract.
	HashKeyBlacklist = common.HexToAddress("0x4200000000000000000000000000000000000100")

	// ERC-7201 namespace for BlackListStorage:
	// slot = keccak256(abi.encode(uint256(keccak256("hashkeychain.rescue-capsule.BlackList")) - 1)) & ~bytes32(uint256(0xff))).
	// computed slot = 0x4ce4469ea4e93d3a899cd1b7c07c17b40c36ac473b8622212e32ec1ea12c3000
	// Field offsets within BlackListStorage:
	// - blackList: slot + 0
	BlackListSlotNumber = func() *big.Int {
		slot, ok := new(big.Int).SetString("0x4ce4469ea4e93d3a899cd1b7c07c17b40c36ac473b8622212e32ec1ea12c3000", 0)
		if !ok {
			panic("initialize BlackListSlotNumber failed")
		}
		return slot
	}()

	// The sandbox policy contracts.
	SandboxPolicyAddress = common.HexToAddress("0x4200000000000000000000000000000000000101")

	// ERC-7201 namespace for CompliancePolicyStorage:
	// slot = keccak256(abi.encode(uint256(keccak256("hashkeychain.compliance.policy.storage")) - 1)) & ~bytes32(uint256(0xff))).
	// computed slot = 0xa05566b51c5978a75b6e841561de44fd654454533c2aa3dac6035f0ba50b6d00
	// Field offsets within CompliancePolicyStorage:
	// - trustContractList: slot + 0
	// - sandboxBoundaryExceptions: slot + 1

	// CompliancePolicyStorage.trustContractList (ERC-7201 namespace slot + 0).
	TrustContractSlot = func() *big.Int {
		slot, ok := new(big.Int).SetString("0xa05566b51c5978a75b6e841561de44fd654454533c2aa3dac6035f0ba50b6d00", 0)
		if !ok {
			panic("initialize TrustContractSlot failed")
		}
		return slot
	}()

	// CompliancePolicyStorage.sandboxBoundaryExceptions (ERC-7201 namespace slot + 1).
	SandboxBoundaryExceptionsSlot = func() *big.Int {
		slot, ok := new(big.Int).SetString("0xa05566b51c5978a75b6e841561de44fd654454533c2aa3dac6035f0ba50b6d01", 0)
		if !ok {
			panic("initialize SandboxBoundaryExceptionsSlot failed")
		}
		return slot
	}()
)
