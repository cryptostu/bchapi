package main

import (
	"github.com/cryptostu/bchapi/api"
)

func main() {
	// address := "17dNxJroBbzyoyoQMTveJwxYpkitAAChtX"
	// message := "0200000001280f134f1cda09b849a44f4b3d038c901e30f6ce0a986768cb10518c9fcb32830000000000ffffffff0180d1f008000000001976a91448b20e254c0677e760bab964aec16818d6b7134a88ac00000000"
	// signature := "ICkfL3WuWIg0PI4ReYYnM1Jq74NxQ7SIc4hFgUf46XRcMUMoIKkzJ8mdJbb+qHR6a1mrowTcFbzuFemU2p3j+bA="
	// result, err := api.VerifyMessage(address, message, signature)
	// if err != nil {
	// 	fmt.Println(err.Error())
	// }
	// fmt.Println(result)

	name := "ankye"
	password := "helloworld"
	password2 := "helloworld2"
	wif := "Kx7BmX9YHcytpoKiv1MoSm47WC4reA5xTYALgzwLk9ESarcGkHMy"
	// 	PrivateKey:    Kx7BmX9YHcytpoKiv1MoSm47WC4reA5xTYALgzwLk9ESarcGkHMy
	// PublicKey:     022a218783f3dacdc2ace28981f84285555c71ab106cb5460697bf167a5883f85b
	// Address:       1Che9VK2N4j5AgRSnVqLT81fHoGT2owjHz
	// CashAddress:   bitcoincash:qzq9jp7pmkpxe5y0fdg4ytr3jrua56tsd53kqcp99u
	account1, err := api.NewAccountByImportWif(name, password, wif)
	if err == nil {

		account1.Dump()

	}
	//PrivateKey:    KzL4jtA7kCdBKNc9uDTZEmC49LxVsqYwNbAczNesFuisD59PYo31
	// PublicKey:     022a218783f3dacdc2ace28981f84285555c71ab106cb5460697bf167a5883f85b
	// Address:       1Che9VK2N4j5AgRSnVqLT81fHoGT2owjHz
	// CashAddress:   bitcoincash:qzq9jp7pmkpxe5y0fdg4ytr3jrua56tsd53kqcp99u
	account2, err := api.NewAccountByImportWif(name, password2, wif)
	if err == nil {

		account2.Dump()

	}

}
