package api

import (
	"fmt"

	"github.com/cryptostu/bchapi"
	"github.com/cryptostu/bchapi/model"
)

//GetTx get tx info
func GetTx(txhash string) (*model.Tx, error) {
	url := fmt.Sprintf(bchapi.TxUrl, txhash)
	result, err := bchapi.HttpGet(url)
	if err != nil {
		return nil, err
	}
	tx, err := model.StringToTx(result)
	return tx, err
}
