package main

import (
	"encoding/json"
	"fmt"
	base_data "game-message-core/grpc/baseData"
	"testing"
)

func TestJsonVector3(t *testing.T) {
	pos := base_data.GrpcVector3{X: 22}
	t.Log(fmt.Sprintf("%+v", pos))

	bs, err := json.Marshal(pos)
	t.Log(bs)
	t.Log(err)

	pos2 := &base_data.GrpcVector3{}
	err = json.Unmarshal(bs, pos2)
	t.Log(pos2)
	t.Log(err)
}
