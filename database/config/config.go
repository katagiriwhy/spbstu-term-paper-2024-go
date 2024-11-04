package config

import (
	"os"
	"time"

	"gopkg.in/yaml.v2"
)

type Config struct {
	DB      DBConfiguration      `yaml:"DBSettings"`
	Program ProgramConfiguration `yaml:"ProgramSettings"`
}

type DBConfiguration struct {
	User     string `yaml:"dbUser"`
	Password string `yaml:"dbPassword"`
	Host     string `yaml:"dbHost"`
	Port     string `yaml:"dbPort"`
	Name     string `yaml:"dbName"`
}

type ProgramConfiguration struct {
	Address           string        `yaml:"bindAddress"`
	Port              string        `yaml:"port"`
	JiraUrl           string        `yaml:"jiraUrl"`
	ThreadCount       uint          `yaml:"threadCount"`
	IssueInOneRequest uint          `yaml:"issueInOneRequest"`
	MaxTimeSleep      time.Duration `yaml:"maxTimeSleep"`
	MinTimeSleep      time.Duration `yaml:"minTimeSleep"`
}

func ReadConfig() (*Config, error) {
	cfg := &Config{}
	file, err := os.Open("C:\\Users\\katagiri\\GolandProjects\\spbstu-term-paper-2024-go\\database\\config\\configDB.yaml")
	if err != nil {
		return nil, err
	}
	defer file.Close()

	decoder := yaml.NewDecoder(file)
	if err = decoder.Decode(cfg); err != nil {
		return nil, err
	}
	return cfg, nil
}
