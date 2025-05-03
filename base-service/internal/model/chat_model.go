package model

type ChatType int

const (
	GetAllUser ChatType = iota
	Disconnect
	Message
)

var chatTypeNames = map[ChatType]string{
	GetAllUser: "GetAllUser",
	Disconnect: "Disconnect",
	Message:    "Message",
}

func (c ChatType) String() string {
	return chatTypeNames[c]
}

type ChatObjc struct {
	ChatClientID string   `json:"chatClientID"`
	Message      string   `json:"message"`
	ChatType     ChatType `json:"chatType"`
}

type User struct {
	Id     string `json:"id"`
	Name   string `json:"name"`
	Avatar string `json:"avatar"`
}
type ChatObjResponse struct {
	ChatType ChatType    `json:"chatType"`
	Data     interface{} `json:"data"`
}
