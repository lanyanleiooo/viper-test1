package config

type Sys struct {
	App *App `yaml:"app"`
	Log *Log `yaml:"log"`
}

type App struct {
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
}

type Log struct {
	Suffix  string `yaml:"suffix"`
	MaxSize int    `yaml:"maxSize"`
}
