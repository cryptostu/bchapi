package api

import "testing"

func TestGetAddress(t *testing.T) {
	addr := "15urYnyeJe3gwbGJ74wcX89Tz7ZtsFDVew"
	if address, err := GetAddress(addr); err != nil {
		t.Error("get address from api failed")
	} else {

		if address.Address != addr {
			t.Errorf("Expect %s but got %s", addr, address.Address)
		}
	}
}
func TestGetMultiAddress(t *testing.T) {
	addr1 := "15urYnyeJe3gwbGJ74wcX89Tz7ZtsFDVew"
	addr2 := "1PErRgFdo757pyyMxFiwB326vuymXC3hev"
	if addressRows, err := GetMultiAddress(addr1, addr2); err != nil {
		t.Error("get address from api failed")
	} else {
		if len(addressRows) != 2 {
			t.Errorf("expect 2 rows got %d", len(addressRows))
		}
		addressRow1 := addressRows[0]
		addressRow2 := addressRows[1]
		if addressRow1.Address != addr1 {
			t.Errorf("Expect %s but got %s", addr1, addressRow1.Address)
		}

		if addressRow2.Address != addr2 {
			t.Errorf("Expect %s but got %s", addr2, addressRow2.Address)
		}
	}
}

func TestGetAddressTx(t *testing.T) {
	addr := "15urYnyeJe3gwbGJ74wcX89Tz7ZtsFDVew"
	page := 1
	pagesize := 3
	if result, err := GetAddressTx(addr, page, pagesize); err != nil {
		t.Errorf("GetAddressTx  api failed,err:%s", err.Error())
	} else {
		hash := "04ffa9c3875b15ceb65c2dd4ee2654c5fb65374123692362e32fac566a6b16aa"
		expect := result.List[0].Hash
		if hash != expect {
			t.Errorf("Expect hash %s but got %s", hash, expect)
		}
		if result.TotalCount >= page*pagesize {
			if len(result.List) != pagesize {
				t.Errorf("get list count failed expect %d ,got %d", pagesize, len(result.List))
			}
		}
	}

}

func TestGetAddressUnspent(t *testing.T) {
	addr := "15urYnyeJe3gwbGJ74wcX89Tz7ZtsFDVew"
	if result, err := GetAddressUnspent(addr, 1, 3); err != nil {
		t.Errorf("GetAddressUnspent  api failed,err:%s", err.Error())
	} else {
		hash := "04ffa9c3875b15ceb65c2dd4ee2654c5fb65374123692362e32fac566a6b16aa"
		expect := result.List[0].TxHash
		if hash != expect {
			t.Errorf("Expect hash %s but got %s", hash, expect)
		}
	}

}
