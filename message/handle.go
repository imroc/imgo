package main

import (
	myrpc "gopush-cluster/rpc"
	"net/http"
	"strconv"
	"time"

	log "code.google.com/p/log4go"
)

//const (
//wsProto  = "1"
//tcpProto = "2"
//)

// getProtoAddr get specified protocol addresss.
//func getProtoAddr(node *myrpc.CometNodeInfo, p string) (addrs []string, ret int) {
//if p == wsProto {
//addrs = node.WsAddr
//} else if p == tcpProto {
//addrs = node.TcpAddr
//} else {
//ret = ParamErr
//return
//}
//if len(addrs) == 0 {
//ret = NotFoundServer
//return
//}
//ret = OK
//return
//}

// GetServer handle for server get
//func GetServer0(w http.ResponseWriter, r *http.Request) {
//if r.Method != "GET" {
//http.Error(w, "Method Not Allowed", 405)
//return
//}
//params := r.URL.Query()
//key := params.Get("key")
//callback := params.Get("callback")
//protoStr := params.Get("proto")
//res := map[string]interface{}{"ret": OK, "msg": "ok"}
//defer retWrite(w, r, res, callback, time.Now())
//if key == "" {
//res["ret"] = ParamErr
//return
//}
//// Match a push-server with the value computed through ketama algorithm
//node := myrpc.GetComet(key)
//if node == nil {
//res["ret"] = NotFoundServer
//return
//}
//addrs, ret := getProtoAddr(node, protoStr)
//if ret != OK {
//res["ret"] = ret
//return
//}
//res["data"] = map[string]interface{}{"server": addrs[0]}
//return
//}

// GetOfflineMsg get offline mesage http handler.
func GetOfflineMsg0(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, "Method Not Allowed", 405)
		return
	}
	params := r.URL.Query()
	key := params.Get("key")
	midStr := params.Get("mid")
	callback := params.Get("callback")

	res := map[string]interface{}{"ret": OK, "msg": "ok"}
	defer retWrite(w, r, res, callback, time.Now())

	//check token
	//if Conf.UseToken {
	//token := params.Get("token")
	//ok, err := IsTokenValid(key, token)
	//if err != nil {
	//res["ret"] = InternalErr
	//return
	//}
	//if !ok {
	//res["ret"] = TokenErr
	//return
	//}
	//}

	if key == "" || midStr == "" {
		res["ret"] = ParamErr
		return
	}
	mid, err := strconv.ParseInt(midStr, 10, 64)
	if err != nil {
		res["ret"] = ParamErr
		log.Error("strconv.ParseInt(\"%s\", 10, 64) error(%v)", midStr, err)
		return
	}
	// RPC get offline messages
	reply := &myrpc.MessageGetResp{}
	args := &myrpc.MessageGetPrivateArgs{MsgId: mid, Key: key}
	client := myrpc.MessageRPC.Get()
	if client == nil {
		res["ret"] = InternalErr
		return
	}
	if err := client.Call(myrpc.MessageServiceGetPrivate, args, reply); err != nil {
		log.Error("myrpc.MessageRPC.Call(\"%s\", \"%v\", reply) error(%v)", myrpc.MessageServiceGetPrivate, args, err)
		res["ret"] = InternalErr
		return
	}
	omsgs := []string{}
	opmsgs := []string{}
	for _, msg := range reply.Msgs {
		omsg, err := msg.OldBytes()
		if err != nil {
			res["ret"] = InternalErr
			return
		}
		omsgs = append(omsgs, string(omsg))
	}

	if len(omsgs) == 0 {
		return
	}

	res["data"] = map[string]interface{}{"msgs": omsgs, "pmsgs": opmsgs}
	return
}

// GetTime get server time http handler.
//func GetTime0(w http.ResponseWriter, r *http.Request) {
//if r.Method != "GET" {
//http.Error(w, "Method Not Allowed", 405)
//return
//}
//params := r.URL.Query()
//callback := params.Get("callback")
//res := map[string]interface{}{"ret": OK, "msg": "ok"}
//now := time.Now()
//defer retWrite(w, r, res, callback, now)
//res["data"] = map[string]interface{}{"timeid": now.UnixNano() / 100}
//return
//}
