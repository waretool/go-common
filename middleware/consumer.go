//go:build test

package middleware

import "github.com/waretool/go-common/domain"

type FakeConsumer struct {
	Uuid    string
	Role    domain.Role
	Enabled bool
	Local   bool
}

func (c *FakeConsumer) GetIdentifier() string {
	return c.Uuid
}

func (c *FakeConsumer) GetRole() domain.Role {
	return c.Role
}

func (c *FakeConsumer) IsEnabled() bool {
	return c.Enabled
}

func (c *FakeConsumer) IsLocal() bool {
	return c.Local
}
