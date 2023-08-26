package domain

import "net/http"

type fakeConsumer struct {
	Uuid    string
	Role    Role
	Enabled bool
	Local   bool
}

func (c *fakeConsumer) GetIdentifier() string {
	return c.Uuid
}

func (c *fakeConsumer) GetRole() Role {
	return c.Role
}

func (c *fakeConsumer) IsEnabled() bool {
	return c.Enabled
}

func (c *fakeConsumer) IsLocal() bool {
	return c.Local
}

type fakeMicroservice struct {
	Name   string
	Server *http.Server
}

func (m *fakeMicroservice) Start() {}

func (m *fakeMicroservice) GetName() string {
	return m.Name
}

func (m *fakeMicroservice) GetServer() *http.Server {
	return m.Server
}
