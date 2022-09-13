package protoTool

import (
	"fmt"
	"game-message-core/proto"
	"testing"
)

func Test_ProtoTool(t *testing.T) {
	t.Log("----------------------")
	msg := &proto.Envelope{
		SeqId: 999,
		Type:  proto.EnvelopeType_SigninPlayer,
		Payload: &proto.Envelope_SigninPlayerRequest{
			SigninPlayerRequest: &proto.SigninPlayerRequest{
				UserId: 7001,
				Token:  "vvvvvvvvvvvbbbbbbbbbbbbbbbb",
			},
		},
	}
	t.Log(fmt.Sprintf("%+v", msg))
	t.Log("******************************************************")

	bs, err := MarshalProto(msg)
	t.Log(err)
	t.Log(bs)
	t.Log(string(bs))
	t.Log("******************************************************")

	msgParse := &proto.Envelope{}
	err = UnmarshalProto(bs, msgParse)
	t.Log(err)
	t.Log(fmt.Sprintf("%+v", msgParse))
	t.Log("******************************************************")

	t.Log(EnvelopeTypeToServiceType(msgParse.Type))

}
