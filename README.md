# Creating-a-simple-Blockchain
Building a simplified Blockchain in Go — basically a linked list of blocks where each block stores a transaction, and is cryptographically linked to the previous one via SHA-256 hashes. If anyone tampers with a block's data, the hash breaks and our `VerifyChain()` function catches it.


    go mod init assignment01

    cd assignment01bca
    go mod init github.com/en1gm4-exe/assignment01bca
    cd ..

    go mod edit -replace github.com/en1gm4-exe/assignment01bca=./assignment01bca

    go mod tidy

    go run main.go
