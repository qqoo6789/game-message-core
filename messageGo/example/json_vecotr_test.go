package main

import (
	"encoding/json"
	"fmt"
	"testing"

	"game-message-core/jsonData"
)

func TestJsonVector3(t *testing.T) {
	pos := jsonData.Vector3{X: 22}
	t.Log(fmt.Sprintf("%+v", pos))

	bs, err := json.Marshal(pos)
	t.Log(bs)
	t.Log(err)

	pos2 := &jsonData.Vector3{}
	err = json.Unmarshal(bs, pos2)
	t.Log(pos2)
	t.Log(err)
}
