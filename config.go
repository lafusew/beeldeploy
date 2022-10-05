package main

import (
	"fmt"
	"reflect"
)

type Secret map[string]string
type Envs map[string]string

// All options that can be forwarded to gcloud deploy cmd
type Config struct {
	AllowUnauthenticated bool              `yaml:"allow-unauthenticated" gcloud:"allow-unauthenticated"`
	DockerRegistery      string            `yaml:"docker-registry" gcloud:"docker-registry"`
	EgressSettings       string            `yaml:"egress" gcloud:"egress-settings"`
	EntryPoint           string            `yaml:"entry-point" glcoud:"entry-point"`
	Gen2                 bool              `yaml:"gen2" gcloud:"gen2"`
	IgnoreFile           string            `yaml:"ignore-file" gcloud:"ignore-file"`
	IngressSettings      string            `yaml:"ingress" gcloud:"ingress-settings"`
	Memory               int               `yaml:"memory" gcloud:"memory"`
	Retry                bool              `yaml:"retry" gcloud:"retry"`
	RunServiceAccount    string            `yaml:"run-service-account" gcloud:"run-service-account"`
	Runtime              string            `yaml:"runtime" gcloud:"runtime"`
	Secrets              map[string]Secret `yaml:"secrets" gcloud:"set-secrets"`
	TriggerHttp          bool              `yaml:"trigger-http" gcloud:"trigger-http"`
	Authorization        string            `yaml:"authorization" gcloud:"authorization"`
	Envs                 Envs              `yaml:"envs" gcloud:"set-env-vars"`
	Source               string            `yaml:"src" gcloud:"source"`
	VpcConnector         string            `yaml:"vpc" gcloud:"vpc-connector"`
	Timeout              int               `yaml:"timeout" gcloud:"timeout"`
	MaxInstances         int               `yaml:"max-instances" gcloud:"max-instances"`
	MinInstances         int               `yaml:"min-instances" gcloud:"min-instances"`
	ServiceAccount       string            `yaml:"service-account" gcloud:"service-account"`
	DockerRepository     string            `yaml:"docker-repository" gcloud:"docker-repository"`
	ServeAllOnLatest     bool              `yaml:"serve-all-traffic-latest-revision" gcloud:"serve-all-traffic-latest-revision"`
	Region               string            `yaml:"region" gcloud:"region"`
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

			cmd += fmt.Sprintf("--%s=%s ", c_type.Field(i).Tag.Get("gcloud"), val)

		case bool:
			if !val {
				continue
			}

			cmd += fmt.Sprintf("--%s ", c_type.Field(i).Tag.Get("gcloud"))

		case int:
			if val == 0 {
				continue
			}

			cmd += fmt.Sprintf("--%s=%v ", c_type.Field(i).Tag.Get("gcloud"), val)

		case map[string]Secret:
			if len(val) == 0 {
				continue
			}

			arg := fmt.Sprintf("--%s ", c_type.Field(i).Tag.Get("gcloud"))

			for k, v := range val {
				arg += fmt.Sprintf("%s=%s:%s ", v["env_var"], k, v["version"])
			}

			cmd += arg

		case Envs:
			if len(val) == 0 {
				continue
			}

			arg := fmt.Sprintf("--%s ", c_type.Field(i).Tag.Get("gcloud"))

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
