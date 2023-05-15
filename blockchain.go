/*	Code to Retrive data from CSV to private blocks in blockchain
*	and to process the most derserving candidate
*
 */
package main

import (
	"crypto/sha256"
	"encoding/csv"
	"encoding/json"
	"fmt"
	"math"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

type Block struct {
	data         map[string]interface{}
	hash         string
	previousHash string
	timestamp    time.Time
	pow          int
}

type Blockchain struct {
	genesisBlock Block
	chain        []Block
	difficulty   int
}

func (b Block) calculateHash() string {
	data, _ := json.Marshal(b.data)
	blockData := b.previousHash + string(data) + b.timestamp.String() + strconv.Itoa(b.pow)
	blockHash := sha256.Sum256([]byte(blockData))
	return fmt.Sprintf("%x", blockHash)
}

func (b *Block) mine(difficulty int) {
	for !strings.HasPrefix(b.hash, strings.Repeat("0", difficulty)) {
		b.pow++
		b.hash = b.calculateHash()
	}
}

func CreateBlockchain(difficulty int) Blockchain {
	genesisBlock := Block{
		hash:      "0",
		timestamp: time.Now(),
	}
	return Blockchain{
		genesisBlock,
		[]Block{genesisBlock},
		difficulty,
	}
}

func (b *Blockchain) addBlock(adhaar_num string, full_name string, ror_data string, bank_data string, income_tax string, land_area string, land_value string) {
	blockData := map[string]interface{}{
		"Adhaar Number":    adhaar_num,
		"Full Name":        full_name,
		"ROR Data":         ror_data,
		"Bank Details":     bank_data,
		"Income Tax":       income_tax,
		"Land in Acres":    land_area,
		"Total Land Value": land_value,
	}
	lastBlock := b.chain[len(b.chain)-1]
	newBlock := Block{
		data:         blockData,
		previousHash: lastBlock.hash,
		timestamp:    time.Now(),
	}
	newBlock.mine(b.difficulty)
	b.chain = append(b.chain, newBlock)
}

func (b Blockchain) isValid() bool {
	for i := range b.chain[1:] {
		previousBlock := b.chain[i]
		currentBlock := b.chain[i+1]
		// fmt.Println(previousBlock, currentBlock)
		if currentBlock.hash != currentBlock.calculateHash() || currentBlock.previousHash != previousBlock.hash {
			return false
		}
	}
	return true
}

// csv or excel files to import and read and store in variable
func csv_call() [][]string {
	// open CSV file
	fd, error := os.Open("data.csv")
	if error != nil {
		fmt.Println(error)
	}
	fmt.Println("Successfully opened the CSV file")
	defer fd.Close()

	// read CSV file
	fileReader := csv.NewReader(fd)
	records, error := fileReader.ReadAll()
	if error != nil {
		fmt.Println(error)
	}
	// fmt.Println(records[0])
	// dataset_extract(records)
	return records

}

var new map[string]interface{}
var records [][]string

func main() {
	records = csv_call()
	// fmt.Println(records)
	// create a new blockchain instance with a mining difficulty of 2
	blockchain := CreateBlockchain(2)

	// record transactions on the blockchain for Alice, Bob, and John
	for i := range records[0:] {
		blockchain.addBlock(records[i][0], records[i][1], records[i][2], records[i][3], records[i][4], records[i][5], records[i][6])
	}

	// check if the blockchain is valid; expecting true
	fmt.Println(blockchain.isValid())

	//extracting the data from blockchain
	for i := range blockchain.chain[0:] {
		fmt.Println("Time Stamp :", blockchain.chain[i].timestamp)
		fmt.Println("Previous Hash :", blockchain.chain[i].previousHash)
		fmt.Println("Current Hash :", blockchain.chain[i].hash)
		fmt.Println("Hashing Power :", blockchain.chain[i].pow)
		fmt.Println("User Data =>")
		new = blockchain.chain[i].data
		for j, k := range new {
			fmt.Println(j, ":", k)
			//if j == "income_tax" {
			//	s1, err := strconv.Atoi(k)
			//	it[i] = s1

		}
		fmt.Println()
	}
	fmt.Println("Start here?")
	//csv reading here
	//as in values in array picked directly from CSV
	var ia int = rand.Intn(4)
	//fmt.Println(ia)
	nm := [4]string{"Ronaldo", "Messi", "Neymar", "Saurez"}
	la := [4]int{56, 76, 46, 39}
	it := [4]int{500000, 250000, 177000, 999000}
	lv := [4]int{5000000, 9000000, 12000000, 3500000}
	var temp int = 100
	var hs, hs1, hs2, fs int
	//in income tax
	var a int = it[ia]
	//fmt.Println(a)
	for ; temp > 0; temp-- {
		var b int = rand.Intn(4)
		//fmt.Println(b)
		if it[b] < a {
			a = it[b]
			hs = b
		} else {
			var loss int = a - it[b]
			var prob float64 = (math.Exp)((float64(loss) / float64(temp)))
			if rand.Float64() <= prob {
				a = it[b]
			}
		}
	}
	//fmt.Println(it[hs])
	temp = 100
	//in land area
	var a1 int = la[ia]
	//fmt.Println(a1)
	for ; temp > 0; temp-- {
		var b int = rand.Intn(4)
		//fmt.Println(b)
		if la[b] < a1 {
			a1 = la[b]
			hs1 = b
		} else {
			var loss int = a1 - la[b]
			var prob float64 = (math.Exp)((float64(loss) / float64(temp)))
			if rand.Float64() <= prob {
				a1 = la[b]
			}
		}
	}
	//fmt.Println(la[hs1])
	temp = 100
	//in land value
	var a2 int = lv[ia]
	//fmt.Println(a2)
	for ; temp > 0; temp-- {
		var b int = rand.Intn(4)
		//fmt.Println(b)
		if lv[b] < a2 {
			a2 = lv[b]
			hs2 = b
		} else {
			var loss int = a2 - lv[b]
			var prob float64 = (math.Exp)((float64(loss) / float64(temp)))
			if rand.Float64() <= prob {
				a2 = lv[b]
			}
		}
	}
	//fmt.Println(lv[hs2])
	fs = 0
	if hs == hs1 {
		fs = hs
	} else if hs1 == hs2 {
		fs = hs1
	} else if hs2 == hs {
		fs = hs2
	} else if (it[hs] + lv[hs] + la[hs]) < (it[hs1] + lv[hs1] + la[hs1]) {
		fs = hs
	} else if (it[hs1] + lv[hs1] + la[hs1]) < (it[hs2] + lv[hs2] + la[hs2]) {
		fs = hs1
	} else if (it[hs2] + lv[hs2] + la[hs2]) < (it[hs] + lv[hs] + la[hs]) {
		fs = hs2
	} else if (it[hs] + lv[hs] + la[hs]) > (it[hs1] + lv[hs1] + la[hs1]) {
		fs = hs
	} else if (it[hs1] + lv[hs1] + la[hs1]) > (it[hs2] + lv[hs2] + la[hs2]) {
		fs = hs1
	} else if (it[hs2] + lv[hs2] + la[hs2]) > (it[hs] + lv[hs] + la[hs]) {
		fs = hs2
	} else if (it[hs] + lv[hs] + la[hs]) == (it[hs1] + lv[hs1] + la[hs1]) {
		fs = hs
	} else if (it[hs1] + lv[hs1] + la[hs1]) == (it[hs2] + lv[hs2] + la[hs2]) {
		fs = hs1
	} else if (it[hs2] + lv[hs2] + la[hs2]) == (it[hs] + lv[hs] + la[hs]) {
		fs = hs2
	}
	fmt.Println(nm[fs])
}
