package control

//Mode contains all the options for modes
type Mode string

//nolint
const (
	MODE_HEAT Mode = "heat"
	MODE_COOL Mode = "cool"
	MODE_AUTO Mode = "auto"
	MODE_FAN  Mode = "fan"
)
