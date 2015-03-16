package postgres

import (
	"time"
	"xavier/storage"

	"github.com/jmoiron/sqlx"
)

type DeviceDatabase struct {
	*sqlx.DB
}

const (
	devicesAll = `
		SELECT token, environment, name, model, os, os_version, app_version, created_at, updated_at, user_id FROM devices d 
		WHERE d.user_id = $1`
	devicesFind = `
		SELECT token, environment, name, model, os, os_version, app_version, created_at, updated_at, user_id FROM devices d
		WHERE d.token = $1 AND d.user_id = $2 
		LIMIT 1`
	devicesCount = `
		SELECT COUNT(*) FROM devices d 
		WHERE d.token = $1 AND d.user_id = $2`
	devicesInsert = `
		INSERT INTO devices (token, environment, name, model, os, os_version, app_version, created_at, updated_at, user_id)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10) 
		RETURNING token, environment, name, model, os, os_version, app_version, created_at, updated_at, user_id`
	devicesUpdate = `
		UPDATE devices SET token=$1, environment=$2, name=$3, model=$4, os=$5, os_version=$6, app_version=$7, updated_at=$8, user_id=$9
		WHERE token = $10 AND user_id = $11
	`
)

func (db *DeviceDatabase) All(userID int) ([]storage.Device, error) {
	d := []storage.Device{}
	return d, db.Select(&d, devicesAll, userID)
}

func (db *DeviceDatabase) Find(token string, userID int) (*storage.Device, error) {
	d := &storage.Device{}
	return d, db.Get(d, devicesFind, token, userID)
}

func (db *DeviceDatabase) Exists(token string, userID int) (bool, error) {
	var count int
	if err := db.Get(&count, devicesCount, token, userID); err != nil {
		return false, err
	}
	return count == 1, nil
}

func (db *DeviceDatabase) InsertOrUpdate(deviceEntry *storage.DeviceEntry) (*storage.Device, bool, error) {
	exists, err := db.Exists(deviceEntry.Token, deviceEntry.UserID)
	if err != nil {
		return nil, false, err
	}

	d := &storage.Device{
		Token:       deviceEntry.Token,
		Environment: deviceEntry.Environment,
		Name:        deviceEntry.Name,
		Model:       deviceEntry.Model,
		Os:          deviceEntry.Os,
		AppVersion:  deviceEntry.AppVersion,
		UserID:      deviceEntry.UserID,
	}
	dateTime := time.Now()
	if exists {
		d.UpdatedAt = dateTime
		_, err := db.Query(devicesUpdate, d.Token, d.Environment, d.Name, d.Model, d.Os, d.OsVersion, d.AppVersion, dateTime, d.UserID, d.Token, d.UserID)
		return d, false, err
	}

	d.CreatedAt = dateTime
	d.UpdatedAt = dateTime

	return d, true, db.Get(d, devicesInsert, d.Token, d.Environment, d.Name, d.Model, d.Os, d.AppVersion, d.CreatedAt, d.UpdatedAt, deviceEntry.UserID)
}
