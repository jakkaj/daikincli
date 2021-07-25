package dcliconfig

//ConfigName is the struct that enforces config names to be set in code
type ConfigName string

//nolint
const (
	DAIKIN_URL      ConfigName = "DAIKIN_URL"
	DAIKIN_PASSWORD ConfigName = "DAIKIN_PASSWORD"
)

var testConfig = ``
