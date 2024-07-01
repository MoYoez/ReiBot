package rei

import (
	"github.com/FloatTech/floatbox/process"
	"github.com/wdvxdr1123/ZeroBot/extension/rate"
	"log"
	"time"
)

// when receiving too much message in times, prevent and use message to stop.

var RateManager = rate.NewManager[int64](time.Second*1, 15)
var PreferStoper = rate.NewManager[int64](time.Second*3, 1)

func init() {
	process.NewCustomOnce(&m).Do(func() {
		OnMessage().SetBlock(false).Handle(func(ctx *Ctx) {
			if PreferStoper.Load(0).Tokens() == 0 {
				log.Print("[Reibot] Receiving too much message in times, stop message handler.")
				ctx.Block()
				return
			}
			if !m.CanResponse(0) && !ctx.Caller.b.RequireAuth {
				m.Response(0)
			}
			RateManager.Load(time.Now().Unix()).Acquire()
			if RateManager.Load(time.Now().Unix()).Tokens() == 0 {
				PreferStoper.Load(0).Acquire()
			}
		})
	})
}
