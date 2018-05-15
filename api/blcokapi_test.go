package api

import (
	"testing"
)

func TestGetBlockTxs(t *testing.T) {

}

func TestGetBlock(t *testing.T) {
	block, err := GetBlock(3)
	if err != nil {
		t.Error(err)
	}
	if block.Hash != "0000000082b5015589a3fdf2d4baff403e6f0be035a5d9742c1cae6295464449" {
		t.Error("blcok hash is error")
	}

}
