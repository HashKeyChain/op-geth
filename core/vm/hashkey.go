package vm

import (
	"fmt"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/log"
	"github.com/ethereum/go-ethereum/params"
)

// CHANGE(hashkey): Check whether the address is in the whitelist if enabled.
func (evm *EVM) SandboxPenetrateCheck(addr common.Address) error {
	// If the address is a precompiled or the sandbox policy address, skip the check.
	_, isPrecompile := evm.precompile(addr)
	if isPrecompile {
		return nil
	}

	preContract := evm.preContract
	evm.preContract = addr
	if evm.isSandboxTrustedContract(addr) {
		if !evm.isInSandbox {
			log.Error("Detect sandbox penetrate in", "from", preContract.Hex(), "to", addr.Hex())
			return fmt.Errorf("detect sandbox penetrate in, from: %s, to: %s", preContract.Hex(), addr.Hex())
		} else {
			return nil
		}
	}
	if evm.isInSandbox {
		log.Error("Detect sandbox penetrate out", "from", preContract.Hex(), "to", addr.Hex())
		return fmt.Errorf("detect sandbox penetrate out, from %s, to %s", preContract.Hex(), addr.Hex())
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

// CHANGE(hashkey): InitSandboxPenetrateCheck init the sandbox penetration check for the tx.
func (evm *EVM) InitSandboxPenetrateCheck(addr common.Address) {
	evm.preContract = addr
	if addr != (common.Address{}) && evm.isSandboxTrustedContract(addr) {
		evm.isInSandbox = true
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
	val := statedb.GetState(params.SandboxPolicyAddress, hash)
	return val != (common.Hash{})
}
