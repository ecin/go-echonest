package echonest

import (
  "os"
  "testing"
)

func TestNewEchoNestClient(t *testing.T) {
  apiKey := "key"
  client := NewEchoNestClient(apiKey)

  if client.ApiKey != apiKey {
    t.Error("EchoNestClient constructor didn't assign API key")
  }
}

func TestGetSongs(t *testing.T) {
  apiKey := os.Getenv("ECHONEST_KEY")
  client := NewEchoNestClient(apiKey)

  ids := []string{
    "spotify:track:3GhntU8mCuMuW5NXHvaTOx",
    "spotify:track:6iejJ6Siz6lHcgcdsGNAaY",
    "spotify:track:1ko1hqVxyzvRlAsbklLIbV",
  }

  catalogs := []string{"spotify"}
  songs, _ := client.GetSongs(ids, catalogs)

  if len(songs) != len(ids) {
    t.Errorf("Expected %d songs to be returned.", len(ids))
  }

  songTests := []struct {
    title string
    artist string
    tempo float64
    id string
  }{
    {"Air", "Rogue", 148.056000, ids[0]},
    {"Prism", "Case & Point", 127.964, ids[1]},
    {"Duality", "Fractal", 159.964, ids[2]},
  }

  for index, test := range songTests {
    song := songs[index]
    title := song.Title
    artist := song.Artist
    tempo := song.AudioSummary.Tempo

    ids := make([]string, len(song.Tracks))
    for _, track := range song.Tracks {
      ids = append(ids, track.Id)
    }

    if title != test.title {
      t.Errorf("song.Title => %q, want %q", title, test.title)
    }

    if artist != test.artist {
      t.Errorf("song.Artist => %q, want %q", artist, test.artist)
    }

    if tempo != test.tempo {
      t.Errorf("song.Tempo => %q, want %q", tempo, test.tempo)
    }

    if !stringInSlice(test.id, ids) {
      t.Errorf("%q should be in song.Tracks %q", test.id, ids)
    }
  }
}

func stringInSlice(value string, strings []string) bool {
  for _, element := range strings {
    if value == element {
      return true
    }
  }

  return false
}
