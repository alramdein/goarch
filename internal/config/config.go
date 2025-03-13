package config

type Config struct {
	Version    int           `yaml:"version"`
	OutputPath string        `yaml:"output_path"`
	Content    ContentConfig `yaml:"content"`
}

type ContentConfig struct {
	Entities []map[string]EntityConfig `yaml:"entities"`
	Layers   []string                  `yaml:"layers"`
}

type EntityConfig struct {
	Methods []MethodConfig `yaml:"methods"`
}

type MethodConfig struct {
	Name    string            `yaml:"name"`
	Params  map[string]string `yaml:"params,omitempty"`
	Returns map[string]string `yaml:"returns,omitempty"`
}
