package lottery

import (
	"context"
	"fmt"
	"log"
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/joho/godotenv"
)

var myEnv map[string]string

type Handler struct {
	Ctx     context.Context
	Session LotterySession
	Client  *ethclient.Client
	Local   bool
}

func loadEnv() {
	var err error
	if myEnv, err = godotenv.Read(".env"); err != nil {
		log.Printf("Could not load env 🙍🏻‍♂️: %v ", err)
	}
}

func (h *Handler) NewHandler(keypass string) LotterySession {
	loadEnv()

	privateKey, err := crypto.HexToECDSA(strings.TrimPrefix(myEnv[keypass], "0x"))
	if err != nil {
		log.Fatalf("Error loading private key from env: %v", err)
	}
	auth := bind.NewKeyedTransactor(privateKey)

	blockNumber, err := h.Client.HeaderByNumber(h.Ctx, nil)
	if err != nil {
		log.Fatalf("block number Error occured: %v", err)
	}

	fmt.Println(blockNumber.Number)

	auth.Nonce = big.NewInt(3)
	auth.GasLimit = 3000000
	auth.GasPrice = big.NewInt(1000000)

	return LotterySession{
		TransactOpts: *auth,
		CallOpts: bind.CallOpts{
			From:    auth.From,
			Context: h.Ctx,
		},
	}
}

func (h *Handler) GetBalance() *big.Int {
	bal, err := h.Client.BalanceAt(h.Ctx, h.Session.CallOpts.From, nil)
	if err != nil {
		log.Fatalf("An error occured: %+v", err)
	}
	return bal
}

func (h *Handler) DeployContract() LotterySession {
	loadEnv()
	address, tx, instance, err := DeployLottery(&h.Session.TransactOpts, h.Client)
	if err != nil {
		log.Fatalf("Deployment error occured: %v", err)
	}
	switch h.Local {
	case true:
		updateEnv("LOCAL_ADDRESS", address.Hex())
	default:
		updateEnv("ADDRESS", address.Hex())
	}
	h.Session.Contract = instance
	log.Println("Contract Deployed🕺🏼🕺🏼! Transaction hash: %s", tx.Hash().Hex())
	return h.Session
}

func (h *Handler) LoadContract() LotterySession {
	var address common.Address
	switch h.Local {
	case true:
		address = common.HexToAddress(myEnv["LOCAL_ADDRESS"])
	default:
		address = common.HexToAddress(myEnv["ADDRESS"])
	}
	instance, err := NewLottery(address, h.Client)
	if err != nil {
		log.Fatalf("Unable to load contract: %v", err)
	}

	h.Session.Contract = instance
	return h.Session
}

func (h *Handler) GetAllPlayer() ([]common.Address, error) {
	return h.Session.AllPlayer()
}

func (h *Handler) JoinLottery() (string, error) {
	tx, err := h.Session.Enter()
	if err != nil {
		return "", err
	}
	return tx.Hash().Hex(), nil
}

func (h *Handler) GetManager() string {
	add, _ := h.Session.Manager()
	return add.Hex()
}

func (h *Handler) SelectWinner() (string, error) {
	tx, err := h.Session.SelectWinner()
	if err != nil {
		return "", err
	}
	return tx.To().Hex(), nil
}

func updateEnv(key string, value string) {
	myEnv[key] = value
	if err := godotenv.Write(myEnv, ".env"); err != nil {
		log.Fatalf("Unable to update env: %v", err)
	}
}
