package ethereum

import (
	"context"
	"crypto/ecdsa"
	"github.com/ethereum/go-ethereum"
	//"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"log"
	"math"
	"math/big"
)

var DecimalMultiplier = new(big.Float).SetFloat64(math.Pow(10, 18))

type Contract struct {
	NetworkID types.EIP155Signer
	Address   common.Address
	//Instance  *airdrop.Airdrop
	Client *ethclient.Client
}

type SingleTransaction struct {
	NetworkID  types.EIP155Signer
	Client     *ethclient.Client
	privateKey string
}

//func ConnectContract(endpoint string, contractAddress string) (*Contract, error) {
//	client, err := ethclient.Dial(endpoint)
//	if err != nil {
//		return nil, err
//	}
//	chainID, err := client.NetworkID(context.Background())
//	if err != nil {
//		return nil, err
//	}
//	address := common.HexToAddress(contractAddress)
//	instance, err := airdrop.NewAirdrop(address, client)
//	if err != nil {
//		return nil, err
//	}
//
//	return &Contract{
//		NetworkID: types.NewEIP155Signer(chainID),
//		Address:   address,
//		Instance:  instance,
//		Client:    client,
//	}, nil
//}

func Connect(endpoint, privateKey string) (*SingleTransaction, error) {
	client, err := ethclient.Dial(endpoint)
	if err != nil {
		return nil, err
	}
	chainID, err := client.NetworkID(context.Background())
	if err != nil {
		return nil, err
	}

	return &SingleTransaction{
		NetworkID:  types.NewEIP155Signer(chainID),
		Client:     client,
		privateKey: privateKey,
	}, nil
}

func (contract *Contract) SignTransaction(hexPrivateKey string, value *big.Int, data []byte) (*types.Transaction, error) {
	privateKey, err := crypto.HexToECDSA(hexPrivateKey)
	if err != nil {
		return nil, err
	}

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
	}
	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	nonce, err := contract.Client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		return nil, err
	}

	gasPrice, err := contract.Client.SuggestGasPrice(context.Background())
	if err != nil {
		return nil, err
	}

	gasLimit, err := contract.Client.EstimateGas(context.Background(), ethereum.CallMsg{
		From:     fromAddress,
		Value:    value,
		To:       &contract.Address,
		GasPrice: gasPrice,
		Data:     data,
	})
	if err != nil {
		return nil, err
	}

	tx := types.NewTransaction(nonce, contract.Address, value, gasLimit, gasPrice, data)
	signedTx, err := types.SignTx(tx, contract.NetworkID, privateKey)
	if err != nil {
		return nil, err
	}
	return signedTx, nil
}

func (st *SingleTransaction) SignTransaction(to string, value *big.Int) (*types.Transaction, error) {
	privateKey, err := crypto.HexToECDSA(st.privateKey)
	if err != nil {
		return nil, err
	}
	publicKeyFrom := privateKey.Public()
	publicKeyECDSAFrom, ok := publicKeyFrom.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
	}
	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSAFrom)
	toAddress := common.HexToAddress(to)

	nonce, err := st.Client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		return nil, err
	}

	gasPrice, err := st.Client.SuggestGasPrice(context.Background())
	if err != nil {
		return nil, err
	}

	gasLimit, err := st.Client.EstimateGas(context.Background(), ethereum.CallMsg{
		From:     fromAddress,
		Value:    value,
		To:       &toAddress,
		GasPrice: gasPrice,
		Data:     nil,
	})
	if err != nil {
		return nil, err
	}

	tx := types.NewTransaction(nonce, toAddress, value, gasLimit, gasPrice, nil)
	signedTx, err := types.SignTx(tx, st.NetworkID, privateKey)
	if err != nil {
		return nil, err
	}
	return signedTx, nil
}

func (contract *Contract) SendSignedTransaction(signedTx *types.Transaction) (string, error) {
	err := contract.Client.SendTransaction(context.Background(), signedTx)
	if err != nil {
		return "", err
	}
	return signedTx.Hash().Hex(), nil
}

func (st *SingleTransaction) SendSignedTransaction(signedTx *types.Transaction) (string, error) {
	err := st.Client.SendTransaction(context.Background(), signedTx)
	if err != nil {
		return "", err
	}
	return signedTx.Hash().Hex(), nil
}

/*func (contract *Contract) GetDataByteCode(methodName string, params ...interface{}) ([]byte, error) {
	airdropAbi, err := abi.JSON(strings.NewReader(airdrop.AirdropABI))
	if err != nil {
		return nil, err
	}
	return airdropAbi.Pack(methodName, params...)
}*/

func HexToAddress(address string) common.Address {
	return common.HexToAddress(address)
}

func EtherToInt(ether float64) *big.Int {
	floatWei := new(big.Float)
	intWei := new(big.Int)
	amount := new(big.Float).SetFloat64(ether)
	floatWei = floatWei.Mul(amount, DecimalMultiplier)
	w, _ := floatWei.Uint64()
	intWei.SetUint64(w)
	return intWei
}
