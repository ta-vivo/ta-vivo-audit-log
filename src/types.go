package main

import (
	"encoding/json"
	"net/http"
	"time"

	"gorm.io/datatypes"
)

type Middleware func(http.HandlerFunc) http.HandlerFunc

type Log struct {
	ID        uint           `json:"id" gorm:"primaryKey"`
	UserId    uint           `json:"userId"`
	Action    string         `json:"action"`
	MetaData  datatypes.JSON `json:"metaData"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (l *Log) ToJson() ([]byte, error) {
	return json.Marshal(l)
}
