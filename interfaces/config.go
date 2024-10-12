package interfaces

type Config struct {
	JWT struct {
		Secret string `yaml:"secret"`
	} `yaml:"jwt"`
}
