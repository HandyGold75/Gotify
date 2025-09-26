package tracks

import (
	"encoding/json"
	"strconv"
	"strings"
	"time"

	"github.com/HandyGold75/gotify/lib"
)

type (
	getTrack lib.TrackObject

	getTracks lib.Tracks

	getUsersSavedTracks struct {
		lib.ItemsHeaders
		Items []struct {
			AddedAt string `json:"added_at"`
			lib.Track
		} `json:"items"`
	}
)

type Tracks struct {
	Send   func(method lib.HTTPMethod, action string, options [][2]string, body []byte) ([]byte, error)
	Market string // An ISO 3166-1 alpha-2 country code, https://en.wikipedia.org/wiki/ISO_3166-1_alpha-2
}

func New(send func(method lib.HTTPMethod, action string, options [][2]string, body []byte) ([]byte, error)) Tracks {
	return Tracks{
		Send:   send,
		Market: "",
	}
}

func (s *Tracks) GetTrack(id string) (getTrack, error) {
	res, err := s.Send(lib.GET, "tracks/"+id, [][2]string{{"market", s.Market}}, []byte{})
	if err != nil {
		return getTrack{}, err
	}
	data := getTrack{}
	err = json.Unmarshal(res, &data)
	return data, err
}

func (s *Tracks) GetSeveralTracks(ids []string) (getTracks, error) {
	res, err := s.Send(lib.GET, "tracks", [][2]string{{"market", s.Market}, {"ids", strings.Join(ids, ",")}}, []byte{})
	if err != nil {
		return getTracks{}, err
	}
	data := getTracks{}
	err = json.Unmarshal(res, &data)
	return data, err
}

// Scopes: `ScopeUserLibraryRead`
func (s *Tracks) GetUsersSavedTracks(limit, offset int) (getUsersSavedTracks, error) {
	res, err := s.Send(lib.GET, "me/tracks", [][2]string{{"market", s.Market}, {"limit", strconv.Itoa(max(1, min(50, limit)))}, {"offset", strconv.Itoa(max(0, offset))}}, []byte{})
	if err != nil {
		return getUsersSavedTracks{}, err
	}
	data := getUsersSavedTracks{}
	err = json.Unmarshal(res, &data)
	return data, err
}

// Scopes: `ScopeUserLibraryModify`
func (s *Tracks) SaveTracksForCurrentUser(ids []string) error {
	body, err := json.Marshal(map[string]any{"ids": ids})
	if err != nil {
		return err
	}
	_, err = s.Send(lib.PUT, "me/tracks", [][2]string{}, body)
	return err
}

// Scopes: `ScopeUserLibraryModify`
func (s *Tracks) SaveTracksForCurrentUserTimestamped(ids []string, timestamp time.Time) error {
	bodyIds := []map[string]any{}
	for _, id := range ids {
		bodyIds = append(bodyIds, map[string]any{"id": id, "added_at": timestamp.Format(time.RFC3339)})
	}
	body, err := json.Marshal(map[string]any{"timestamped_ids": bodyIds})
	if err != nil {
		return err
	}
	_, err = s.Send(lib.PUT, "me/tracks", [][2]string{}, body)
	return err
}

// Scopes: `ScopeUserLibraryModify`
func (s *Tracks) RemoveUsersSavedTracks(ids []string) error {
	body, err := json.Marshal(map[string]any{"ids": ids})
	if err != nil {
		return err
	}
	_, err = s.Send(lib.DELETE, "me/tracks", [][2]string{}, body)
	return err
}

// Scopes: `ScopeUserLibraryRead`
func (s *Tracks) CheckUsersSavedTracks(ids []string) ([]bool, error) {
	res, err := s.Send(lib.GET, "me/tracks/contains", [][2]string{{"ids", strings.Join(ids, ",")}}, []byte{})
	if err != nil {
		return []bool{}, err
	}
	data := []bool{}
	err = json.Unmarshal(res, &data)
	return data, err
}
