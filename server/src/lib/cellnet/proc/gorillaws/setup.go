package gorillaws

import (
	"lib/cellnet"
	"lib/cellnet/proc"
)

func init() {

	proc.RegisterProcessor("gorillaws.ltv", func(bundle proc.ProcessorBundle, userCallback cellnet.EventCallback, args ...interface{}) {

		bundle.SetTransmitter(new(WSMessageTransmitter))
		bundle.SetHooker(new(MsgHooker))
		bundle.SetCallback(proc.NewQueuedEventCallback(userCallback))

	})
}
