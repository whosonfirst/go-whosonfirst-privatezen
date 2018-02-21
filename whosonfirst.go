package privatezen

import (
	"github.com/whosonfirst/go-whosonfirst-geojson-v2"
	"github.com/whosonfirst/go-whosonfirst-geojson-v2/properties/whosonfirst"
	"github.com/whosonfirst/go-whosonfirst-placetypes"
	"strconv"
)

func WOFPlaceFromFeature(f geojson.Feature) (*WOFPlaces, error) {

	id := whosonfirst.Id(f)

	if id == -1 {
		return nil, errors.New("Invalid WOF ID")
	}

	str_pt := whosonfirst.Placeptype(f)

	if !placetypes.IsValidPlacetype(str_pr) {
		return nil, errors.New("Invalid WOF placetype")
	}

	pt, err := placetypes.GetPlacetypeByName(str_pl)

	if err != nil {
		return nil, err
	}

	pl := WOFPlace{
		Id:        id,
		Placetype: pt,
	}

	return &pl, nil
}

type WOFPlace struct {
	Id        int64
	Placetype *placetypes.WOFPlacetype
}

func (w *WOFPlace) String() string {
	return strconv.Itoa(w.Id)
}
