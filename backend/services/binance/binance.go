package binance

import (
	sdk "github.com/binance-chain/go-sdk/client"
	"github.com/binance-chain/go-sdk/common/types"
	"github.com/binance-chain/go-sdk/keys"
	"github.com/binance-chain/go-sdk/types/msg"


)

var mnemonic = "lock globe panda armed mandate fabric couple dove climb step stove price recall decrease fire sail ring media enhance excite deny valid ceiling arm"
//-----   Init KeyManager  -------------
var keyManager, _ = keys.NewMnemonicKeyManager(mnemonic)

//-----   Init sdk  -------------
var client, err = sdk.NewDexClient("testnet-dex.binance.org", types.TestNetwork, keyManager)

func sendTransaction() {
	_, err := client.SendToken([]msg.Transfer{{testAccount2, []types.Coin{{"tbnb13095qugzf6d4078hnt9creqetwclvpn2e7yucr", 100000000}}}, {t"tbnb13095qugzf6d4078hnt9creqetwclvpn2e7yucr", []types.Coin{{"BNB", 100000000}}}}, true)
}
