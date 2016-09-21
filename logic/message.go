package main

import (
	inet "imgo/libs/net"
	"imgo/libs/net/xrpc"
	"imgo/libs/proto"

	log "github.com/thinkboy/log4go"
)

const (
	messageService             = "MessageRPC"
	messageServicePing         = "MessageRPC.Ping"
	messageServiceGetPrivate   = "MessageRPC.GetPrivate"
	messageServiceSavePrivate  = "MessageRPC.SavePrivate"
	messageServiceSavePrivates = "MessageRPC.SavePrivates"
	messageServiceDelPrivate   = "MessageRPC.DelPrivate"
	messageServiceSaveToken    = "MessageRPC.SaveToken"
	messageServiceGetUid       = "MessageRPC.GetUid"
)

var (
	rpcClient *xrpc.Client
)

//example:tcp@localhost:8072
func InitMessage(bind string) (err error) {
	var (
		network, addr string
	)

	if network, addr, err = inet.ParseNetwork(bind); err != nil {
		log.Error("inet.ParseNetwork() error(%v)", err)
		return
	}

	option := xrpc.ClientOptions{
		Proto: network,
		Addr:  addr,
	}

	rpcClient = xrpc.Dial(option)

	go rpcClient.Ping(messageServicePing)
	return
}

func saveToken(t *proto.Token) error {
	ret := 0
	if err := rpcClient.Call(messageServiceSaveToken, t, &ret); err != nil {
		log.Error("c.Call(\"%s\",\"%v\") error(%v)", messageServiceSaveToken, t, err)
		return err
	}
	return nil
}

func getUid(token string) (uid int64, err error) {
	if err = rpcClient.Call(messageServiceGetUid, token, &uid); err != nil {
		log.Error("c.Call(\"%s\",\"%v\") error(%v)", messageServiceGetUid, token, err)
		return
	}
	return
}
