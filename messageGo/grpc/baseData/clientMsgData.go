package base_data

// 对应 proto.Vector3
type MultiClientMsgData struct {
	UserList []int64 `json:"userList"` //
	MsgId    int32   `json:"msgId"`
	MsgBody  []byte  `json:"msgBody"` // proto message marshal bytes
}
