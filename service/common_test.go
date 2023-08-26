package service

import "github.com/waretool/go-common/domain"

type fakeConsumer struct {
	Uuid    string
	Role    domain.Role
	Enabled bool
	Local   bool
}

func (c *fakeConsumer) GetIdentifier() string {
	return c.Uuid
}

func (c *fakeConsumer) GetRole() domain.Role {
	return c.Role
}

func (c *fakeConsumer) IsEnabled() bool {
	return c.Enabled
}

func (c *fakeConsumer) IsLocal() bool {
	return c.Local
}
