package assignment01bca

import (
	"crypto/sha256"
	"fmt"
	"strconv"
	"time"
)

type block struct {
	Transaction  string
	Nonce        int
	PreviousHash string
	CurrentHash  string
	Index        int
	Timestamp    string
}

var Blockchain []block

func CalculateHash(stringToHash string) string {
	sum := sha256.Sum256([]byte(stringToHash))
	return fmt.Sprintf("%x", sum)
}

func NewBlock(transaction string, nonce int, previousHash string) *block {
	index := len(Blockchain)
	timestamp := time.Now().String()
	hashInput := transaction + strconv.Itoa(nonce) + previousHash + strconv.Itoa(index) + timestamp
	currentHash := CalculateHash(hashInput)
	newBlock := block{
		Transaction:  transaction,
		Nonce:        nonce,
		PreviousHash: previousHash,
		CurrentHash:  currentHash,
		Index:        index,
		Timestamp:    timestamp,
	}
	Blockchain = append(Blockchain, newBlock)
	return &Blockchain[len(Blockchain)-1]
}

func ListBlocks() {
	fmt.Println("=================================================================")
	fmt.Println("                        BLOCKCHAIN                              ")
	fmt.Println("=================================================================")
	for _, b := range Blockchain {
		fmt.Printf("Index        : %d\n", b.Index)
		fmt.Printf("Transaction  : %s\n", b.Transaction)
		fmt.Printf("Nonce        : %d\n", b.Nonce)
		fmt.Printf("Timestamp    : %s\n", b.Timestamp)
		fmt.Printf("PreviousHash : %s\n", b.PreviousHash)
		fmt.Printf("CurrentHash  : %s\n", b.CurrentHash)
		fmt.Println("-----------------------------------------------------------------")
	}
}

func ChangeBlock(index int, newTransaction string) {
	if index < 0 || index >= len(Blockchain) {
		fmt.Println("Error: Block index out of range!")
		return
	}
	Blockchain[index].Transaction = newTransaction
	fmt.Printf("[TAMPER] Block %d transaction changed to: \"%s\"\n", index, newTransaction)
}

func VerifyChain() bool {
	fmt.Println("=================================================================")
	fmt.Println("                     VERIFYING BLOCKCHAIN                       ")
	fmt.Println("=================================================================")
	valid := true
	for i, b := range Blockchain {
		hashInput := b.Transaction + strconv.Itoa(b.Nonce) + b.PreviousHash + strconv.Itoa(b.Index) + b.Timestamp
		expectedHash := CalculateHash(hashInput)
		if b.CurrentHash != expectedHash {
			fmt.Printf("[TAMPERED] Block %d: Hash mismatch detected!\n", i)
			fmt.Printf("           Stored   : %s\n", b.CurrentHash)
			fmt.Printf("           Expected : %s\n", expectedHash)
			valid = false
		} else {
			fmt.Printf("[VALID]    Block %d: Hash is correct.\n", i)
		}
		if i > 0 {
			if b.PreviousHash != Blockchain[i-1].CurrentHash {
				fmt.Printf("[BROKEN]   Block %d: PreviousHash does not match Block %d's hash.\n", i, i-1)
				valid = false
			}
		}
	}
	fmt.Println("-----------------------------------------------------------------")
	if valid {
		fmt.Println("RESULT: Blockchain is VALID and INTACT.")
	} else {
		fmt.Println("RESULT: Blockchain INTEGRITY COMPROMISED. Tampering detected!")
	}
	fmt.Println("=================================================================")
	return valid
}