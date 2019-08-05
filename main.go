package main

import (
	"block"
	"fmt"
)

func main() {
	var b1 block.Block // instance of Block type

	b1 = block.MiningBlock(block.GenesisBlock, []string{"foo", "data"})

	fmt.Println(b1.Index)
	fmt.Println(b1.Timestamp)
	fmt.Println(b1.Data)
	fmt.Println(b1.PrevHash)
	fmt.Println(b1.Nonce)
	fmt.Println(b1.Target)
	fmt.Printf("%X\n", b1.CurrHash)

}
