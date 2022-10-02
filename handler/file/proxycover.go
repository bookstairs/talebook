package file

import (
	"fmt"
	"time"

	"github.com/gofiber/fiber/v2"

	"github.com/bookstairs/talebook/handler/common"
)

const (
	acceptLanguage = "zh-CN,zh;q=0.8,zh-TW;q=0.6"
	accept         = "text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,*/*;q=0.8"
	userAgent      = "Mozilla/5.0 (Windows NT 10.0; WOW64) AppleWebKit/537.36 (KHTML, like Gecko)" +
		"Chrome/66.0.3359.139 Safari/537.36"
)

func ProxyCover(ctx *fiber.Ctx) error {
	url := common.GetQueryString(ctx, "url", "")
	if url == "" {
		return common.NotFound(ctx, fmt.Errorf("no url was provided"))
	}

	// Start downloading the image through the fiber agent.
	agent := fiber.AcquireAgent()
	agent.Timeout(time.Second)
	r := agent.Request()
	r.Header.SetMethod("GET")
	r.Header.Add("Accept", accept)
	r.Header.Add("Accept-Language", acceptLanguage)
	r.Header.Add("User-Agent", userAgent)
	r.SetRequestURI(url)
	agent.SetResponse(ctx.Response())

	return agent.Parse()
}
