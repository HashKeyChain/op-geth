package txpool

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/state"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/params"
)

func IsBlacklisted(statedb *state.StateDB, addr common.Address) bool {
	var buf [64]byte
	// Left-pad address to 32 bytes (address is 20 bytes, pad first 12 bytes with 0)
	copy(buf[12:32], addr[:])
	// Slot as 32-byte big-endian
	slot := params.BlackListSlotNumber.Bytes()
	copy(buf[64-len(slot):], slot)

	hash := crypto.Keccak256Hash(buf[:])
	val := statedb.GetState(params.HashKeyBlacklist, hash)
	return val != (common.Hash{})
}
