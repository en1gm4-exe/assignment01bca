# Creating-a-simple-Blockchain
Building a simplified Blockchain in Go — basically a linked list of blocks where each block stores a transaction, and is cryptographically linked to the previous one via SHA-256 hashes. If anyone tampers with a block's data, the hash breaks and our `VerifyChain()` function catches it.
