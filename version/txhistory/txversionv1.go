package txhistory

import (
	"github.com/elastos/Elastos.ELA/core/types"
	"github.com/elastos/Elastos.ELA/version"

	. "github.com/elastos/Elastos.ELA.Utility/common"
)

type TxVersionV1 struct {
	version.TxVersionMain
}

func (v *TxVersionV1) GetVersion() byte {
	return 1
}

func (v *TxVersionV1) CheckOutputPayload(output *types.Output) error {
	return nil
}

func (v *TxVersionV1) CheckCoinbaseMinerReward(tx *types.Transaction, totalReward Fixed64) error {
	return nil
}

func (v *TxVersionV1) CheckCoinbaseArbitratorsReward(coinbase *types.Transaction, rewardInCoinbase Fixed64) error {
	return nil
}

func (v *TxVersionV1) CheckVoteProducerOutputs(outputs []*types.Output, references map[*types.Input]*types.Output) error {
	return nil
}

func (v *TxVersionV1) CheckTxHasNoProgramsAndAttributes(tx *types.Transaction) error {
	return nil
}
