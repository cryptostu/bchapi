package api

import (
	"fmt"

	"github.com/cryptostu/bchapi"
)

//TxDecode decode transaction
//rawhex Hex-encoded bytes of the serialized transaction
func TxDecode(rawhex string) (str string, err error) {
	data := fmt.Sprintf("rawhex=%s", rawhex)
	result, err := bchapi.HttpPost(bchapi.ToolsTxDecodeUrl, data, bchapi.ConnTimeoutMS, bchapi.ServeTimeoutMS)
	if err != nil {
		return
	}
	str = result
	return
}

//TxPublish Submits the serialized, hex-encoded transaction to the btc.com peer and relays it to the network.
//rawhex Hex-encoded bytes of the serialized transaction
func TxPublish(rawhex string) (str string, err error) {
	data := fmt.Sprintf("rawhex=%s", rawhex)
	result, err := bchapi.HttpPost(bchapi.ToolsTxPublishUrl, data, bchapi.ConnTimeoutMS, bchapi.ServeTimeoutMS)
	if err != nil {
		return
	}
	str = result
	return
}

//VerifyMessage Verify a signed message.
//address The bitcoin address to use for the signature
//message The signed message
//signature The base-64 encoded signature provided by the signer
func VerifyMessage(address string, message string, signature string) (str string, err error) {
	data := fmt.Sprintf("address=%s&message=%s&signature=%s", address, message, signature)
	result, err := bchapi.HttpPost(bchapi.ToolsVerifyMessageUrl, data, bchapi.ConnTimeoutMS, bchapi.ServeTimeoutMS)
	if err != nil {
		return
	}
	str = result
	return
}
