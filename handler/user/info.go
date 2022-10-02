package user

import (
	"github.com/gofiber/fiber/v2"
	"github.com/golang-module/carbon/v2"

	"github.com/bookstairs/talebook/calibre"
	"github.com/bookstairs/talebook/config"
	"github.com/bookstairs/talebook/handler/common"
	"github.com/bookstairs/talebook/model"
)

// Info would serve /api/user/info
func Info(ctx *fiber.Ctx) error {
	bookCount, err := calibre.QueryBookCount(ctx.UserContext())
	if err != nil {
		return common.ErrResp(ctx, err)
	}
	tagCount, err := calibre.QueryTagCount(ctx.UserContext())
	if err != nil {
		return common.ErrResp(ctx, err)
	}
	authorCount, err := calibre.QueryAuthorCount(ctx.UserContext())
	if err != nil {
		return common.ErrResp(ctx, err)
	}
	publisherCount, err := calibre.QueryPublisherCount(ctx.UserContext())
	if err != nil {
		return common.ErrResp(ctx, err)
	}
	seriesCount, err := calibre.QuerySeriesCount(ctx.UserContext())
	if err != nil {
		return common.ErrResp(ctx, err)
	}

	// TODO Query friend site.
	friends := []model.FriendSite{
		{
			Name: "芒果读书",
			Href: "http://diumx.com/",
		},
		{
			Name: "鸠摩搜索",
			Href: "https://www.jiumodiary.com/",
		},
		{
			Name: "追更神器",
			Href: "https://github.com/hectorqin/reader",
		},
		{
			Name: "阅读链",
			Href: "https://www.yuedu.pro/",
		},
		{
			Name: "苦瓜书盘",
			Href: "https://www.kgbook.com",
		},
		{
			Name: "三秋书屋",
			Href: "https://www.sanqiu.cc/",
		},
	}

	// TODO Query system allow.
	allow := model.SysAllow{
		Register: true,
		Download: true,
		Push:     true,
		Read:     true,
	}

	// TODO Query system information.
	sys := model.SysInfo{
		Books:      int(bookCount),
		Tags:       int(tagCount),
		Authors:    int(authorCount),
		Publishers: int(publisherCount),
		Series:     int(seriesCount),
		Mtime:      carbon.Now().ToDateString(),
		Users:      1,
		Active:     1,
		Version:    config.TalebookVersion().SemanticString(),
		Title:      "藏经阁",
		Friends:    friends,
		Footer:     "本站基于Calibre构建，感谢开源界的力量。所有资源搜集于互联网，如有侵权请邮件联系。",
		Allow:      allow,
	}

	// TODO Query user information.
	user := model.User{
		Avatar:      "https://tva1.sinaimg.cn/default/images/default_avatar_male_50.gif",
		IsLogin:     false,
		IsAdmin:     false,
		Nickname:    "",
		Email:       "",
		KindleEmail: "",
		Extra:       map[string]any{},
	}

	return common.SuccResp(ctx, map[string]any{
		"cdn":  "",
		"sys":  sys,
		"user": user,
	})
}
