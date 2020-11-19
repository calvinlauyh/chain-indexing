package model

import (
	"github.com/crypto-com/chainindex/usecase/coin"
)

type MsgBeginRedelegateParams struct {
	DelegatorAddress    string    `json:"delegatorAddress"`
	ValidatorSrcAddress string    `json:"validatorSrcAddress"`
	ValidatorDstAddress string    `json:"validatorDstAddress"`
	Amount              coin.Coin `json:"amount"`
}
