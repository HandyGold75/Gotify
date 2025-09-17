package playlists

import (
	"encoding/base64"
	"encoding/json"
	"strconv"
	"strings"

	"github.com/HandyGold75/gotify/lib"
)

// TODO: All responses

type Playlists struct {
	Send   func(method lib.HTTPMethod, action string, options [][2]string, body []byte) ([]byte, error)
	Market string // An ISO 3166-1 alpha-2 country code, https://en.wikipedia.org/wiki/ISO_3166-1_alpha-2
}

func New(send func(method lib.HTTPMethod, action string, options [][2]string, body []byte) ([]byte, error)) Playlists {
	return Playlists{Send: send}
}

func (s *Playlists) GetPlaylist(id string, fields []string) error {
	_, err := s.Send(lib.GET, "playlists/"+id+"", [][2]string{{"market", s.Market}, {"fields", strings.Join(fields, ",")}}, []byte{})
	return err
}

// Scopes: `ScopePlaylistModifyPublic`, `ScopePlaylistModifyPrivate`
func (s *Playlists) ChangePlaylistDetails(id, name string, public, collaborative bool, description string) error {
	body, err := json.Marshal(map[string]any{"name": name, "public": public, "collaborative": collaborative, "description": description})
	if err != nil {
		return err
	}
	_, err = s.Send(lib.PUT, "playlists/"+id+"", [][2]string{}, body)
	return err
}

// Scopes: `ScopePlaylistReadPrivate`
func (s *Playlists) GetPlaylistItems(id string, fields []string, limit, offset int) error {
	_, err := s.Send(lib.GET, "playlists/"+id+"/tracks", [][2]string{{"market", s.Market}, {"fields", strings.Join(fields, ",")}, {"limit", strconv.Itoa(max(1, min(50, limit)))}, {"offset", strconv.Itoa(max(0, offset))}}, []byte{})
	return err
}

// Scopes: `ScopePlaylistModifyPublic`, `ScopePlaylistModifyPrivate`
func (s *Playlists) UpdatePlaylistItemsReoder(id string, start, before, length int, snapshot string) (string, error) {
	body, err := json.Marshal(map[string]any{"range_start": start, "insert_before": before, "range_length": length, "snapshot_id": snapshot})
	if err != nil {
		return "", err
	}
	res, err := s.Send(lib.PUT, "playlists/"+id+"/tracks", [][2]string{}, body)
	if err != nil {
		return "", err
	}
	data := struct {
		SnapshotID string `json:"snapshot_id"`
	}{}
	err = json.Unmarshal(res, &data)
	return data.SnapshotID, err
}

// Scopes: `ScopePlaylistModifyPublic`, `ScopePlaylistModifyPrivate`
func (s *Playlists) UpdatePlaylistItemsReplace(id string, uris []lib.URI) (string, error) {
	body, err := json.Marshal(map[string]any{"uris": uris})
	if err != nil {
		return "", err
	}
	res, err := s.Send(lib.PUT, "playlists/"+id+"/tracks", [][2]string{}, body)
	if err != nil {
		return "", err
	}
	data := struct {
		SnapshotID string `json:"snapshot_id"`
	}{}
	err = json.Unmarshal(res, &data)
	return data.SnapshotID, err
}

// Scopes: `ScopePlaylistModifyPublic`, `ScopePlaylistModifyPrivate`
func (s *Playlists) AddItemsToPlaylist(id string, uris []lib.URI, position int) (string, error) {
	body, err := json.Marshal(map[string]any{"uris": uris, "position": max(0, position)})
	if err != nil {
		return "", err
	}
	res, err := s.Send(lib.POST, "playlists/"+id+"/tracks", [][2]string{}, body)
	if err != nil {
		return "", err
	}
	data := struct {
		SnapshotID string `json:"snapshot_id"`
	}{}
	err = json.Unmarshal(res, &data)
	return data.SnapshotID, err
}

// Scopes: `ScopePlaylistModifyPublic`, `ScopePlaylistModifyPrivate`
func (s *Playlists) RemovePlaylistItems(id string, tracks []lib.URI, snapshot string) (string, error) {
	bodyTracks := []map[string]lib.URI{}
	for _, track := range tracks {
		bodyTracks = append(bodyTracks, map[string]lib.URI{"uri": track})
	}
	body, err := json.Marshal(map[string]any{"tracks": bodyTracks, "snapshot_id": snapshot})
	if err != nil {
		return "", err
	}
	res, err := s.Send(lib.DELETE, "playlists/"+id+"/tracks", [][2]string{}, body)
	if err != nil {
		return "", err
	}
	data := struct {
		SnapshotID string `json:"snapshot_id"`
	}{}
	err = json.Unmarshal(res, &data)
	return data.SnapshotID, err
}

// Scopes: `ScopePlaylistReadPrivate`
func (s *Playlists) GetCurrentUsersPlaylists(limit, offset int) error {
	_, err := s.Send(lib.GET, "me/playlists", [][2]string{{"limit", strconv.Itoa(max(1, min(50, limit)))}, {"offset", strconv.Itoa(max(0, offset))}}, []byte{})
	return err
}

// Scopes: `ScopePlaylistReadPrivate`, `ScopePlaylistReadCollaborative`
func (s *Playlists) GetUsersPlaylists(id string, limit, offset int) error {
	_, err := s.Send(lib.GET, "users/"+id+"/playlists", [][2]string{{"limit", strconv.Itoa(max(1, min(50, limit)))}, {"offset", strconv.Itoa(max(0, offset))}}, []byte{})
	return err
}

// Scopes: `ScopePlaylistModifyPublic`, `ScopePlaylistModifyPrivate`
func (s *Playlists) CreatePlaylist(id, name string, public, collaborative bool, description string) error {
	body, err := json.Marshal(map[string]any{"name": name, "public": public, "collaborative": collaborative, "description": description})
	if err != nil {
		return err
	}
	_, err = s.Send(lib.POST, "users/"+id+"/playlists", [][2]string{}, body)
	return err
}

func (s *Playlists) GetPlaylistCoverImage(id string) error {
	_, err := s.Send(lib.GET, "playlists/"+id+"/images", [][2]string{}, []byte{})
	return err
}

// Scopes: `ScopeUgcImageUpload`, `ScopePlaylistModifyPublic`, `ScopePlaylistModifyPrivate`
//
// `img` should be Base64 encoded JPEG image data, maximum payload size is 256 KB.
func (s *Playlists) AddCustomPlaylistCoverImage(id string, img string) error {
	body, err := base64.StdEncoding.DecodeString(img)
	if err != nil {
		return err
	}
	_, err = s.Send(lib.PUT, "playlists/"+id+"/images", [][2]string{}, body)
	return err
}
