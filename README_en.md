imgo
==============
`imroc/imgo` is a IM and push notification server cluster.

---------------------------------------
  * [Features](#features)
  * [Installing](#installing)
  * [Configurations](#configurations)
  * [Examples](#examples)
  * [Documents](#documents)
  * [More](#more)

---------------------------------------

## Features
 * Light weight
 * High performance
 * Pure Golang
 * Supports single push, multiple push, room push and broadcasting
 * Supports offline message
 * Supports one key to multiple subscribers (Configurable maximum subscribers count)
 * Supports heartbeats (Application heartbeats, TCP, KeepAlive)
 * Supports authentication (Unauthenticated user can't subscribe)
 * Supports multiple protocols (WebSocket，TCP）
 * Scalable architecture (Unlimited dynamic job and logic modules)
 * Asynchronous push notification based on Kafka

## Installing
### Dependencies
```sh
$ yum -y install java-1.7.0-openjdk
```

### Install Kafka

Please follow the official quick start [here](http://kafka.apache.org/documentation.html#quickstart).

### Install Golang environment

Please follow the official quick start [here](https://golang.org/doc/install).

### Deploy imgo
1.Download imgo
```sh
$ yum install git
$ cd $GOPATH/src
$ git clone https://github.com/imroc/imgo.git
$ cd $GOPATH/src/imgo
$ go get ./...
```

2.Install router、logic、comet、job modules(You might need to change the configuration files based on your servers)
```sh
$ cd $GOPATH/src/imgo/router
$ go install
$ cp router-example.conf $GOPATH/bin/router.conf
$ cp router-log.xml $GOPATH/bin/
$ cd ../message/
$ go install
$ cp message.conf $GOPATH/bin/message.conf
$ cp message-log.xml $GOPATH/bin/
$ cd ../logic/
$ go install
$ cp logic-example.conf $GOPATH/bin/logic.conf
$ cp logic-log.xml $GOPATH/bin/
$ cd ../comet/
$ go install
$ cp comet-example.conf $GOPATH/bin/comet.conf
$ cp comet-log.xml $GOPATH/bin/
$ cd ../logic/job/
$ go install
$ cp job-example.conf $GOPATH/bin/job.conf
$ cp job-log.xml $GOPATH/bin/
```

Everything is DONE!

### Run imgo
You may need to change the log files location.
```sh
$ cd /$GOPATH/bin
$ nohup $GOPATH/bin/message -c $GOPATH/bin/message.conf 2>&1 > /data/logs/imgo/panic-message.log &
$ nohup $GOPATH/bin/router -c $GOPATH/bin/router.conf 2>&1 > /data/logs/imgo/panic-router.log &
$ nohup $GOPATH/bin/logic -c $GOPATH/bin/logic.conf 2>&1 > /data/logs/imgo/panic-logic.log &
$ nohup $GOPATH/bin/comet -c $GOPATH/bin/comet.conf 2>&1 > /data/logs/imgo/panic-comet.log &
$ nohup $GOPATH/bin/job -c $GOPATH/bin/job.conf 2>&1 > /data/logs/imgo/panic-job.log &
```

If it fails, please check the logs for debugging.

### Testing

Check the push protocols here[push HTTP protocols](./doc/push.md)

## Configurations
TODO

## Examples
Websocket: [Websocket Client Demo](https://github.com/imroc/imgo/tree/master/examples/javascript)

Android: [Android SDK](https://github.com/roamdy/imgo-sdk)

iOS: [iOS](https://github.com/roamdy/imgo-oc-sdk)

## Documents
[push HTTP protocols](./doc/en/push.md)

[Comet client protocols](./doc/en/proto.md)

##More
TODO
