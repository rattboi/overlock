package registry

import (
	"context"

	"github.com/charmbracelet/log"
	"github.com/kndpio/kndp/internal/registry"
	"k8s.io/client-go/rest"
)

type deleteCmd struct {
	Name    string `required:"" help:"Registry name."`
	Default bool   `help:"Remove from default."`
	Local   bool   `help:"Remove associated local registry."`
}

func (c deleteCmd) Run(ctx context.Context, config *rest.Config, logger *log.Logger) error {
	reg := registry.Registry{}
	reg.Name = c.Name
	reg.SetDefault(c.Default)
	reg.SetLocal(c.Local)
	err := reg.Delete(ctx, config, logger)
	if err != nil {
		logger.Error(err)
	} else {
		logger.Info("Registry was removed successfully.")
	}
	return nil
}
