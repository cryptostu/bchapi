package model

import (
	"encoding/json"
	"errors"
)

//Address BitcoinCash address model
type Address struct {
	Address             string
	Received            int
	Sent                int
	Balance             int
	TxCount             int `json:"tx_count"`
	UnconfirmedTxCount  int `json:"unconfirmed_tx_count"`
	UnconfirmedReceived int `json:"unconfirmed_received"`
	UnconfirmedSent     int `json:"unconfirmed_sent"`
	UnspentTxCount      int `json:"unspent_tx_count"`
}
type responseCode struct {
	ErrNo  int    `json:"err_no"`
	ErrMsg string `json:"err_msg"`
}

type paginator struct {
	TotalCount int `json:"total_count"`
	Page       int
	Pagesize   int
}

type multiAddressResult struct {
	Address []Address `json:"data"`
	responseCode
}

type addressResult struct {
	Address Address `json:"data"`
	responseCode
}

//AddressTx address
type AddressTx struct {
	paginator
	Txs []Tx `json:"list"`
}

type addressTxResult struct {
	responseCode
	AddressTx AddressTx `json:"data"`
}

//StringToAddress json convert to model
func StringToAddress(str string) (*Address, error) {

	var data addressResult
	json.Unmarshal([]byte(str), &data)

	if data.ErrNo != 0 {
		return nil, errors.New(data.ErrMsg)
	}
	return &data.Address, nil

}

//StringToMultiAddress json convert to model
func StringToMultiAddress(str string) (*[]Address, error) {

	var data multiAddressResult
	json.Unmarshal([]byte(str), &data)

	if data.ErrNo != 0 {
		return nil, errors.New(data.ErrMsg)
	}
	return &data.Address, nil

}

//StringToAddressTx json convert to model
func StringToAddressTx(str string) (*AddressTx, error) {

	var data addressTxResult
	json.Unmarshal([]byte(str), &data)

	if data.ErrNo != 0 {
		return nil, errors.New(data.ErrMsg)
	}
	return &data.AddressTx, nil

}
