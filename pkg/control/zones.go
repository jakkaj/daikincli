package control

import (
	"net/url"
	"strings"
)

func (s *Manager) GetZones() (zones map[string]bool, err error) {

	z, err := s.get("get_zone_setting")

	if err != nil {
		return nil, err
	}

	pairs := make(map[string]string)

	for _, group := range strings.Split(z, ",") {
		nameVal := strings.Split(group, "=")

		if nameVal[0] == "ret" {
			continue
		}

		decoded, err := url.QueryUnescape(nameVal[1])

		if err != nil {
			return nil, err
		}

		pairs[nameVal[0]] = decoded
	}

	zones = make(map[string]bool)

	zoneSplit := strings.Split(pairs["zone_name"], ";")
	zoneValsSplit := strings.Split(pairs["zone_onoff"], ";")

	for i, val := range zoneSplit {
		trimVal := strings.TrimSpace(val)
		if strings.ContainsAny(trimVal, "Zone") {
			continue
		}
		zones[trimVal] = zoneValsSplit[i] == "1"
	}

	return
}
