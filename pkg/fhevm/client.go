package fhevm

import (
	"context"
	"crypto/ecdsa"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

type Client struct {
	rpcUrl string
	c      *ethclient.Client
	key    *ecdsa.PrivateKey
	from   common.Address
}

func NewClient(rpcUrl string, privHex string) (*Client, error) {
	c, err := ethclient.Dial(rpcUrl)
	if err != nil {
		return nil, err
	}
	key, err := crypto.HexToECDSA(privHex)
	if err != nil {
		return nil, err
	}
	from := crypto.PubkeyToAddress(key.PublicKey)
	return &Client{rpcUrl: rpcUrl, c: c, key: key, from: from}, nil
}

func (cl *Client) SendRawDataToContract(to common.Address, data []byte) (string, error) {
	// Build simple tx (no gas price tuning). For testnet/local only.
	nonce, err := cl.c.PendingNonceAt(context.Background(), cl.from)
	if err != nil {
		return "", err
	}
	gasPrice, err := cl.c.SuggestGasPrice(context.Background())
	if err != nil {
		return "", err
	}
	tx := types.NewTransaction(nonce, to, big.NewInt(0), uint64(300000), gasPrice, data)
	chainID, err := cl.c.NetworkID(context.Background())
	if err != nil {
		return "", err
	}
	signedTx, err := types.SignTx(tx, types.NewEIP155Signer(chainID), cl.key)
	if err != nil {
		return "", err
	}
	err = cl.c.SendTransaction(context.Background(), signedTx)
	if err != nil {
		return "", err
	}
	return signedTx.Hash().Hex(), nil
}

// Helper to create transactor if need call contract methods via bindings
func (cl *Client) NewTransactor() (*bind.TransactOpts, error) {
	auth, err := bind.NewKeyedTransactorWithChainID(cl.key, big.NewInt(1337)) // adjust chainID
	if err != nil {
		return nil, err
	}
	auth.GasLimit = uint64(300000)
	return auth, nil
}
