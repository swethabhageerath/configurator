package configurationmanager

type FileEncoding int

const (
	YAML FileEncoding = iota
	JSON
)

func (f FileEncoding) String() string {
	switch f {
	case YAML:
		return "Yaml"
	case JSON:
		return "Json"
	default:
		return "Yaml"
	}
}
