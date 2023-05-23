package protoTool

import (
	"encoding/binary"
	"fmt"
	"game-message-core/proto"

	googleProto "google.golang.org/protobuf/proto"
)

const (
	MSG_TYPE_SIZE   = 4
	MSG_LENGTH_SIZE = 2
)

func EnvelopeTypeToServiceType(envelopeType proto.EnvelopeType) proto.ServiceType {
	return proto.ServiceType(int(envelopeType) >> 16)
}

func MarshalProto(message googleProto.Message) ([]byte, error) {
	return googleProto.Marshal(message)
}

func UnmarshalProto(data []byte, message googleProto.Message) error {
	return googleProto.Unmarshal(data, message)
}

// func UnMarshalToEnvelope(data []byte) (*proto.Envelope, error) {
// 	msg := &proto.Envelope{}
// 	err := UnmarshalProto(data, msg)
// 	return msg, err
// }

// func MarshalEnvelope(msg *proto.Envelope) ([]byte, error) {
// 	return MarshalProto(msg)
// }

// 编码
func EnCodeNetData(envelopeType proto.EnvelopeType, data []byte) []byte {
	dataLength := len(data)
	body := make([]byte, MSG_TYPE_SIZE+MSG_LENGTH_SIZE+dataLength)

	var index int
	binary.LittleEndian.PutUint32(body[:], uint32(envelopeType))
	index += MSG_TYPE_SIZE

	binary.LittleEndian.PutUint16(body[index:], uint16(dataLength))
	index += MSG_LENGTH_SIZE

	copy(body[index:], data)
	return body
}

// 解码
func DeCodeNetData(body []byte) (proto.EnvelopeType, []byte, error) {
	if len(body) < (MSG_TYPE_SIZE + MSG_LENGTH_SIZE) {
		return proto.EnvelopeType_Unknown, nil, fmt.Errorf("net data length invalid")
	}

	typeBody := body[:MSG_TYPE_SIZE]
	typeId := binary.LittleEndian.Uint32(typeBody)

	body = body[MSG_TYPE_SIZE:]
	l := binary.LittleEndian.Uint16(body[:MSG_LENGTH_SIZE])

	body = body[MSG_LENGTH_SIZE:]
	if len(body) < int(l) {
		return proto.EnvelopeType_Unknown, nil, fmt.Errorf("invalid net data")
	}

	return proto.EnvelopeType(typeId), body[:l], nil
}
