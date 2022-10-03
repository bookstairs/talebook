package file

import (
	"encoding/binary"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/coocood/freecache"
	"github.com/gofiber/fiber/v2"

	"github.com/bookstairs/talebook/calibre"
	"github.com/bookstairs/talebook/config"
	"github.com/bookstairs/talebook/handler/common"
)

const (
	defaultCoverExpireTime = 1800 // 30 minutes.
	coverContentType       = "image/jpeg"
)

var (
	cache *freecache.Cache
)

// InitCoverCache will create a memory cache for caching the book covers.
func InitCoverCache(c *config.ServerConfig) error {
	if cache != nil {
		return fmt.Errorf("cover cache has been initailized")
	}
	if c.CoverCache == 0 {
		log.Println("cover cache is not enabled. You can set it for better performance.")
		return nil
	}

	cache = freecache.NewCache(c.CoverCache * 1024 * 1024)
	return nil
}

// ImageCover will get the request book cover and cached it for better performance.
func ImageCover(ctx *fiber.Ctx) error {
	id, err := common.GetParamInt(ctx, "id")
	if err != nil {
		return common.ErrResp(ctx, err)
	}

	kind := ctx.Params("kind")
	switch {
	case kind == "thumb" || strings.HasPrefix(kind, "thumb_"):
		return getThumb(ctx, kind, id)
	case kind == "cover":
		return getCover(ctx, id)
	case kind == "opf":
		return getOpf(ctx, id)
	default:
		return common.NotFound(ctx, fmt.Errorf("bad url"))
	}
}

// getCover will load the cover from the file system and cache it if it needs.
func getCover(ctx *fiber.Ctx, id int64) error {
	// Try to read file from cache.
	cacheKey := make([]byte, 8)
	binary.LittleEndian.PutUint64(cacheKey, uint64(id))

	if cache != nil {
		image, err := cache.Get(cacheKey)
		if err == nil {
			return sendFile(ctx, image, coverContentType)
		}
	}

	// Read the cover path from the calibre library.
	cover, err := calibre.QueryBookCover(ctx.UserContext(), id)
	if err != nil {
		return common.ErrResp(ctx, err)
	}
	if cover == "" {
		return ctx.Redirect(config.DefaultCoverPath, 302)
	}

	// Read file and save it into cache.
	if cache != nil {
		file, err := os.ReadFile(cover)
		if err != nil {
			// Cover isn't found.
			return ctx.Redirect(config.DefaultCoverPath, 302)
		}
		_ = cache.Set(cacheKey, file, defaultCoverExpireTime) // No need to care this error.

		return sendFile(ctx, file, coverContentType)
	}

	// Manually serve images, this would be used only if you disable the cache.
	return ctx.SendFile(cover, false)
}

// getThumb will resize the jpegs, we don't cache the resized jpeg, instead, we will cache the original file.
func getThumb(ctx *fiber.Ctx, kind string, id int64) error {
	// TODO Finish thumbnail generation.
	return nil
}

func getOpf(ctx *fiber.Ctx, id int64) error {
	// TODO Read the book metadata file.
	return nil
}

// sendFile will put the file bytes into the http response.
func sendFile(ctx *fiber.Ctx, file []byte, contextType string) error {
	ctx.Response().SetBodyRaw(file)
	ctx.Response().SetStatusCode(200)
	ctx.Response().Header.SetContentType(contextType)

	return nil
}
