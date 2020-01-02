package toolshell

import (
	"encoding/hex"
	"fmt"
	"github.com/hacash/cmdwallet/toolshell/ctx"
	"github.com/hacash/core/transactions"
)

func putTx(ctx ctx.Context, params []string) {
	if len(params) < 1 {
		fmt.Println("params not enough")
		return
	}
	txbody, err := hex.DecodeString(params[0])
	if err != nil {
		fmt.Println(err)
		return
	}
	// 解析交易
	newTrs, _, err2 := transactions.ParseTransaction(txbody, 0)
	if err2 != nil {
		fmt.Println(err2)
		return
	}

	// 交易加入

	// ok
	ctx.Println("transaction append success! ")
	ctx.Println("hash: <" + hex.EncodeToString(newTrs.Hash()) + ">, hash_with_fee: <" + hex.EncodeToString(newTrs.HashWithFee()) + ">")

	// 记录
	ctx.SetTxToRecord(newTrs.Hash(), newTrs)

}
