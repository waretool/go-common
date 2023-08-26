package domain

import (
	"github.com/waretool/go-common/domain"
	"github.com/waretool/go-common/env"
	"github.com/waretool/go-common/logger"
	"net/http"
)

type microservice struct {
	name   string
	server *http.Server
}

func NewMicroservice(name string, router domain.Router) domain.Microservice {
	srv := &http.Server{
		Addr:    ":" + env.GetEnv("PORT", "8080"),
		Handler: router,
	}

	return &microservice{
		name:   name,
		server: srv,
	}
}

func (m *microservice) Start() {
	err := m.server.ListenAndServe()
	if err != nil {
		logger.Fatalf("unable to start microservice due to: %s", err)
	}
	logger.Infof("microservice %s started", m.name)
}

func (m *microservice) GetName() string {
	return m.name
}

func (m *microservice) GetServer() *http.Server {
	return m.server
}
