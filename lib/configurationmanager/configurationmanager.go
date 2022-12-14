package configurationmanager

type ConfigurationReader interface {
	Read(data []byte, target interface{}) error
}

type configurationManager struct {
	configurationReader ConfigurationReader
}

func NewconfigurationManager(configurationReader ConfigurationReader) configurationManager {
	return configurationManager{
		configurationReader: configurationReader,
	}
}

func (c configurationManager) GetConfiguration(data []byte, target interface{}) error {
	return c.configurationReader.Read(data, target)
}
