package player

import (
	"encoding/json"
	"strconv"
	"time"

	"github.com/HandyGold75/gotify/lib"
)

type (
	Player struct {
		Send     func(method lib.HTTPMethod, action string, options [][2]string, body []byte) ([]byte, error)
		DeviceID string
		Market   string // An ISO 3166-1 alpha-2 country code, https://en.wikipedia.org/wiki/ISO_3166-1_alpha-2
	}

	getPlaybackState struct {
		Device       lib.Device `json:"device"`
		RepeatState  string     `json:"repeat_state"`
		ShuffleState bool       `json:"shuffle_state"`
		lib.Context
		Timestamp  int  `json:"timestamp"`
		ProgressMs int  `json:"progress_ms"`
		IsPlaying  bool `json:"is_playing"`
		Item       struct {
			lib.TrackObject
			lib.EpisodeObject
		} `json:"item"`
		CurrentlyPlayingType string      `json:"currently_playing_type"`
		Actions              lib.Actions `json:"actions"`
	}

	getAvailableDevices struct {
		Devices []lib.Device `json:"devices"`
	}

	getCurrentlyPlayingTrack getPlaybackState

	getRecentlyPlayedTracks struct {
		lib.ItemsCursorsHeaders
		Items []struct {
			lib.Track
			PlayedAt string `json:"played_at"`
			lib.Context
		} `json:"items"`
	}

	getTheUsersQueue struct {
		CurrentlyPlaying struct {
			lib.TrackObject
			lib.EpisodeObject
		} `json:"currently_playing"`
		Queue []struct {
			lib.TrackObject
			lib.EpisodeObject
		} `json:"queue"`
	}
)

func New(send func(method lib.HTTPMethod, action string, options [][2]string, body []byte) ([]byte, error)) Player {
	return Player{
		Send:     send,
		DeviceID: "", Market: "",
	}
}

// Scopes: `ScopeUserReadPlaybackState`
func (s *Player) GetPlaybackState() (getPlaybackState, error) {
	res, err := s.Send(lib.GET, "player", [][2]string{{"market", s.Market}}, []byte{})
	if err != nil {
		return getPlaybackState{}, err
	}
	data := getPlaybackState{}
	err = json.Unmarshal(res, &data)
	return data, err
}

// Requires premium.
//
// Scopes: `ScopeUserModifyPlaybackState`
func (s *Player) TransferPlayback(deviceID string, play bool) error {
	body, err := json.Marshal(map[string]any{"device_ids": deviceID, "play": play})
	if err != nil {
		return err
	}
	_, err = s.Send(lib.PUT, "player", [][2]string{}, body)
	return err
}

// Scopes: `ScopeUserReadPlaybackState`
func (s *Player) GetAvailableDevices() (getAvailableDevices, error) {
	res, err := s.Send(lib.GET, "player/devices", [][2]string{}, []byte{})
	if err != nil {
		return getAvailableDevices{}, err
	}
	data := getAvailableDevices{}
	err = json.Unmarshal(res, &data)
	return data, err
}

// Scopes: `ScopeUserReadCurrentlyPlaying`
func (s *Player) GetCurrentlyPlayingTrack() (getCurrentlyPlayingTrack, error) {
	res, err := s.Send(lib.GET, "player/currently-playing", [][2]string{{"market", s.Market}}, []byte{})
	if err != nil {
		return getCurrentlyPlayingTrack{}, err
	}
	data := getCurrentlyPlayingTrack{}
	err = json.Unmarshal(res, &data)
	return data, err
}

// Requires premium.
//
// Scopes: `ScopeUserModifyPlaybackState`
//
// Use `time.Duration(-1)` to disable this filter.
func (s *Player) StartResumePlayback(position time.Duration) error {
	value := ""
	if time.Duration(0) > position {
		value = strconv.Itoa(int(position.Milliseconds()))
	}
	body, err := json.Marshal(map[string]any{"position_ms": value})
	if err != nil {
		return err
	}
	_, err = s.Send(lib.PUT, "player/play", [][2]string{{"device_id", s.DeviceID}}, body)
	return err
}

// Requires premium.
//
// Scopes: `ScopeUserModifyPlaybackState`
//
// Body:
//
//	{
//	    "context_uri": "spotify:album:5ht7ItJgpBH7W6vJ5BqpPr",
//	    "uris": ["spotify:track:4iV5W9uYEdYUVa79Axb7Rh", "spotify:track:1301WleyT98MSxVHPZCA6M"],
//	    "offset": {
//	        "position": 5,
//	        "uri": "spotify:track:1301WleyT98MSxVHPZCA6M"
//	    },
//	    "position_ms": 0
//	}
func (s *Player) StartResumePlaybackRaw(req map[string]any) error {
	body, err := json.Marshal(req)
	if err != nil {
		return err
	}
	_, err = s.Send(lib.PUT, "player/play", [][2]string{{"device_id", s.DeviceID}}, body)
	return err
}

// Requires premium.
//
// Scopes: `ScopeUserModifyPlaybackState`
func (s *Player) PausePlayback() error {
	_, err := s.Send(lib.PUT, "player/pause", [][2]string{{"device_id", s.DeviceID}}, []byte{})
	return err
}

// Requires premium.
//
// Scopes: `ScopeUserModifyPlaybackState`
func (s *Player) SkipToNext() error {
	_, err := s.Send(lib.POST, "player/next", [][2]string{{"device_id", s.DeviceID}}, []byte{})
	return err
}

// Requires premium.
//
// Scopes: `ScopeUserModifyPlaybackState`
func (s *Player) SkipToPrevious() error {
	_, err := s.Send(lib.POST, "player/previous", [][2]string{{"device_id", s.DeviceID}}, []byte{})
	return err
}

// Requires premium.
//
// Scopes: `ScopeUserModifyPlaybackState`
func (s *Player) SeekToPosition(position time.Duration) error {
	_, err := s.Send(lib.PUT, "player/seek", [][2]string{{"device_id", s.DeviceID}, {"position_ms", strconv.Itoa(int(position.Milliseconds()))}}, []byte{})
	return err
}

// Requires premium.
//
// Scopes: `ScopeUserModifyPlaybackState`
func (s *Player) SetRepeatMode(state lib.RepeatMode) error {
	_, err := s.Send(lib.PUT, "player/repeat", [][2]string{{"device_id", s.DeviceID}, {"state", string(state)}}, []byte{})
	return err
}

// Requires premium.
//
// Scopes: `ScopeUserModifyPlaybackState`
func (s *Player) SetPlaybackVolume(volume int) error {
	_, err := s.Send(lib.PUT, "player/volume", [][2]string{{"device_id", s.DeviceID}, {"volume_percent", strconv.Itoa(max(0, min(100, volume)))}}, []byte{})
	return err
}

// Requires premium.
//
// Scopes: `ScopeUserModifyPlaybackState`
func (s *Player) TogglePlaybackShuffle(state bool) error {
	_, err := s.Send(lib.PUT, "player/shuffle", [][2]string{{"device_id", s.DeviceID}, {"state", strconv.FormatBool(state)}}, []byte{})
	return err
}

// Scopes: `ScopeUserReadRecentlyPlayed`
//
// Return items after stamp if after is true, otherwise returns items before time.
// Use `time.Time{}` to disable this filter.
func (s *Player) GetRecentlyPlayedTracks(limit int, stamp time.Time, after bool) (getRecentlyPlayedTracks, error) {
	key, value := "before", strconv.Itoa(int(stamp.Unix()))
	if stamp.Unix() == (time.Time{}.Unix()) {
		value = ""
	} else if after {
		key = "after"
	}
	res, err := s.Send(lib.GET, "player/recently-played", [][2]string{{"limit", strconv.Itoa(max(1, min(50, limit)))}, {key, value}}, []byte{})
	if err != nil {
		return getRecentlyPlayedTracks{}, err
	}
	data := getRecentlyPlayedTracks{}
	err = json.Unmarshal(res, &data)
	return data, err
}

// Scopes: `ScopeUserReadCurrentlyPlaying`, `ScopeUserReadPlaybackState`
func (s *Player) GetTheUsersQueue() (getTheUsersQueue, error) {
	res, err := s.Send(lib.GET, "player/queue", [][2]string{}, []byte{})
	if err != nil {
		return getTheUsersQueue{}, err
	}
	data := getTheUsersQueue{}
	err = json.Unmarshal(res, &data)
	return data, err
}

// Requires premium.
//
// Scopes: `ScopeUserModifyPlaybackState`
func (s *Player) AddItemToPlaybackQueue(uri lib.URI) error {
	_, err := s.Send(lib.POST, "player", [][2]string{{"device_id", s.DeviceID}, {"uri", string(uri)}}, []byte{})
	return err
}
