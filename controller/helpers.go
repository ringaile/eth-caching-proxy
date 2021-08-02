package controller

import "github.com/ethereum/go-ethereum/common/hexutil"

func checkIfBlockIsWithinLast20(searchKey string, latestKey string) bool {
	latestNumber, _ := hexutil.DecodeUint64(latestKey)

	blockNumber, _ := hexutil.DecodeUint64(searchKey)

	if (latestNumber - blockNumber) <= 20 {
		return true
	}
	return false
}
