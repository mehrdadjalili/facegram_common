package models

import "time"

type (
	Info struct {
		Title       string `json:"title"`
		Body        string `json:"body"`
		IsEncrypted bool   `json:"is_encrypted"`
	}
	Service struct {
		Ip   string `json:"ip"`
		Name string `json:"name"`
	}
	Field struct {
		Type        string    `json:"type"` //warning error info
		Section     string    `json:"section"`
		Function    string    `json:"function"`
		Time        time.Time `json:"time"`
		Message     string    `json:"message"`
		Service     Service   `json:"service"`
		Information []Info    `json:"data"`
	}
)
