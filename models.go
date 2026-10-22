package transservice

import "github.com/ethereum/go-ethereum/common"

type Transaction struct{
	SenderAddr		string
	ReceiverAddr	string
	AmountWei 		string
}

type NewTransactionRequest struct{
	SenderWallet	*common.Address
	ReceiverWallet  *common.Address
	Amount 			string
}

