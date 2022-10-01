package main

import (
	"fmt"
	"reflect"
)

type Secret map[string]string
type Envs map[string]string

// All options that can be forwarded to gcloud deploy cmd
type Config struct {
	AllowUnauthenticated bool              `yaml:"allow-unauthenticated"`
	EntryPoint           string            `yaml:"entry-point"`
	Secrets              map[string]Secret `yaml:"secrets"`
	Gen2                 bool              `yaml:"gen2"`
	TriggerHttp          bool              `yaml:"trigger-http"`
	Authorization        string            `yaml:"authorization"`
	Runtime              string            `yaml:"runtime"`
	Envs                 Envs              `yaml:"envs"`
	Source               string            `yaml:"source"`
	VpcConnector         string            `yaml:"vpc-connector"`
	Timeout              int               `yaml:"timeout"`
	MaxInstances         int               `yaml:"max-instances"`
	MinInstances         int               `yaml:"min-instances"`
	Memory               int               `yaml:"memory"`
	ServiceAccount       string            `yaml:"service-account"`
	DockerRepository     string            `yaml:"docker-repository"`
	ServeAllOnLatest     bool              `yaml:"serve-all-traffic-latest-revision"`
	Region               string            `yaml:"region"`
}

// Return cmd line args and flags from the .yaml config file
func Command(config Config) string {
	var cmd string

	c_values := reflect.ValueOf(config)
	c_type := c_values.Type()

	for i := 0; i < c_values.NumField(); i++ {
		switch val := c_values.Field(i).Interface().(type) {
		case string:
			if val == "" {
				continue
			}

			cmd += fmt.Sprintf("--%s=%s ", c_type.Field(i).Tag.Get("yaml"), val)

		case bool:
			if !val {
				continue
			}

			cmd += fmt.Sprintf("--%s ", c_type.Field(i).Tag.Get("yaml"))

		case int:
			if val == 0 {
				continue
			}

			cmd += fmt.Sprintf("--%s=%v ", c_type.Field(i).Tag.Get("yaml"), val)

		case map[string]Secret:
			if len(val) == 0 {
				continue
			}

			arg := fmt.Sprintf("--%s ", c_type.Field(i).Tag.Get("yaml"))

			for k, v := range val {
				arg += fmt.Sprintf("%s=%s:%s ", v["env_var"], k, v["version"])
			}

			cmd += arg

		case Envs:
			if len(val) == 0 {
				continue
			}

			arg := fmt.Sprintf("--%s ", c_type.Field(i).Tag.Get("yaml"))

			for k, v := range val {
				arg += fmt.Sprintf("%s=%s ", k, v)
			}

			cmd += arg

		default:
			continue
		}
	}

	return cmd
}
