package domain

type Consumer interface {
	GetIdentifier() string
	GetRole() Role
	IsEnabled() bool
	IsLocal() bool
}
