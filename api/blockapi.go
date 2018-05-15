package api

import (
	"fmt"
	"github.com/cryptostu/bchapi"
	"github.com/cryptostu/bchapi/model"

)



func GetBlockTxs(height int) ([]*model.Tx, error) {
	url := fmt.Sprintf(bchapi.BlockTxsUrl, height)
	result, err := bchapi.HttpGet(url)
	if err != nil {
		fmt.Println(err)
	}
	txs, err := model.StringToTxs(result)
	return txs, err

}

func GetBlock(height int) (*model.Block, error) {
	url := fmt.Sprintf(bchapi.BlockUrl, height)
	result, err := bchapi.HttpGet(url)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	block, err := model.StringToBlock(result)
	return block, err

}
