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
