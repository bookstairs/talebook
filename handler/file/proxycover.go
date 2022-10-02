package file

import (
	"github.com/gofiber/fiber/v2"
	"time"
)

const maxBookQuery = 30

func proxyCover(ctx *fiber.Ctx) error {
	url := ctx.Query("url", "")
	agent := fiber.AcquireAgent()
	agent.Timeout(time.Second)
	r := agent.Request()
	r.Header.SetMethod("GET")
	r.Header.Add("Accept-Language", "zh-CN,zh;q=0.8,zh-TW;q=0.6")
	r.Header.Add("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,*/*;q=0.8")
	r.Header.Add("User-Agent", "Mozilla/5.0 (Windows NT 10.0; WOW64) AppleWebKit/537.36 (KHTML, like Gecko)"+"Chrome/66.0.3359.139 Safari/537.36")
	r.SetRequestURI(url)
	agent.SetResponse(ctx.Response())
	return agent.Parse()
}
