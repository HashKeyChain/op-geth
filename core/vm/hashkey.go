package vm

import (
	"fmt"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/params"
)

// CHANGE(hashkey): Check whether the address is in the whitelist if enabled.
func (evm *EVM) sandboxPenetrateCheck(addr common.Address) error {
	_, isPrecompile := evm.precompile(addr)
	// If the address is a precompiled or the sandbox policy address, skip the check.
	if isPrecompile || !evm.enableSandboxPenetrateCheck {
		return nil
	}
	// If the address is not in the whitelist, return an error.
	if !evm.isSandboxTrustedContract(addr) {
		return fmt.Errorf("%s should be in the white list", addr)
	}
	return nil
}

// CHANGE(hashkey): isSandboxTrustedContract checks whether the address is in the whitelist.
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

// CHANGE(hashkey): EnableSandboxPenetrateCheck enables the whitelist check for sandbox penetration.
func (evm *EVM) EnableSandboxPenetrateCheck(addr common.Address) {
	if addr != (common.Address{}) && evm.isSandboxTrustedContract(addr) {
		evm.enableSandboxPenetrateCheck = true
	}
}

// CHANGE(hashkey): check whether the address is in the blacklist or not.
func IsBlackListAddress(statedb StateDB, addr common.Address) bool {
	// Correct storage slot calculation for mapping(address => bool) at slot BlackListSlotNumber
	var buf [64]byte
	// Left-pad address to 32 bytes (address is 20 bytes, pad first 12 bytes with 0)
	copy(buf[12:32], addr[:])
	// Slot as 32-byte big-endian
	slot := params.BlackListSlotNumber.Bytes()
	copy(buf[64-len(slot):], slot)
	hash := crypto.Keccak256Hash(buf[:])
	val := statedb.GetState(params.BlackListAddress, hash)
	return val != (common.Hash{})
}
