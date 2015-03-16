package client

import (
	"net/http"
	"xavier/storage"
)

type ThemeService struct {
	client *Client
}

func (t *ThemeService) All() ([]storage.Theme, *http.Response, error) {
	request, err := t.client.NewRequest("GET", "themes", false, nil)
	if err != nil {
		return nil, nil, err
	}

	themes := []storage.Theme{}
	response, err := t.client.Do(request, "themes", &themes)
	if err != nil {
		return nil, response, err
	}

	return themes, response, nil
}

func (t *ThemeService) Find(uuid string) (*storage.Theme, *http.Response, error) {
	request, err := t.client.NewRequest("GET", "/themes/"+uuid, false, nil)
	if err != nil {
		return nil, nil, err
	}

	theme := &storage.Theme{}
	response, err := t.client.Do(request, "themes", theme)
	if err != nil {
		return nil, response, err
	}

	return theme, response, nil
}
