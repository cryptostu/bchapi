package crypto

import (
	"fmt"

	"testing"

	"github.com/jchavannes/jgo/jerr"
)

const Password = "bl&8aZ&74CK!vMBt"
const SomePlaintext = "some plaintext"

func TestEncryption(t *testing.T) {
	key, err := GenerateEncryptionKeyFromPassword(Password)
	if err != nil {
		t.Fatal(jerr.Get("error generating key from password", err))
	}
	fmt.Printf("Key: %x\n", key)
	encrypted, err := Encrypt([]byte(SomePlaintext), key)
	if err != nil {
		t.Fatal(jerr.Get("failed to encrypt", err))
	}
	fmt.Printf("Encrypted: %x\n", encrypted)
	decrypted, err := Decrypt(encrypted, key)
	if err != nil {
		t.Fatal(jerr.Get("failed to decrypt", err))
	}
	fmt.Printf("Decrypted: %s\n", decrypted)
	if string(decrypted) != SomePlaintext {
		t.Fatal(jerr.New("decrypted value does not match original"))
	}
}
