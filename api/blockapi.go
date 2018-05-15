package api

import (
	"fmt"
	"github.com/cryptostu/bchapi"
	"github.com/cryptostu/bchapi/model"
)

const (
	BLOCK_URL    = "block/%d/"
	BLOCK_TX_URL = "block/%d/tx"
)

func GetBlockTxs(height int) (error, []*model.Tx) {
	url := fmt.Sprintf(BLOCK_TX_URL, height)
	result, err := bchapi.HttpGet(url)
	if err != nil {
		fmt.Println(err)
	}
	txs, err := model.StringToTxs(result)
	return err, txs

}
