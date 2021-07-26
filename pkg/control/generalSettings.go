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

func (s *Manager) get(endpoint string) (string, error) {

	baseUrl := s.config.GetString(dcliconfig.DAIKIN_URL)
	password := s.config.GetString(dcliconfig.DAIKIN_PASSWORD)

	url := fmt.Sprintf("%v/skyfi/aircon/%v?lpw=%v", baseUrl, endpoint, password)

	resp, err := http.Get(url)

	if err != nil {
		return "", err
	}

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		return "", err
	}
	strBody := string(body)

	return strBody, nil

}

func (s *Manager) set(endpoint string, qs string) error {

	baseUrl := s.config.GetString(dcliconfig.DAIKIN_URL)
	password := s.config.GetString(dcliconfig.DAIKIN_PASSWORD)

	url := fmt.Sprintf("%v/skyfi/aircon/%v?lpw=%v%v", baseUrl, endpoint, password, qs)

	_, err := http.Get(url)

	if err != nil {
		return err
	}

	return nil

}

func (s *Manager) SetState(temp string, mode string, fanSpeed string, power string) (err error) {

	if temp == "" && mode == "" && fanSpeed == "" && power == "" {
		return
	}

	b, err := s.get("get_control_info")

	if err != nil {
		return
	}

	pairs := make(map[string]string)

	for _, group := range strings.Split(b, ",") {
		nameVal := strings.Split(group, "=")
		pairs[nameVal[0]] = nameVal[1]
	}

	if temp != "" {
		pairs["stemp"] = temp
	}

	if power != "" {
		if power == "on" {
			pairs["pow"] = "1"
		} else {
			pairs["pow"] = "0"
		}

	}

	switch fanSpeed {
	case "1":
		pairs["f_rate"] = "1"
	case "2":
		pairs["f_rate"] = "3"
	case "3":
		pairs["f_rate"] = "5"
	}

	switch Mode(mode) {
	case MODE_FAN:
		pairs["mode"] = "0"

	case MODE_COOL:
		pairs["mode"] = "2"

	case MODE_HEAT:
		pairs["mode"] = "1"

	case MODE_AUTO:
		pairs["mode"] = "3"

	}

	qs := ""

	for pair, val := range pairs {
		qs = fmt.Sprintf("%v&%v=%v", qs, pair, val)
	}

	err = s.set("set_control_info", qs)

	return
}

func (s *Manager) GetState() (*Settings, error) {
	b, err := s.get("get_control_info")

	if err != nil {
		return &Settings{}, err
	}

	settings, err := s.parseState(b)

	return settings, err

}

func (s *Manager) parseState(b string) (*Settings, error) {
	pairs := make(map[string]string)

	for _, group := range strings.Split(b, ",") {
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

	var err error

	settings.FanSpeed, err = strconv.Atoi(pairs["f_rate"])

	switch settings.FanSpeed {
	case 3:
		settings.FanSpeed = 2
	case 5:
		settings.FanSpeed = 3
	}

	if err != nil {
		return &Settings{}, err
	}

	settings.Temp, err = strconv.Atoi(pairs["stemp"])

	if err != nil {
		return &Settings{}, err
	}

	settings.Power = pairs["pow"] == "1"

	return &settings, nil
}

func NewManager(logger *dclilog.Logger) *Manager {
	m := &Manager{
		logger: logger,
		config: dcliconfig.GetConfig(),
	}

	return m
}
