package configuration

type (
	ServiceConfig struct {
		ServicePort string
		ServiceHost string
	}
)

// NewServiceConfig is constructor
func NewServiceConfig(data map[string]string) *ServiceConfig {

	return &ServiceConfig{
		ServicePort: data[ServicePort],
		ServiceHost: data[ServiceHost],
	}
}
