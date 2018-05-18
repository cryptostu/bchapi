package api

import "testing"

func TestTxDecode(t *testing.T) {
	rawhex := "907c2bc503ade11cc3b04eb2918b6f547b0630ab569273824748c87ea14b0696526c66ba740200000004ab65ababfd1f9bdd4ef073c7afc4ae00da8a66f429c917a0081ad1e1dabce28d373eab81d8628de802000000096aab5253ab52000052ad042b5f25efb33beec9f3364e8a9139e8439d9d7e26529c3c30b6c3fd89f8684cfd68ea0200000009ab53526500636a52ab599ac2fe02a526ed040000000008535300516352515164370e010000000003006300ab2ec229"
	if _, err := TxDecode(rawhex); err != nil {
		t.Errorf("TxDecode  api failed,err:%s", err.Error())
	}

}

func TestTxPublish(t *testing.T) {
	rawhex := "020000000121b00cc355677dc6d7114f62bfaccd5d76ec50fae5d62c86d50096c4817a4b93820000006b483045022100adaab40f2c9fe7250e9f6b3930e764091fcaf1a0efa86e61236cd0c427a4eb9f02200ef7d395a90d55f804f2fbdd6eabe4b345e43202f7ce107ccfde15aad6b3a46141210384dd3ad997f2e10980e755236b474f986c519599946027876cdeb4eb5a30a09fffffffff0110270000000000001976a9140c4660bc6acf0daba88b95e43411d17162b2466588ac00000000"
	if _, err := TxPublish(rawhex); err != nil {
		t.Errorf("TxPublish  api failed,err:%s", err.Error())
	}
}

func TestVerifyMessage(t *testing.T) {
	address := "17dNxJroBbzyoyoQMTveJwxYpkitAAChtX"
	message := "0200000001280f134f1cda09b849a44f4b3d038c901e30f6ce0a986768cb10518c9fcb32830000000000ffffffff0180d1f008000000001976a91448b20e254c0677e760bab964aec16818d6b7134a88ac00000000"
	signature := "ICkfL3WuWIg0PI4ReYYnM1Jq74NxQ7SIc4hFgUf46XRcMUMoIKkzJ8mdJbb+qHR6a1mrowTcFbzuFemU2p3j+bA="

	if result, err := VerifyMessage(address, message, signature); err != nil {
		t.Errorf("VerifyMessage  api failed,err:%s", err.Error())
	} else {
		if !result {
			t.Error(result)
		}

	}
}
