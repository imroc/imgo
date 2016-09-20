package main

import "imgo/libs/define"

// developer could implement "Auth" interface for decide how get userId, or roomId
type Auther interface {
	Auth(token string) (AuthRes, error)
}

type AuthRes struct {
	Uid    int64
	RoomId int32
}

type DefaultAuther struct {
}

func NewDefaultAuther() *DefaultAuther {
	return &DefaultAuther{}
}

func (a *DefaultAuther) Auth(token string) (res AuthRes, err error) {
	res.RoomId = define.NoRoom
	res.Uid, err = getUid(token)
	return
}
