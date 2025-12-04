package whatsmeow

import (
	"fmt"
	waBinary "go.mau.fi/whatsmeow/binary"
	"go.mau.fi/whatsmeow/types"
	"math/rand"
	"time"
)

// 模拟发送消息前发送键盘事件的模拟行为
func (cli *Client) sendChatState1(jid types.JID) {
	// 1、发送自己在线状态
	// <presence type="available" name="Tank" />
	cli.SendPresence(types.PresenceAvailable)
	// 2、发送订阅请求
	// <presence type="subscribe" to="639757430046@s.whatsapp.net"><tctoken>0401173767940d8cc2be16</tctoken></presence>
	cli.SubscribePresence(jid)

	// 3、开始输入
	// <chatstate to="639757430046@s.whatsapp.net"><composing /></chatstate>
	cli.SendChatPresence(jid, types.ChatPresenceComposing, types.ChatPresenceMediaText)

	// 随机延迟1-3秒
	SleepRandom1to3()

	// 4、输入结束
	// <chatstate to="639757430046@s.whatsapp.net"><paused /></chatstate>
	cli.SendChatPresence(jid, types.ChatPresencePaused, types.ChatPresenceMediaText)
}

// 模拟发送消息前发送键盘事件的模拟行为
func (cli *Client) sendChatState2(jid types.JID) {
	// 6、建立信任
	// <iq to="s.whatsapp.net" type="set" xmlns="privacy" id="29294.52599-149"><tokens><token jid="639757430046@s.whatsapp.net" t="1761622030" type="trusted_contact" /></tokens></iq>
	cli.SetTrustedContact(jid.String())
}

/*
设置信任用户
<iq to="s.whatsapp.net" type="set" xmlns="privacy" id="29294.52599-149">

	<tokens>
	   <token jid="639757430046@s.whatsapp.net" t="1761622030" type="trusted_contact" />
	</tokens>

</iq>
*/
func (cli *Client) SetTrustedContact(jid string) error {
	if cli == nil {
		return ErrClientIsNil
	}
	timeStamp := time.Now().Unix()

	_, err := cli.sendIQ(infoQuery{
		Namespace: "privacy",
		Type:      "set",
		To:        types.ServerJID,
		Content: []waBinary.Node{{
			Tag: "tokens",
			Content: []waBinary.Node{{
				Tag: "token",
				Attrs: waBinary.Attrs{
					"jid":  jid,
					"t":    timeStamp,
					"type": "trusted_contact",
				},
			}},
		}},
	})
	if err != nil {
		return fmt.Errorf("error SetTrustedContact: %w", err)
	}
	return nil
}

// SleepRandom1to3 随机休眠 1 到 3 秒
func SleepRandom1to3() {
	// rand.Intn(3) 会生成 0,1,2 之一，所以 +1 得到 1,2,3
	secs := rand.Intn(3) + 1
	time.Sleep(time.Duration(secs) * time.Second)
}
