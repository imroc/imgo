package main

import (
	"flag"
	"imgo/libs/perf"
	"runtime"

	log "github.com/thinkboy/log4go"
)

func main() {
	flag.Parse()
	if err := InitConfig(); err != nil {
		panic(err)
	}
	runtime.GOMAXPROCS(Conf.MaxProc)
	log.LoadConfiguration(Conf.Log)
	defer log.Close()
	log.Info("logic[%s] start", Ver)
	perf.Init(Conf.PprofAddrs)
	// router rpc
	if err := InitRouter(Conf.RouterRPCAddrs); err != nil {
		log.Warn("router rpc current can't connect, retry")
	}
	MergeCount()
	go SyncCount()
	// logic rpc
	if err := InitRPC(NewDefaultAuther()); err != nil {
		panic(err)
	}
	if err := InitKafka(Conf.KafkaAddrs); err != nil {
		panic(err)
	}
	log.Debug("init message")
	if err := InitMessage(Conf.MessageAddr); err != nil {
		panic(err)
	}
	log.Info("ready http...")
	if err := InitHTTP(); err != nil {
		log.Error("damn it,http failed to start...")
		panic(err)
	}
	// block until a signal is received.
	InitSignal()
}
