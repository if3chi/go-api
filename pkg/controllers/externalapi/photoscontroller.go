package externalapi

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"sync"

	"github.com/if3chi/go-api/pkg/kernel"
	"github.com/if3chi/go-api/pkg/services/photo"
)

func HandleGetPhotos(app *kernel.Application) http.HandlerFunc {
	service := photo.NewPhotoService()

	type photo struct {
		AlbumID      int    `json:"albumId"`
		ID           int    `json:"id"`
		Title        string `json:"title"`
		URL          string `json:"url"`
		ThumbnailURL string `json:"thumbnailUrl"`
	}

	return func(rw http.ResponseWriter, req *http.Request) {
		var (
			waitGroup sync.WaitGroup
			response  *http.Response
			apiErr    error
			photos    []photo
		)
		waitGroup.Add(1)
		go func() {
			response, apiErr = service.Get(
				"https://jsonplaceholder.typicode.com/photos",
				&waitGroup,
			)

			if apiErr != nil {
				app.Logger.Fatal(apiErr.Error())
				panic(apiErr)
			}
		}()
		waitGroup.Wait()

		body, readError := ioutil.ReadAll(response.Body)
		if readError != nil {
			app.Logger.Fatal(readError.Error())
			panic(readError)
		}

		if unmarshalErr := json.Unmarshal([]byte(body), &photos); unmarshalErr != nil {
			app.Logger.Fatal(unmarshalErr.Error())
			panic(unmarshalErr)
		}

		app.Respond(rw, req, photos, response.StatusCode)
	}
}
