package photo

import (
	"net/http"
	"sync"
	"time"
)

type PhotoService struct {
	Client *http.Client
}

func NewPhotoService() *PhotoService {
	return &PhotoService{
		Client: &http.Client{
			Timeout: 15 * time.Second,
		},
	}
}

func (photoService *PhotoService) Get(url string, waitGroup *sync.WaitGroup) (*http.Response, error) {
	response, err := photoService.Client.Get(url)

	waitGroup.Done()

	return response, err
}
