package transservice

import (
	"context"
	"crypto/ecdsa"
	"math/big"

	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)
 
func  CreateTransaction(req NewTransactionRequest) *Transaction{ 
	return &Transaction{
		SenderAddr: req.SenderWallet,
		ReceiverAddr: req.ReceiverWallet,
		AmountWei: req.Amount,
	}
}

func SignTransaction(t *Transaction, p privateKey, chainID *big.Int) (*types.Transaction,error){
	tx := types.NewTx(&types.DynamicFeeTx{
		ChainID : chainID,
		GasFeeCap : gasFeeCap,
		GasTipCap : tipFeeCap,
		GasCap	  : gasLimit,
		Nonce 	  : nonce,
		To 		  : t.ReceiverAddr,
		Value	  : t.AmountWei,
	})

	signer := types.LatestSignerForChainID(chainID)
	signedTx,err := types.SignTx(tx,signer,p)
	if err != nil{
		return nil,err
	}
	return signedTx,nil
} 


func BroadcastTransaction(ctx context.Context,signedTx *types.Transaction,client *ethclient.Client) error{
	err := client.SendTransaction(ctx,signedTx)
	
	if err != nil{
		return err
	}

	return nil
}
