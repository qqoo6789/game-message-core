package main

import (
	"game-message-core/proto"
	"game-message-core/protoTool"
	"testing"
)

func TestMessageIdToServiceType(t *testing.T) {
	serviceType := protoTool.EnvelopeTypeToServiceType(proto.EnvelopeType_ItemGet)
	t.Log(serviceType)

	serviceType = protoTool.EnvelopeTypeToServiceType(proto.EnvelopeType_UnloadAvatar)
	t.Log(serviceType)

	serviceType = protoTool.EnvelopeTypeToServiceType(proto.EnvelopeType_QueryPlayer)
	t.Log(serviceType)

	serviceType = protoTool.EnvelopeTypeToServiceType(proto.EnvelopeType_EnterMap)
	t.Log(serviceType)

	serviceType = protoTool.EnvelopeTypeToServiceType(proto.EnvelopeType_UseSkill)
	t.Log(serviceType)

	serviceType = protoTool.EnvelopeTypeToServiceType(proto.EnvelopeType_AcceptTask)
	t.Log(serviceType)

	serviceType = protoTool.EnvelopeTypeToServiceType(proto.EnvelopeType_BroadCastChatMessages)
	t.Log(serviceType)

	serviceType = protoTool.EnvelopeTypeToServiceType(proto.EnvelopeType_BroadCastUpdateChatState)
	t.Log(serviceType)

	serviceType = protoTool.EnvelopeTypeToServiceType(proto.EnvelopeType_Ping)
	t.Log(serviceType)

}
