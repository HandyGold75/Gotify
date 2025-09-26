package users

import (
	"encoding/json"
	"strconv"
	"strings"

	"github.com/HandyGold75/gotify/lib"
)

type (
	Users struct {
		Send     func(method lib.HTTPMethod, action string, options [][2]string, body []byte) ([]byte, error)
		DeviceID string
	}

	getCurrentUsersProfile lib.Profile

	getUsersTopArtists struct {
		lib.ItemsHeaders
		Items []lib.ArtistObject `json:"items"`
	}

	getUsersTopTracks struct {
		lib.ItemsHeaders
		Items []lib.TrackObject `json:"items"`
	}

	getUsersProfile lib.ProfilePublic

	getFollowedArtists struct {
		Artists struct {
			lib.ItemsCursorsHeaders
			Items []lib.ArtistObject `json:"items"`
		} `json:"artists"`
	}
)

func New(send func(method lib.HTTPMethod, action string, options [][2]string, body []byte) ([]byte, error)) Users {
	return Users{Send: send, DeviceID: ""}
}

// Scopes: `ScopeUserReadPrivate`, `ScopeUserReadEmail`
func (s *Users) GetCurrentUsersProfile() (getCurrentUsersProfile, error) {
	res, err := s.Send(lib.GET, "me", [][2]string{}, []byte{})
	if err != nil {
		return getCurrentUsersProfile{}, err
	}
	data := getCurrentUsersProfile{}
	err = json.Unmarshal(res, &data)
	return data, err
}

// Scopes: `ScopeUserTopRead`
func (s *Users) GetUsersTopArtists(time lib.TimeRange, limit, offset int) (getUsersTopArtists, error) {
	res, err := s.Send(lib.GET, "me/top/artists", [][2]string{{"time_range", string(time)}, {"limit", strconv.Itoa(max(1, min(50, limit)))}, {"offset", strconv.Itoa(max(0, offset))}}, []byte{})
	if err != nil {
		return getUsersTopArtists{}, err
	}
	data := getUsersTopArtists{}
	err = json.Unmarshal(res, &data)
	return data, err
}

// Scopes: `ScopeUserTopRead`
func (s *Users) GetUsersTopTracks(time lib.TimeRange, limit, offset int) (getUsersTopTracks, error) {
	res, err := s.Send(lib.GET, "me/top/tracks", [][2]string{{"time_range", string(time)}, {"limit", strconv.Itoa(max(1, min(50, limit)))}, {"offset", strconv.Itoa(max(0, offset))}}, []byte{})
	if err != nil {
		return getUsersTopTracks{}, err
	}
	data := getUsersTopTracks{}
	err = json.Unmarshal(res, &data)
	return data, err
}

func (s *Users) GetUsersProfile(id string) (getUsersProfile, error) {
	res, err := s.Send(lib.GET, "users/"+id, [][2]string{}, []byte{})
	if err != nil {
		return getUsersProfile{}, err
	}
	data := getUsersProfile{}
	err = json.Unmarshal(res, &data)
	return data, err
}

// Scopes: `ScopePlaylistModifyPublic`, `ScopePlaylistModifyPrivate`
func (s *Users) FollowPlaylist(id string, public bool) error {
	body, err := json.Marshal(map[string]any{"public": public})
	if err != nil {
		return err
	}
	_, err = s.Send(lib.PUT, "playlists/"+id+"/followers", [][2]string{}, body)
	return err
}

// Scopes: `ScopePlaylistModifyPublic`, `ScopePlaylistModifyPrivate`
func (s *Users) UnfollowPlaylist(id string) error {
	_, err := s.Send(lib.DELETE, "playlists/"+id+"/followers", [][2]string{}, []byte{})
	return err
}

// Scopes: `ScopeUserFollowRead`
func (s *Users) GetFollowedArtists(after string, limit int) (getFollowedArtists, error) {
	res, err := s.Send(lib.GET, "me/following", [][2]string{{"type", "artists"}, {"after", after}, {"limit", strconv.Itoa(max(1, min(50, limit)))}}, []byte{})
	if err != nil {
		return getFollowedArtists{}, err
	}
	data := getFollowedArtists{}
	err = json.Unmarshal(res, &data)
	return data, err
}

// Scopes: `ScopeUserFollowModify`
func (s *Users) FollowArtists(ids []string) error {
	body, err := json.Marshal(map[string]any{"ids": ids})
	if err != nil {
		return err
	}
	_, err = s.Send(lib.PUT, "me/following", [][2]string{{"type", "artist"}}, body)
	return err
}

// Scopes: `ScopeUserFollowModify`
func (s *Users) FollowUsers(ids []string) error {
	body, err := json.Marshal(map[string]any{"ids": ids})
	if err != nil {
		return err
	}
	_, err = s.Send(lib.PUT, "me/following", [][2]string{{"type", "user"}}, body)
	return err
}

// Scopes: `ScopeUserFollowModify`
func (s *Users) UnfollowArtists(ids []string) error {
	body, err := json.Marshal(map[string]any{"ids": ids})
	if err != nil {
		return err
	}
	_, err = s.Send(lib.DELETE, "me/following", [][2]string{{"type", "artist"}}, body)
	return err
}

// Scopes: `ScopeUserFollowModify`
func (s *Users) UnfollowUsers(ids []string) error {
	body, err := json.Marshal(map[string]any{"ids": ids})
	if err != nil {
		return err
	}
	_, err = s.Send(lib.DELETE, "me/following", [][2]string{{"type", "user"}}, body)
	return err
}

// Scopes: `ScopeUserFollowRead`
func (s *Users) CheckIfUserFollowsArtists(ids []string) ([]bool, error) {
	res, err := s.Send(lib.GET, "me/following/contains", [][2]string{{"type", "artist"}, {"ids", strings.Join(ids, ",")}}, []byte{})
	if err != nil {
		return []bool{}, err
	}
	data := []bool{}
	err = json.Unmarshal(res, &data)
	return data, err
}

// Scopes: `ScopeUserFollowRead`
func (s *Users) CheckIfUserFollowsUsers(ids []string) ([]bool, error) {
	res, err := s.Send(lib.GET, "me/following/contains", [][2]string{{"type", "user"}, {"ids", strings.Join(ids, ",")}}, []byte{})
	if err != nil {
		return []bool{}, err
	}
	data := []bool{}
	err = json.Unmarshal(res, &data)
	return data, err
}

func (s *Users) CheckIfCurrentUserFollowsPlaylist(id string) (bool, error) {
	res, err := s.Send(lib.GET, "playlists/"+id+"/followers/contains", [][2]string{}, []byte{})
	if err != nil {
		return false, err
	}
	data := []bool{}
	err = json.Unmarshal(res, &data)
	if len(data) < 1 {
		return false, lib.Errors.UnexpectedResponse
	}
	return data[0], err
}
