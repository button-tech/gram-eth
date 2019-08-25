package waves

import (
	"fmt"
	"github.com/wavesplatform/gowaves/pkg/wallet"
	"github.com/wavesplatform/gowaves/pkg/client"
	"os"
)

func SendWavesToAddress(address string) {

	seed := os.Getenv("WAVES_MNEMONIC")

	w, err := wallet.NewWalletFromSeed([]byte(seed))
	if err != nil {
		fmt.Println(err)
		return
	}

	if err != nil {
		fmt.Println(err)
		return
	}

	s, p , err := w.GenPair()


	_, err = client.NewClient(client.Options{
		BaseUrl: "https://testnode1.wavesnodes.com",
	})
	if err != nil {
		fmt.Println(err)
		return
	}


	//tx := proto.NewUnsignedPayment(p , address, 100000, 100000, uint64(time.Now().Unix()))

	//err = tx.Sign(s)
	//if err != nil{
	//	fmt.Println(err)
	//	return
	//}


}