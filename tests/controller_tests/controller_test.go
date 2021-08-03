package controller_tests

import (
	"rest-api/controller"
	"rest-api/tests/controller_tests/mocks"
	"testing"
)

func TestBlockController_GetBlock(t *testing.T) {
	for i := range testCasesGetBlock {

		mockProxy := mocks.NewProxy(5, 5)
		mockEthClient := mocks.NewEthClient("", 5)
		controller := controller.NewBlockController(&mockEthClient, &mockProxy)

		block, err := controller.GetBlock(testCasesGetBlock[i].key)

		if err != nil {
			if testCasesGetBlock[i].expectedError == nil {
				t.Fatalf("Unexpected error occured: %v", err)
			}

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

func TestBlockController_GetTransaction(t *testing.T) {
	for i := range testCasesGetTransaction {
		mockProxy := mocks.NewProxy(5, 5)
		mockEthClient := mocks.NewEthClient("", 5)
		controller := controller.NewBlockController(&mockEthClient, &mockProxy)

		trx, err := controller.GetTransaction(testCasesGetTransaction[i].blockKey, testCasesGetTransaction[i].trxKey)

		if err != nil {
			if testCasesGetTransaction[i].expectedError == nil {
				t.Fatalf("Unexpected error occured: %v", err)
			}

			if err.Error() != testCasesGetTransaction[i].expectedError.Error() {
				t.Fatalf("FAIL Error: %s\nExpected: %#v\nActual: %#v\nTestcase: %#v", testCasesGetTransaction[i].description, testCasesGetTransaction[i].expectedError, err, i)
			}
		}

		if trx != nil && testCasesGetTransaction[i].expectedTransaction != nil {
			//BlockHash
			if trx.BlockHash != testCasesGetTransaction[i].expectedTransaction.BlockHash {
				t.Fatalf("FAIL BlockHash: %s\nExpected: %#v\nActual: %#v\nTestcase: %#v", testCasesGetTransaction[i].description, testCasesGetTransaction[i].expectedTransaction.BlockHash, trx.BlockHash, i)
			}

			//BlockNumber
			if trx.BlockNumber != testCasesGetTransaction[i].expectedTransaction.BlockNumber {
				t.Fatalf("FAIL BlockNumber: %s\nExpected: %#v\nActual: %#v\nTestcase: %#v", testCasesGetTransaction[i].description, testCasesGetTransaction[i].expectedTransaction.BlockNumber, trx.BlockNumber, i)
			}

			//From
			if trx.From != testCasesGetTransaction[i].expectedTransaction.From {
				t.Fatalf("FAIL From: %s\nExpected: %#v\nActual: %#v\nTestcase: %#v", testCasesGetTransaction[i].description, testCasesGetTransaction[i].expectedTransaction.From, trx.From, i)
			}

			//To
			if trx.To != testCasesGetTransaction[i].expectedTransaction.To {
				t.Fatalf("FAIL To: %s\nExpected: %#v\nActual: %#v\nTestcase: %#v", testCasesGetTransaction[i].description, testCasesGetTransaction[i].expectedTransaction.To, trx.To, i)
			}

			//Gas
			if trx.Gas != testCasesGetTransaction[i].expectedTransaction.Gas {
				t.Fatalf("FAIL Gas: %s\nExpected: %#v\nActual: %#v\nTestcase: %#v", testCasesGetTransaction[i].description, testCasesGetTransaction[i].expectedTransaction.Gas, trx.Gas, i)
			}

			//GasPrice
			if trx.GasPrice != testCasesGetTransaction[i].expectedTransaction.GasPrice {
				t.Fatalf("FAIL GasPrice: %s\nExpected: %#v\nActual: %#v\nTestcase: %#v", testCasesGetTransaction[i].description, testCasesGetTransaction[i].expectedTransaction.GasPrice, trx.GasPrice, i)
			}

			//Hash
			if trx.Hash != testCasesGetTransaction[i].expectedTransaction.Hash {
				t.Fatalf("FAIL Hash: %s\nExpected: %#v\nActual: %#v\nTestcase: %#v", testCasesGetTransaction[i].description, testCasesGetTransaction[i].expectedTransaction.Hash, trx.Hash, i)
			}

			//Input
			if trx.Input != testCasesGetTransaction[i].expectedTransaction.Input {
				t.Fatalf("FAIL Input: %s\nExpected: %#v\nActual: %#v\nTestcase: %#v", testCasesGetTransaction[i].description, testCasesGetTransaction[i].expectedTransaction.Input, trx.Input, i)
			}

			//Nonce
			if trx.Nonce != testCasesGetTransaction[i].expectedTransaction.Nonce {
				t.Fatalf("FAIL Nonce: %s\nExpected: %#v\nActual: %#v\nTestcase: %#v", testCasesGetTransaction[i].description, testCasesGetTransaction[i].expectedTransaction.Nonce, trx.Nonce, i)
			}

			//TransactionIndex
			if trx.TransactionIndex != testCasesGetTransaction[i].expectedTransaction.TransactionIndex {
				t.Fatalf("FAIL TransactionIndex: %s\nExpected: %#v\nActual: %#v\nTestcase: %#v", testCasesGetTransaction[i].description, testCasesGetTransaction[i].expectedTransaction.TransactionIndex, trx.TransactionIndex, i)
			}

			//Value
			if trx.Value != testCasesGetTransaction[i].expectedTransaction.Value {
				t.Fatalf("FAIL Value: %s\nExpected: %#v\nActual: %#v\nTestcase: %#v", testCasesGetTransaction[i].description, testCasesGetTransaction[i].expectedTransaction.Value, trx.Value, i)
			}

			//Type
			if trx.Type != testCasesGetTransaction[i].expectedTransaction.Type {
				t.Fatalf("FAIL Type: %s\nExpected: %#v\nActual: %#v\nTestcase: %#v", testCasesGetTransaction[i].description, testCasesGetTransaction[i].expectedTransaction.Type, trx.Type, i)
			}

			//V
			if trx.V != testCasesGetTransaction[i].expectedTransaction.V {
				t.Fatalf("FAIL V: %s\nExpected: %#v\nActual: %#v\nTestcase: %#v", testCasesGetTransaction[i].description, testCasesGetTransaction[i].expectedTransaction.V, trx.V, i)
			}

			//R
			if trx.R != testCasesGetTransaction[i].expectedTransaction.R {
				t.Fatalf("FAIL R: %s\nExpected: %#v\nActual: %#v\nTestcase: %#v", testCasesGetTransaction[i].description, testCasesGetTransaction[i].expectedTransaction.R, trx.R, i)
			}

			//S
			if trx.S != testCasesGetTransaction[i].expectedTransaction.S {
				t.Fatalf("FAIL S: %s\nExpected: %#v\nActual: %#v\nTestcase: %#v", testCasesGetTransaction[i].description, testCasesGetTransaction[i].expectedTransaction.S, trx.S, i)
			}
		}

		t.Logf("Pass: %s", testCasesGetTransaction[i].description)
	}
}
