package waves
/*
import (
	"fmt"
	"github.com/wavesplatform/gowaves/pkg/proto"
	"github.com/wavesplatform/gowaves/pkg/wallet"
	"github.com/wavesplatform/gowaves/pkg/client"
	"github.com/wavesplatform/gowaves/pkg/libs/serializer"
	"time"
)

func send() {
	seed := "author robust mixture despair head mind behave resemble code swift into bird inner spike gravity"
	w, err := wallet.NewWalletFromSeed([]byte(seed))
	if err != nil {
		fmt.Println(err)
		return
	}
	s, p , err := w.GenPair()

	c, err := client.NewClient(client.Options{
		BaseUrl: "https://testnode1.wavesnodes.com",
	})
	if err != nil {
		fmt.Println(err)
		return
	}
	tx := proto.NewUnsignedPayment(p , proto.Address{[]byte("sdffeg")}, 100000, 100000, uint64(time.Now().Unix()))
	err := tx.Sign(s)
}
*/