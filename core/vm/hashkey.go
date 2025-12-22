package vm

import (
	"fmt"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/log"
	"github.com/ethereum/go-ethereum/params"
)

// CHANGE(hashkey): Sandbox type constants for penetration detection.
const (
	SandboxTypeTransparent uint8 = 0 // Transparent: doesn't affect detection
	SandboxTypeTrustList   uint8 = 1 // TrustList: sandbox core
	SandboxTypeOther       uint8 = 2 // Other: outside sandbox
)

// CHANGE(hashkey): Check whether the call would penetrate sandbox boundary.
func (evm *EVM) SandboxPenetrateCheck(addr common.Address) error {
	// If the address is a precompiled, skip the check.
	_, isPrecompile := evm.precompile(addr)
	if isPrecompile {
		return nil
	}

	// If no code at address, skip the check.
	if len(evm.resolveCode(addr)) == 0 {
		return nil
	}

	preContract := evm.preContract
	evm.preContract = addr

	// Determine the type of target address
	addrType := evm.getSandboxType(addr)

	// Apply OR operation (GreyList=0 is transparent)
	evm.sandboxFlag |= addrType

	// If sandboxFlag == 3, both TrustList(1) and Other(2) appeared in the call chain
	if evm.sandboxFlag == 3 {
		log.Error("Detect sandbox penetrate", "from", preContract.Hex(), "to", addr.Hex(), "flag", evm.sandboxFlag)
		return fmt.Errorf("detect sandbox penetrate, from: %s, to: %s", preContract.Hex(), addr.Hex())
	}

	return nil
}

// CHANGE(hashkey): getSandboxType returns the sandbox type of an address.
func (evm *EVM) getSandboxType(addr common.Address) uint8 {
	if evm.isSandboxTrustedContract(addr) {
		return SandboxTypeTrustList
	}
	if evm.isGreyListContract(addr) {
		return SandboxTypeTransparent
	}
	return SandboxTypeOther
}

// CHANGE(hashkey): isSandboxTrustedContract checks whether the address is in the TrustContractList.
func (evm *EVM) isSandboxTrustedContract(addr common.Address) bool {
	// Correct storage slot calculation for mapping(address => bool) at slot TrustContractSlot.
	var buf [64]byte
	// Left-pad address to 32 bytes (address is 20 bytes, pad first 12 bytes with 0)
	copy(buf[12:32], addr[:])
	// Slot as 32-byte big-endian
	slot := params.TrustContractSlot.Bytes()
	copy(buf[64-len(slot):], slot)
	hash := crypto.Keccak256Hash(buf[:])
	val := evm.StateDB.GetState(params.SandboxPolicyAddress, hash)
	return val != (common.Hash{})
}

// CHANGE(hashkey): isGreyListContract checks whether the address is in the GreyList.
func (evm *EVM) isGreyListContract(addr common.Address) bool {
	// Correct storage slot calculation for mapping(address => bool) at slot GreyContractSlot.
	var buf [64]byte
	// Left-pad address to 32 bytes (address is 20 bytes, pad first 12 bytes with 0)
	copy(buf[12:32], addr[:])
	// Slot as 32-byte big-endian
	slot := params.SandboxBoundaryExceptionsSlot.Bytes()
	copy(buf[64-len(slot):], slot)
	hash := crypto.Keccak256Hash(buf[:])
	val := evm.StateDB.GetState(params.SandboxPolicyAddress, hash)
	return val != (common.Hash{})
}

// CHANGE(hashkey): InitSandboxPenetrateCheck init the sandbox penetration check for the tx.
func (evm *EVM) InitSandboxPenetrateCheck(addr common.Address) {
	evm.preContract = addr
	// Set initial sandboxFlag based on entry address type
	evm.sandboxFlag = evm.getSandboxType(addr)
}

// CHANGE(hashkey): check whether the address is in the blacklist or not.
func (evm *EVM) IsBlackListAddress(addr common.Address) bool {
	// Correct storage slot calculation for mapping(address => bool) at slot BlackListSlotNumber
	var buf [64]byte
	// Left-pad address to 32 bytes (address is 20 bytes, pad first 12 bytes with 0)
	copy(buf[12:32], addr[:])
	// Slot as 32-byte big-endian
	slot := params.BlackListSlotNumber.Bytes()
	copy(buf[64-len(slot):], slot)
	hash := crypto.Keccak256Hash(buf[:])
	val := evm.StateDB.GetState(params.HashKeyBlacklist, hash)
	return val != (common.Hash{})
}
