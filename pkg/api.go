package pkg

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"sync"
)

type Artist struct {
	Id             int                 `json:"id"`
	Name           string              `json:"name"`
	Members        []string            `json:"members"`
	Image          string              `json:"image"`
	CreationDate   int                 `json:"creationDate"`
	FirstAlbum     string              `json:"firstAlbum"`
	Relations      string              `json:"relations"`
	DatesLocations map[string][]string `json:"-"`
}

func (a *Artist) Clone() *Artist {
	clone := *a
	clone.Members = make([]string, len(a.Members))
	copy(clone.Members, a.Members)

	clone.DatesLocations = make(map[string][]string)
	for k, v := range a.DatesLocations {
		clone.DatesLocations[k] = append([]string{}, v...)
	}

	return &clone
}

var instance *APIClient
var once sync.Once

type APIClient struct {
	artists []Artist
}

func GetAPIClient() *APIClient {
	once.Do(func() {
		instance = &APIClient{}
		instance.fetchArtists()
	})
	return instance
}

func (client *APIClient) fetchArtists() error {
	response, err := http.Get("https://groupietrackers.herokuapp.com/api/artists")
	if err != nil {
		return err
	}
	defer response.Body.Close()
	return json.NewDecoder(response.Body).Decode(&client.artists)
}

func GetAPI() error {
	response, err := http.Get("https://groupietrackers.herokuapp.com/api/artists")
	if err != nil {
		return err
	}
	defer response.Body.Close()

	if err := json.NewDecoder(response.Body).Decode(&artists); err != nil {
		return err
	}

	return nil
}

type ArtistFacade struct {
	client *APIClient
}

func NewArtistFacade() *ArtistFacade {
	return &ArtistFacade{
		client: GetAPIClient(),
	}
}

func (af *ArtistFacade) GetCompleteArtistInfo(id int) (*Artist, error) {
	if id < 1 || id > len(af.client.artists) {
		return nil, fmt.Errorf("artist not found")
	}
	artist := &af.client.artists[id-1]
	err := fetchRelationsForArtist(artist)
	return artist, err
}

func fetchRelationsForArtist(artist *Artist) error {
	resp, err := http.Get(artist.Relations)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	type RelationsResponse struct {
		ID             int                 `json:"id"`
		DatesLocations map[string][]string `json:"datesLocations"`
	}
	var relationsResp RelationsResponse
	if err := json.NewDecoder(resp.Body).Decode(&relationsResp); err != nil {
		log.Printf("Error decoding relations data: %v", err)
		return err
	}

	artist.DatesLocations = relationsResp.DatesLocations

	return nil
}
