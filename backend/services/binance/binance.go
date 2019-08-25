package main

import (
	"encoding/json"
	"fmt"
	"os"

	sdk "github.com/binance-chain/go-sdk/client"
	"github.com/imroc/req"
	"golang.org/x/net/websocket"
	"log"
	"time"

	ctypes "github.com/binance-chain/go-sdk/common/types"
	"github.com/binance-chain/go-sdk/keys"
	"github.com/binance-chain/go-sdk/types/msg"
)

type Tx struct {
	Code   int    `json:"code"`
	Hash   string `json:"hash"`
	Height string `json:"height"`
	Log    string `json:"log"`
	Ok     bool   `json:"ok"`
	Tx     struct {
		Type  string `json:"type"`
		Value struct {
			Data interface{} `json:"data"`
			Memo string      `json:"memo"`
			Msg  []struct {
				Type  string `json:"type"`
				Value struct {
					Inputs []struct {
						Address string `json:"address"`
						Coins   []struct {
							Amount string `json:"amount"`
							Denom  string `json:"denom"`
						} `json:"coins"`
					} `json:"inputs"`
					Outputs []struct {
						Address string `json:"address"`
						Coins   []struct {
							Amount string `json:"amount"`
							Denom  string `json:"denom"`
						} `json:"coins"`
					} `json:"outputs"`
				} `json:"value"`
			} `json:"msg"`
			Signatures []struct {
				AccountNumber string `json:"account_number"`
				PubKey        struct {
					Type  string `json:"type"`
					Value string `json:"value"`
				} `json:"pub_key"`
				Sequence  string `json:"sequence"`
				Signature string `json:"signature"`
			} `json:"signatures"`
			Source string `json:"source"`
		} `json:"value"`
	} `json:"tx"`
}

type WssResponse struct {
	Stream string `json:"stream"`
	Data   struct {
		Estr string `json:"e"`
		Eint int    `json:"E"`
		H    string `json:"H"`
		M    string `json:"M"`
		F    string `json:"f"`
		T    []struct {
			O string `json:"o"`
			C []struct {
				Alil string `json:"a"`
				Abig string `json:"A"`
			} `json:"c"`
		} `json:"t"`
	} `json:"data"`
}

var (
	mnemonic = "critic mouse category pig visit kidney weasel coin media price suspect next art model soul shuffle welcome slot thrive sign train large wild submit"
	origin   = "http://localhost.localdomain/"
	url      = "wss://testnet-dex.binance.org/api/ws"
)

// 0.10000001 BNB
// sendTransaction("tbnb15qfcd5863pgf9qevefn5sj056cyk4r9mtcktnn", 10000001)

func sendTransaction(address string, sum int64) {
	keyManager, err := keys.NewMnemonicKeyManager(mnemonic)
	if err != nil {
		log.Fatal(err)
	}
	client, err := sdk.NewDexClient("testnet-dex.binance.org", ctypes.TestNetwork, keyManager)
	acc, err := ctypes.AccAddressFromBech32(address)
	if err != nil {
		log.Fatal(err)
	}
	send, err := client.SendToken([]msg.Transfer{{acc, []ctypes.Coin{{msg.NativeToken, sum}}}}, true)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(send)
}

func listenAndSay() {
	if os.Getenv("ENV") == "PROD" {
		origin = "https://berlin.buttonwallet.tech/"
	}
	for {
		ws, err := websocket.Dial(url, "", origin)
		if err != nil {
			log.Fatal(err)
		}
		message := []byte(`{"method":"subscribe","topic":"transfers","address":"tbnb13095qugzf6d4078hnt9creqetwclvpn2e7yucr"}`)
		_, err = ws.Write(message)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("Send: %s\n", message)
		var msg = make([]byte, 512)
		readLen, err := ws.Read(msg)
		if err != nil {
			log.Fatal(err)
		}

		var data WssResponse
		err = json.Unmarshal(msg[:readLen], &data)
		if err != nil {
			log.Fatal(err)
		}

		res := returnMemo(data.Data.H)
		if res == "GRAMETH" {

		}
		time.Sleep(time.Second)
	}
}

func returnMemo(tx string) string {
	url := "https://testnet-dex.binance.org/api/v1/tx/" + tx

	header := req.Header{
		"Content-Type": "application/json",
	}

	response, err := req.Get(url+"?format=json", header)
	if err != nil {
		fmt.Println(err)
	}

	var data Tx

	err = response.ToJSON(&data)
	if err != nil {
		return ""
	}

	return data.Tx.Value.Memo
}
