# Creating-a-simple-Blockchain
This project implements a simplified blockchain in Go, as part of a Blockchain course assignment. It demonstrates the core concepts of a blockchain: blocks linked via cryptographic hashes, immutability, tamper detection, and chain verification.

The blockchain is essentially a linked list where each block contains:
- A transaction string
- A nonce (arbitrary integer)
- The hash of the previous block
- Its own hash (SHA‑256 of `transaction + nonce + previousHash + index + timestamp`)
- Index (position in the chain)
- Timestamp

If any block’s data is altered, its hash changes, breaking the link to the next block. The `VerifyChain()` function detects such tampering.


## Setup

### 1. Initialise Go modules

        go mod init assignment01
        cd assignment01bca
        go mod init github.com/en1gm4-exe/assignment01bca
        cd ..

### 2. Set up the local replace directive (so the main program can find the package):
        go mod edit -replace github.com/en1gm4-exe/assignment01bca=./assignment01bca


### 3. Tidy up the modules:
        go mod tidy

### 4. Execution
        go run main.go
