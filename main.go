package main

import (
  "bufio"
  "crypto/sha256"
  "encoding/hex"
  "encoding/json"
  "io"
  "log"
  "net"
  "os"
  "strconv"
  "time"

  "github.com/davecgh/go-spew/spew"
	"github.com/joho/godotenv"
)

type Block struct {
  Idx int
  Timestamp string
  Val int
  Hash string
  Prevhash string
}

var bcServer chan []Block
var Blockchain []Block

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}
	bcServer = make(chan []Block)
  t := time.Now()
  block_0 := Block{0, t.String(), 0, "",""}
  spew.Dump(block_0)
	Blockchain = append(Blockchain, block_0)
  //start TCP
  server, err := net.Listen("tcp", ":"+os.Getenv("ADDR"))
  if err != nil {
    log.Fatal(err)
  }
  defer server.Close()
  for {
		conn, err := server.Accept()
		if err != nil {
			log.Fatal(err)
		}
		go handleConn(conn)
	}
}

func handleConn(conn net.Conn) {
	defer conn.Close()
  io.WriteString(conn, "Enter a new value:")
  scanner := bufio.NewScanner(conn)
  // add val to blockchain
  go func() {
  	for scanner.Scan() {
  		val, err := strconv.Atoi(scanner.Text())
  		if err != nil {
  			log.Printf("%v not a number: %v", scanner.Text(), err)
  			continue
  		}
  		newBlock, err := generateBlock(Blockchain[len(Blockchain)-1], val)
  		if err != nil {
  			log.Println(err)
  			continue
  		}
  		if isBlockValid(newBlock, Blockchain[len(Blockchain)-1]) {
  			newBlockchain := append(Blockchain, newBlock)
  			replaceChain(newBlockchain)
  		}

  		bcServer <- Blockchain
  		io.WriteString(conn, "\nEnter a new value:")
  	}
  }()
  // simulate receiving broadcast
	go func() {
		for {
			time.Sleep(30 * time.Second)
			output, err := json.Marshal(Blockchain)
			if err != nil {
				log.Fatal(err)
			}
			io.WriteString(conn, string(output))
		}
	}()

	for _ = range bcServer {
		spew.Dump(Blockchain)
	}
}

func calculateHash(block Block) string {
  record := string(block.Idx) + block.Timestamp + string(block.Val) + block.Prevhash
  h := sha256.New()
  h.Write([]byte(record))
  hashed := h.Sum(nil)
  return hex.EncodeToString(hashed)
}

func generateBlock(oldBlock Block, Val int) (Block, error) {
  var newBlock Block
  newBlock.Idx = oldBlock.Idx +1
  newBlock.Timestamp = time.Now().String()
  newBlock.Val = Val
  newBlock.Prevhash = oldBlock.Hash
  newBlock.Hash = calculateHash(newBlock)
  return newBlock, nil
}

func isBlockValid(newBlock, oldBlock Block) bool{
  if oldBlock.Idx+1 != newBlock.Idx {
    return false
  }
  if oldBlock.Hash != newBlock.Prevhash{
    return false
  }
  if calculateHash(newBlock) != newBlock.Hash{
    return false
  }
  return true
}

func replaceChain(newBlocks []Block){
  if len(newBlocks) > len(Blockchain){
    Blockchain = newBlocks
  }
}
