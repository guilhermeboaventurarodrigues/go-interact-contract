package main

import (
	"fmt"
	"log"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/guilhermeboaventurarodrigues/token-go/tokens"
)

func main() {
	client, err := ethclient.Dial("HTTP://127.0.0.1:7545") // IP GANACHE
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("we have a connection")
	_ = client

	contract, err := token.NewToken(common.HexToAddress("0xf94Fe9E829014e3e8Dc86b51b1a5D3e6E19eB069"), client) //ADDRESS CONTRACT DEPLOYED
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("contract is loaded")
	_ = contract

	name, err := contract.Name(&bind.CallOpts{})
	if err != nil {
		log.Fatal(err)
	}

	symbol, err := contract.Symbol(&bind.CallOpts{})
	if err != nil {
		log.Fatal(err)
	}

	decimals, err := contract.Decimals(&bind.CallOpts{})
	if err != nil {
		log.Fatal(err)
	}

	totalsupply, err := contract.TotalSupply(&bind.CallOpts{})
	if err != nil{
		log.Fatal(err)
	}

	balanceof, err := contract.BalanceOf(&bind.CallOpts{}, common.HexToAddress("0xfeE13d798f3B9cbad32D134b8BE5F973d7e32568")) //ADDRESS OWNER CONTRACT
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("------------------------")
	fmt.Println("Name token:",name)
	fmt.Println("------------------------")
	fmt.Println("Symbol token:",symbol)
	fmt.Println("------------------------")
	fmt.Println("Decimals token:",decimals)
	fmt.Println("------------------------")
	fmt.Println("Total Supply:", totalsupply)
	fmt.Println("------------------------")
	fmt.Println("BalanceOf Owner:", balanceof)
	fmt.Println("------------------------")
	fmt.Println("Me contrate :)")
}