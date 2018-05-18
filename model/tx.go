package model

import (
	"encoding/json"
	"errors"
)

type Input struct {
	PrevAddresses []string `json:"prev_addresses"` // Array<String> 输入地址
	PrevPosition  int      `json:"prev_position"`  //前向交易的输出位置
	PrevTxHash    string   `json:"prev_tx_hash"`   // 前向交易哈希
	PrevValue     int      `json:"prev_value"`     // 前向交易输入金额
	ScriptAsm     string   `json:"script_asm"`     // Script Asm
	ScriptHex     string   `json:"script_hex"`     // Script Hex
	Sequence      int      //Sequence
}
type Output struct {
	Addresses []string // Array<String> 输出地址
	Value     int      // 输出金额
}
type Tx struct {
	BlockHeight  int      `json:"block_height"` //所在块高度
	BlockTime    int      `json:"block_time"`   //所在块时间
	BlockHash    string   `json:"block_hash"`
	CreateAt     int      `json:"created_at"` //该记录系统处理时间，没有业务含义
	Fee          int      //该交易的手续费 在表示金额时，为避免浮点数产生精度问题，所有的金额单位均为聪
	Hash         string   //交易哈希
	Inputs       []Input  //输入
	InputsCount  int      `json:"inputs_count"` //输入数量
	InputsValue  int      `json:"inputs_value"` //输入金额
	IsCoinbase   bool     `json:"is_coinbase"`  // 是否为 coinbase 交易
	LockTime     int      `json:"lock_time"`    //lock time
	Outputs      []Output //输出
	OutputsCount int      `json:"outputs_count"` //输出数量
	OutputsValue int      `json:"outputs_value"` //输出金额
	Size         int      //交易体积
	Version      int      //交易版本号
	PreFee       float64
}

func StringToTxs(s string) ([]*Tx, error) {

	json, err := JsonToMap(s)
	if err != nil {
		return nil, err
	}
	var txs []*Tx
	data := json["data"].(map[string]interface{})
	list := data["list"].([]interface{})
	for _, txJson := range list {
		txMap := txJson.(map[string]interface{})
		fee := txMap["fee"].(float64)
		size := txMap["size"].(float64)
		tx := Tx{Hash: txMap["hash"].(string), Fee: int(fee), Size: int(size), PreFee: fee / size}
		txs = append(txs, &tx)
	}
	return txs, err

}

type txResult struct {
	responseCode
	Tx Tx `json:"data"`
}

//StringToTx json convert to AddressTx model
func StringToTx(str string) (*Tx, error) {

	var data txResult
	json.Unmarshal([]byte(str), &data)

	if data.ErrNo != 0 {
		return nil, errors.New(data.ErrMsg)
	}
	data.Tx.PreFee = float64(data.Tx.Fee) / float64(data.Tx.Size)
	return &data.Tx, nil

}

func JsonToMap(s string) (map[string]interface{}, error) {

	var result map[string]interface{}
	if err := json.Unmarshal([]byte(s), &result); err != nil {
		return nil, err
	}
	return result, nil
}

type txUnconfirmedResult struct {
	responseCode
	List []string `json:"data"`
}

//StringToUnconfirmedHex json convert to  model
func StringToUnconfirmedHex(str string) ([]string, error) {

	var data txUnconfirmedResult
	json.Unmarshal([]byte(str), &data)

	if data.ErrNo != 0 {
		return nil, errors.New(data.ErrMsg)
	}
	return data.List, nil

}

//TxUncomfirmedSummary size and count
type TxUncomfirmedSummary struct {
	Size  int
	Count int
}

type resultTxUncomfirmedSummary struct {
	responseCode
	Summary TxUncomfirmedSummary `json:"data"`
}

//StringToUncomfirmedSummary json convert to  model
func StringToUncomfirmedSummary(str string) (*TxUncomfirmedSummary, error) {

	var data resultTxUncomfirmedSummary
	json.Unmarshal([]byte(str), &data)

	if data.ErrNo != 0 {
		return nil, errors.New(data.ErrMsg)
	}
	return &data.Summary, nil

}
