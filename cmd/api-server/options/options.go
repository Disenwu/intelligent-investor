package apiserver

import (
	"errors"
	"time"

	"github.com/spf13/pflag"
	utilerrors "k8s.io/apimachinery/pkg/util/errors"
)

type ServerOptions struct {
	JWTSecret  string        `json:"jwt-secret" mapstructure:"jwt-secret"`
	Expiration time.Duration `json:"expiration" mapstructure:"expiration"`
	Port       int           `json:"port" mapstructure:"port"`
}

func NewServerOptions() *ServerOptions {
	return &ServerOptions{
		JWTSecret:  "Rtg8BPKNEf2mB4mgvKONGPZZQSaJWNLijxR42qRgq0iBb5",
		Expiration: time.Hour * 12,
		Port:       9000,
	}
}

func (o *ServerOptions) AddFlags(fs *pflag.FlagSet) {
	fs.StringVar(&o.JWTSecret, "jwt-secret", o.JWTSecret, "JWT signing secret. Must be at least 6 characters long.")
	fs.DurationVar(&o.Expiration, "expiration", o.Expiration, "The expiration duration of JWT tokens.")
	fs.IntVar(&o.Port, "port", o.Port, "The port number on which the server will listen.")
}

func (o *ServerOptions) Validate() error {
	errs := []error{}
	if len(o.JWTSecret) < 6 {
		errs = append(errs, errors.New("--jwt-key must be at least 6 characters long"))
	}
	if o.Port <= 0 {
		errs = append(errs, errors.New("--port must be a positive integer"))
	}
	// 端口号必须在 0-65535 之间
	if o.Port < 0 || o.Port > 65535 {
		errs = append(errs, errors.New("--port must be between 0 and 65535"))
	}
	return utilerrors.NewAggregate(errs)
}
