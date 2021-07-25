package control

import (
	"daikincli/internal/dcliconfig"
	"daikincli/internal/dclilog"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
)

type Manager struct {
	logger *dclilog.Logger
	config *dcliconfig.Config
}

type Settings struct {
	Mode     Mode
	Temp     int
	FanSpeed int
	Power    bool
}

func (s *Manager) get(endpoint string) (Settings, error) {

	baseUrl := s.config.GetString(dcliconfig.DAIKIN_URL)
	password := s.config.GetString(dcliconfig.DAIKIN_PASSWORD)

	url := fmt.Sprintf("%v/skyfi/aircon/%v?lpw=%v", baseUrl, endpoint, password)

	resp, err := http.Get(url)

	if err != nil {
		return Settings{}, err
	}

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		return Settings{}, err
	}
	strBody := string(body)

	pairs := make(map[string]string)

	for _, group := range strings.Split(strBody, ",") {
		nameVal := strings.Split(group, "=")
		pairs[nameVal[0]] = nameVal[1]
	}

	settings := Settings{}

	switch modeNumber := pairs["mode"]; modeNumber {

	case "0":
		settings.Mode = MODE_FAN
	case "1":
		settings.Mode = MODE_HEAT
	case "2":
		settings.Mode = MODE_COOL
	case "3":
		settings.Mode = MODE_AUTO

	}

	settings.FanSpeed, err = strconv.Atoi(pairs["f_rate"])

	switch settings.FanSpeed {
	case 3:
		settings.FanSpeed = 2
	case 5:
		settings.FanSpeed = 3
	}

	if err != nil {
		return Settings{}, err
	}

	settings.Temp, err = strconv.Atoi(pairs["stemp"])

	if err != nil {
		return Settings{}, err
	}

	settings.Power = pairs["pow"] == "1"

	return settings, nil

}

func (s *Manager) GetState() (Settings, error) {
	b, err := s.get("get_control_info")
	if err != nil {
		return Settings{}, err
	}

	return b, nil
}

func NewManager(logger *dclilog.Logger) *Manager {
	m := &Manager{
		logger: logger,
		config: dcliconfig.GetConfig(),
	}

	return m
}
