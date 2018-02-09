package privatezen

import (
	"strconv"
)

type Whosonfirst struct {
	Id        int64
	Placetype string // should be go-whosonfirst-placetype...
}

func (w *Whosonfirst) String() string {
	return strconv.Itoa(w.Id)
}
