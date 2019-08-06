package main

import (
	"blockchain"
	"fmt"
)

func main() {

	var bc1 blockchain.Blockchain

	bc1.AddBlock(blockchain.GenesisBlock, "test0")

	fmt.Println(bc1.Chain[0].Index)
	fmt.Println(bc1.Chain[0].Timestamp)
	fmt.Println(bc1.Chain[0].Data)
	fmt.Printf("%X\n", bc1.Chain[0].PrevHash)
	fmt.Printf("%X\n", bc1.Chain[0].CurrHash)
	fmt.Println("lang of the change:", len(bc1.Chain))

	bc1.AddBlock(bc1.Chain[0], "test1")
	fmt.Println(bc1.Chain[1].Index)
	fmt.Println(bc1.Chain[1].Timestamp)
	fmt.Println(bc1.Chain[1].Data)
	fmt.Printf("%X\n", bc1.Chain[1].PrevHash)
	fmt.Printf("%X\n", bc1.Chain[1].CurrHash)
	fmt.Println("lang of the change:", len(bc1.Chain))

}
