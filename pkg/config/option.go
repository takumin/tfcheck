package config

type Option interface {
	Apply(*Config)
}

type WorkingDirectory string

func (o WorkingDirectory) Apply(c *Config) {
	c.WorkingDirectory = string(o)
}

type ExcludeDirectory string

func (o ExcludeDirectory) Apply(c *Config) {
	c.ExcludeDirectories = append(c.ExcludeDirectories, string(o))
}
