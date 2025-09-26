package lib

import (
	"errors"
)

type (
	HTTPMethod  string
	RepeatMode  string
	TimeRange   string
	AlbumGroup  string
	URIResource string

	URI string
)

const (
	GET    HTTPMethod = "GET"
	PUT    HTTPMethod = "PUT"
	POST   HTTPMethod = "POST"
	DELETE HTTPMethod = "DELETE"

	RepeatTrack   RepeatMode = "track"   // repeat the current track.
	RepeatContext RepeatMode = "context" // repeat the current context.
	RepeatOff     RepeatMode = "off"     // repeat off.

	TimeRangeLongTerm   TimeRange = "long_term"   // calculated from ~1 year of data and including all new data as it becomes available
	TimeRangeMediumTerm TimeRange = "medium_term" // approximately last 6 months.
	TimeRangeShortTerm  TimeRange = "short_term"  // approximately last 4 weeks.

	AlbumGroupAlbum       AlbumGroup = "album"
	AlbumGroupSingle      AlbumGroup = "single"
	AlbumGroupApearsOn    AlbumGroup = "apears_on"
	AlbumGroupCompilation AlbumGroup = "compilation"

	URIResourceTrack     URIResource = "track"
	URIResourceArtist    URIResource = "artist"
	URIResourceAlbum     URIResource = "album"
	URIResourcePlaylist  URIResource = "playlist"
	URIResourceShow      URIResource = "show"
	URIResourceEpisode   URIResource = "episode"
	URIResourceAudiobook URIResource = "audiobook"
	URIResourceUser      URIResource = "user"
)

var Errors = struct{ UnexpectedResponse error }{
	UnexpectedResponse: errors.New("unexpected response"),
}

func NewURI(resource URIResource, id string) URI {
	return URI("spotify:" + string(resource) + ":" + id)
}

type (
	Context struct {
		Context struct {
			Type string `json:"type"`
			Href string `json:"href"`
			externalUrls
			URI string `json:"uri"`
		} `json:"context"`
	}
	Cursors struct {
		Cursors struct {
			After  string `json:"after"`
			Before string `json:"before"`
		} `json:"cursors"`
	}
	ItemsHeaders struct {
		Href     string `json:"href"`
		Limit    int    `json:"limit"`
		Next     string `json:"next"`
		Offset   int    `json:"offset"`
		Previous string `json:"previous"`
		Total    int    `json:"total"`
	}
	ItemsCursorsHeaders struct {
		Href  string `json:"href"`
		Limit int    `json:"limit"`
		Next  string `json:"next"`
		Cursors
		Total int `json:"total"`
	}
	Profile struct {
		Country         string `json:"country"`
		DisplayName     string `json:"display_name"`
		Email           string `json:"email"`
		ExplicitContent struct {
			FilterEnabled bool `json:"filter_enabled"`
			FilterLocked  bool `json:"filter_locked"`
		} `json:"explicit_content"`
		externalUrls
		followers
		Href string `json:"href"`
		ID   string `json:"id"`
		images
		Product string `json:"product"`
		Type    string `json:"type"`
		URI     string `json:"uri"`
	}
	ProfilePublic struct {
		DisplayName string `json:"display_name"`
		externalUrls
		followers
		Href string `json:"href"`
		ID   string `json:"id"`
		images
		Type string `json:"type"`
		URI  string `json:"uri"`
	}
	Actions struct {
		InterruptingPlayback  bool `json:"interrupting_playback"`
		Pausing               bool `json:"pausing"`
		Resuming              bool `json:"resuming"`
		Seeking               bool `json:"seeking"`
		SkippingNext          bool `json:"skipping_next"`
		SkippingPrev          bool `json:"skipping_prev"`
		TogglingRepeatContext bool `json:"toggling_repeat_context"`
		TogglingShuffle       bool `json:"toggling_shuffle"`
		TogglingRepeatTrack   bool `json:"toggling_repeat_track"`
		TransferringPlayback  bool `json:"transferring_playback"`
	}
	Device struct {
		ID               string `json:"id"`
		IsActive         bool   `json:"is_active"`
		IsPrivateSession bool   `json:"is_private_session"`
		IsRestricted     bool   `json:"is_restricted"`
		Name             string `json:"name"`
		Type             string `json:"type"`
		VolumePercent    int    `json:"volume_percent"`
		SupportsVolume   bool   `json:"supports_volume"`
	}
	Categorie struct {
		Href  string `json:"href"`
		Icons []struct {
			URL    string `json:"url"`
			Height int    `json:"height"`
			Width  int    `json:"width"`
		} `json:"icons"`
		ID   string `json:"id"`
		Name string `json:"name"`
	}

	images struct {
		Images []struct {
			URL    string `json:"url"`
			Height int    `json:"height"`
			Width  int    `json:"width"`
		} `json:"images"`
	}
	owner struct {
		Owner struct {
			ExternalUrls struct {
				Spotify string `json:"spotify"`
			} `json:"external_urls"`
			Href        string `json:"href"`
			ID          string `json:"id"`
			Type        string `json:"type"`
			URI         string `json:"uri"`
			DisplayName string `json:"display_name"`
		} `json:"owner"`
	}
	restrictions struct {
		Restrictions struct {
			Reason string `json:"reason"`
		} `json:"restrictions"`
	}
	followers struct {
		Followers struct {
			Href  string `json:"href"`
			Total int    `json:"total"`
		} `json:"followers"`
	}
	externalUrls struct {
		ExternalUrls struct {
			Spotify string `json:"spotify"`
		} `json:"external_urls"`
	}
	copyrights []struct {
		Copyrights []struct {
			Text string `json:"text"`
			Type string `json:"type"`
		} `json:"copyrights"`
	}
	externalIds struct {
		ExternalIds struct {
			Isrc string `json:"isrc"`
			Ean  string `json:"ean"`
			Upc  string `json:"upc"`
		} `json:"external_ids"`
	}
	resumePoint struct {
		ResumePoint struct {
			FullyPlayed      bool `json:"fully_played"`
			ResumePositionMs int  `json:"resume_position_ms"`
		} `json:"resume_point"`
	}
	linkedFrom struct {
		LinkedFrom struct {
			externalUrls
			Href string `json:"href"`
			ID   string `json:"id"`
			Type string `json:"type"`
			URI  string `json:"uri"`
		} `json:"linked_from"`
	}
)

type (
	trackSimple struct {
		ArtistsSimple
		AvailableMarkets []string `json:"available_markets"`
		DiscNumber       int      `json:"disc_number"`
		DurationMs       int      `json:"duration_ms"`
		Explicit         bool     `json:"explicit"`
		externalUrls
		Href       string `json:"href"`
		ID         string `json:"id"`
		IsPlayable bool   `json:"is_playable"`
		linkedFrom
		restrictions
		Name        string `json:"name"`
		PreviewURL  string `json:"preview_url"`
		TrackNumber int    `json:"track_number"`
		Type        string `json:"type"`
		URI         string `json:"uri"`
		IsLocal     bool   `json:"is_local"`
	}
	TrackSimpleObject trackSimple

	track struct {
		AlbumSimple
		ArtistsSimple
		AvailableMarkets []string `json:"available_markets"`
		DiscNumber       int      `json:"disc_number"`
		DurationMs       int      `json:"duration_ms"`
		Explicit         bool     `json:"explicit"`
		externalIds
		externalUrls
		Href       string `json:"href"`
		ID         string `json:"id"`
		IsPlayable bool   `json:"is_playable"`
		linkedFrom
		restrictions
		Name        string `json:"name"`
		Popularity  int    `json:"popularity"`
		PreviewURL  string `json:"preview_url"`
		TrackNumber int    `json:"track_number"`
		Type        string `json:"type"`
		URI         string `json:"uri"`
		IsLocal     bool   `json:"is_local"`
	}
	Track struct {
		Track track `json:"track"`
	}
	Tracks struct {
		Tracks []track `json:"tracks"`
	}
	TrackObject track

	artistSimple struct {
		externalUrls
		Href string `json:"href"`
		ID   string `json:"id"`
		Name string `json:"name"`
		Type string `json:"type"`
		URI  string `json:"uri"`
	}
	ArtistsSimple struct {
		Artists []artistSimple `json:"artists"`
	}

	artist struct {
		externalUrls
		followers
		Genres []string `json:"genres"`
		Href   string   `json:"href"`
		ID     string   `json:"id"`
		images
		Name       string `json:"name"`
		Popularity int    `json:"popularity"`
		Type       string `json:"type"`
		URI        string `json:"uri"`
	}
	Artists struct {
		Artists []artist `json:"artists"`
	}
	ArtistObject artist

	albumSimple struct {
		AlbumType        string   `json:"album_type"`
		TotalTracks      int      `json:"total_tracks"`
		AvailableMarkets []string `json:"available_markets"`
		externalUrls
		Href string `json:"href"`
		ID   string `json:"id"`
		images
		Name                 string `json:"name"`
		ReleaseDate          string `json:"release_date"`
		ReleaseDatePrecision string `json:"release_date_precision"`
		restrictions
		Type string `json:"type"`
		URI  string `json:"uri"`
		ArtistsSimple
		AlbumGroup string `json:"album_group"`
	}
	AlbumSimple struct {
		Album albumSimple `json:"album"`
	}
	AlbumSimpleObject albumSimple

	album struct {
		AlbumType        string   `json:"album_type"`
		TotalTracks      int      `json:"total_tracks"`
		AvailableMarkets []string `json:"available_markets"`
		externalUrls
		Href string `json:"href"`
		ID   string `json:"id"`
		images
		Name                 string `json:"name"`
		ReleaseDate          string `json:"release_date"`
		ReleaseDatePrecision string `json:"release_date_precision"`
		restrictions
		Type string `json:"type"`
		URI  string `json:"uri"`
		ArtistsSimple
		Tracks struct {
			ItemsHeaders
			Items []trackSimple `json:"items"`
		} `json:"tracks"`
		copyrights
		externalIds
		Genres     []string `json:"genres"`
		Label      string   `json:"label"`
		Popularity int      `json:"popularity"`
	}
	Album struct {
		Album album `json:"album"`
	}
	Albums struct {
		Albums []album `json:"albums"`
	}

	playlistSimple struct {
		Collaborative bool   `json:"collaborative"`
		Description   string `json:"description"`
		externalUrls
		Href string `json:"href"`
		ID   string `json:"id"`
		images
		owner
		Public     bool   `json:"public"`
		SnapshotID string `json:"snapshot_id"`
		Tracks     struct {
			Href  string `json:"href"`
			Total int    `json:"total"`
		} `json:"tracks"`
		Type string `json:"type"`
		URI  string `json:"uri"`
	}
	PlaylistSimpleObject playlistSimple

	showSimple struct {
		AvailableMarkets []string `json:"available_markets"`
		copyrights
		Description     string `json:"description"`
		HTMLDescription string `json:"html_description"`
		Explicit        bool   `json:"explicit"`
		externalUrls
		Href string `json:"href"`
		ID   string `json:"id"`
		images
		IsExternallyHosted bool     `json:"is_externally_hosted"`
		Languages          []string `json:"languages"`
		MediaType          string   `json:"media_type"`
		Name               string   `json:"name"`
		Publisher          string   `json:"publisher"`
		Type               string   `json:"type"`
		URI                string   `json:"uri"`
		TotalEpisodes      int      `json:"total_episodes"`
	}
	ShowSimple struct {
		Show showSimple `json:"show"`
	}
	ShowSimpleObject showSimple

	episodeSimple struct {
		AudioPreviewURL string `json:"audio_preview_url"`
		Description     string `json:"description"`
		HTMLDescription string `json:"html_description"`
		DurationMs      int    `json:"duration_ms"`
		Explicit        bool   `json:"explicit"`
		externalUrls
		Href string `json:"href"`
		ID   string `json:"id"`
		images
		IsExternallyHosted   bool     `json:"is_externally_hosted"`
		IsPlayable           bool     `json:"is_playable"`
		Language             string   `json:"language"`
		Languages            []string `json:"languages"`
		Name                 string   `json:"name"`
		ReleaseDate          string   `json:"release_date"`
		ReleaseDatePrecision string   `json:"release_date_precision"`
		resumePoint
		Type string `json:"type"`
		URI  string `json:"uri"`
		restrictions
	}
	EpisodeSimpleObject episodeSimple

	episode struct {
		AudioPreviewURL string `json:"audio_preview_url"`
		Description     string `json:"description"`
		HTMLDescription string `json:"html_description"`
		DurationMs      int    `json:"duration_ms"`
		Explicit        bool   `json:"explicit"`
		externalUrls
		Href string `json:"href"`
		ID   string `json:"id"`
		images
		IsExternallyHosted   bool     `json:"is_externally_hosted"`
		IsPlayable           bool     `json:"is_playable"`
		Language             string   `json:"language"`
		Languages            []string `json:"languages"`
		Name                 string   `json:"name"`
		ReleaseDate          string   `json:"release_date"`
		ReleaseDatePrecision string   `json:"release_date_precision"`
		resumePoint
		Type string `json:"type"`
		URI  string `json:"uri"`
		restrictions
		ShowSimple
	}
	Episode struct {
		Episode episode `json:"episode"`
	}
	Episodes struct {
		Episodes []episode `json:"episodes"`
	}
	EpisodeObject episode

	audiobookSimple struct {
		Authors []struct {
			Name string `json:"name"`
		} `json:"authors"`
		AvailableMarkets []string `json:"available_markets"`
		copyrights
		Description     string `json:"description"`
		HTMLDescription string `json:"html_description"`
		Edition         string `json:"edition"`
		Explicit        bool   `json:"explicit"`
		externalUrls
		Href string `json:"href"`
		ID   string `json:"id"`
		images
		Languages []string `json:"languages"`
		MediaType string   `json:"media_type"`
		Name      string   `json:"name"`
		Narrators []struct {
			Name string `json:"name"`
		} `json:"narrators"`
		Publisher     string `json:"publisher"`
		Type          string `json:"type"`
		URI           string `json:"uri"`
		TotalChapters int    `json:"total_chapters"`
	}
	AudiobookSimple struct {
		Audiobook audiobookSimple `json:"audiobook"`
	}
	AudiobookSimpleObject audiobookSimple

	audiobook struct {
		Authors []struct {
			Name string `json:"name"`
		} `json:"authors"`
		AvailableMarkets []string `json:"available_markets"`
		copyrights
		Description     string `json:"description"`
		HTMLDescription string `json:"html_description"`
		Edition         string `json:"edition"`
		Explicit        bool   `json:"explicit"`
		externalUrls
		Href string `json:"href"`
		ID   string `json:"id"`
		images
		Languages []string `json:"languages"`
		MediaType string   `json:"media_type"`
		Name      string   `json:"name"`
		Narrators []struct {
			Name string `json:"name"`
		} `json:"narrators"`
		Publisher     string `json:"publisher"`
		Type          string `json:"type"`
		URI           string `json:"uri"`
		TotalChapters int    `json:"total_chapters"`
		Chapters      struct {
			ItemsHeaders
			Items []chapterSimple `json:"items"`
		} `json:"chapters"`
	}
	Audiobooks struct {
		Audiobooks []audiobook `json:"audiobooks"`
	}
	AudiobookObject audiobook

	chapterSimple struct {
		AudioPreviewURL  string   `json:"audio_preview_url"`
		AvailableMarkets []string `json:"available_markets"`
		ChapterNumber    int      `json:"chapter_number"`
		Description      string   `json:"description"`
		HTMLDescription  string   `json:"html_description"`
		DurationMs       int      `json:"duration_ms"`
		Explicit         bool     `json:"explicit"`
		externalUrls
		Href string `json:"href"`
		ID   string `json:"id"`
		images
		IsPlayable           bool     `json:"is_playable"`
		Languages            []string `json:"languages"`
		Name                 string   `json:"name"`
		ReleaseDate          string   `json:"release_date"`
		ReleaseDatePrecision string   `json:"release_date_precision"`
		resumePoint
		Type string `json:"type"`
		URI  string `json:"uri"`
		restrictions
	}
	ChapterSimpleObject chapterSimple

	chapter struct {
		AudioPreviewURL  string   `json:"audio_preview_url"`
		AvailableMarkets []string `json:"available_markets"`
		ChapterNumber    int      `json:"chapter_number"`
		Description      string   `json:"description"`
		HTMLDescription  string   `json:"html_description"`
		DurationMs       int      `json:"duration_ms"`
		Explicit         bool     `json:"explicit"`
		externalUrls
		Href string `json:"href"`
		ID   string `json:"id"`
		images
		IsPlayable           bool     `json:"is_playable"`
		Languages            []string `json:"languages"`
		Name                 string   `json:"name"`
		ReleaseDate          string   `json:"release_date"`
		ReleaseDatePrecision string   `json:"release_date_precision"`
		resumePoint
		Type string `json:"type"`
		URI  string `json:"uri"`
		restrictions
		AudiobookSimple
	}
	Chapters struct {
		Chapters []chapter `json:"chapters"`
	}
	ChapterObject chapter
)
