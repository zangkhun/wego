package conf

type SysConf struct {
	IP    string `yaml:"ip"`
	Port  string `yaml:"port"`
	Debug bool   `yaml:"debug"`
}

type ServerConf struct {
	SysConf `yaml:"server"`
}
