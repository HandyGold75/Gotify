package artists

import (
	"encoding/json"
	"strconv"
	"strings"

	"github.com/HandyGold75/gotify/lib"
)

type (
	Artists struct {
		Send   func(method lib.HTTPMethod, action string, options [][2]string, body []byte) ([]byte, error)
		Market string // An ISO 3166-1 alpha-2 country code, https://en.wikipedia.org/wiki/ISO_3166-1_alpha-2
	}

	getArtist lib.ArtistObject

	getSeveralArtists lib.Artists

	getArtistsAlbums struct {
		lib.ItemsHeaders
		Items []lib.AlbumSimpleObject `json:"items"`
	}

	getArtistsTopTracks lib.Tracks
)

func New(send func(method lib.HTTPMethod, action string, options [][2]string, body []byte) ([]byte, error)) Artists {
	return Artists{Send: send, Market: ""}
}

func (s *Artists) GetArtist(id string) (getArtist, error) {
	res, err := s.Send(lib.GET, "artists/"+id, [][2]string{}, []byte{})
	if err != nil {
		return getArtist{}, err
	}
	data := getArtist{}
	err = json.Unmarshal(res, &data)
	return data, err
}

func (s *Artists) GetSeveralArtists(ids []string) (getSeveralArtists, error) {
	res, err := s.Send(lib.GET, "artists", [][2]string{{"ids", strings.Join(ids, ",")}}, []byte{})
	if err != nil {
		return getSeveralArtists{}, err
	}
	data := getSeveralArtists{}
	err = json.Unmarshal(res, &data)
	return data, err
}

func (s *Artists) GetArtistsAlbums(id string, groups []lib.AlbumGroup, limit, offset int) (getArtistsAlbums, error) {
	grps := []string{}
	for _, grp := range groups {
		grps = append(grps, string(grp))
	}
	res, err := s.Send(lib.GET, "artists/"+id+"/albums", [][2]string{{"include_groups", strings.Join(grps, ",")}, {"market", s.Market}, {"limit", strconv.Itoa(max(1, min(50, limit)))}, {"offset", strconv.Itoa(max(0, offset))}}, []byte{})
	if err != nil {
		return getArtistsAlbums{}, err
	}
	data := getArtistsAlbums{}
	err = json.Unmarshal(res, &data)
	return data, err
}

func (s *Artists) GetArtistsTopTracks(id string) (getArtistsTopTracks, error) {
	res, err := s.Send(lib.GET, "artists/"+id+"/top-tracks", [][2]string{{"market", s.Market}}, []byte{})
	if err != nil {
		return getArtistsTopTracks{}, err
	}
	data := getArtistsTopTracks{}
	err = json.Unmarshal(res, &data)
	return data, err
}
