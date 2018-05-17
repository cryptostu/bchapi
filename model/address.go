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

type multiAddressResult struct {
	Address []Address `json:"data"`
	ErrNo   int       `json:"err_no"`
	ErrMsg  string    `json:"err_msg"`
}

type addressResult struct {
	Address Address `json:"data"`
	ErrNo   int     `json:"err_no"`
	ErrMsg  string  `json:"err_msg"`
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
