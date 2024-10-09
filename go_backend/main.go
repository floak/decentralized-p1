package main

import (
	"goproject1/identity" 
	"context"
    "fmt"
    "log"
	"math/big"

    "github.com/ethereum/go-ethereum"
    "github.com/ethereum/go-ethereum/rpc"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
    "github.com/ethereum/go-ethereum/common"
    "github.com/ethereum/go-ethereum/ethclient"
    // Import your generated Go bindings
)

func main() {
    // Connect to the local Ganache Ethereum client
    client, err := rpc.Dial("http://127.0.0.1:7545")
    if err != nil {
        log.Fatalf("Failed to connect to the Ethereum client: %v", err)
    }
    //fmt.Println("Connected to Ethereum client")

    // Add your smart contract interaction code here...
	// Specify the deployed contract address
	contractAddress := common.HexToAddress("0xf82e171d3ff6f140371e5d199bc29506c35a4941") // Replace with your contract address
	instance, err := identity.NewIdentity(contractAddress, client)
	if err != nil {
		log.Fatalf("Failed to instantiate the contract: %v", err)
	}

	// Call getIdentity function
	accountAddress := common.HexToAddress("0x190DdBD790E3b58D9337560A9391cc65fFF5aF6B") // Replace with the account address
	did, pubKey, err := instance.GetIdentity(&bind.CallOpts{}, accountAddress)
	if err != nil {
		log.Fatalf("Failed to call contract function: %v", err)
	}

	fmt.Printf("DID: %s, Public Key: %s\n", did, pubKey)

	// Create a transactor to send a transaction
	privateKey := "0x895f8b320a57818c9c9bc210448e151c4ac1721fd84d5313606c4fd8fd394e23" // Replace with your account's private key
	auth, err := bind.NewTransactorWithChainID(strings.NewReader(privateKey), "password", big.NewInt(1337))
	if err != nil {
		log.Fatalf("Failed to create transactor: %v", err)
	}

	// Call registerIdentity function
	tx, err := instance.RegisterIdentity(auth, "did:example:123", "0xPublicKeyHere")
	if err != nil {
		log.Fatalf("Failed to execute transaction: %v", err)
	}

	fmt.Printf("Transaction submitted: %s\n", tx.Hash().Hex())

}
