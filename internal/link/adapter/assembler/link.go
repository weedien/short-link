package assembler

import (
	"database/sql"
	"shortlink/internal/base/database"
	"shortlink/internal/link/adapter/po"
	"shortlink/internal/link/domain/link"
	"time"
)

type LinkAssembler struct {
	linkFactory *link.Factory
}

func NewLinkAssembler(linkFactory *link.Factory) LinkAssembler {
	return LinkAssembler{linkFactory: linkFactory}
}

func (a *LinkAssembler) LinkEntityToLinkPo(lk *link.Link) po.Link {

	startDate := sql.NullTime{}
	if lk.ValidDate().StartDate() == nil {
		startDate.Time = time.Now()
		startDate.Valid = true
	} else {
		startDate.Time = *lk.ValidDate().StartDate()
		startDate.Valid = true
	}

	endDate := sql.NullTime{}
	if lk.ValidDate().EndDate() == nil {
		endDate.Valid = false
	} else {
		endDate.Time = *lk.ValidDate().EndDate()
		endDate.Valid = true
	}

	return po.Link{
		BaseModel: database.BaseModel{
			ID: lk.ID(),
		},
		Gid:         lk.Gid(),
		ShortUri:    lk.ShortUri(),
		OriginalUrl: lk.OriginalUrl(),
		Favicon:     lk.Favicon(),
		Status:      lk.Status(),
		CreateType:  lk.CreateType(),
		ValidType:   lk.ValidDate().ValidType(),
		StartDate:   startDate,
		EndDate:     endDate,
		Desc:        lk.Desc(),
	}
}

func (a *LinkAssembler) LinkPoToLinkEntity(po po.Link) *link.Link {

	start, end := new(time.Time), new(time.Time)
	if po.StartDate.Valid {
		*start = po.StartDate.Time
	}
	if po.EndDate.Valid {
		*end = po.EndDate.Time
	}

	validDate, err := link.NewValidDate(po.ValidType, start, end)
	if err != nil {
		return nil
	}

	lk := &link.Link{}
	if lk, err = a.linkFactory.NewLinkFromDB(
		po.ID,
		po.Gid,
		po.ShortUri,
		po.OriginalUrl,
		po.Status,
		po.CreateType,
		po.Favicon,
		po.Desc,
		validDate,
	); err != nil {
		return nil
	}
	return lk
}

func (a *LinkAssembler) LinkEntityToLinkGotoPo(lk *link.Link) *po.LinkGoto {
	return &po.LinkGoto{
		Gid:      lk.Gid(),
		ShortUri: lk.ShortUri(),
	}
}
