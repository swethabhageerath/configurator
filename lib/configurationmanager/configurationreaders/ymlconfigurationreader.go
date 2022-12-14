package configurationreaders

import (
	"gopkg.in/yaml.v3"
)

type YamlConfigurationReader struct{}

func (y YamlConfigurationReader) Read(data []byte, target interface{}) error {
	err := yaml.Unmarshal([]byte(data), target)
	if err != nil {
		return err
	}
	return nil
}
