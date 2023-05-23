package protoTool

import (
	"game-message-core/proto"
	"testing"
)

func makeSingInMsgBody() ([]byte, error) {
	msg := &proto.SigninPlayerReq{
		ReqTitle:          &proto.ReqTitle{SeqId: 99},
		UserId:            7001,
		Token:             "vvvvvvvvvvvbbbbbbbbbbbbbbbb",
		SceneServiceAppId: "game-service-world",
	}

	return MarshalProto(msg)
}
func UnmarshalSingInMsg(body []byte, t *testing.T) {
	msg := &proto.SigninPlayerReq{}
	err := UnmarshalProto(body, msg)
	t.Log(err)
	t.Log(msg)
}

func Test_NetDataCode(t *testing.T) {
	eType := proto.EnvelopeType_SigninPlayer
	data, err := makeSingInMsgBody()
	t.Log(data)
	t.Log(err)

	println()
	println()

	body := EnCodeNetData(eType, data)
	t.Log(body)

	println()
	println()

	rType, bs, err := DeCodeNetData(body)
	t.Log(rType)
	t.Log(bs)
	t.Log(err)

	UnmarshalSingInMsg(bs, t)
}
