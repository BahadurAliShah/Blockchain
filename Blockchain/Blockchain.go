package main

import (
	"bufio"
	"crypto/sha256"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func CalculateHash(stringToHash string) string {
	return fmt.Sprintf("%x", sha256.Sum256([]byte(stringToHash)))
}

type data struct {
	previousHash string
	transaction  string
	nonce        int
}

func (b *data) AddBlock(previousHash string, transaction string, nonce int) *data {
	b.transaction = transaction
	b.nonce = nonce
	if previousHash == "" {
		b.previousHash = "0000000000000000000000000000000000000000000000000000000000000000"
	} else {
		b.previousHash = previousHash
	}
	b.previousHash = previousHash
	return b
}

func (b *data) print() {
	fmt.Println("Nonce: ", b.nonce)
	fmt.Println("Transaction: ", b.transaction)
	fmt.Println("Previous Hash: ", b.previousHash)
}

type block struct {
	nodes []data
}

var BlockChain = new(block)

func NewBlock(transaction string, nonce int, previousHash string) *block {
	BlockChain.nodes = append(BlockChain.nodes, *new(data).AddBlock(previousHash, transaction, nonce))
	return BlockChain
}

func ListBlocks() {
	if len(BlockChain.nodes) == 0 {
		fmt.Println("No Blocks")
		return
	}
	for i, v := range BlockChain.nodes {
		fmt.Println("Node: ", i)
		v.print()
		fmt.Println(strings.Repeat("-", 50))
	}
}

func ChangeBlock(ID int, transaction string, nonce int, previousHash string) {
	if ID >= len(BlockChain.nodes) || ID < 0 {
		fmt.Println("Invalid ID")
		return
	} else {
		BlockChain.nodes[ID].AddBlock(previousHash, transaction, nonce)
	}
}

func VerifyChain() bool {
	for i, v := range BlockChain.nodes {
		if i == 0 {
			continue
		}
		if v.previousHash != CalculateHash(BlockChain.nodes[i-1].transaction+strconv.Itoa(BlockChain.nodes[i-1].nonce)) {
			return false
		}
	}
	return true
}

func printMenu() int {
	fmt.Println("1. Add Block")
	fmt.Println("2. List Blocks")
	fmt.Println("3. Change Block")
	fmt.Println("4. Verify Chain")
	fmt.Println("5. Exit")
	var choice int
	fmt.Scanln(&choice)
	for choice < 1 || choice > 5 {
		fmt.Println("Invalid Choice")
		fmt.Scanln(&choice)
	}
	return choice
}

func getTopHash() string {
	if len(BlockChain.nodes) == 0 {
		return "0000000000000000000000000000000000000000000000000000000000000000"
	}
	return CalculateHash(BlockChain.nodes[len(BlockChain.nodes)-1].transaction + strconv.Itoa(BlockChain.nodes[len(BlockChain.nodes)-1].nonce))
}

func main() {
	for {
		switch printMenu() {
		case 1:
			var transaction string
			var nonce int
			fmt.Println("Enter Transaction: ")
			scanner := bufio.NewScanner(os.Stdin)
			scanner.Scan()
			transaction = scanner.Text()
			fmt.Println("Enter Nonce: ")
			fmt.Scanln(&nonce)
			NewBlock(transaction, nonce, getTopHash())
		case 2:
			ListBlocks()
		case 3:
			var ID int
			var transaction string
			var nonce int
			if len(BlockChain.nodes) == 0 {
				fmt.Println("No Blocks")
				break
			}
			fmt.Println("Enter ID: ")
			fmt.Scanln(&ID)
			fmt.Println("Enter Transaction: ")
			scanner := bufio.NewScanner(os.Stdin)
			scanner.Scan()
			transaction = scanner.Text()
			fmt.Println("Enter Nonce: ")
			fmt.Scanln(&nonce)
			ChangeBlock(ID, transaction, nonce, "")
		case 4:
			if VerifyChain() {
				fmt.Println("Chain is Valid")
			} else {
				fmt.Println("Chain is Invalid")
			}
		case 5:
			return
		}
	}
}
