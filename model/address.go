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
	TxCount             int
	UnconfirmedTxCount  int
	UnconfirmedReceived int
	UnconfirmedSent     int
	UnspentTxCount      int
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
