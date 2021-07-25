package control

import (
	"daikincli/internal/dcliconfig"
	"daikincli/internal/dclilog"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Manager struct {
	logger *dclilog.Logger
	config *dcliconfig.Config
}

func (s *Manager) get(endpoint string) (*[]byte, error) {

	baseUrl := s.config.GetString(dcliconfig.DAIKIN_URL)
	password := s.config.GetString(dcliconfig.DAIKIN_PASSWORD)

	url := fmt.Sprintf("%v/skyfi/aircon/%v?lpw=%v", baseUrl, endpoint, password)

	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		return nil, err
	}

	return &body, nil

}

func (s *Manager) GetState() (string, error) {
	b, err := s.get("get_control_info")
	if err != nil {
		return "", err
	}

	return string(*b), nil
}

func NewManager(logger *dclilog.Logger) *Manager {
	m := &Manager{
		logger: logger,
		config: dcliconfig.GetConfig(),
	}

	return m
}
