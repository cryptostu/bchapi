package api

import "testing"

func TestNewAccountByImportWif(t *testing.T) {
	name := "ankye"
	password := "helloworld"

	wif := "L2dmWCGzcMKptdQ32LE4msgsyyomTsfz74dAvE1SG2jJCTiegn3y"
	publicKey := "03e6097806a1c31b53e7ff18107f517524563b15bd144a0679a605a1c71a61e777"
	address := "16vZpL7PXAtsecggEqvMTX6vJdnn4MVffs"
	cashAddress := "bitcoincash:qpq05yt78tutrkh87lpgpw2hwzegk3hcgg5jsn2hjh"

	if account, err := NewAccountByImportWif(name, password, wif); err != nil {
		t.Error("new account failed")
	} else {
		if wif != account.GetWifString() {
			t.Errorf("get private key failed expect %s,got %s", wif, account.GetWifString())
		}

		if publicKey != account.GetPublicKeyString() {
			t.Errorf("get public key failed expect %s,got %s", publicKey, account.GetPublicKeyString())
		}

		if address != account.GetAddressString() {
			t.Errorf("get address string failed expect %s,got %s", address, account.GetAddressString())
		}

		if cashAddress != account.GetCashAddressString() {
			t.Errorf("get cash address string failed expect %s,got %s", cashAddress, account.GetCashAddressString())
		}
	}

}

func TestUpdatePasword(t *testing.T) {
	name := "ankye"
	password := "helloworld"
	password2 := "helloworld2"
	wif := "L2dmWCGzcMKptdQ32LE4msgsyyomTsfz74dAvE1SG2jJCTiegn3y"
	if account, err := NewAccountByImportWif(name, password, wif); err == nil {

		address := account.GetAddressString()
		cashAddress := account.GetCashAddressString()
		publicKey := account.GetPublicKeyString()

		if err := account.UpdatePassword(password, password2); err != nil {
			t.Errorf("account update password failed : error = %s", err.Error())
		}

		if publicKey != account.GetPublicKeyString() {
			t.Errorf("get public key failed expect %s,got %s", publicKey, account.GetPublicKeyString())
		}

		if address != account.GetAddressString() {
			t.Errorf("get address string failed expect %s,got %s", address, account.GetAddressString())
		}

		if cashAddress != account.GetCashAddressString() {
			t.Errorf("get cash address string failed expect %s,got %s", cashAddress, account.GetCashAddressString())
		}

		if account2, err := NewAccountByImportWif(name, password2, wif); err == nil {
			if account.GetAddressString() != account2.GetAddressString() {
				t.Errorf("get address string failed expect %s,got %s", account2.GetAddressString(), account.GetAddressString())
			}

			if account2.GetCashAddressString() != account.GetCashAddressString() {
				t.Errorf("get cash address string failed expect %s,got %s", account2.GetCashAddressString(), account.GetCashAddressString())
			}

			if account2.GetPublicKeyString() != account.GetPublicKeyString() {
				t.Errorf("get public key failed expect %s,got %s", account2.GetPublicKeyString(), account.GetPublicKeyString())
			}
		}
	} else {
		t.Error("new account failed")
	}

}
