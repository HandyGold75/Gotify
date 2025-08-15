package Gotify

import (
	"time"

	"github.com/HandyGold75/GOLib/gotify/lib"
)

func (gp *GotifyPlayer) Play() error                       { return gp.Player.StartResumePlayback(time.Duration(-1)) }
func (gp *GotifyPlayer) Pause() error                      { return gp.Player.PausePlayback() }
func (gp *GotifyPlayer) Next() error                       { return gp.Player.SkipToNext() }
func (gp *GotifyPlayer) Previous() error                   { return gp.Player.SkipToPrevious() }
func (gp *GotifyPlayer) Seek(position time.Duration) error { return gp.Player.SeekToPosition(position) }
func (gp *GotifyPlayer) Repeat(state lib.RepeatMode) error { return gp.Player.SetRepeatMode(state) }
func (gp *GotifyPlayer) Volume(volume int) error           { return gp.Player.SetPlaybackVolume(volume) }
func (gp *GotifyPlayer) Shuffle(state bool) error          { return gp.Player.TogglePlaybackShuffle(state) }
