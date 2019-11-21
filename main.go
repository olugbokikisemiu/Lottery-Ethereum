package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/joho/godotenv"
	"github.com/olugbokikisemiu/Lottery/lottery"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatalf("Could not load env: %v", err)
	}

	c := context.Background()

	client, err := ethclient.Dial(os.Getenv("LOCAL_GATEWAY"))
	if err != nil {
		log.Fatalf("Could not load env: %v", err)
	}

	defer client.Close()

	lotterySession := &lottery.Handler{
		Ctx:    c,
		Client: client,
		Local:  true,
	}

	lotterySession.Session = lotterySession.NewHandler("KEYSTORE2", "KEYPASS2")
	sess := lotterySession.LoadContract()
	fmt.Println("From Address", sess.CallOpts.From.Hex())

	address, err := lotterySession.GetAllPlayer()
	if err != nil {
		fmt.Println("Error %v\n ", err)
	}
	fmt.Println("address", address)
}
