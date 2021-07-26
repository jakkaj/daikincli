package get

import (
	"daikincli/internal/cli"
	"daikincli/internal/dclilog"
	"daikincli/pkg/control"
	"fmt"

	"github.com/spf13/cobra"
)

var (
	getCmd = &cobra.Command{
		Use:   "get",
		Short: "Get current settings from the controller",
		Long:  "Use the various options to set values on the Daikin",
		RunE: func(cmd *cobra.Command, args []string) error {
			logger := dclilog.GetInstance()
			manager := control.NewManager(logger)

			state, err := manager.GetState()

			if err != nil {
				return fmt.Errorf("could not read controller state: %w", err)
			}

			zone, err := manager.GetZones()

			if err != nil {
				return fmt.Errorf("could not read controller zone: %w", err)
			}

			cli.RenderSettings(state)

			cli.RenderZones(zone)

			return nil
		},
	}
)

// New initialises 'set' command
func New() *cobra.Command {

	return getCmd
}
