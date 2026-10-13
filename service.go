package transservice
 
func  CreateTransaction(sender string, receiver string, amount string) *Transaction{ 
	return &Transaction{
		SenderAddr: sender,
		ReceiverAddr: receiver,
		AmountWei: amount,
	}
}