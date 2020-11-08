package manpage

import "github.com/yuin/goldmark/renderer"

type Config struct {
	Format uint
}

func NewConfig() Config {
	return Config{
		Format: SCDOC,
	}
}

func (c *Config) SetOption(name renderer.OptionName, value interface{}) {
	switch name {
	case optFormat:
		c.Format = value.(uint)
	}
}

const optFormat = "manpage_output_format"

const (
	SCDOC uint = iota
)

func WithOutputFormat(fmt uint) renderer.Option {
	return renderer.WithOption(optFormat, fmt)
}
