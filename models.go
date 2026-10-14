package transservice

type Transaction struct{
	SenderAddr		string
	ReceiverAddr	string
	AmountWei 		string
}

type NewTransactionRequest struct{
	SenderWallet	string
	ReceiverWallet  string
	Amount 			string
}

