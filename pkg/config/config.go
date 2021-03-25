package config

type Config struct {
	WorkingDirectory   string
	ExcludeDirectories []string
}

func NewConfig(opts ...Option) *Config {
	c := &Config{}
	for _, o := range opts {
		o.Apply(c)
	}
	return c
}
