package audiobooks

import (
	"encoding/json"
	"strconv"
	"strings"

	"github.com/HandyGold75/gotify/lib"
)

type (
	Audiobooks struct {
		Send   func(method lib.HTTPMethod, action string, options [][2]string, body []byte) ([]byte, error)
		Market string // An ISO 3166-1 alpha-2 country code, https://en.wikipedia.org/wiki/ISO_3166-1_alpha-2
	}

	getAnAudiobook lib.AudiobookObject

	getSeveralAudiobooks lib.Audiobooks

	getAudiobookChapters struct {
		lib.ItemsHeaders
		Items []lib.ChapterSimpleObject `json:"items"`
	}

	getUsersSavedAudiobooks struct {
		lib.ItemsHeaders
		Items []lib.AudiobookObject `json:"items"`
	}
)

func New(send func(method lib.HTTPMethod, action string, options [][2]string, body []byte) ([]byte, error)) Audiobooks {
	return Audiobooks{Send: send, Market: ""}
}

func (s *Audiobooks) GetAnAudiobook(id string) (getAnAudiobook, error) {
	res, err := s.Send(lib.GET, "audiobooks/"+id, [][2]string{{"market", s.Market}}, []byte{})
	if err != nil {
		return getAnAudiobook{}, err
	}
	data := getAnAudiobook{}
	err = json.Unmarshal(res, &data)
	return data, err
}

func (s *Audiobooks) GetSeveralAudiobooks(ids []string) (getSeveralAudiobooks, error) {
	res, err := s.Send(lib.GET, "audiobooks", [][2]string{{"market", s.Market}, {"ids", strings.Join(ids, ",")}}, []byte{})
	if err != nil {
		return getSeveralAudiobooks{}, err
	}
	data := getSeveralAudiobooks{}
	err = json.Unmarshal(res, &data)
	return data, err
}

func (s *Audiobooks) GetAudiobookChapters(id string, limit, offset int) (getAudiobookChapters, error) {
	res, err := s.Send(lib.GET, "audiobooks/"+id+"/chapters", [][2]string{{"market", s.Market}, {"limit", strconv.Itoa(max(1, min(50, limit)))}, {"offset", strconv.Itoa(max(0, offset))}}, []byte{})
	if err != nil {
		return getAudiobookChapters{}, err
	}
	data := getAudiobookChapters{}
	err = json.Unmarshal(res, &data)
	return data, err
}

// Scopes: `ScopeUserLibraryRead`
func (s *Audiobooks) GetUsersSavedAudiobooks(limit, offset int) (getUsersSavedAudiobooks, error) {
	res, err := s.Send(lib.GET, "me/audiobooks", [][2]string{{"limit", strconv.Itoa(max(1, min(50, limit)))}, {"offset", strconv.Itoa(max(0, offset))}}, []byte{})
	if err != nil {
		return getUsersSavedAudiobooks{}, err
	}
	data := getUsersSavedAudiobooks{}
	err = json.Unmarshal(res, &data)
	return data, err
}

// Scopes: `ScopeUserLibraryModify`
func (s *Audiobooks) SaveAudiobooksForCurrentUser(ids []string) error {
	_, err := s.Send(lib.PUT, "me/audiobooks", [][2]string{{"ids", strings.Join(ids, ",")}}, []byte{})
	return err
}

// Scopes: `ScopeUserLibraryModify`
func (s *Audiobooks) RemoveUsersSavedAudiobooks(ids []string) error {
	_, err := s.Send(lib.DELETE, "me/audiobooks", [][2]string{{"ids", strings.Join(ids, ",")}}, []byte{})
	return err
}

// Scopes: `ScopeUserLibraryRead`
func (s *Audiobooks) CheckUsersSavedAudiobooks(ids []string) ([]bool, error) {
	res, err := s.Send(lib.GET, "me/audiobooks/contains", [][2]string{{"ids", strings.Join(ids, ",")}}, []byte{})
	if err != nil {
		return []bool{}, err
	}
	data := []bool{}
	err = json.Unmarshal(res, &data)
	return data, err
}
