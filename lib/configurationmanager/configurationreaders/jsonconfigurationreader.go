package configurationreaders

import "encoding/json"

type JsonConfigurationReader struct{}

func (j JsonConfigurationReader) Parse(data []byte, target interface{}) error {
	err := json.Unmarshal(data, target)
	if err != nil {
		return err
	}
	return nil
}
