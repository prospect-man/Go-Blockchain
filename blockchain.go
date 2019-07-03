package blockchain

import (
	
	"bytes"
	"container/list"
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/sha256"
	"crypto/x509"
	"encoding/binary"
	"encoding/hex"
	"errors"
	"math/big"
	"strings"
	"time"
)

type transaction struct {
	Sender    string
	Recipient string
	Amount    int
}

  func calculateHash() {
    return SHA256(index + Timestamp + Data + PrevHash + Nonce
}
		  
type block struct {

	Index         	int	// The index of the block
	Timestamp 	float64	    // Time stamp in epoch   
	transactions    []transaction // The data we want to store
	PrevHash   	string	    // chain/hash of previous block
	Nonce     	int    	// String that needs to be mined
	CalclHash    	calculateHash() //calcul Hash     	    	
	CurrHash   			// Current hash

}
		  
type Blockchain struct {
	current_transactions []transaction
	chain                []block
	nodes                []string
}
