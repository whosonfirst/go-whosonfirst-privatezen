package privatezen

import (
	"time"
)

type Visit struct {
	Id              int64
	Name            string
	Date            *time.Time
	Latitude        Float64
	Longitude       Float64
	WOFId           *WOFPlace
	NeighbourhoodId *WOFPlace
	MacrohoodId     *WOFPlace
	LocalityId      *WOFPlace
	MetroareaId     *WOFPlace
	RegionId        *WOFPlace
	CountryId       *WOFPlace
	FeelingsId      *Feeling
}
