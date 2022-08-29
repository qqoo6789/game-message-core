package protoTool

import (
	googleProto "google.golang.org/protobuf/proto"
)

func MarshalProto(message googleProto.Message) ([]byte, error) {
	return googleProto.Marshal(message)
}

func UnmarshalProto(data []byte, message googleProto.Message) error {
	return googleProto.Unmarshal(data, message)
}
