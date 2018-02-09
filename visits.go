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
	WOFId           *Whosonfirst
	NeighbourhoodId *Whosonfirst
	MacrohoodId     *Whosonfirst
	LocalityId      *Whosonfirst
	MetroareaId     *Whosonfirst
	RegionId        *Whosonfirst
	CountryId       *Whosonfirst
	FeelingsId      *Feeling
}
