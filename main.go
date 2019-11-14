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

	lotterySession.Session = lotterySession.NewHandler("KEYSTORE", "KEYPASS")
	lotterySession.LoadContract()
	address, _ := lotterySession.GetAllPlayer()
	// tx, err := lotterySession.JoinLottery()
	// if err != nil {
	// 	fmt.Printf("%v\n", err)
	// }
	// fmt.Println("Transaction hash ", tx)
	// fmt.Println("From address ", lotterySession.Session.CallOpts.From.Hex())
	fmt.Println(" address ", address)
}
