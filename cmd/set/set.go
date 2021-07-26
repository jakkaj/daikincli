package set

import (
	"daikincli/internal/dclilog"
	"daikincli/pkg/control"

	"github.com/spf13/cobra"
)

var (
	mode string

	setCmd = &cobra.Command{
		Use:   "set",
		Short: "Set values",
		Long:  "Use the various options to set values on the Daikin",
		RunE: func(cmd *cobra.Command, args []string) error {
			logger := dclilog.GetInstance()
			manager := control.NewManager(logger)

			state, err := manager.GetState()

			if err != nil {
				return err
			}

			return nil
		},
	}
)

// New initialises 'set' command
func New() *cobra.Command {
	setCmd.Flags().StringVarP(&mode, "mode", "m", "", "The mode to set the unit to - options are heat, cool, auto, fan")
	_ = setCmd.MarkFlagRequired("mode")

	return setCmd
}
