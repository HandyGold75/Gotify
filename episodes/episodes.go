package episodes

import (
	"encoding/json"
	"strconv"
	"strings"

	"github.com/HandyGold75/gotify/lib"
)

type (
	Episodes struct {
		Send   func(method lib.HTTPMethod, action string, options [][2]string, body []byte) ([]byte, error)
		Market string // An ISO 3166-1 alpha-2 country code, https://en.wikipedia.org/wiki/ISO_3166-1_alpha-2
	}

	getEpisode lib.EpisodeObject

	getSeveralEpisodes lib.Episodes

	getUsersSavedEpisodes struct {
		lib.ItemsHeaders
		Items []struct {
			AddedAt string `json:"added_at"`
			lib.Episode
		} `json:"items"`
	}
)

func New(send func(method lib.HTTPMethod, action string, options [][2]string, body []byte) ([]byte, error)) Episodes {
	return Episodes{Send: send, Market: ""}
}

// Scopes: `ScopeUserReadPlaybackPosition`
func (s *Episodes) GetEpisode(id string) (getEpisode, error) {
	res, err := s.Send(lib.GET, "episodes/"+id, [][2]string{{"market", s.Market}}, []byte{})
	if err != nil {
		return getEpisode{}, err
	}
	data := getEpisode{}
	err = json.Unmarshal(res, &data)
	return data, err
}

// Scopes: `ScopeUserReadPlaybackPosition`
func (s *Episodes) GetSeveralEpisodes(ids []string) (getSeveralEpisodes, error) {
	res, err := s.Send(lib.GET, "episodes", [][2]string{{"market", s.Market}, {"ids", strings.Join(ids, ",")}}, []byte{})
	if err != nil {
		return getSeveralEpisodes{}, err
	}
	data := getSeveralEpisodes{}
	err = json.Unmarshal(res, &data)
	return data, err
}

// Scopes: `ScopeUserLibraryRead`, `ScopeUserReadPlaybackPosition`
func (s *Episodes) GetUsersSavedEpisodes(limit, offset int) (getUsersSavedEpisodes, error) {
	res, err := s.Send(lib.GET, "me/episodes", [][2]string{{"market", s.Market}, {"limit", strconv.Itoa(max(1, min(50, limit)))}, {"offset", strconv.Itoa(max(0, offset))}}, []byte{})
	if err != nil {
		return getUsersSavedEpisodes{}, err
	}
	data := getUsersSavedEpisodes{}
	err = json.Unmarshal(res, &data)
	return data, err
}

// Scopes: `ScopeUserLibraryModify`
func (s *Episodes) SaveEpisodesForCurrentUser(ids []string) error {
	body, err := json.Marshal(map[string]any{"ids": ids})
	if err != nil {
		return err
	}
	_, err = s.Send(lib.PUT, "me/episodes", [][2]string{}, body)
	return err
}

// Scopes: `ScopeUserLibraryModify`
func (s *Episodes) RemoveUsersSavedEpisodes(ids []string) error {
	body, err := json.Marshal(map[string]any{"ids": ids})
	if err != nil {
		return err
	}
	_, err = s.Send(lib.DELETE, "me/episodes", [][2]string{}, body)
	return err
}

// Scopes: `ScopeUserLibraryRead`
func (s *Episodes) CheckUsersSavedEpisodes(ids []string) ([]bool, error) {
	res, err := s.Send(lib.GET, "me/episodes/contains", [][2]string{{"ids", strings.Join(ids, ",")}}, []byte{})
	if err != nil {
		return []bool{}, err
	}
	data := []bool{}
	err = json.Unmarshal(res, &data)
	return data, err
}
