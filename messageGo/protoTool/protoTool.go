package protoTool

import (
	"game-message-core/proto"

	"github.com/golang/snappy"
	googleProto "google.golang.org/protobuf/proto"
)

func MarshalProto(message googleProto.Message) ([]byte, error) {
	serialized, err := googleProto.Marshal(message)
	if err != nil {
		return nil, err
	}

	// 使用snappy压缩字节切片
	compressed := snappy.Encode(nil, serialized)
	return compressed, nil
}

func UnmarshalProto(data []byte, message googleProto.Message) error {
	decompressed, err := snappy.Decode(nil, data)
	if err != nil {
		return err
	}
	err = googleProto.Unmarshal(decompressed, message)
	if err != nil {
		return err
	}
	return nil
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
