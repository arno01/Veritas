package main

import(
  "os"
  "flag"
  "fmt"
  "strconv"
)

type CLI struct {
  bc *Blockchain
}

func (cli *CLI) printUsage() {
	fmt.Println("Usage:")
  fmt.Println("To add blocks to the chain: \naddblock -data BLOCK_DATA_HERE")
  fmt.Println("To print the blockchain:\nprintchain")
}

func (cli *CLI) validateArgs() {
	if len(os.Args) < 2 {
		cli.printUsage()
		os.Exit(1)
	}
}

func (cli *CLI) addBlock(data string) {
  cli.bc.AddBlock(data)
  fmt.Println("SUCCESS")
}

func (cli *CLI) printChain() {
  bci := cli.bc.Iterator()
  for{
    block := bci.Next()
    fmt.Printf("Prev. hash: %x\n", block.PrevHash)
		fmt.Printf("Data: %s\n", block.Data)
		fmt.Printf("Hash: %x\n", block.Hash)
		pow := NewProofOfWork(block)
		fmt.Printf("PoW: %s\n", strconv.FormatBool(pow.Validate()))
		fmt.Println()
    if len(block.PrevHash) ==0 {
      break
    }
  }
}

func (cli *CLI) Run(){
  cli.validateArgs()
  addBlockCmd := flag.NewFlagSet("addblock", flag.ExitOnError)
  printChainCmd := flag.NewFlagSet("printchain", flag.ExitOnError)
  addBlockData := addBlockCmd.String("data","","Block data")
  switch os.Args[1]{
  case "addblock":
    _ = addBlockCmd.Parse(os.Args[2:])
  case "printchain":
    _ = printChainCmd.Parse(os.Args[2:])
  default:
    cli.printUsage()
    os.Exit(1)
  }

  if addBlockCmd.Parsed(){
    if *addBlockData == "" {
      addBlockCmd.Usage()
      os.Exit(1)
    }
    cli.addBlock(*addBlockData)
  }
  if printChainCmd.Parsed() {
    cli.printChain()
  }
}
