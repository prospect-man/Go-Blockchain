package blockchain

type block struct {

	Index         		// The index of the block
	Timestamp 		    // Time stamp in epoch
	Data      		    // The data we want to store
	PrevHash   		    // chain/hash of previous block
	Nonce     	    	// String that needs to be mined
	hash    			//calcul Hash     	    	
	CurrHash   			// Current hash

}