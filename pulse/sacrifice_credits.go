package pulse

import (
	_ "embed"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/state"
	"github.com/ethereum/go-ethereum/log"
	"github.com/ethereum/go-ethereum/params"
	"github.com/holiman/uint256"
)

// The testnet credits are approximate and not final for mainnet
// see https://gitlab.com/pulsechaincom/compressed-allocations/-/tree/Testnet-R2-Credits
//
//go:embed sacrifice_credits.bin
var rawCredits []byte

// Applies the sacrifice credits for the PrimordialPulse fork.
func ApplySacrificeCredits(state *state.StateDB, treasury *params.Treasury) {
	if treasury != nil {
		log.Info("Applying PrimordialPulse treasury allocation 💸")
		state.AddBalance(common.HexToAddress(treasury.Addr), uint256.MustFromBig((*big.Int)(treasury.Balance)))
	}

	log.Info("Applying PrimordialPulse sacrifice credits 💸")
	for ptr := 0; ptr < len(rawCredits); {
		byteCount := int(rawCredits[ptr])
		ptr++

		record := rawCredits[ptr : ptr+byteCount]
		ptr += byteCount

		addr := common.BytesToAddress(record[:20])
		credit := new(uint256.Int).SetBytes(record[20:])
		state.AddBalance(addr, credit)
	}

	log.Info("Finished applying PrimordialPulse sacrifice credits 🤑")
}
