package controller_tests

import (
	"fmt"
	"rest-api/model"
)

var testCasesGetBlock = []struct {
	description   string
	key           string
	expectedBlock *model.Block
	expectedError error
}{
	// #0
	{
		description: "[Successfully received a block.]",
		key:         "0xc58c82",
		expectedBlock: &model.Block{
			Difficulty:       "0x19f7ee142d8fe6",
			ExtraData:        "0xd883010a06846765746888676f312e31362e36856c696e7578",
			GasLimit:         "0xe4e1c0",
			GasUsed:          "0xe492cf",
			Hash:             "0x0ae1d8201dc947c228ee793854ed2013d3a1142bb175890cef052b06dac6c165",
			LogsBloom:        "0x98e3430383a1308304a2ed6aa0531a57c5097e505c08a0ea1d43a07abbcbb1230d3c336a10201568fd4455d22f093957cea6fb27da11b5e522485ae62a36e908f478b2a03c7f1c7bcdb4335f178a56e156660510cc761422695c90e9c86bc2895fb02a659ade49440da81a5841cc9d9380d8747520eda5724383c79628cccad2c953e161370c3f7681fe090cb197db5d6052325503868c48af79ea4cbad03e34abf1c03f07cc286e34e82883d007f9800807480caa4661e389f80e42e8c3bc4d2c9992a6fc4cc9eb0ccb886147921dedeb431a0f167f8496371d71a666a6e6912f19641bd90f7c48a295242f1884363407c154d0937a90c529ae9a328c804225",
			Miner:            "0xe206e3dca498258f1b7eec1c640b5aee7bb88fd0",
			MixHash:          "0x4e6d0db8b7fc0e84824ccd050ab0b8944b3f2c6146f2582fac305f45eb7a1be5",
			Nonce:            "0x72d3e69f3d8f7144",
			Number:           "0xc58c82",
			ParentHash:       "0x7b47c400189632d9ffc264f5041098df1340c91ef44a80f368f1d53fe2235453",
			ReceiptsRoot:     "0x870a7e2b73147e27f67c1264470120139a16eb72d10dea52953688d2884b8247",
			Sha3Uncles:       "0x7088b239e79654ee4aa177f859324dac6c276a42efb75280115ac397978786ea",
			Size:             "0xf11d",
			StateRoot:        "0x7a9cc96255fa66cf7722f1fbe25ce8e96182a3629622c1c9eae0a35be5ea82e9",
			Timestamp:        "0x61080628",
			TotalDifficulty:  "0x601118734b1f13d176a",
			TransactionsRoot: "0xde9793fbcde0d3ba6d03a07d506102413469a2ada20689f5b08b27eebcf2ad50",
			Uncles:           []string{"0x3db4fb3f2474daa42bf21ef83558465b393faef878ae5c365b12dbb08a96e1ad"},
		},
		expectedError: nil,
	},

	// #1
	{
		description:   "[No block found.]",
		key:           "12345",
		expectedBlock: nil,
		expectedError: fmt.Errorf("Error: invalid argument 0: hex string without 0x prefix"),
	},
}

var testCasesGetTransaction = []struct {
	description         string
	blockKey            string
	trxKey              string
	expectedTransaction *model.Transaction
	expectedError       error
}{
	// #0
	{
		description: "[Successfully received a transaction.]",
		blockKey:    "0xc59e6a",
		trxKey:      "0",
		expectedTransaction: &model.Transaction{
			BlockHash:        "0xe9d4061bd545a6006b28bf906edeee714820cb3094dca1ab3eeeb6f3eef1b830",
			BlockNumber:      "0xc59e6a",
			From:             "0xec3c1059e889703fa809cf34beb8588ef1c0279c",
			Gas:              "0x3a629",
			GasPrice:         "0x7ea8ed400",
			Hash:             "0x57dce8d244fc9996e74270009c9ce2193057c26f61f7010ff2b8be6869c24546",
			Input:            "0x38ed1739000000000000000000000000000000000000000000001303334ceffb07e00000000000000000000000000000000000000000000000000000000000041ca2a05800000000000000000000000000000000000000000000000000000000000000a0000000000000000000000000ec3c1059e889703fa809cf34beb8588ef1c0279c000000000000000000000000000000000000000000000000000000006108ffc50000000000000000000000000000000000000000000000000000000000000003000000000000000000000000d084b83c305dafd76ae3e1b4e1f1fe2ecccb3988000000000000000000000000c02aaa39b223fe8d0a0e5c4f27ead9083c756cc2000000000000000000000000dac17f958d2ee523a2206206994597c13d831ec7",
			Nonce:            "0x14",
			To:               "0x7a250d5630b4cf539739df2c5dacb4c659f2488d",
			TransactionIndex: "0x1",
			Value:            "0x0",
			Type:             "0x0",
			V:                "0x25",
			R:                "0x4f953906b4751f05a75995e9ce3500072df1a26dfecd1c19cd70b31a098813d9",
			S:                "0x437c5d8bec7f2068946dd95dcb48fe93716aa3fb79da9f3bac2c59c3a112c8bb",
		},
		expectedError: nil,
	},
	// #1
	{
		description:         "[No transaction found.]",
		blockKey:            "12345",
		trxKey:              "0",
		expectedTransaction: nil,
		expectedError:       fmt.Errorf("Error: invalid argument 0: hex string without 0x prefix"),
	},
}
