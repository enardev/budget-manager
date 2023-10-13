package configutil

import (
	"fmt"
	"os"
	"strings"

	"slices"

	configv2 "github.com/gookit/config/v2"
	"github.com/gookit/config/v2/yamlv3"
)

func LoadConfig() {
	profiles := getActiveProfiles()
	files := []string{"config/application.yaml"}
	for _, prof := range profiles {
		files = append(files, fmt.Sprintf("config/application-%s.yaml", prof))
	}
	configv2.WithOptions(configv2.ParseEnv)
	configv2.AddDriver(yamlv3.Driver)
	configv2.LoadFiles(files...)
}

func getActiveProfiles() []string {
	profiles := []string{}
	for _, arg := range os.Args {
		splitted := strings.Split(arg, "=")
		if splitted[0] == "--profiles" {
			profiles = append(profiles, strings.Split(splitted[1], ",")...)
		}
	}

	fromEnv := os.Getenv("SERVICE_ACTIVE_PROFILES")
	if fromEnv != "" {
		splitted := strings.Split(fromEnv, ",")
		for _, prof := range splitted {
			if !slices.Contains(profiles, prof) {
				profiles = append(profiles, prof)
			}
		}
	}

	return profiles
}
