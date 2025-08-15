# Gotify

This module aims to implement Spotify Web API bindings in an easy and simple while keeping advanced control possible.

The documentation from [Spotify Web API](https://developer.spotify.com/documentation/web-api) can be referenced for the base implementation.

Next to the base implementation some helper function are present for easier use.  
For all helper functions please refer to the files in the top level directory of this project.

Note that most of this project is untested, as such not all functions might not work as expected.  
Do you're own testing and see if the functionality you need actually works.

## Usage

Creating a GotifyPlayer for controlling a Spotify session can be done with this method:

```go
gp := Gotify.NewGotifyPlayer(
    "ClientID",                         // Use the ClientID of you're application.
    "http://127.0.0.1"                  // Use the redirect URL of you're application.
    gotify.ScopeUserModifyPlaybackState // Add all scopes you're application will use.
)
```

Authentication might be required depending on the functions you're applications will use.
During authentication the user is directed to Spotify for oauth2 authentication, after this the user is redirected to the redirect URL.
Authenticating the GotifyPlayer can be done by any of these methods:

```go
err := gp.AuthenticateHTTP(5050)   // Listen on the given port at uri /spotify_auth_callback for http calls, redirect URL should point here.
err := gp.AuthenticateStdin()      // Listen on stdin, user should manualy paste the post authentication URL here.
err := gp.AuthenticateToken(token) // Authenticate using oauth2 token.

// Get current oauth2 token, this token can be stored and used later for authentication.
//
// Note that while the client is active the token will automatically be refreshed, changing the token in the process.
token, err := gp.Token()
```

After a GotifyPlayer is successfully created and authenticated the associated Spotify session can be controlled.  
This can be done using either the Spotify references (base implementation) or the helpers.

The available Spotify references are:

- [gp.Albums](/Albums/Albums.go)
- [gp.Artists](/Artists/Artists.go)
- [gp.Audiobooks](/Audiobooks/Audiobooks.go)
- [gp.Categories](/Categories/Categories.go)
- [gp.Chapters](/Chapters/Chapters.go)
- [gp.Episodes](/Episodes/Episodes.go)
- [gp.Genres](/Genres/Genres.go)
- [gp.Markets](/Markets/Markets.go)
- [gp.Player](/Player/Player.go)
- [gp.Playlists](/Playlists/Playlists.go)
- [gp.Search](/Search/Search.go)
- [gp.Shows](/Shows/Shows.go)
- [gp.Tracks](/Tracks/Tracks.go)
- [gp.Users](/Users/Users.go)

## Examples

Some examples for controlling a Spotify session using the Spotify references:

```go
_ = gp.Albums.TODO() // TODO
_ = gp.Artists.TODO() // TODO
_ = gp.Audiobooks.TODO() // TODO
_ = gp.Categories.TODO() // TODO
_ = gp.Chapters.TODO() // TODO
_ = gp.Episodes.TODO() // TODO
_ = gp.Genres.TODO() // TODO
_ = gp.Markets.TODO() // TODO
_ = gp.Player.TODO() // TODO
_ = gp.Playlists.TODO() // TODO
_ = gp.Search.TODO() // TODO
_ = gp.Shows.TODO() // TODO
_ = gp.Tracks.TODO() // TODO
_ = gp.Users.TODO() // TODO
```

Some examples for controlling a Spotify session using the helpers:

```go
// TODO
```

## Structure

This project is structured as follows:

- Gotify ([/Gotify.go](/Gotify.go); Entrypoint and main functions to get started)
- lib ([/lib/lib.go](/lib/lib.go); Contains functions and variables that are used throughout the project)
- Spotify References ([/\*/\*.go](/Player/Player.go); Implements base as documented in [Spotify Web API](https://developer.spotify.com/documentation/web-api))
- Sonos Reference Helpers (Ex: [/\*.go](/Player.go); Build upon the base implementation for easier use)
