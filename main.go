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

	lotterySession.Session = lotterySession.NewHandler(lottery.PrivateKeys()[0])

	lotterySession.LoadContract()

	address, err := lotterySession.GetAllPlayer()
	if err != nil {
		fmt.Printf("Error %v\n ", err)
	}

	fmt.Println("Address", address)
	fmt.Println("Balance ", lotterySession.GetBalance())
}
