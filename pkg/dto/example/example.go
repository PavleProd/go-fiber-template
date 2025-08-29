package example

import v "github.com/invopop/validation"

type (
	LogRequest struct {
		Message string `json:"message"`
	}
)

func (r LogRequest) Validate() error {
	return v.ValidateStruct(&r,
		v.Field(&r.Message, v.Required),
	)
}
