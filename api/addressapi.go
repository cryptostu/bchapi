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
		fmt.Println(err)
		return nil, err
	}
	address, err := model.StringToAddress(result)
	return address, err
}

//GetMultiAddress get multi bch address model by address
func GetMultiAddress(args ...string) ([]model.Address, error) {
	addrs := strings.Join(args, ",")
	fmt.Println(addrs)
	url := fmt.Sprintf(bchapi.AddressUrl, addrs)
	result, err := bchapi.HttpGet(url)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	addressRows, err := model.StringToMultiAddress(result)
	return *addressRows, err
}
