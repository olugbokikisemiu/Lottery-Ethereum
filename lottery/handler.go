package lottery

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"log"
	"math/big"
	"os"
	"sort"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/joho/godotenv"
)

var myEnv map[string]string

// Handler handles common needed values
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

// NewHandler binds transactor through private key
// and returns the session
func (h *Handler) NewHandler(key string) LotterySession {
	loadEnv()

	privateKey, err := crypto.HexToECDSA(strings.TrimPrefix(key, "0x"))
	if err != nil {
		log.Fatalf("Error loading private key from env: %v", err)
	}
	auth := bind.NewKeyedTransactor(privateKey)

	auth.Nonce = big.NewInt(0)
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

// GetBalance returns the balance of specified address
func (h *Handler) GetBalance() *big.Int {
	bal, err := h.Client.BalanceAt(h.Ctx, h.Session.CallOpts.From, nil)
	if err != nil {
		log.Fatalf("An error occured: %+v", err)
	}
	return bal
}

// DeployContract deploys smart contract
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
	log.Printf("Contract Deployed🕺🏼🕺🏼! Transaction hash: %s", tx.Hash().Hex())
	return h.Session
}

// LoadContract loads smart contract
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

// GetAllPlayer return []address of all players in lottery
func (h *Handler) GetAllPlayer() ([]common.Address, error) {
	return h.Session.AllPlayer()
}

// JoinLottery adds user to the lottery system([]address)
func (h *Handler) JoinLottery() (string, error) {
	h.Session.TransactOpts.Value = big.NewInt(1000000000000000000)
	tx, err := h.Session.Enter()
	if err != nil {
		return "", err
	}
	return tx.Hash().Hex(), nil
}

// GetManager returns the address that deployed the contract
func (h *Handler) GetManager() string {
	add, _ := h.Session.Manager()
	return add.Hex()
}

// SelectWinner choose an address at random and send
// the coins to it
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

// PrivateKeys gets the private key of acccounts from
// ganache json file
func PrivateKeys() []string {
	jsonFile, err := os.Open("keystore/ganache-accounts.json")
	if err != nil {
		log.Fatalf("Error opening file %+v ", err)
	}

	defer jsonFile.Close()

	var ganacheConfig map[string]map[string]interface{}
	var privateKeysList []string

	value, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		log.Fatalf("Error reading file %+v ", err)
	}
	if err := json.Unmarshal(value, &ganacheConfig); err != nil {
		log.Fatalf("Error unmarshalling json %+v ", err)
	}

	for _, v := range ganacheConfig["private_keys"] {
		privateKeysList = append(privateKeysList, v.(string))
	}

	sort.Strings(privateKeysList)

	return privateKeysList
}
