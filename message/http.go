package main

import (
	"encoding/json"
	"fmt"
	"net"
	"net/http"
	"time"

	log "code.google.com/p/log4go"
)

// StartHTTP start listen http.
func StartHTTP() {
	// external
	httpServeMux := http.NewServeMux()
	// 2
	//httpServeMux.HandleFunc("/2/server/get", GetServer2)
	// 1.0
	//httpServeMux.HandleFunc("/1/server/get", GetServer)
	httpServeMux.HandleFunc("/1/msg/get", GetOfflineMsg)
	//httpServeMux.HandleFunc("/1/time/get", GetTime)
	// old
	//httpServeMux.HandleFunc("/server/get", GetServer0)
	httpServeMux.HandleFunc("/msg/get", GetOfflineMsg0)
	//httpServeMux.HandleFunc("/time/get", GetTime0)
	// internal
	//httpAdminServeMux := http.NewServeMux()
	// 1.0
	//httpAdminServeMux.HandleFunc("/1/admin/push/private", PushPrivate)
	//httpAdminServeMux.HandleFunc("/1/admin/push/mprivate", PushMultiPrivate)
	//httpAdminServeMux.HandleFunc("/1/admin/msg/del", DelPrivate)
	//httpAdminServeMux.HandleFunc("/1/admin/token/new", NewTokenPrivate)
	// old
	//httpAdminServeMux.HandleFunc("/admin/push", PushPrivate)
	//httpAdminServeMux.HandleFunc("/admin/msg/clean", DelPrivate)

	for _, bind := range Conf.HttpBind {
		log.Info("start http listen addr:\"%s\"", bind)
		go httpListen(httpServeMux, bind)
	}
	//for _, bind := range Conf.AdminBind {
	//log.Info("start admin http listen addr:\"%s\"", bind)
	//go httpListen(httpAdminServeMux, bind)
	//}
}

func httpListen(mux *http.ServeMux, bind string) {
	server := &http.Server{Handler: mux, ReadTimeout: Conf.HttpServerTimeout, WriteTimeout: Conf.HttpServerTimeout}
	server.SetKeepAlivesEnabled(false)
	l, err := net.Listen("tcp", bind)
	if err != nil {
		log.Error("net.Listen(\"tcp\", \"%s\") error(%v)", bind, err)
		panic(err)
	}
	if err := server.Serve(l); err != nil {
		log.Error("server.Serve() error(%v)", err)
		panic(err)
	}
}

// retWrite marshal the result and write to client(get).
func retWrite(w http.ResponseWriter, r *http.Request, res map[string]interface{}, callback string, start time.Time) {
	data, err := json.Marshal(res)
	if err != nil {
		log.Error("json.Marshal(\"%v\") error(%v)", res, err)
		return
	}
	dataStr := ""
	if callback == "" {
		// Normal json
		dataStr = string(data)
	} else {
		// Jsonp
		dataStr = fmt.Sprintf("%s(%s)", callback, string(data))
	}
	if n, err := w.Write([]byte(dataStr)); err != nil {
		log.Error("w.Write(\"%s\") error(%v)", dataStr, err)
	} else {
		log.Debug("w.Write(\"%s\") write %d bytes", dataStr, n)
	}
	log.Info("req: \"%s\", res:\"%s\", ip:\"%s\", time:\"%fs\"", r.URL.String(), dataStr, r.RemoteAddr, time.Now().Sub(start).Seconds())
}

// retPWrite marshal the result and write to client(post).
//func retPWrite(w http.ResponseWriter, r *http.Request, res map[string]interface{}, body *string, start time.Time) {
//data, err := json.Marshal(res)
//if err != nil {
//log.Error("json.Marshal(\"%v\") error(%v)", res, err)
//return
//}
//dataStr := string(data)
//if n, err := w.Write([]byte(dataStr)); err != nil {
//log.Error("w.Write(\"%s\") error(%v)", dataStr, err)
//} else {
//log.Debug("w.Write(\"%s\") write %d bytes", dataStr, n)
//}
//log.Info("req: \"%s\", post: \"%s\", res:\"%s\", ip:\"%s\", time:\"%fs\"", r.URL.String(), *body, dataStr, r.RemoteAddr, time.Now().Sub(start).Seconds())
//}
