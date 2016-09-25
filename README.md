Imgo
==============
Imgo is a distributed and high performance push server written in golang based on [goim](https://github.com/Terry-Mao/goim).
compared to goim,it added offline message support,add will support IM server later on.

Note:imgo is in developing,not stable yet,do not use it in production until its' first release version come out.

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
Client connect to server will be like this:

![arch](https://github.com/imroc/imgo/blob/master/doc/connect.gif)

Push message will be like this:

![arch](https://github.com/imroc/imgo/blob/master/doc/push.gif)

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
