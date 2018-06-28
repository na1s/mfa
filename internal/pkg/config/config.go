package config

type Configuration interface {
	GetDatabaseName() string
}

type DummyConfiguration struct {
}

func (r *DummyConfiguration) GetDatabaseName() string {
	return "mfa"
}
