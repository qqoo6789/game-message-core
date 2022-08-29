package main

import (
	"fmt"

	"game-message-core/proto"
	`game-message-core/protoTool`
)

func main() {
	pbMy := &proto.Vector3{X: 1.0, Y: 2.0, Z: 3.0}

	fmt.Println(fmt.Sprintf("%+v", pbMy))

	out, err := protoTool.MarshalProto(pbMy)
	fmt.Println(out)
	fmt.Println(err)

	// 反序列化user结构体
	user2 := proto.Vector3{}
	err = protoTool.UnmarshalProto(out, &user2)
	fmt.Println(fmt.Sprintf("user2 = %+v", user2))
	fmt.Println(err)
}
