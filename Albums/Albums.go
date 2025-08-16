package Albums

import (
	"encoding/json"
	"strconv"
	"strings"

	"github.com/HandyGold75/Gotify/lib"
)

type (
	Albums struct {
		Send   func(method lib.HttpMethod, action string, options [][2]string, body []byte) ([]byte, error)
		Market string // An ISO 3166-1 alpha-2 country code, https://en.wikipedia.org/wiki/ISO_3166-1_alpha-2
	}

	getAlbum lib.Album

	getSeveralAlbums lib.Albums

	getAlbumTracks struct {
		Href     string                  `json:"href"`
		Limit    int                     `json:"limit"`
		Next     string                  `json:"next"`
		Offset   int                     `json:"offset"`
		Previous string                  `json:"previous"`
		Total    int                     `json:"total"`
		Items    []lib.TrackSimpleObject `json:"items"`
	}

	getUsersSavedAlbums struct {
		Href     string `json:"href"`
		Limit    int    `json:"limit"`
		Next     string `json:"next"`
		Offset   int    `json:"offset"`
		Previous string `json:"previous"`
		Total    int    `json:"total"`
		Items    []struct {
			AddedAt string `json:"added_at"`
			lib.Album
		} `json:"items"`
	}

	getNewReleases struct {
		Albums struct {
			Href     string            `json:"href"`
			Limit    int               `json:"limit"`
			Next     string            `json:"next"`
			Offset   int               `json:"offset"`
			Previous string            `json:"previous"`
			Total    int               `json:"total"`
			Items    []lib.AlbumSimple `json:"items"`
		} `json:"albums"`
	}
)

func New(send func(method lib.HttpMethod, action string, options [][2]string, body []byte) ([]byte, error)) Albums {
	return Albums{Send: send, Market: ""}
}

func (s *Albums) GetAlbum(id string) (getAlbum, error) {
	res, err := s.Send(lib.GET, "albums/"+id, [][2]string{{"market", s.Market}}, []byte{})
	if err != nil {
		return getAlbum{}, err
	}
	data := getAlbum{}
	err = json.Unmarshal(res, &data)
	return data, err
}

func (s *Albums) GetSeveralAlbums(ids []string) (getSeveralAlbums, error) {
	res, err := s.Send(lib.GET, "albums", [][2]string{{"market", s.Market}, {"ids", strings.Join(ids, ",")}}, []byte{})
	if err != nil {
		return getSeveralAlbums{}, err
	}
	data := getSeveralAlbums{}
	err = json.Unmarshal(res, &data)
	return data, err
}

func (s *Albums) GetAlbumTracks(id string, limit, offset int) (getAlbumTracks, error) {
	res, err := s.Send(lib.GET, "albums/"+id+"/tracks", [][2]string{{"market", s.Market}, {"limit", strconv.Itoa(max(1, min(50, limit)))}, {"offset", strconv.Itoa(max(0, offset))}}, []byte{})
	if err != nil {
		return getAlbumTracks{}, err
	}
	data := getAlbumTracks{}
	err = json.Unmarshal(res, &data)
	return data, err
}

// Scopes: `ScopeUserLibraryRead`
func (s *Albums) GetUsersSavedAlbums(limit, offset int) (getUsersSavedAlbums, error) {
	res, err := s.Send(lib.GET, "me/albums", [][2]string{{"market", s.Market}, {"limit", strconv.Itoa(max(1, min(50, limit)))}, {"offset", strconv.Itoa(max(0, offset))}}, []byte{})
	if err != nil {
		return getUsersSavedAlbums{}, err
	}
	data := getUsersSavedAlbums{}
	err = json.Unmarshal(res, &data)
	return data, err
}

// Scopes: `ScopeUserLibraryModify`
func (s *Albums) SaveAlbumsForCurrentUser(ids []string) error {
	_, err := s.Send(lib.PUT, "me/albums", [][2]string{{"ids", strings.Join(ids, ",")}}, []byte{})
	return err
}

// Scopes: `ScopeUserLibraryModify`
func (s *Albums) RemoveUsersSavedAlbums(ids []string) error {
	_, err := s.Send(lib.DELETE, "me/albums", [][2]string{{"ids", strings.Join(ids, ",")}}, []byte{})
	return err
}

// Scopes: `ScopeUserLibraryRead`
func (s *Albums) CheckUsersSavedAlbums(ids []string) ([]bool, error) {
	res, err := s.Send(lib.GET, "me/albums/contains", [][2]string{{"ids", strings.Join(ids, ",")}}, []byte{})
	data := []bool{}
	err = json.Unmarshal(res, &data)
	return data, err
}

func (s *Albums) GetNewReleases(limit, offset int) (getNewReleases, error) {
	res, err := s.Send(lib.GET, "browse/new-releases", [][2]string{{"limit", strconv.Itoa(max(1, min(50, limit)))}, {"offset", strconv.Itoa(max(0, offset))}}, []byte{})
	if err != nil {
		return getNewReleases{}, err
	}
	data := getNewReleases{}
	err = json.Unmarshal(res, &data)
	return data, err
}
