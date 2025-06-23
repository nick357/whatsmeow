package whatsmeow

import (
	waBinary "go.mau.fi/whatsmeow/binary"
	"go.mau.fi/whatsmeow/types"
	"go.mau.fi/whatsmeow/types/events"
)

func (cli *Client) handleMessageAck(node *waBinary.Node) {
	if node.Attrs["class"] == "message" {
		var msgId = node.Attrs["id"].(types.MessageID)
		cli.dispatchEvent(&events.MessageAck{
			MessageID: msgId,
		})
	}
}
