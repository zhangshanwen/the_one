package conf

type Conf struct {
	Port          string        `yaml:"port"`
	DB            DB            `yaml:"db"`
	Authorization Authorization `yaml:"authorization"`
	Level         string        `yaml:"level"`
}
