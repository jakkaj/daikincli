package zones

import (
	"daikincli/internal/cli"
	"daikincli/internal/dclilog"
	"daikincli/pkg/control"
	"strings"

	"github.com/spf13/cobra"
)

var (
	br string
	st string

	zoneCmd = &cobra.Command{
		Use:   "zone",
		Short: "Get and set Zones",
		Long:  "See the current state of zones",
		RunE: func(cmd *cobra.Command, args []string) error {
			logger := dclilog.GetInstance()
			manager := control.NewManager(logger)

			br = strings.ToLower(br)
			st = strings.ToLower(st)

			if br != "" || st != "" {
				err := manager.SetZones(br == "on", st == "on")
				if err != nil {
					return err
				}
			}

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
	zoneCmd.Flags().StringVarP(&br, "bedroom", "b", "", "bedroom enabled (on/off)")
	zoneCmd.Flags().StringVarP(&st, "study", "s", "", "study enabled (on/off)")

	return zoneCmd
}
