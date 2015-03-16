package redis

import (
	"errors"
	"xavier/storage"
)

type AppCache struct {
	// *redis.Client
	Apps []storage.App
}

func (a *AppCache) Find(uuid string) (*storage.App, error) {
	for _, app := range a.Apps {
		if app.UUID == uuid {
			return &app, nil
		}
	}
	return nil, errors.New("App not found.")

}

// func (c *ClientCache) Find(uuid string) (*storage.Client, error) {
// 	hmap, err := c.Cmd("hgetall", "client:"+uuid).Hash()
// 	if err != nil {
// 		return nil, err
// 	}

// 	clientID, err := strconv.Atoi(hmap["Id"])
// 	if err != nil {
// 		return nil, err
// 	}

// 	client := &storage.Client{
// 		UUID:   uuid,
// 		ID:     clientID,
// 		Name:   hmap["Name"],
// 		Url:    hmap["Url"],
// 		Secret: hmap["Secret"],
// 	}
// 	return client, err
// }

func (a *AppCache) Insert(apps []storage.App) {
	a.Apps = apps
}
