package domain

import "net/http"

type Microservice interface {
	Start()
	GetName() string
	GetServer() *http.Server
}
