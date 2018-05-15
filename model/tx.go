package model

import (
	"encoding/json"
)

type Tx struct {
	Hash   string
	Fee    float64
	Size   float64
	PreFee float64
}

func StringToTxs(s string) ([]*Tx, error) {

	json, err := JsonToMap(s)
	if err != nil {
		return nil, err
	}
	var txs []*Tx;
	data := json["data"].(map[string]interface{})
	list := data["list"].([]interface{})
	for _, txJson := range list {
		txMap := txJson.(map[string]interface{})
		fee := txMap["fee"].(float64)
		size := txMap["size"].(float64)
		tx := Tx{Hash: txMap["hash"].(string), Fee: fee, Size: size, PreFee: fee / size}
		txs = append(txs, &tx)
	}
	return txs, err

}

func JsonToMap(s string) (map[string]interface{}, error) {

	var result map[string]interface{}
	if err := json.Unmarshal([]byte(s), &result); err != nil {
		return nil, err
	}
	return result, nil
}
