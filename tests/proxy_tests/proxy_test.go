package proxy_tests

import (
	"rest-api/proxy"
	"testing"
)

const defaultExpiration = 5
const cleanupExpiration = 5

func TestProxy_GetCache(t *testing.T) {
	for i := range testCasesGetCache {
		proxy := proxy.New(defaultExpiration, cleanupExpiration)

		if testCasesGetCache[i].isBlockSaved {
			proxy.SetCache(testCasesGetCache[i].key, testCasesGetCache[i].expectedBlock)
		}
		block, found := proxy.GetCache(testCasesGetCache[i].key)

		//Found
		if found != testCasesGetCache[i].found {
			t.Fatalf("FAIL Found: %s\nExpected: %#v\nActual: %#v\nTestcase: %#v", testCasesGetCache[i].description, testCasesGetCache[i].found, found, i)
		}

		if block != nil && testCasesGetCache[i].expectedBlock != nil {

			//Difficulty
			if block.Difficulty != testCasesGetCache[i].expectedBlock.Difficulty {
				t.Fatalf("FAIL Difficulty: %s\nExpected: %#v\nActual: %#v\nTestcase: %#v", testCasesGetCache[i].description, testCasesGetCache[i].expectedBlock.Difficulty, block.Difficulty, i)
			}

			//ExtraData
			if block.ExtraData != testCasesGetCache[i].expectedBlock.ExtraData {
				t.Fatalf("FAIL ExtraData: %s\nExpected: %#v\nActual: %#v\nTestcase: %#v", testCasesGetCache[i].description, testCasesGetCache[i].expectedBlock.ExtraData, block.ExtraData, i)
			}

			//GasLimit
			if block.GasLimit != testCasesGetCache[i].expectedBlock.GasLimit {
				t.Fatalf("FAIL GasLimit: %s\nExpected: %#v\nActual: %#v\nTestcase: %#v", testCasesGetCache[i].description, testCasesGetCache[i].expectedBlock.GasLimit, block.GasLimit, i)
			}

			//GasLimit
			if block.GasUsed != testCasesGetCache[i].expectedBlock.GasUsed {
				t.Fatalf("FAIL GasUsed: %s\nExpected: %#v\nActual: %#v\nTestcase: %#v", testCasesGetCache[i].description, testCasesGetCache[i].expectedBlock.GasUsed, block.GasUsed, i)
			}

			//Hash
			if block.Hash != testCasesGetCache[i].expectedBlock.Hash {
				t.Fatalf("FAIL Hash: %s\nExpected: %#v\nActual: %#v\nTestcase: %#v", testCasesGetCache[i].description, testCasesGetCache[i].expectedBlock.Hash, block.Hash, i)
			}

			//LogsBloom
			if block.LogsBloom != testCasesGetCache[i].expectedBlock.LogsBloom {
				t.Fatalf("FAIL LogsBloom: %s\nExpected: %#v\nActual: %#v\nTestcase: %#v", testCasesGetCache[i].description, testCasesGetCache[i].expectedBlock.LogsBloom, block.LogsBloom, i)
			}

			//Miner
			if block.Miner != testCasesGetCache[i].expectedBlock.Miner {
				t.Fatalf("FAIL Miner: %s\nExpected: %#v\nActual: %#v\nTestcase: %#v", testCasesGetCache[i].description, testCasesGetCache[i].expectedBlock.Miner, block.Miner, i)
			}

			//MixHash
			if block.MixHash != testCasesGetCache[i].expectedBlock.MixHash {
				t.Fatalf("FAIL MixHash: %s\nExpected: %#v\nActual: %#v\nTestcase: %#v", testCasesGetCache[i].description, testCasesGetCache[i].expectedBlock.MixHash, block.MixHash, i)
			}

			//Nonce
			if block.Nonce != testCasesGetCache[i].expectedBlock.Nonce {
				t.Fatalf("FAIL Nonce: %s\nExpected: %#v\nActual: %#v\nTestcase: %#v", testCasesGetCache[i].description, testCasesGetCache[i].expectedBlock.Nonce, block.Nonce, i)
			}

			//Number
			if block.Number != testCasesGetCache[i].expectedBlock.Number {
				t.Fatalf("FAIL Number: %s\nExpected: %#v\nActual: %#v\nTestcase: %#v", testCasesGetCache[i].description, testCasesGetCache[i].expectedBlock.Number, block.Number, i)
			}

			//ParentHash
			if block.ParentHash != testCasesGetCache[i].expectedBlock.ParentHash {
				t.Fatalf("FAIL ParentHash: %s\nExpected: %#v\nActual: %#v\nTestcase: %#v", testCasesGetCache[i].description, testCasesGetCache[i].expectedBlock.ParentHash, block.ParentHash, i)
			}

			//ReceiptsRoot
			if block.ReceiptsRoot != testCasesGetCache[i].expectedBlock.ReceiptsRoot {
				t.Fatalf("FAIL ReceiptsRoot: %s\nExpected: %#v\nActual: %#v\nTestcase: %#v", testCasesGetCache[i].description, testCasesGetCache[i].expectedBlock.ReceiptsRoot, block.ReceiptsRoot, i)
			}

			//Sha3Uncles
			if block.Sha3Uncles != testCasesGetCache[i].expectedBlock.Sha3Uncles {
				t.Fatalf("FAIL Sha3Uncles: %s\nExpected: %#v\nActual: %#v\nTestcase: %#v", testCasesGetCache[i].description, testCasesGetCache[i].expectedBlock.Sha3Uncles, block.Sha3Uncles, i)
			}

			//Size
			if block.Size != testCasesGetCache[i].expectedBlock.Size {
				t.Fatalf("FAIL Size: %s\nExpected: %#v\nActual: %#v\nTestcase: %#v", testCasesGetCache[i].description, testCasesGetCache[i].expectedBlock.Size, block.Size, i)
			}

			//StateRoot
			if block.StateRoot != testCasesGetCache[i].expectedBlock.StateRoot {
				t.Fatalf("FAIL StateRoot: %s\nExpected: %#v\nActual: %#v\nTestcase: %#v", testCasesGetCache[i].description, testCasesGetCache[i].expectedBlock.StateRoot, block.StateRoot, i)
			}

			//Timestamp
			if block.Timestamp != testCasesGetCache[i].expectedBlock.Timestamp {
				t.Fatalf("FAIL Timestamp: %s\nExpected: %#v\nActual: %#v\nTestcase: %#v", testCasesGetCache[i].description, testCasesGetCache[i].expectedBlock.Timestamp, block.Timestamp, i)
			}

			//TotalDifficulty
			if block.TotalDifficulty != testCasesGetCache[i].expectedBlock.TotalDifficulty {
				t.Fatalf("FAIL TotalDifficulty: %s\nExpected: %#v\nActual: %#v\nTestcase: %#v", testCasesGetCache[i].description, testCasesGetCache[i].expectedBlock.TotalDifficulty, block.TotalDifficulty, i)
			}

			//TotalDifficulty
			if block.TransactionsRoot != testCasesGetCache[i].expectedBlock.TransactionsRoot {
				t.Fatalf("FAIL TransactionsRoot: %s\nExpected: %#v\nActual: %#v\nTestcase: %#v", testCasesGetCache[i].description, testCasesGetCache[i].expectedBlock.TransactionsRoot, block.TransactionsRoot, i)
			}

			//Uncles
			if len(block.Uncles) != len(testCasesGetCache[i].expectedBlock.Uncles) {
				t.Fatalf("FAIL Uncles: %s\nExpected: %#v\nActual: %#v\nTestcase: %#v", testCasesGetCache[i].description, testCasesGetCache[i].expectedBlock.Uncles, block.Uncles, i)
			}

		}

		t.Logf("Pass: %s", testCasesGetCache[i].description)
	}
}
