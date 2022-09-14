package protoTool

import (
	"game-message-core/proto"

	googleProto "google.golang.org/protobuf/proto"
)

func MarshalProto(message googleProto.Message) ([]byte, error) {
	return googleProto.Marshal(message)
}

func UnmarshalProto(data []byte, message googleProto.Message) error {
	return googleProto.Unmarshal(data, message)
}

func UnMarshalToEnvelope(data []byte) (*proto.Envelope, error) {
	msg := &proto.Envelope{}
	err := UnmarshalProto(data, msg)
	return msg, err
}

func MarshalEnvelope(msg *proto.Envelope) ([]byte, error) {
	return MarshalProto(msg)
}

func EnvelopeTypeToServiceType(envelopeType proto.EnvelopeType) proto.ServiceType {
	return proto.ServiceType(int(envelopeType) >> 16)
}
