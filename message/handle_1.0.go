package main

import (
	myrpc "gopush-cluster/rpc"
	"net/http"
	"strconv"
	"time"

	log "code.google.com/p/log4go"
)

// GetServer handle for server get
//func GetServer(w http.ResponseWriter, r *http.Request) {
//if r.Method != "GET" {
//http.Error(w, "Method Not Allowed", 405)
//return
//}
//params := r.URL.Query()
//key := params.Get("k")
//callback := params.Get("cb")
//protoStr := params.Get("p")
//res := map[string]interface{}{"ret": OK}
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

//// GetServer2 handle for server get.
//func GetServer2(w http.ResponseWriter, r *http.Request) {
//if r.Method != "GET" {
//http.Error(w, "Method Not Allowed", 405)
//return
//}
//params := r.URL.Query()
//key := params.Get("k")
//callback := params.Get("cb")
//protoStr := params.Get("p")
//res := map[string]interface{}{"ret": OK}
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
//// give client a addr list, client do the best choice
//res["data"] = map[string]interface{}{"server": addrs}
//return
//}

// GetOfflineMsg get offline mesage http handler.
func GetOfflineMsg(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, "Method Not Allowed", 405)
		return
	}
	params := r.URL.Query()
	uidStr := params.Get("uid")
	midStr := params.Get("mid")
	callback := params.Get("cb")
	token := params.Get("token")
	res := map[string]interface{}{"ret": OK}
	defer retWrite(w, r, res, callback, time.Now())

	log.Debug("离线消息参数,uid=\"%s\",mid=%s,cb=%s,token=%s\n", uidStr, midStr, callback, midStr)

	if uidStr == "" || midStr == "" || token == "" {
		res["ret"] = ParamErr
		return
	}

	uid, err := strconv.ParseInt(uidStr, 10, 64)
	if err != nil {
		log.Info("userid不是整数:%s", uidStr)
		res["ret"] = ParamErr
		return
	}

	//check token
	//token := params.Get("token")
	log.Debug("使用token:%s\n", token)
	ok, err := IsTokenValid(uid, token)
	if err != nil {
		res["ret"] = InternalErr
		return
	}
	if !ok {
		res["ret"] = TokenErr
		return
	}

	mid, err := strconv.ParseInt(midStr, 10, 64)
	if err != nil {
		res["ret"] = ParamErr
		log.Error("strconv.ParseInt(\"%s\", 10, 64) error(%v)", midStr, err)
		return
	}

	var msgs []*myrpc.Message
	log.Debug("准备获取离线消息\n")
	msgs, err = UseStorage.GetPrivate(uidStr, mid)
	if err != nil {
		res["ret"] = InternalErr
		log.Error("UseStorage.GetPrivate(\"%s\",%v) error(%v)", uidStr, mid, err)
		return
	}
	log.Debug("离线消息为:%v\n", msgs)

	if len(msgs) == 0 {
		return
	}

	// RPC get offline messages
	//reply := &myrpc.MessageGetResp{}
	//args := &myrpc.MessageGetPrivateArgs{MsgId: mid, Key: key}
	//client := myrpc.MessageRPC.Get()
	//if client == nil {
	//log.Error("no message node found")
	//res["ret"] = InternalErr
	//return
	//}
	//if err := client.Call(myrpc.MessageServiceGetPrivate, args, reply); err != nil {
	//log.Error("myrpc.MessageRPC.Call(\"%s\", \"%v\", reply) error(%v)", myrpc.MessageServiceGetPrivate, args, err)
	//res["ret"] = InternalErr
	//return
	//}
	//if len(reply.Msgs)
	//return
	//}
	//res["data"] = map[string]interface{}{"msgs": reply.Msgs}
	res["data"] = map[string]interface{}{"msgs": msgs}
	return
}

// GetTime get server time http handler.
//func GetTime(w http.ResponseWriter, r *http.Request) {
//if r.Method != "GET" {
//http.Error(w, "Method Not Allowed", 405)
//return
//}
//params := r.URL.Query()
//callback := params.Get("cb")
//res := map[string]interface{}{"ret": OK}
//now := time.Now()
//defer retWrite(w, r, res, callback, now)
//res["data"] = map[string]interface{}{"timeid": now.UnixNano() / 100}
//return
//}
