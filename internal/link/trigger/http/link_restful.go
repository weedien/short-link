package http

import (
	"errors"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/jinzhu/copier"
	"shortlink/internal/base/errno"
	"shortlink/internal/base/server/validator"
	"shortlink/internal/base/toolkit"
	"shortlink/internal/link/app"
	"shortlink/internal/link/app/command"
	"shortlink/internal/link/app/query"
	"shortlink/internal/link/common/config"
	"shortlink/internal/link/domain/event"
	"shortlink/internal/link/trigger/http/dto/req"
	"shortlink/internal/link/trigger/http/dto/resp"
	"strings"
	"time"
)

type LinkApi struct {
	app app.Application
}

func NewLinkApi(app app.Application, router fiber.Router) {
	api := &LinkApi{
		app: app,
	}

	prefix := config.Get().AppLink.BaseRoutePrefix

	// 短链接未找到
	router.All("/page/notfound", func(c *fiber.Ctx) error {
		return c.SendFile("../../templates/notfound.html")
	})
	// 短链接跳转到原始链接
	router.Get("/:shortUri", api.Redirect)
	// 创建短链接
	router.Post(prefix+"/create", api.CreateLink)
	// 通过分布式锁创建短链接
	router.Post(prefix+"/create/with-lock", api.CreateLinkWithLock)
	// 批量创建短链接
	router.Post(prefix+"/create-batch", api.BatchCreateLink)
	// 更新短链接
	router.Put(prefix+"/update", api.UpdateLink)
	// 分页查询短链接
	router.Get(prefix+"/page", api.PageQueryLink)
	// 查询短链接分组内数量
	router.Get(prefix+"/group-links-count", api.ListGroupLinkCount)
}

// Redirect 短链接跳转到原始链接
func (h LinkApi) Redirect(c *fiber.Ctx) error {

	shortUri := c.Params("shortUri")
	if shortUri == "" {
		return errno.LinkNotExists
	}

	os, browser, device, network := toolkit.GetRequestInfo(c)

	uv := c.Cookies("uv")
	if uv == "" {
		uv = uuid.NewString()
		c.Cookie(&fiber.Cookie{
			Name:    "uv",
			Value:   uv,
			Expires: time.Now().AddDate(0, 1, 0),
		})
	}

	userVisitInfo := event.UserVisitInfo{
		ShortUri:    shortUri,
		RemoteAddr:  c.IP(),
		OS:          os,
		Browser:     browser,
		Device:      device,
		Network:     network,
		UV:          uv,
		CurrentDate: time.Now(),
	}

	q := query.GetOriginalUrl{
		ShortUri:      shortUri,
		UserVisitInfo: userVisitInfo,
	}

	// 获取原始链接
	originalUrl, err := h.app.Queries.GetOriginalUrl.Handle(c.Context(), q)
	if err != nil {
		return err
	}

	if originalUrl == "" {
		return errno.LinkNotExists
	}

	return c.Redirect(originalUrl)
}

// CreateLink 创建短链接
func (h LinkApi) CreateLink(c *fiber.Ctx) (err error) {

	reqParam := req.LinkCreateReq{}
	if err = c.BodyParser(&reqParam); err != nil {
		return err
	}

	if err = validator.Get().Validate(reqParam); err != nil {
		return err
	}

	cmd := &command.CreateLink{
		OriginalUrl: reqParam.OriginalUrl,
		Gid:         reqParam.Gid,
		CreateType:  reqParam.CreateType,
		ValidType:   reqParam.ValidType,
		StartDate:   reqParam.StartDate.ToTime(),
		EndDate:     reqParam.EndDate.ToTime(),
		Desc:        reqParam.Desc,
		WithLock:    false,
	}

	if err = h.app.Commands.CreateLink.Handle(c.Context(), cmd); err != nil {
		return err
	}

	if res := cmd.ExecutionResult(); res != nil {
		response := resp.LinkCreateResp{
			Gid:          res.Gid,
			OriginalUrl:  res.OriginalUrl,
			FullShortUrl: res.FullShortUrl,
		}
		return c.JSON(response)
	}
	return
}

// CreateLinkWithLock 加锁创建短链接
func (h LinkApi) CreateLinkWithLock(c *fiber.Ctx) (err error) {

	reqParam := req.LinkCreateReq{}
	if err = c.BodyParser(&reqParam); err != nil {
		return err
	}

	if err = validator.Get().Validate(reqParam); err != nil {
		return err
	}

	cmd := &command.CreateLink{
		OriginalUrl: reqParam.OriginalUrl,
		Gid:         reqParam.Gid,
		CreateType:  reqParam.CreateType,
		ValidType:   reqParam.ValidType,
		StartDate:   reqParam.StartDate.ToTime(),
		EndDate:     reqParam.EndDate.ToTime(),
		Desc:        reqParam.Desc,
		WithLock:    true,
	}

	if err = h.app.Commands.CreateLink.Handle(c.Context(), cmd); err != nil {
		return err
	}

	if res := cmd.ExecutionResult(); res != nil {
		response := resp.LinkCreateResp{
			Gid:          res.Gid,
			OriginalUrl:  res.OriginalUrl,
			FullShortUrl: res.FullShortUrl,
		}
		return c.JSON(response)
	}
	return
}

// BatchCreateLink 批量创建短链接
func (h LinkApi) BatchCreateLink(c *fiber.Ctx) error {

	reqParam := req.LinkBatchCreateReq{}
	if err := c.BodyParser(&reqParam); err != nil {
		return err
	}

	if err := validator.Get().Validate(reqParam); err != nil {
		return err
	}

	cmd := command.CreateLinkBatch{
		OriginalUrls: reqParam.OriginalUrls,
		Descs:        reqParam.Descs,
		Gid:          reqParam.Gid,
		CreateType:   reqParam.CreateType,
		ValidType:    reqParam.ValidType,
		StartDate:    reqParam.StartDate.ToTime(),
		EndDate:      reqParam.EndDate.ToTime(),
	}

	if err := h.app.Commands.CreateLinkBatch.Handle(c.Context(), &cmd); err != nil {
		return err
	}

	if res := cmd.ExecutionResult(); res != nil {
		response := resp.LinkBatchCreateResp{}
		if err := copier.Copy(&response, res); err != nil {
			return err
		}
		return c.JSON(response)
	}

	return nil
}

// UpdateLink 更新短链接
func (h LinkApi) UpdateLink(c *fiber.Ctx) error {

	reqParam := req.LinkUpdateReq{}
	if err := c.BodyParser(&reqParam); err != nil {
		return err
	}

	err := h.app.Commands.UpdateLink.Handle(c.Context(), command.UpdateLink{
		ShortUri:       reqParam.ShortUri,
		OriginalUrl:    reqParam.OriginalUrl,
		Gid:            reqParam.Gid,
		Status:         reqParam.Status,
		ValidType:      reqParam.ValidType,
		ValidStartDate: reqParam.StartDate.ToTime(),
		ValidEndDate:   reqParam.EndDate.ToTime(),
		Desc:           reqParam.Desc,
	})
	if err != nil {
		return err
	}

	c.Status(fiber.StatusNoContent)

	return nil
}

// PageQueryLink 分页查询短链接
func (h LinkApi) PageQueryLink(c *fiber.Ctx) error {

	reqParam := req.LinkPageReq{}
	if err := c.QueryParser(&reqParam); err != nil {
		return err
	}

	res, err := h.app.Queries.PageLink.Handle(c.Context(), query.PageLink{
		PageReq:  reqParam.PageReq,
		Gid:      reqParam.Gid,
		OrderTag: reqParam.OrderTag,
	})
	if err != nil {
		return err
	}

	//response := resp.LinkPageResp{}
	//if err = copier.Copy(&response, &res); err != nil {
	//	return err
	//}

	return c.JSON(res)
}

// ListGroupLinkCount 查询短链接分组内数量
func (h LinkApi) ListGroupLinkCount(c *fiber.Ctx) error {

	gidStr := c.Query("gid")
	if gidStr == "" {
		return errors.New("gid is required")
	}
	gids := strings.Split(gidStr, ",")

	res, err := h.app.Queries.ListGroupCount.Handle(c.Context(), query.ListGroupLinksCount{Gids: gids})
	if err != nil {
		return err
	}

	var response interface{}
	if len(res) == 1 {
		response = resp.GroupCountDTO{}
		if err = copier.Copy(&response, &res[0]); err != nil {
			return err
		}
	} else {
		response = resp.LinkGroupCountQueryResp{}
		if err = copier.Copy(&response, &res); err != nil {
			return err
		}
	}

	return c.JSON(response)
}
