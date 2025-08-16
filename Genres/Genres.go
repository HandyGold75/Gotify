package Genres

import (
	"github.com/HandyGold75/Gotify/lib"
)

type Genres struct {
	Send func(method lib.HttpMethod, action string, options [][2]string, body []byte) ([]byte, error)
}

func New(send func(method lib.HttpMethod, action string, options [][2]string, body []byte) ([]byte, error)) Genres {
	return Genres{Send: send}
}
