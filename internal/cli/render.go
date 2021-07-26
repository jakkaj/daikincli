package cli

import (
	"daikincli/pkg/control"
	"fmt"
	"strings"

	"github.com/fatih/color"
	"github.com/kyokomi/emoji/v2"
)

func RenderZones(zones map[string]bool) {
	white := color.New(color.FgHiWhite).PrintfFunc()

	green := color.New(color.FgHiGreen, color.Bold).PrintfFunc()
	red := color.New(color.FgRed, color.Bold).PrintfFunc()

	for i, v := range zones {
		white(strings.ToLower(i))

		if v {
			green(" ON ")
		} else {
			red(" OFF ")
		}

	}
	fmt.Println()
}

func RenderSettings(state *control.Settings) {
	white := color.New(color.FgHiWhite).PrintfFunc()

	green := color.New(color.FgHiGreen, color.Bold).PrintfFunc()
	red := color.New(color.FgRed, color.Bold).PrintfFunc()
	blue := color.New(color.FgBlue, color.Bold).PrintfFunc()
	otherblue := color.New(color.FgHiBlue, color.Bold).PrintfFunc()
	yellow := color.New(color.FgYellow, color.Bold).PrintfFunc()

	if state.Power {
		green("ON ")
	} else {
		white("OFF ")
	}

	switch state.Mode {
	case control.MODE_AUTO:
		yellow(string(state.Mode))
	case control.MODE_COOL:
		otherblue(string(state.Mode))
	case control.MODE_FAN:
		blue(string(state.Mode))
	case control.MODE_HEAT:
		red(string(state.Mode))
	}

	fmt.Printf(" %v", state.Temp)

	emoji.Print(" :dash:")
	fmt.Printf("%d", state.FanSpeed)
	fmt.Println("")
}
