package categories

import (
	"encoding/json"
	"strconv"

	"github.com/HandyGold75/gotify/lib"
)

type (
	Categories struct {
		Send   func(method lib.HTTPMethod, action string, options [][2]string, body []byte) ([]byte, error)
		Locale string // an ISO 639-1 language code, http://en.wikipedia.org/wiki/ISO_639-1 and an ISO 3166-1 alpha-2 country code, http://en.wikipedia.org/wiki/ISO_3166-1_alpha-2 joined by an underscore.
	}

	getSeveralBrowseCategories struct {
		Categories struct {
			lib.ItemsHeaders
			Items []lib.CategorieObject `json:"items"`
		} `json:"categories"`
	}

	getSingleBrowseCategory lib.CategorieObject
)

func New(send func(method lib.HTTPMethod, action string, options [][2]string, body []byte) ([]byte, error)) Categories {
	return Categories{Send: send, Locale: ""}
}

func (s *Categories) GetSeveralBrowseCategories(limit, offset int) (getSeveralBrowseCategories, error) {
	res, err := s.Send(lib.GET, "browse/categories", [][2]string{{"locale", s.Locale}, {"limit", strconv.Itoa(max(1, min(50, limit)))}, {"offset", strconv.Itoa(max(0, offset))}}, []byte{})
	if err != nil {
		return getSeveralBrowseCategories{}, err
	}
	data := getSeveralBrowseCategories{}
	err = json.Unmarshal(res, &data)
	return data, err
}

func (s *Categories) GetSingleBrowseCategory(id string) (getSingleBrowseCategory, error) {
	res, err := s.Send(lib.GET, "browse/categories/"+id, [][2]string{{"locale", s.Locale}}, []byte{})
	if err != nil {
		return getSingleBrowseCategory{}, err
	}
	data := getSingleBrowseCategory{}
	err = json.Unmarshal(res, &data)
	return data, err
}
