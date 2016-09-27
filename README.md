Imgo
==============
Imgo is a distributed and high performance push server written in golang based on [goim](https://github.com/Terry-Mao/goim).
compared to goim,it added offline message support,add will support IM server later on.


## Features
 * Light weight and high performance
 * Supports single push, multiple push and broadcasting
 * Supports one key to multiple subscribers (Configurable maximum subscribers count)
 * Supports authentication (Unauthenticated user can't subscribe)
 * Supports multiple protocols (WebSocket，TCP）
 * Supports offline message (you can push even if user is not online)
 * Scalable architecture (Unlimited dynamic comet,logic,router,job modules)
 * Asynchronous push notification based on Kafka

## Architecture
Client connect to server:

![arch](https://github.com/imroc/imgo/blob/master/doc/connect.gif)

A client wants to subscribe a channel on comet through tcp or websocket,comet tells logic: "Here comes a guy,shall I keep a connection with him ?".
Logic take the token from comet and showed it to store: "Is this token valid? If it is,tell me the user id".
Logic got the user id,told router this user is online and keeps a connection with that comet,and told comet:"yes, you shall".
Then comet keeps connect to that client,and matains a heartbeat with him.
Logic knowed that comet and client was keep a connection, he ask store:"Is there any offline message of that user ?","yes,three of it",store answered and gave these to logic.
Logic packed the message into an envelope and thrown to kafka.
Job found a new envelope in kafka,fetch it and read the address:"comet 1,user 123456",then he told comet 1:"tell this to user 123456".
At last,Comet told this message to user 123456.


Server push message to client:

![arch](https://github.com/imroc/imgo/blob/master/doc/push.gif)

Caller(usually a bussiness system) tells logic:"I want to send hello to a person,his user id is 123456".
Logic got the user id,ask router:"Is user 123456 online ?","yes,he is keeping a connection with comet 1" router replied.
Logic packed the message into an envelope and thrown to kafka.
Job found a new envelope in kafka,fetch it and read the address:"comet 1,user 123456",then he told comet 1:"tell this to user 123456".
At last,Comet told this message to user 123456.

Protocol:

[proto](https://github.com/imroc/imgo/blob/master/doc/protocol.png)

## Document
[中文](./README_cn.md)

## Examples
Websocket: [Websocket Client Demo](https://github.com/imroc/imgo/tree/master/examples/javascript)

Java: [Java](https://github.com/imroc/imgo-java-sdk)

## Benchmark
![benchmark](./doc/benchmark.jpg)

### Benchmark Server
| CPU | Memory | OS | Instance |
| :---- | :---- | :---- | :---- |
| Intel(R) Xeon(R) CPU E5-2630 v2 @ 2.60GHz  | DDR3 32GB | Debian GNU/Linux 8 | 1 |

### Benchmark Case
* Online: 1,000,000
* Duration: 15min
* Push Speed: 40/s (broadcast room)
* Push Message: {"test":1}
* Received calc mode: 1s per times, total 30 times

### Benchmark Resource
* CPU: 2000%~2300%
* Memory: 14GB
* GC Pause: 504ms
* Network: Incoming(450MBit/s), Outgoing(4.39GBit/s)

### Benchmark Result
* Received: 35,900,000/s

[中文](./doc/benchmark_cn.md)

[English](./doc/benchmark_en.md)

## LICENSE
imgo is is distributed under the terms of the MIT License.
