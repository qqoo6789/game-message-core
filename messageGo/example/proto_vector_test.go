package main

import (
	`fmt`
	`testing`

	`game-message-core/proto`
	`game-message-core/protoTool`
)

// go test -v -run TestProtoVector3 messageGo/example/proto_vector_test.go

func TestProtoVector3(t *testing.T) {
	pos := &proto.Vector3{X: 1.3, Y: 2.1, Z: 3.9}
	t.Log(fmt.Sprintf("%+v", pos))

	out, err := protoTool.MarshalProto(pos)
	t.Log(out, "	", err)

	// 反序列化user结构体
	pos2 := proto.Vector3{}
	err = protoTool.UnmarshalProto(out, &pos2)
	t.Log(pos2, "	", err)
	t.Log(pos2.String())
}
