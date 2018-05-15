package model

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

func StringToBlock(str string) ([]*Block, error) {
	return nil, nil

}
