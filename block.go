package block

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
	Nonce     string   // String that needs to be mined
	Target    int64    // Number of leading zeros
	CurrHash  string   // Current hash
}

//GenesisBlock exported
var GenesisBlock = Block{
	Index:     0,
	Timestamp: 0,
	Data:      []string{"genesis Data"},
	PrevHash:  "No prevHash",
	Nonce:     "000000",
	Target:    10,
	CurrHash:  "genesis hash",
}

//MiningBlock exported
func MiningBlock(lb Block, data []string) (cb Block) {
	cb.Index = lb.Index + 1
	cb.Timestamp = time.Now().Unix()
	cb.Data = append(lb.Data, data...)
	cb.PrevHash = lb.CurrHash
	cb.Nonce = lb.Nonce + "111111"
	cb.Target = 10 + lb.Target
	cb.CurrHash = hash(lb)
	return cb
}

func hash(d Block) string {
	index := []byte(strconv.FormatInt(d.Index, 10))
	timestamp := []byte(strconv.FormatInt(d.Timestamp, 10))
	// Convert []string to []byte e
	buffer := &bytes.Buffer{}
	gob.NewEncoder(buffer).Encode(d.Data)
	data := buffer.Bytes()
	prevHash := []byte(d.PrevHash)
	nonce := []byte(d.Nonce)
	target := []byte(strconv.FormatInt(d.Target, 10))
	dataStr := bytes.Join([][]byte{prevHash, timestamp, index, data, nonce, target}, []byte{})
	hash := sha256.Sum256(dataStr)
	hashStr := string(hash[:])
	return hashStr
}
