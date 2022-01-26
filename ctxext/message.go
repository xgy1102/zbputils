package ctxext

import (
	"strconv"

	zero "github.com/wdvxdr1123/ZeroBot"
	"github.com/wdvxdr1123/ZeroBot/message"
)

type (
	NoCtxGetMsg  = func(int64) zero.Message
	NoCtxSendMsg = func(interface{}) int64
)

func GetMessage(ctx *zero.Ctx) NoCtxGetMsg {
	return func(id int64) zero.Message {
		return ctx.GetMessage(message.NewMessageID(strconv.FormatInt(id, 10)))
	}
}

func SendMessage(ctx *zero.Ctx, user int64) NoCtxSendMsg {
	return func(msg interface{}) int64 {
		return ctx.SendPrivateMessage(user, msg)
	}
}
