package echonest

import (
  "encoding/json"
  "fmt"
  "io/ioutil"
  "net/http"
  "net/url"
)

const (
  SONG_PROFILE_ENDPOINT = "https://developer.echonest.com/api/v4/song/profile"
)

type SongRequest struct {
  ApiKey string `json:"api_key"`
  Id string     `json:"id"`

  Buckets string `json:"bucket,omitempty"`
  format string `json:"format"`
}

type Song struct {
  Status string `json:"status"`
  Artist string `json:"artist_name"`
  AudioSummary struct {
    Tempo float64 `json:"tempo"`
  } `json:"audio_summary"`
  Title string `json:"title"`
  Tracks []struct {
    Catalog string `json:"catalog"`
    Id string `json:"foreign_id"`
  } `json:"tracks"`
}

type SongResponse struct {
  Response struct {
    Songs []Song `json:"songs"`
  } `json:"response"`
}

type EchoNestClient struct {
  ApiKey string
}

func NewEchoNestClient(apiKey string) EchoNestClient {
  client := EchoNestClient{
    apiKey,
  }

  return client
}

func (client EchoNestClient) GetSongs(trackIds []string, catalogs []string) ([]Song, error) {
  params := url.Values{}
  params.Add("api_key", client.ApiKey)
  params.Add("bucket", "audio_summary")
  params.Add("format", "json")

  if len(catalogs) != 0 {
    params.Add("bucket", "tracks")
    params.Add("limit", "true")
  }

  for _, catalog := range catalogs {
    params.Add("bucket", fmt.Sprintf("id:%s", catalog))
  }

  for _, trackId := range trackIds {
    params.Add("track_id", trackId)
  }

  endpoint := fmt.Sprintf("%s?%s", SONG_PROFILE_ENDPOINT, params.Encode())

  // TODO: handle error codes
  response, err := http.Get(endpoint)
  if err != nil {
    return []Song{}, err
  }


  body, err := ioutil.ReadAll(response.Body)
  if err != nil {
    return []Song{}, err
  }

  var songResponse SongResponse
  json.Unmarshal(body, &songResponse)

  return songResponse.Response.Songs, nil
}
