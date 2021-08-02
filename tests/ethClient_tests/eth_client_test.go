package ethClient_tests

import (
	"rest-api/ethclient"
	"testing"
)

const url = "https://cloudflare-eth.com"
const timeout = 5

func TestEthClient_GetBllock(t *testing.T) {
	for i := range testCasesGetBlock {
		ethClient := ethclient.NewCloudflareEthGateway(url, timeout)

		block, err := ethClient.GetBlock(testCasesGetBlock[i].key)

		if err != nil {
			if testCasesGetBlock[i].expectedError == nil {
				t.Fatalf("Unexpected error occured: %v", err)
			}

			t.Fatal(err)
			if err.Error() != testCasesGetBlock[i].expectedError.Error() {
				t.Fatalf("FAIL Error: %s\nExpected: %#v\nActual: %#v\nTestcase: %#v", testCasesGetBlock[i].description, testCasesGetBlock[i].expectedError, err, i)
			}
		}

		if block != nil && testCasesGetBlock[i].expectedBlock != nil {

			//Difficulty
			if block.Difficulty != testCasesGetBlock[i].expectedBlock.Difficulty {
				t.Fatalf("FAIL Difficulty: %s\nExpected: %#v\nActual: %#v\nTestcase: %#v", testCasesGetBlock[i].description, testCasesGetBlock[i].expectedBlock.Difficulty, block.Difficulty, i)
			}

			//ExtraData
			if block.ExtraData != testCasesGetBlock[i].expectedBlock.ExtraData {
				t.Fatalf("FAIL ExtraData: %s\nExpected: %#v\nActual: %#v\nTestcase: %#v", testCasesGetBlock[i].description, testCasesGetBlock[i].expectedBlock.ExtraData, block.ExtraData, i)
			}

			//GasLimit
			if block.GasLimit != testCasesGetBlock[i].expectedBlock.GasLimit {
				t.Fatalf("FAIL GasLimit: %s\nExpected: %#v\nActual: %#v\nTestcase: %#v", testCasesGetBlock[i].description, testCasesGetBlock[i].expectedBlock.GasLimit, block.GasLimit, i)
			}

			//GasLimit
			if block.GasUsed != testCasesGetBlock[i].expectedBlock.GasUsed {
				t.Fatalf("FAIL GasUsed: %s\nExpected: %#v\nActual: %#v\nTestcase: %#v", testCasesGetBlock[i].description, testCasesGetBlock[i].expectedBlock.GasUsed, block.GasUsed, i)
			}

			//Hash
			if block.Hash != testCasesGetBlock[i].expectedBlock.Hash {
				t.Fatalf("FAIL Hash: %s\nExpected: %#v\nActual: %#v\nTestcase: %#v", testCasesGetBlock[i].description, testCasesGetBlock[i].expectedBlock.Hash, block.Hash, i)
			}

			//LogsBloom
			if block.LogsBloom != testCasesGetBlock[i].expectedBlock.LogsBloom {
				t.Fatalf("FAIL LogsBloom: %s\nExpected: %#v\nActual: %#v\nTestcase: %#v", testCasesGetBlock[i].description, testCasesGetBlock[i].expectedBlock.LogsBloom, block.LogsBloom, i)
			}

			//Miner
			if block.Miner != testCasesGetBlock[i].expectedBlock.Miner {
				t.Fatalf("FAIL Miner: %s\nExpected: %#v\nActual: %#v\nTestcase: %#v", testCasesGetBlock[i].description, testCasesGetBlock[i].expectedBlock.Miner, block.Miner, i)
			}

			//MixHash
			if block.MixHash != testCasesGetBlock[i].expectedBlock.MixHash {
				t.Fatalf("FAIL MixHash: %s\nExpected: %#v\nActual: %#v\nTestcase: %#v", testCasesGetBlock[i].description, testCasesGetBlock[i].expectedBlock.MixHash, block.MixHash, i)
			}

			//Nonce
			if block.Nonce != testCasesGetBlock[i].expectedBlock.Nonce {
				t.Fatalf("FAIL Nonce: %s\nExpected: %#v\nActual: %#v\nTestcase: %#v", testCasesGetBlock[i].description, testCasesGetBlock[i].expectedBlock.Nonce, block.Nonce, i)
			}

			//Number
			if block.Number != testCasesGetBlock[i].expectedBlock.Number {
				t.Fatalf("FAIL Number: %s\nExpected: %#v\nActual: %#v\nTestcase: %#v", testCasesGetBlock[i].description, testCasesGetBlock[i].expectedBlock.Number, block.Number, i)
			}

			//ParentHash
			if block.ParentHash != testCasesGetBlock[i].expectedBlock.ParentHash {
				t.Fatalf("FAIL ParentHash: %s\nExpected: %#v\nActual: %#v\nTestcase: %#v", testCasesGetBlock[i].description, testCasesGetBlock[i].expectedBlock.ParentHash, block.ParentHash, i)
			}

			//ReceiptsRoot
			if block.ReceiptsRoot != testCasesGetBlock[i].expectedBlock.ReceiptsRoot {
				t.Fatalf("FAIL ReceiptsRoot: %s\nExpected: %#v\nActual: %#v\nTestcase: %#v", testCasesGetBlock[i].description, testCasesGetBlock[i].expectedBlock.ReceiptsRoot, block.ReceiptsRoot, i)
			}

			//Sha3Uncles
			if block.Sha3Uncles != testCasesGetBlock[i].expectedBlock.Sha3Uncles {
				t.Fatalf("FAIL Sha3Uncles: %s\nExpected: %#v\nActual: %#v\nTestcase: %#v", testCasesGetBlock[i].description, testCasesGetBlock[i].expectedBlock.Sha3Uncles, block.Sha3Uncles, i)
			}

			//Size
			if block.Size != testCasesGetBlock[i].expectedBlock.Size {
				t.Fatalf("FAIL Size: %s\nExpected: %#v\nActual: %#v\nTestcase: %#v", testCasesGetBlock[i].description, testCasesGetBlock[i].expectedBlock.Size, block.Size, i)
			}

			//StateRoot
			if block.StateRoot != testCasesGetBlock[i].expectedBlock.StateRoot {
				t.Fatalf("FAIL StateRoot: %s\nExpected: %#v\nActual: %#v\nTestcase: %#v", testCasesGetBlock[i].description, testCasesGetBlock[i].expectedBlock.StateRoot, block.StateRoot, i)
			}

			//Timestamp
			if block.Timestamp != testCasesGetBlock[i].expectedBlock.Timestamp {
				t.Fatalf("FAIL Timestamp: %s\nExpected: %#v\nActual: %#v\nTestcase: %#v", testCasesGetBlock[i].description, testCasesGetBlock[i].expectedBlock.Timestamp, block.Timestamp, i)
			}

			//TotalDifficulty
			if block.TotalDifficulty != testCasesGetBlock[i].expectedBlock.TotalDifficulty {
				t.Fatalf("FAIL TotalDifficulty: %s\nExpected: %#v\nActual: %#v\nTestcase: %#v", testCasesGetBlock[i].description, testCasesGetBlock[i].expectedBlock.TotalDifficulty, block.TotalDifficulty, i)
			}

			//TotalDifficulty
			if block.TransactionsRoot != testCasesGetBlock[i].expectedBlock.TransactionsRoot {
				t.Fatalf("FAIL TransactionsRoot: %s\nExpected: %#v\nActual: %#v\nTestcase: %#v", testCasesGetBlock[i].description, testCasesGetBlock[i].expectedBlock.TransactionsRoot, block.TransactionsRoot, i)
			}

			//Uncles
			if len(block.Uncles) != len(testCasesGetBlock[i].expectedBlock.Uncles) {
				t.Fatalf("FAIL Uncles: %s\nExpected: %#v\nActual: %#v\nTestcase: %#v", testCasesGetBlock[i].description, testCasesGetBlock[i].expectedBlock.Uncles, block.Uncles, i)
			}

		}

		t.Logf("Pass: %s", testCasesGetBlock[i].description)
	}
}
