package main

import (
	inet "imgo/libs/net"
	"imgo/libs/net/xrpc"
	"imgo/libs/proto"
	"strconv"

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

func getOfflineMsg(uid int64) (msgs [][]byte, err error) {
	arg := &proto.MessageGetPrivateArgs{
		MsgId: 0,
		Key:   strconv.FormatInt(uid, 10),
	}
	res := &proto.MessageGetResp{}
	if err = rpcClient.Call(messageServiceGetPrivate, arg, res); err != nil {
		log.Error("c.Call(\"%s\",\"%v\") error(%v)", messageServiceGetPrivate, arg, err)
		return
	}
	if len(res.Msgs) == 0 {
		return
	}
	msgs = make([][]byte, 0, len(res.Msgs))

	for _, msg := range res.Msgs {
		msgs = append(msgs, msg.Msg)
	}
	return
}
