package api

import (
	"fmt"
	"strings"

	"github.com/cryptostu/bchapi"
	"github.com/cryptostu/bchapi/model"
)

//GetAddress get bch address model by address
func GetAddress(addr string) (*model.Address, error) {

	url := fmt.Sprintf(bchapi.AddressUrl, addr)
	result, err := bchapi.HttpGet(url)
	if err != nil {
		return nil, err
	}
	address, err := model.StringToAddress(result)
	return address, err
}

//GetMultiAddress get multi bch address model by address
func GetMultiAddress(args ...string) ([]model.Address, error) {
	addrs := strings.Join(args, ",")
	url := fmt.Sprintf(bchapi.AddressUrl, addrs)
	result, err := bchapi.HttpGet(url)
	if err != nil {
		return nil, err
	}
	addressRows, err := model.StringToMultiAddress(result)
	return *addressRows, err
}

//GetAddressTx get Address tx list
func GetAddressTx(addr string, page int, pagesize int) (*model.AddressTx, error) {
	url := fmt.Sprintf(bchapi.AddressTxUrl, addr, page, pagesize)
	result, err := bchapi.HttpGet(url)
	if err != nil {
		return nil, err
	}
	addressTx, err := model.StringToAddressTx(result)
	return addressTx, err
}

//GetAddressUnspent get address unspent struct
func GetAddressUnspent(addr string, page int, pagesize int) (*model.AddressUnspent, error) {
	url := fmt.Sprintf(bchapi.AddressUnspentUrl, addr, page, pagesize)
	result, err := bchapi.HttpGet(url)
	if err != nil {
		return nil, err
	}
	addressUnspent, err := model.StringToAddressUnspent(result)
	return addressUnspent, err
}
