package protoTool

import (
	"game-message-core/proto"

	googleProto "google.golang.org/protobuf/proto"
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
