package model

import (
	"encoding/json"
	"errors"
)

type Block struct {
	Height        int
	Version       int
	MrklRoot      string `json:"mrkl_root"`
	Timestamp     uint32
	Bits          uint32
	Nonce         uint32
	Hash          string
	PrevBlockHash string `json:"next_block_hashs"`

	NextBlockHash    string `json:"next_block_hash"`
	Size             int32
	PoolDifficulty   int32  `json:"pool_difficulty"`
	Difficulty       int32
	TxCount          int32  `json:"tx_count"`
	RewardBlock      int64
	RewardFees       int64  `json:"reward_fees"`
	CreatedAt        uint32 `json:"created_at"`
	Confirmations    int32
	IsOrphan         bool   `json:"is_orphan"`
	CurrMaxTimestamp uint32 `json:"curr_max_timestamp"`
}
type Extras struct {
	PoolName string `json:"pool_name"`
	PoolLink string `json:"pool_link"`
}
type jsonResult struct {
	Block  Block  `json:"data"`
	ErrNo  int    `json:"err_no"`
	ErrMsg string `json:"err_msg"`
}

func StringToBlock(str string) (*Block, error) {

	var data jsonResult
	json.Unmarshal([]byte(str), &data)

	if data.ErrNo != 0 {
		return nil, errors.New(data.ErrMsg)
	}
	return &data.Block, nil

}
