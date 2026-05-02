package main

import (
	"fmt"
	"strings"

	"github.com/en1gm4-exe/assignment01bca"
)

func main() {
	fmt.Println(strings.Repeat("=", 65))
	fmt.Println("   Blockchain Assignment 1 - Go Implementation")
	fmt.Println("   Name  : Uzzam Arif")
	fmt.Println("   Roll No  : 22I-1748")
	fmt.Println(strings.Repeat("=", 65))

	
	// GENESIS BLOCK Nonce = Sum of digits of roll number: 2+2+1+7+4+8 = 24
	genesisNonce := 2 + 2 + 1 + 7 + 4 + 8
	genesisBlock := assignment01bca.NewBlock("Genesis Block - 22I-1748", genesisNonce, "0")
	fmt.Println("Genesis Block created.")

	
	// BLOCK 1
	block1 := assignment01bca.NewBlock("Uzzam sends 748 coins to Alice", 749, genesisBlock.CurrentHash)
	fmt.Println("Block 1 created.")

	
	// BLOCK 2 
	block2 := assignment01bca.NewBlock("Alice sends 50 coins to Bob", 750, block1.CurrentHash)
	fmt.Println("Block 2 created.")

	
	// BLOCK 3
	assignment01bca.NewBlock("Bob sends 100 coins to Charlie", 751, block2.CurrentHash)
	fmt.Println("Block 3 created.")

	
	// Print the blockchain
	fmt.Println("\n--- Printing the Full Blockchain ---")
	assignment01bca.ListBlocks()

	
	// Verify the blockchain (should be VALID)
	fmt.Println("\n--- Verify Blockchain (Before Tampering) ---")
	assignment01bca.VerifyChain()

	
	// Tamper with Block 1 using ChangeBlock()
	fmt.Println("\n--- Tampering with Block 1 ---")
	assignment01bca.ChangeBlock(1, "HACKER sends 9999 coins to Hacker Wallet")

	
	// Verify blockchain again (should detect TAMPERING)
	fmt.Println("\n--- Verify Blockchain (After Tampering) ---")
	assignment01bca.VerifyChain()
}
