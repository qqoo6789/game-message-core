package protoTool

import (
	"game-message-core/proto"
	"testing"
	"time"
)

func Test_ProtoTool(t *testing.T) {
	msg := &proto.AllLandDataResponse{}
	var num int32 = 350
	for i := int32(0); i < num; i++ {
		msg.Lands = append(msg.Lands, &proto.LandData{
			Id:      i + 1,
			X:       i - 1,
			Z:       i + 1,
			Owner:   300000 + i,
			Timeout: time.Now().UTC().UnixMilli(),
		})
	}

	bs, err := MarshalProto(msg)
	t.Log(err)
	t.Log(bs)
	t.Log("[", num, "]个land, pb交互数据length= [", float32(len(bs))/1000.0, "]KB")

	//10 19 8 1 16 255 255 255 255 255 255 255 255 255 1 24 1 32 225 167 18

	// msg := &proto.Envelope{
	// 	SeqId: 999,
	// 	Type:  proto.EnvelopeType_SigninPlayer,
	// 	Payload: &proto.Envelope_SigninPlayerRequest{
	// 		SigninPlayerRequest: &proto.SigninPlayerRequest{
	// 			UserId: 7001,
	// 			Token:  "vvvvvvvvvvvbbbbbbbbbbbbbbbb",
	// 		},
	// 	},
	// }
	// t.Log(fmt.Sprintf("%+v", msg))
	// t.Log("******************************************************")

	// bs, err := MarshalProto(msg)
	// t.Log(err)
	// t.Log(bs)
	// t.Log(string(bs))
	// t.Log("******************************************************")

	// msgParse := &proto.Envelope{}
	// err = UnmarshalProto(bs, msgParse)
	// t.Log(err)
	// t.Log(fmt.Sprintf("%+v", msgParse))
	// t.Log("******************************************************")

	// t.Log(EnvelopeTypeToServiceType(msgParse.Type))

}
