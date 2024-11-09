package http

import (
	"github.com/gofiber/fiber/v2"
	"github.com/jinzhu/copier"
	"shortlink/internal/link_stats/app"
	"shortlink/internal/link_stats/app/query"
	"shortlink/internal/link_stats/trigger/http/dto/req"
	"shortlink/internal/link_stats/trigger/http/dto/resp"
)

type LinkStatsApi struct {
	app app.Application
}

func NewLinkStatsApi(app app.Application, router fiber.Router) {
	api := &LinkStatsApi{
		app: app,
	}

	// 访问单个短链接指定时间内监控数据
	router.Get("/stats", api.GetLinkStats)
	// 访问分组短链接指定时间内监控数据
	router.Get("/stats/group", api.GroupLinkStats)
	// 访问单个短链接指定时间内访问记录监控数据
	router.Get("/stats/access-record", api.GetLinkStatsAccessRecord)
	// 访问分组短链接指定时间内访问记录监控数据
	router.Get("/stats/access-record/group", api.GroupLinkStatsAccessRecord)
}

// GetLinkStats 获取短链接统计信息
func (h LinkStatsApi) GetLinkStats(c *fiber.Ctx) (err error) {
	reqParam := req.LinkStatsReq{}
	if err = c.QueryParser(&reqParam); err != nil {
		return err
	}

	res := &query.LinkStats{}
	res, err = h.app.Queries.GetLinkStats.Handle(c.Context(), query.GetLinkStats{
		FullShortUrl: reqParam.FullShortUrl,
		Gid:          reqParam.Gid,
		StartDate:    reqParam.StartTime,
		EndDate:      reqParam.EndTime,
		Status:       reqParam.Status,
	})
	if err != nil {
		return err
	}

	response := resp.LinkStatsResp{}
	if err = copier.Copy(&response, res); err != nil {
		return err
	}

	return c.JSON(res)
}

// GroupLinkStats 分组获取短链接统计信息
func (h LinkStatsApi) GroupLinkStats(c *fiber.Ctx) (err error) {
	reqParam := req.LinkGroupStatsReq{}
	if err := c.QueryParser(&reqParam); err != nil {
		return err
	}

	res := &query.LinkStats{}
	res, err = h.app.Queries.GroupLinkStats.Handle(c.Context(), query.GroupLinkStats{
		Gid:       reqParam.Gid,
		StartDate: reqParam.StartTime,
		EndDate:   reqParam.EndTime,
	})
	if err != nil {
		return err
	}

	response := resp.LinkStatsResp{}
	if err = copier.Copy(&response, &res); err != nil {
		return err
	}

	return c.JSON(res)
}

// GetLinkStatsAccessRecord 获取短链接访问记录
func (h LinkStatsApi) GetLinkStatsAccessRecord(c *fiber.Ctx) (err error) {

	reqParam := req.LinkStatsAccessRecordReq{}
	if err = c.QueryParser(&reqParam); err != nil {
		return err
	}

	res := &query.LinkStatsAccessRecord{}
	res, err = h.app.Queries.GetLinkStatsAccessRecord.Handle(c.Context(), query.GetLinkStatsAccessRecord{
		PageReq:      reqParam.PageReq,
		FullShortUrl: reqParam.FullShortUrl,
		Gid:          reqParam.Gid,
		StartDate:    reqParam.StartTime,
		EndDate:      reqParam.EndTime,
		Status:       reqParam.Status,
	})
	if err != nil {
		return err
	}

	response := resp.LinkStatsAccessRecordResp{}
	if err = copier.Copy(&response, res); err != nil {
		return err
	}

	return c.JSON(response)
}

// GroupLinkStatsAccessRecord 分组获取短链接访问记录
func (h LinkStatsApi) GroupLinkStatsAccessRecord(c *fiber.Ctx) (err error) {
	reqParam := req.LinkGroupStatsAccessRecordReq{}
	if err = c.QueryParser(&reqParam); err != nil {
		return err
	}

	res := &query.LinkStatsAccessRecord{}
	res, err = h.app.Queries.GroupLinkStatsAccessRecord.Handle(c.Context(), query.GroupLinkStatsAccessRecord{
		PageReq:   reqParam.PageReq,
		Gid:       reqParam.Gid,
		StartDate: reqParam.StartTime,
		EndDate:   reqParam.EndTime,
	})
	if err != nil {
		return err
	}

	var response resp.LinkStatsAccessRecordResp
	if err = copier.Copy(&response, &res); err != nil {
		return err
	}

	return c.JSON(response)
}
