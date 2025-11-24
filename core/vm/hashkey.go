package vm

import (
	"fmt"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/params"
)

// CHANGE(hashkey): Check whether the address is in the whitelist if enabled.
func (evm *EVM) checkWhiteListAddress(addr common.Address) error {
	_, isPrecompile := evm.precompile(addr)
	// If the address is a precompiled or the sandbox policy address, skip the check.
	if isPrecompile {
		return nil
	}
	isWhiteAddr := evm.isWhiteAddress(addr)
	if evm.enableWhiteList {
		// If the address is not in the whitelist, return an error.
		if !isWhiteAddr {
			return fmt.Errorf("%s should be in the white list", addr)
		}
	} else if isWhiteAddr {
		// If the whitelist is not enabled, and the address is in the whitelist, enable it.
		evm.enableWhiteList = true
	}
	return nil
}

// CHANGE(hashkey): isWhiteAddress checks whether the address is in the whitelist.
func (evm *EVM) isWhiteAddress(addr common.Address) bool {
	// Correct storage slot calculation for mapping(address => bool) at slot WhiteListSlotNumber
	var buf [64]byte
	// Left-pad address to 32 bytes (address is 20 bytes, pad first 12 bytes with 0)
	copy(buf[12:32], addr[:])
	// Slot as 32-byte big-endian
	slot := params.WhiteListSlotNumber.Bytes()
	copy(buf[64-len(slot):], slot)
	hash := crypto.Keccak256Hash(buf[:])
	val := evm.StateDB.GetState(params.SandboxPolicyAddress, hash)
	return val != (common.Hash{})
}
