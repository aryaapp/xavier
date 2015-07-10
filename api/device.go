package api

import (
	"net/http"

	"github.com/labstack/echo"
)

func (a *AppContext) FindAllDevices(c *echo.Context) error {
	userID := c.Get("user.id").(int)
	d, err := a.DeviceStorage.FindAll(userID)
	if err == nil {
		return c.JSON(http.StatusOK, Data{d})
	}
	return c.JSON(http.StatusBadRequest, Error{"Devices could not be found."})
}

/*
func (a *AppContext) EditDevice(c *echo.Context) error {

	return c.JSON(http.StatusOk, Data{""})

	// var params DeviceParams
	// if err := c.BindParamsAndValidate(&params); err != nil {
	// 	c.LogError(err)
	// 	return &api.Error{422, "Device could not be created/updated. Invalid parameters"}
	// }

	// token := c.URLParams.ByName("device")
	// if len([]rune(token)) != 64 {
	// 	return &api.Error{422, "Device could not be created/updated. Invalid token"}
	// }

	// device, inserted, err := c.DeviceStorage.InsertOrUpdate(&storage.DeviceEntry{
	// 	Token:       token,
	// 	Environment: params.Environment,
	// 	Name:        params.Name,
	// 	Model:       params.Model,
	// 	Os:          params.Os,
	// 	OsVersion:   params.OsVersion,
	// 	AppVersion:  params.AppVersion,
	// 	UserID:      c.GetUserID(),
	// })
	// if err != nil {
	// 	c.LogError(err)
	// 	return &api.Error{500, "Devices could not be created/updated."}
	// }
	// if inserted {
	// 	return c.JSON(201, "devices", device)
	// }
	// return c.JSON(200, "devices", device)

}
*/
