# go_blockchain

alpha version (0.5) of the implementation of a simple blockchain with Golang.
Uses `Module` structure.
Uses  `BadgerDB` for persistance and `CLI` (command line interface)

## TODO:

1) add Transactions
2) add a walet module
3) build a digital signature
4) add UTXO persistence layer  (Fansy!!)
5) adding the Merkel tree (even fansier! )
6) adding the network module
7) Refactor ...

## Usage

1) for printing out elements of your blockchain in your terminal run

`% go run main.go print`

2) for adding block to your blockchain in your terminal run

`% go run main.go add -block "some string"`

