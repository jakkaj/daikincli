package zones

import (
	"daikincli/internal/cli"
	"daikincli/internal/dclilog"
	"daikincli/pkg/control"

	"github.com/spf13/cobra"
)

var (
	zone string
	val  string

	zoneCmd = &cobra.Command{
		Use:   "zone",
		Short: "Get Zones",
		Long:  "See the current state of zones",
		RunE: func(cmd *cobra.Command, args []string) error {
			logger := dclilog.GetInstance()
			manager := control.NewManager(logger)

			zones, err := manager.GetZones()

			if err != nil {
				return err
			}

			cli.RenderZones(zones)

			return nil
		},
	}
)

// New initialises 'set' command
func New() *cobra.Command {
	zoneCmd.Flags().StringVarP(&zone, "zone", "z", "", "zone to change")
	zoneCmd.Flags().StringVarP(&val, "onoff", "o", "", "zone on or off boolean value")

	return zoneCmd
}
