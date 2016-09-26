package main

import (
	"imgo/libs/proto"
	"net/http"
	"strconv"
	"time"

	log "code.google.com/p/log4go"
)

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

	if uidStr == "" || midStr == "" || token == "" {
		res["ret"] = ParamErr
		return
	}

	uid, err := strconv.ParseInt(uidStr, 10, 64)
	if err != nil {
		res["ret"] = ParamErr
		return
	}

	//check token
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

	var msgs []*proto.Message
	msgs, err = UseStorage.GetPrivate(uidStr, mid)
	if err != nil {
		res["ret"] = InternalErr
		log.Error("UseStorage.GetPrivate(\"%s\",%v) error(%v)", uidStr, mid, err)
		return
	}

	if len(msgs) == 0 {
		return
	}

	res["data"] = map[string]interface{}{"msgs": msgs}
	return
}
