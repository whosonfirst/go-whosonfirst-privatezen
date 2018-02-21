package privatezen

import (
	"github.com/whosonfirst/go-whosonfirst-placetypes"
	"strconv"
)

type WOFPlace struct {
	Id        int64
	Placetype *placetypes.WOFPlacetype
}

func (w *WOFPlace) String() string {
	return strconv.Itoa(w.Id)
}
