package set

import (
	"daikincli/internal/cli"
	"daikincli/internal/dclilog"
	"daikincli/pkg/control"
	"fmt"
	"strings"
	"time"

	"github.com/spf13/cobra"
)

var (
	mode  string
	fan   string
	power string
	temp  string
	br    string
	st    string

	setCmd = &cobra.Command{
		Use:   "set",
		Short: "Set values",
		Long:  "Use the various options to set values on the Daikin",
		RunE: func(cmd *cobra.Command, args []string) error {
			logger := dclilog.GetInstance()
			manager := control.NewManager(logger)

			if mode != "" || power != "" || fan != "" || temp == "" {
				mode = strings.ToLower(mode)
				power = strings.ToLower(power)

				err := manager.SetState(temp, mode, fan, power)

				if err != nil {
					return err
				}

			}

			if br != "" || st != "" {
				err := manager.SetZones(br == "on", st == "on")
				if err != nil {
					return err
				}
			}

			time.Sleep(2 * time.Second)

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
	setCmd.Flags().StringVarP(&mode, "mode", "m", "", "mode options are heat, cool, auto, fan")
	setCmd.Flags().StringVarP(&fan, "fan", "f", "", "fan speed options are 1, 2 or 3")
	setCmd.Flags().StringVarP(&power, "power", "p", "", "power options are on or off")
	setCmd.Flags().StringVarP(&temp, "temp", "t", "", "temp options are up to you")
	setCmd.Flags().StringVarP(&br, "bedroom", "b", "", "bedroom enabled (on/off)")
	setCmd.Flags().StringVarP(&st, "study", "s", "", "study enabled (on/off)")

	return setCmd
}
