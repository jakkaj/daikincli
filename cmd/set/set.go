package set

import (
	"daikincli/internal/cli"
	"daikincli/internal/dclilog"
	"daikincli/pkg/control"
	"fmt"

	"github.com/spf13/cobra"
)

var (
	mode  string
	fan   string
	power string
	temp  string

	setCmd = &cobra.Command{
		Use:   "set",
		Short: "Set values",
		Long:  "Use the various options to set values on the Daikin",
		RunE: func(cmd *cobra.Command, args []string) error {
			logger := dclilog.GetInstance()
			manager := control.NewManager(logger)

			before, after, err := manager.SetState(temp, mode, fan, power)

			fmt.Println("Before")
			cli.RenderSettings(before)
			fmt.Println("After")
			cli.RenderSettings(after)

			if err != nil {
				return err
			}

			return nil
		},
	}
)

// New initialises 'set' command
func New() *cobra.Command {
	setCmd.Flags().StringVarP(&mode, "mode", "m", "", "mode options are heat, cool, auto, fan")
	setCmd.Flags().StringVarP(&fan, "fan", "f", "", "fan speed options are 1, 2 or 3")
	setCmd.Flags().StringVarP(&power, "power", "p", "", "power options are on or off")
	setCmd.Flags().StringVarP(&temp, "temp", "t", "", "temp options are up to you")

	return setCmd
}
