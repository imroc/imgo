// Copyright © 2014 Terry Mao, LiuDing All rights reserved.
// This file is part of gopush-cluster.

// gopush-cluster is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.

// gopush-cluster is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU General Public License for more details.

// You should have received a copy of the GNU General Public License
// along with gopush-cluster.  If not, see <http://www.gnu.org/licenses/>.

package main

import (
	"errors"
	"imgo/libs/proto"
	"net"
	"net/rpc"

	log "code.google.com/p/log4go"
)

// RPC For receive offline messages
type MessageRPC struct {
}

// InitRPC start accept rpc call.
func InitRPC() error {
	msg := &MessageRPC{}
	rpc.Register(msg)
	for _, bind := range Conf.RPCBind {
		log.Info("start rpc listen addr: \"%s\"", bind)
		go rpcListen(bind)
	}
	return nil
}

func rpcListen(bind string) {
	l, err := net.Listen("tcp", bind)
	if err != nil {
		log.Error("net.Listen(\"tcp\", \"%s\") error(%v)", bind, err)
		panic(err)
	}
	defer func() {
		if err := l.Close(); err != nil {
			log.Error("listener.Close() error(%v)", err)
		}
	}()
	rpc.Accept(l)
}

// SavePrivate rpc interface save user private message.
func (r *MessageRPC) SavePrivate(m *proto.MessageSavePrivateArgs, ret *int) error {
	if m == nil || m.Msg == nil || m.MsgId < 0 {
		return errors.New("parameter error")
	}
	if err := UseStorage.SavePrivate(m.Key, m.Msg, m.MsgId, m.Expire); err != nil {
		log.Error("UseStorage.SavePrivate(\"%s\", \"%s\", %d, %d) error(%v)", m.Key, string(m.Msg), m.MsgId, m.Expire, err)
		return err
	}
	log.Debug("UseStorage.SavePrivate(\"%s\", \"%s\", %d, %d) ok", m.Key, string(m.Msg), m.MsgId, m.Expire)
	return nil
}

// SavePrivates rpc interface save user private messages.
func (r *MessageRPC) SavePrivates(m *proto.MessageSavePrivatesArgs, rw *proto.MessageSavePrivatesResp) error {
	if m == nil || m.Msg == nil || m.MsgId < 0 {
		return errors.New("parameter error")
	}
	fkeys, err := UseStorage.SavePrivates(m.Keys, m.Msg, m.MsgId, m.Expire)
	if err != nil {
		log.Error("UseStorage.SavePrivates(\"%v\", \"%s\", %d, %d) error(%v)", m.Keys, string(m.Msg), m.MsgId, m.Expire, err)
	}
	rw.FKeys = fkeys
	log.Debug("UseStorage.SavePrivates(\"%v\", \"%s\", %d, %d) ok fkeys len(%d)", m.Keys, string(m.Msg), m.MsgId, m.Expire, len(fkeys))
	return nil
}

//// GetPrivate rpc interface get user private message.
//func (r *MessageRPC) GetPrivate(m *proto.MessageGetPrivateArgs, rw *proto.MessageGetResp) error {
//log.Debug("messageRPC.GetPrivate key:\"%s\" mid:\"%d\"", m.Key, m.MsgId)
//if m == nil || m.Key == "" || m.MsgId < 0 {
//return proto.ErrParam
//}
//msgs, err := UseStorage.GetPrivate(m.Key, m.MsgId)
//if err != nil {
//log.Error("UseStorage.GetPrivate(\"%s\", %d) error(%v)", m.Key, m.MsgId, err)
//return err
//}
//rw.Msgs = msgs
//log.Debug("UserStorage.GetPrivate(\"%s\", %d) ok", m.Key, m.MsgId)
//return nil
//}

//// DelPrivate rpc interface delete user private message.
//func (r *MessageRPC) DelPrivate(key string, ret *int) error {
//if key == "" {
//return proto.ErrParam
//}
//if err := UseStorage.DelPrivate(key); err != nil {
//log.Error("UserStorage.DelPrivate(\"%s\") error(%v)", key, err)
//return err
//}
//log.Debug("UserStorage.DelPrivate(\"%s\") ok", key)
//return nil
//}

/*
// SavePublish rpc interface save publish message.
func (r *MessageRPC) SavePublish(m *proto.MessageSaveGroupArgs, ret *int) error {
	return nil
}

// GetPublish rpc interface get publish message.
func (r *MessageRPC) GetPublish(m *proto.MessageGetGroupArgs, rw *proto.MessageGetResp) error {
	return nil
}

// DelPublish rpc interface delete publish message.
func (r *MessageRPC) DelPublish(key string, ret *int) error {
	return nil
}

// SaveGroup rpc interface save publish message.
func (r *MessageRPC) SaveGroup(m *proto.MessageSaveGroupArgs, ret *int) error {
	return nil
}

// GetPublish rpc interface get publish message.
func (r *MessageRPC) GetGroup(m *proto.MessageGetGroupArgs, rw *proto.MessageGetResp) error {
	return nil
}

// DelPublish rpc interface delete publish message.
func (r *MessageRPC) DelGroup(key string, ret *int) error {
	return nil
}
*/

// Server Ping interface
func (r *MessageRPC) Ping(arg *proto.NoArg, reply *proto.NoReply) error {
	//log.Debug("ping ok")
	return nil
}

func (r *MessageRPC) SaveToken(t *proto.Token, ret *int) error {
	return UseStorage.SaveToken(t.Uid, t.Token, t.Expire)
}

func (r *MessageRPC) GetUid(token string, uid *int64) (err error) {
	*uid, err = UseStorage.GetUid(token)
	log.Debug("MessageRPC.GetUid(%s)=%d", token, *uid)
	return
}
