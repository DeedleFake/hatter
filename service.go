package main

import (
	"os"
	"os/exec"
	"path/filepath"
)

var ServiceBase = filepath.Join(Base, "services")

type Service struct {
	Name   string
	Config *ServiceConfig

	cmd *exec.Cmd

	start func() error
	stop  func() error
}

func LoadService(name string) (*Service, error) {
	path := filepath.Join(ServiceBase, name)
	startPath := filepath.Join(path, "start")
	stopPath := filepath.Join(path, "stop")
	//configPath := filepath.Join(path, "config")

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

	return svc, nil
}

func (svc *Service) defaultStop() error {
	panic("Not implemented.")
}

func (svc *Service) customStop() error {
	panic("Not implemented.")
}

type ServiceConfig struct {
	Deps []string
}
