package redis

import (
	"errors"
	"xavier/storage"
)

type ClientCache struct {
	// *redis.Client
	Clients []storage.Client
}

func (c *ClientCache) Find(uuid string) (*storage.Client, error) {
	for _, client := range c.Clients {
		if client.UUID == uuid {
			return &client, nil
		}
	}
	return nil, errors.New("Client not found.")

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

func (c *ClientCache) Insert(clients []storage.Client) {
	c.Clients = clients
}
