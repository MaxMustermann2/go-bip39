package bip39

import (
	"fmt"
	"testing"
)

func TestEncryptMnemonic(t *testing.T) {
	mnmonic := "all hour make first leader extend hole alien behind guard gospel lava path output census museum junior mass reopen famous sing advance salt reform"
	password := "securePassword"
	encMnemonic, err := EncryptMnemonic(mnmonic, password)
	if err != nil {
		fmt.Println("Got error", err)
	}
	fmt.Println("Encrypted mnemonic is", encMnemonic)
}
