package proto

// Message SavePrivate args
import "encoding/json"

type MessageSavePrivateArgs struct {
	Key    string // subscriber key
	Msg    []byte // message content
	MsgId  int64  // message id
	Expire uint   // message expire second
}

// Message SavePrivates args
type MessageSavePrivatesArgs struct {
	Keys   []string        // subscriber keys
	Msg    json.RawMessage // message content
	MsgId  int64           // message id
	Expire uint            // message expire second
}

// Message SavePrivates response
type MessageSavePrivatesResp struct {
	FKeys []string // failed key
}

// Message SavePublish args
type MessageSavePublishArgs struct {
	MsgID  int64  // message id
	Msg    string // message content
	Expire int64  // message expire second
}

//type MessageSaveTokenArgs struct {
//Key, Value string
//Expire     int64
//}
type Token struct {
	Token       string
	Uid, Expire int64
}

// Message Get args
type MessageGetPrivateArgs struct {
	MsgId int64  // message id
	Key   string // subscriber key
}

// Message Get Response
type MessageGetResp struct {
	Msgs []*Message // messages
}

// The Message struct
type Message struct {
	Msg     []byte `json:"msg"` // message content
	MsgId   int64  `json:"mid"` // message id
	GroupId uint   `json:"gid"` // group id
}

const (
	// group id
	PrivateGroupId = 0
	PublicGroupId  = 1
)
