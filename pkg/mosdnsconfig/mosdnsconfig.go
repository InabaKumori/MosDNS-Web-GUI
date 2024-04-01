package mosdnsconfig

import (
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

type MosDNSConfig struct {
	Enabled        bool   `yaml:"enabled"`
	ConfigFile     string `yaml:"configfile"`
	ListenPort     int    `yaml:"listen_port"`
	LogLevel       string `yaml:"log_level"`
	LogFile        string `yaml:"log_file"`
	// ... other configuration options
}

func ReadConfig(configFile string) (*MosDNSConfig, error) {
	data, err := ioutil.ReadFile(configFile)
	if err != nil {
		return nil, err
	}

	config := &MosDNSConfig{}
	err = yaml.Unmarshal(data, config)
	if err != nil {
		return nil, err
	}

	return config, nil
}

func WriteConfig(configFile string, config *MosDNSConfig) error {
	data, err := yaml.Marshal(config)
	if err != nil {
		return err
	}

	return ioutil.WriteFile(configFile, data, 0644)
}
