package api

import "testing"

func TestGetTx(t *testing.T) {
	txhash := "0eab89a271380b09987bcee5258fca91f28df4dadcedf892658b9bc261050d96"
	if result, err := GetTx(txhash); err != nil {
		t.Errorf("GetTx  api failed,err:%s", err.Error())
	} else {
		blockHash := "000000000000000005cb6f6e2f09e84a353ab91756a38aa50fbaf25059f76666"
		expect := result.BlockHash
		if blockHash != expect {
			t.Errorf("Expect hash %s but got %s", blockHash, expect)
		}
	}
}
