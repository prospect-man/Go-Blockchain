package blockchain

import (
	"bytes"
	"crypto/sha256"
	"encoding/gob"
	"strconv"
	"time"
)

//Block is a exported variable
type Block struct {
	Index     int64    // The index of the block
	Timestamp int64    // Time stamp in epoch
	Data      []string // The data we want to store
	PrevHash  string   // chain/hash of previous block
	CurrHash  string   // Current hash
}

//GenesisBlock exported
var GenesisBlock = Block{
	Index:     1,
	Timestamp: time.Now().Unix(),
	Data:      []string{"genesis Data"},
	PrevHash:  "",
	CurrHash:  "genesis hash",
}

//Blockchain exported
type Blockchain struct {
	Chain []Block
}

//NewBlock exported
func (bc *Blockchain) NewBlock(lb Block, dt string) *Block {

	blockinstance := &Block{}

	blockinstance.Index = int64(len(bc.Chain) + 1)
	blockinstance.Timestamp = time.Now().Unix()
	blockinstance.Data = append(lb.Data, dt)
	blockinstance.PrevHash = lb.CurrHash
	blockinstance.CurrHash = Hash(lb)

	//fmt.Printf("index:%d \n", blockinstance.Index)
	//fmt.Println("lang of the change:", len(bc.Chain))
	//bc.Chain = append(bc.Chain, *blockinstance)

	return blockinstance
}

//Hash exported
func Hash(d Block) string {
	index := []byte(strconv.FormatInt(d.Index, 10))
	timestamp := []byte(strconv.FormatInt(d.Timestamp, 10))
	// Convert []string to []byte e
	buffer := &bytes.Buffer{}
	gob.NewEncoder(buffer).Encode(d.Data)
	data := buffer.Bytes()
	prevHash := []byte(d.PrevHash)
	dataStr := bytes.Join([][]byte{prevHash, timestamp, index, data}, []byte{})
	hash := sha256.Sum256(dataStr)
	hashStr := string(hash[:])
	return hashStr
}

//AddBlock Exported
func (bc *Blockchain) AddBlock(lsb Block, dat string) {
	bc.Chain = append(bc.Chain, *bc.NewBlock(lsb, dat))
	//return bc.Chain[len(bc.Chain)-1]
}
