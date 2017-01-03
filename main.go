package main

import (
	"fmt"

	"github.com/FactomProject/factom"
	"time"
)

var _ = factom.AddressLength
var _ = fmt.Sprintf("")

// EC2EcU4v4f82vbKYJhbbEJdKvhLBRJhjKKZwZmZwNhVZHBL3NpJ4
var ESKey string = "Es3gZoQbNd2p2nDDRtULkUaneoSJY1WTCQ7LSyNqHWZ2UkttuS1o"

// FA3EPZYqodgyEGXNMbiZKE5TS2x2J9wF8J9MvPZb52iGR78xMgCb
var FSKey string = "Fs2DNirmGDtnAZGXqca3XHkukTNMxoMGFFQxJA3bAjJnKzzsZBMH"

func main() {
	// Set up factom library
	factom.SetFactomdServer("localhost:8088")
	factom.SetWalletServer("localhost:8089")

	// Get our entry credit key ready
	ec, err := factom.GetECAddress(ESKey)
	if err != nil {
		panic(err)
	}

	// Make a Chain
	// A chain's first entry determines it's chainID
	exIds := make([][]byte, 2)
	exIds[0] = []byte("Unique Identifier")
	exIds[1] = []byte("Make as many of these as you want")

	message := []byte("Hello World")

	firstEntry := CreateEntry(exIds, message)
	chain := factom.NewChain(firstEntry)

	fmt.Println("ChainID:", chain.ChainID)

	// Now we have the chain ready
	// Commit an Chain
	cCom, err := factom.CommitChain(chain, ec)
	if err != nil {
		panic(err)
	}
	fmt.Println("Chain commit successful:", cCom)

	//time.Sleep(10 * time.Second)
	// Reveal an Chain
	eHash, err := factom.RevealChain(chain)
	if err != nil {
		panic(err)
	}
	fmt.Println("Chain reveal successful:", eHash)

	time.Sleep(10 * time.Second)
	ent, err := factom.GetEntry(eHash)
	if err != nil {
		panic(err)
	}

	fmt.Println("Got it back!")
	fmt.Println(ent.String())
}

func CreateEntry(extIds [][]byte, content []byte) *factom.Entry {
	entry := new(factom.Entry)
	entry.ExtIDs = extIds
	entry.Content = content

	return entry
}
