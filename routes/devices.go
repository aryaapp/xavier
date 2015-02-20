package routes

import (
	"xavier/app"
	"xavier/storage"
)

type DeviceParams struct {
	Environment string `json:"environment" validate:"nonzero"`
	Name        string `json:"name" validate:"nonzero"`
	Model       string `json:"model" validate:"nonzero"`
	Os          string `json:"os" validate:"nonzero"`
	OsVersion   string `json:"os_version" validate:"nonzero"`
	AppVersion  string `json:"app_version" validate:"nonzero"`
}

func UserDevicesIndex(c *app.Context) *app.Error {
	d, err := c.DeviceStorage.All(c.GetUserID())
	if err == nil {
		return c.JSON(200, "devices", d)
	}
	c.LogError(err)
	return &app.Error{404, "Devices could not be found."}
}

func UserDevicesUpdate(c *app.Context) *app.Error {
	var params DeviceParams
	if err := c.BindParamsAndValidate(&params); err != nil {
		c.LogError(err)
		return &app.Error{422, "Device could not be created/updated. Invalid parameters"}
	}

	token := c.URLParams.ByName("device")
	if len([]rune(token)) != 64 {
		return &app.Error{422, "Device could not be created/updated. Invalid token"}
	}

	device := &storage.Device{}
	device.Token = token
	device.Environment = params.Environment
	device.Name = params.Name
	device.Model = params.Model
	device.Os = params.Os
	device.OsVersion = params.OsVersion
	device.AppVersion = params.AppVersion
	device.UserID = c.GetUserID()

	inserted, err := c.DeviceStorage.InsertOrUpdate(device)
	if err != nil {
		c.LogError(err)
		return &app.Error{500, "Devices could not be created/updated."}
	}
	if inserted {
		return c.JSON(201, "devices", device)
	}
	return c.JSON(200, "devices", device)

}
