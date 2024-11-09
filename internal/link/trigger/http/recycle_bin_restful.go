package http

import (
	"github.com/gofiber/fiber/v2"
	"shortlink/internal/base/types"
	"shortlink/internal/link/app"
	"shortlink/internal/link/app/query"
	"shortlink/internal/link/common/config"
	"shortlink/internal/link/domain/link"
	"shortlink/internal/link/trigger/http/dto/req"
)

type RecycleBinApi struct {
	app app.Application
}

func NewLinkRecycleBinApi(app app.Application, router fiber.Router) {
	api := &RecycleBinApi{
		app: app,
	}

	prefix := config.Get().AppLink.BaseRoutePrefix

	recycleBin := router.Group(prefix + "/recycle-bin")
	// 保存到回收站
	recycleBin.Post("/save", api.SaveToRecycleBin)
	// 分页查询回收站短链接
	recycleBin.Get("/page", api.PageQueryRecycleBin)
	// 恢复短链接
	recycleBin.Post("/recover", api.RecoverLink)
	// 从回收站移除短链接
	recycleBin.Delete("/remove", api.RemoveFromRecycleBin)
}

// SaveToRecycleBin 保存到回收站
func (h RecycleBinApi) SaveToRecycleBin(c *fiber.Ctx) error {
	reqParam := req.RecycleBinSaveReq{}
	if err := c.BodyParser(&reqParam); err != nil {
		return err
	}

	err := h.app.Commands.SaveToRecycleBin.Handle(c.Context(), link.Identifier{
		Gid:      reqParam.Gid,
		ShortUri: reqParam.ShortUri,
	})
	if err != nil {
		return err
	}

	c.Status(fiber.StatusNoContent)

	return nil
}

// PageQueryRecycleBin 分页查询回收站短链接
func (h RecycleBinApi) PageQueryRecycleBin(c *fiber.Ctx) (err error) {
	reqParam := req.RecycleBinPageReq{}
	if err = c.QueryParser(&reqParam); err != nil {
		return err
	}

	res := &types.PageResp[query.Link]{}
	res, err = h.app.Queries.PageRecycleBin.Handle(c.Context(), query.PageRecycleBin{
		PageReq: reqParam.PageReq,
		Gids:    reqParam.Gids,
	})
	if err != nil {
		return err
	}

	// 查询用户拥有的 gids
	// 判断gid是否在用户拥有的gids中
	// 如果不在，返回错误
	// 如果在，继续执行

	// 如果输入的gids为空
	// 则查询用户拥有的所有gids

	//response := resp.LinkPageResp{}
	//if err = copier.Copy(&response, &res); err != nil {
	//	return err
	//}

	return c.JSON(res)
}

// RecoverLink 恢复短链接
func (h RecycleBinApi) RecoverLink(c *fiber.Ctx) error {
	reqParam := req.RecycleBinRecoverReq{}
	if err := c.BodyParser(&reqParam); err != nil {
		return err
	}

	err := h.app.Commands.RecoverFromRecycleBin.Handle(c.Context(), link.Identifier{
		Gid:      reqParam.Gid,
		ShortUri: reqParam.ShortUri,
	})
	if err != nil {
		return err
	}

	c.Status(fiber.StatusNoContent)

	return nil
}

// RemoveFromRecycleBin 从回收站移除短链接
func (h RecycleBinApi) RemoveFromRecycleBin(c *fiber.Ctx) error {
	reqParam := req.RecycleBinDeleteReq{}
	if err := c.BodyParser(&reqParam); err != nil {
		c.Status(fiber.StatusBadRequest)
		return err
	}

	err := h.app.Commands.RemoveFromRecycleBin.Handle(c.Context(), link.Identifier{
		Gid:      reqParam.Gid,
		ShortUri: reqParam.ShortUri,
	})
	if err != nil {
		c.Status(fiber.StatusInternalServerError)
		return err
	}

	c.Status(fiber.StatusNoContent)

	return nil
}
