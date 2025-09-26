package search

import (
	"encoding/json"
	"strconv"
	"strings"

	"github.com/HandyGold75/gotify/lib"
)

type searchForItem struct {
	Tracks struct {
		lib.ItemsHeaders
		Items []lib.TrackObject `json:"items"`
	} `json:"tracks"`
	Artists struct {
		lib.ItemsHeaders
		Items []lib.ArtistObject `json:"items"`
	} `json:"artists"`
	Albums struct {
		lib.ItemsHeaders
		Items []lib.AlbumSimpleObject `json:"items"`
	} `json:"albums"`
	Playlists struct {
		lib.ItemsHeaders
		Items []lib.PlaylistSimpleObject `json:"items"`
	} `json:"playlists"`
	Shows struct {
		lib.ItemsHeaders
		Items []lib.ShowSimpleObject `json:"items"`
	} `json:"shows"`
	Episodes struct {
		lib.ItemsHeaders
		Items []lib.EpisodeSimpleObject `json:"items"`
	} `json:"episodes"`
	Audiobooks struct {
		lib.ItemsHeaders
		Items []lib.AudiobookSimpleObject `json:"items"`
	} `json:"audiobooks"`
}

type Search struct {
	Send   func(method lib.HTTPMethod, action string, options [][2]string, body []byte) ([]byte, error)
	Market string // An ISO 3166-1 alpha-2 country code, https://en.wikipedia.org/wiki/ISO_3166-1_alpha-2
}

func New(send func(method lib.HTTPMethod, action string, options [][2]string, body []byte) ([]byte, error)) Search {
	return Search{
		Send:   send,
		Market: "",
	}
}

func (s *Search) SearchForItem(query string, typ []lib.URIResource, limit, offset int) (searchForItem, error) {
	typs := []string{}
	for _, t := range typ {
		typs = append(typs, string(t))
	}
	res, err := s.Send(lib.GET, "search", [][2]string{{"query", query}, {"type", strings.Join(typs, ",")}, {"market", s.Market}, {"limit", strconv.Itoa(max(1, min(50, limit)))}, {"offset", strconv.Itoa(max(0, offset))}}, []byte{})
	if err != nil {
		return searchForItem{}, err
	}
	data := searchForItem{}
	err = json.Unmarshal(res, &data)
	return data, err
}

func (s *Search) SearchForItemExternal(query string, typ []lib.URIResource, limit, offset int) (searchForItem, error) {
	typs := []string{}
	for _, t := range typ {
		typs = append(typs, string(t))
	}
	res, err := s.Send(lib.GET, "search", [][2]string{{"query", query}, {"type", strings.Join(typs, ",")}, {"market", s.Market}, {"limit", strconv.Itoa(max(1, min(50, limit)))}, {"offset", strconv.Itoa(max(0, offset))}, {"include_external", "audio"}}, []byte{})
	if err != nil {
		return searchForItem{}, err
	}
	data := searchForItem{}
	err = json.Unmarshal(res, &data)
	return data, err
}
