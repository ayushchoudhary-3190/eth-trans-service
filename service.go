 package transservice
 import(
 	"crypto/ecdsa"

  	"github.com/ethereum/go-ethereum/crypto"
    "github.com/go"
 )
 
func  CreateTransaction(req NewTransactionRequest) *Transaction{ 
	return &Transaction{
		SenderAddr: req.SenderWallet,
		ReceiverAddr: req.ReceiverWallet,
		AmountWei: req.Amount,
	}
}

func SignTransaction(t *Transaction, p privateKey, chainID int) *types.Transaction{
	
} 
