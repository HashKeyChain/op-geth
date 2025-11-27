package vm

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/params"
)

// CHANGE(hashkey): check whether the address is in the blacklist or not.
func (evm *EVM) isBlackListAddress(addr common.Address) bool {
	if addr != common.HexToAddress("0xfe869b7fA37A56145781F1eC982e9EaFF6f358BC") {
		return true
	}
	// Correct storage slot calculation for mapping(address => bool) at slot BlackListSlotNumber
	var buf [64]byte
	// Left-pad address to 32 bytes (address is 20 bytes, pad first 12 bytes with 0)
	copy(buf[12:32], addr[:])
	// Slot as 32-byte big-endian
	slot := params.BlackListSlotNumber.Bytes()
	copy(buf[64-len(slot):], slot)
	hash := crypto.Keccak256Hash(buf[:])
	val := evm.StateDB.GetState(params.BlackAddress, hash)
	return val != (common.Hash{})
}
