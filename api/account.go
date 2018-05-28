package api

import (
	"fmt"

	"github.com/btcsuite/btcutil"
	"github.com/cpacia/bchutil"
	"github.com/cryptostu/bchapi/crypto"
	"github.com/memocash/memo/app/bitcoin/wallet"
)

func createKey(name string, password string, privateKey wallet.PrivateKey, key []byte) (*Account, error) {
	encryptedSecret, err := crypto.Encrypt(privateKey.Secret, key)
	if err != nil {
		return nil, fmt.Errorf("failed to encrypt,error:%s", err.Error())
	}
	var account = &Account{
		Name:      name,
		Password:  password,
		Secret:    encryptedSecret,
		Wif:       privateKey.GetBase58Compressed(),
		PublicKey: privateKey.GetPublicKey().GetSerialized(),
		PkHash:    privateKey.GetPublicKey().GetAddress().GetScriptAddress(),
	}

	return account, nil
}

//NewAccount create account by name and password,generate random private key
func NewAccount(name string, password string) (*Account, error) {
	key, err := crypto.GenerateEncryptionKeyFromPassword(password)
	if err != nil {
		return nil, fmt.Errorf("error generating key from password ,error :%s", err.Error())
	}
	privateKey := wallet.GeneratePrivateKey()

	return createKey(name, password, privateKey, key)
}

//NewAccountByImportWif create account by name and password,import wif to generate privatekey
//wif is a Base58 compressed private key
func NewAccountByImportWif(name string, password string, wif string) (*Account, error) {
	key, err := crypto.GenerateEncryptionKeyFromPassword(password)
	if err != nil {
		return nil, fmt.Errorf("error generating key from password,error :%s", err.Error())
	}
	privateKey, err := wallet.ImportPrivateKey(wif)
	if err != nil {
		return nil, fmt.Errorf("error importing key from wif,error :%s", err.Error())
	}
	return createKey(name, password, privateKey, key)
}

//Account account info
type Account struct {
	Name      string
	Password  string
	Wif       string
	Secret    []byte
	PublicKey []byte
	PkHash    []byte
	Balance   float64
}

//GetPublicKeyString get serialized public key
func (k *Account) GetPublicKeyString() string {
	return k.GetPublicKey().GetSerializedString()
}

//GetWifString get base58 compressed serialized private key
func (k *Account) GetWifString() string {
	return k.Wif
}

//GetAddressString get address string
func (k *Account) GetAddressString() string {
	addr, err := btcutil.NewAddressPubKeyHash(k.PkHash, &wallet.MainNetParamsOld)
	if err != nil {
		return ""
	}
	return addr.String()
}

//GetCashAddressString get cash address string
func (k *Account) GetCashAddressString() string {

	addr, err := btcutil.NewAddressPubKeyHash(k.PkHash, &wallet.MainNetParamsOld)
	if err != nil {
		return ""
	}
	cashAddr, err := bchutil.NewCashAddressPubKeyHash(addr.ScriptAddress(), &wallet.MainNetParamsOld)
	if err != nil {
		return ""
	}
	return cashAddr.String()
}

//GetPrivateKey get private key by password
func (k *Account) GetPrivateKey(password string) (*wallet.PrivateKey, error) {
	key, err := crypto.GenerateEncryptionKeyFromPassword(password)
	if err != nil {
		return nil, fmt.Errorf("error generating key from password,error :%s", err.Error())
	}
	decrypted, err := crypto.Decrypt(k.Secret, key)
	if err != nil {
		return nil, fmt.Errorf("failed to decrypt,error :%s", err.Error())
	}
	privateKey := wallet.PrivateKey{
		Secret: decrypted,
	}

	return &privateKey, nil
}

//UpdatePassword update password and generate new privatekey and publickey,address not change
func (k *Account) UpdatePassword(oldPassword string, newPassword string) error {
	privateKey, err := k.GetPrivateKey(oldPassword)
	if err != nil {
		return fmt.Errorf("error getting key from password,error :%s", err.Error())
	}
	encryptionKey, err := crypto.GenerateEncryptionKeyFromPassword(newPassword)
	if err != nil {
		return fmt.Errorf("error generating key from password,error :%s", err.Error())
	}
	encryptedSecret, err := crypto.Encrypt(privateKey.Secret, encryptionKey)
	if err != nil {
		return fmt.Errorf("failed to encrypt,error :%s", err.Error())
	}
	k.Secret = encryptedSecret
	k.Password = newPassword

	return nil
}

//GetPublicKey get wallet public key struct
func (k *Account) GetPublicKey() wallet.PublicKey {
	return wallet.GetPublicKey(k.PublicKey)
}

//GetAddress get wallet address struct
func (k *Account) GetAddress() wallet.Address {
	return k.GetPublicKey().GetAddress()
}

//GetBalance get Account balance
func (k *Account) GetBalance() (int, error) {
	ret, err := GetAddress(k.GetAddressString())
	if err != nil {
		return 0, err
	}

	return ret.Balance, nil

}

//Dump dump account info
func (k *Account) Dump() {
	fmt.Println("/////////////Account Dump ///////////////////////////////////////")
	fmt.Printf("Wif:           %s\r\n", k.GetWifString())
	fmt.Printf("PublicKey:     %s\r\n", k.GetPublicKeyString())
	fmt.Printf("Address:       %s\r\n", k.GetAddressString())
	fmt.Printf("CashAddress:   %s\r\n", k.GetCashAddressString())
	fmt.Printf("Secret:        %x\r\n", k.Secret)
	fmt.Printf("Name:          %s\r\n", k.Name)
	fmt.Printf("Password:      %s\r\n", k.Password)
	if balance, err := k.GetBalance(); err == nil {
		fmt.Printf("Balance:       %d\r\n", balance)
	} else {
		fmt.Printf("Get Balance Failed:%s", err.Error())
	}
	fmt.Println("////////////////////////////////////////////////////")

}
