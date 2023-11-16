package blockchain

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"time"
)

type Block struct {
	Timestamp     int64
	Transactions  []Transaction
	PrevBlockHash string
	Hash          string
}

func NewBlock(transactions []Transaction, prevBlockHash string) *Block {
	block := &Block{Timestamp: time.Now().Unix(), Transactions: transactions, PrevBlockHash: prevBlockHash}
	block.calculateHash()
	return block
}

func (b *Block) calculateHash() {
	data := string(b.Timestamp) + b.PrevBlockHash
	for _, tx := range b.Transactions {
		data += tx.Sender + tx.Receiver + fmt.Sprintf("%f", tx.Amount)
	}
	hash := sha256.Sum256([]byte(data))
	b.Hash = hex.EncodeToString(hash[:])
}
