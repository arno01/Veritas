package main

import (
	"bytes"
	"crypto/sha256"
	"time"
  "encoding/gob"
)

type Block struct {
	Timestamp int64
	Transactions []*Transaction
	PrevHash  []byte
	Hash      []byte
	Nonce    int
}

func NewBlock(transactions []*Transaction, prevHash []byte) *Block {
	timestamp := time.Now().Unix()
	block := &Block{timestamp, transactions, prevHash, []byte{}, 0}
	pow := NewProofOfWork(block)
	nonce, hash := pow.Run()
	block.Hash = hash[:]
	block.Nonce = nonce
	return block
}

func NewGenesisBlock(base *Transaction) *Block {
	return NewBlock([]*Transaction{base}, []byte{})
}

//all transactions in a block identified by a single hash
func (b *Block) HashTransactions() []byte {
  var txHashes [][]byte
  var txHash [32]byte
  for _,tx := range b.Transactions{
    txHashes = append(txHashes, tx.ID)
  }
  txHash = sha256.Sum256(bytes.Join(txHashes,[]byte{}))
  return txHash[:]
}

func (b *Block) Serialize() []byte {
  var result bytes.Buffer
  encoder := gob.NewEncoder(&result)
  _ = encoder.Encode(b)
  return result.Bytes()
}

func DeserializeBlock(d []byte) *Block {
  var block Block
  decoder := gob.NewDecoder(bytes.NewReader(d))
  _ = decoder.Decode(&block)
  return &block
}
