package main

import (
	"context"
	"log"

	"github.com/fox-one/mint-withdraw"
	jsoniter "github.com/json-iterator/go"
)

func main() {
	t, err := mint.ReadTransaction(transaction)
	if err != nil {
		panic(err)
	}

	ctx := context.Background()
	s := Store{}
	k := NewKey(0, signerAPIBases...)

	o, err := mint.WithdrawTransaction(ctx, t, k, s, Receiver, ReceiverExtra)
	if err != nil {
		panic(err)
	}

	{
		bts, _ := jsoniter.Marshal(o)
		log.Println(string(bts))
	}
}