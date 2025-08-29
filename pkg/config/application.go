package config

import v "github.com/invopop/validation"

type (
	Application struct {
		Name string `mapstructure:"name"`
		Host string `mapstructure:"host"`
		Port int    `mapstructure:"port"`
	}
)

func (a Application) Validate() error {
	return v.ValidateStruct(&a,
		v.Field(&a.Name, v.Required),
		v.Field(&a.Host, v.Required),
		v.Field(&a.Port, v.Required),
	)
}
