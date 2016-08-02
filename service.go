package main

import (
	"io"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/naoina/toml"
)

var ServiceBase = filepath.Join(Base, "services")

type Service struct {
	Name   string
	Config *ServiceConfig

	cmd *exec.Cmd

	stop func() error
}

func LoadService(name string) (*Service, error) {
	path := filepath.Join(ServiceBase, name)
	startPath := filepath.Join(path, "start")
	stopPath := filepath.Join(path, "stop")
	configPath := filepath.Join(path, "config")

	svc := &Service{
		Name: name,
	}

	if start, err := os.Stat(startPath); (err != nil) || (start.Mode()&0111 == 0) {
		return nil, err
	}

	svc.stop = svc.defaultStop
	if stop, err := os.Stat(stopPath); err == nil {
		if stop.Mode()&0111 != 0 {
			svc.stop = svc.customStop
		}
	}

	if config, err := os.Open(configPath); err == nil {
		defer config.Close()

		svc.Config, err = LoadServiceConfig(config)
		if err != nil {
			return nil, err
		}
	}

	return svc, nil
}

func (svc *Service) defaultStop() error {
	panic("Not implemented.")
}

func (svc *Service) customStop() error {
	panic("Not implemented.")
}

func (svc *Service) Start() error {
	panic("Not implemented.")
}

func (svc *Service) Stop() error {
	return svc.stop()
}

type ServiceConfig struct {
	Deps []string
}

func LoadServiceConfig(r io.Reader) (*ServiceConfig, error) {
	d := toml.NewDecoder(r)

	var config ServiceConfig
	err := d.Decode(&config)
	return &config, err
}
