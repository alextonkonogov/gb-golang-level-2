package config

import (
	"flag"
	"fmt"
)

var (
	path             = flag.String("path", ".", "a path for the app to find dublicates of files")
	workers          = flag.Int("workers", 5, "amount of workers")
	deleteDublicates = flag.Bool("delete", false, "delete the found dublicates?")
)

type AppConfig struct {
	Path             string
	Workers          int
	DeleteDublicates bool
}

func (c *AppConfig) Check() error {
	if c.Workers < 1 || c.Workers > 50 {
		return fmt.Errorf("Amount of workers is limited from 1 to 50")
	}

	return nil
}

func NewAppConfig() (*AppConfig, error) {
	flag.Parse()
	config := &AppConfig{*path, *workers, *deleteDublicates}
	return config, config.Check()
}
