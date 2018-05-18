package api

import (
	"fmt"

	"github.com/cryptostu/bchapi"
	"github.com/cryptostu/bchapi/model"
)

//GetTx get tx info
func GetTx(txhash string) (*model.Tx, error) {
	url := fmt.Sprintf(bchapi.TxUrl, txhash)
	result, err := bchapi.HttpGet(url, bchapi.ConnTimeoutMS, bchapi.ServeTimeoutMS)
	if err != nil {
		return nil, err
	}
	tx, err := model.StringToTx(result)
	return tx, err
}

//GetTxUnconfirmed 获取未确认交易的哈希
func GetTxUnconfirmed() ([]string, error) {

	result, err := bchapi.HttpGet(bchapi.TxUnconfirmedUrl, bchapi.ConnTimeoutMS, bchapi.ServeTimeoutMS)
	if err != nil {
		return nil, err
	}
	tx, err := model.StringToUnconfirmedHex(result)
	return tx, err
}

//GetTxUnconfirmedSummary 获取未确认交易的信息，包括体积和数量。
func GetTxUnconfirmedSummary() (*model.TxUncomfirmedSummary, error) {

	result, err := bchapi.HttpGet(bchapi.TxUnconfirmedSummaryUrl, bchapi.ConnTimeoutMS, bchapi.ServeTimeoutMS)
	if err != nil {
		return nil, err
	}
	summary, err := model.StringToUncomfirmedSummary(result)
	return summary, err
}
