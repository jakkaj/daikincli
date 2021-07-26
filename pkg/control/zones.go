package control

import (
	"fmt"
	"net/url"
	"strings"
)

func (s *Manager) SetZones(bedroom bool, study bool) (err error) {

	z, err := s.get("get_zone_setting")

	if err != nil {
		return err
	}

	zones := strings.Split(z, ",")[1]

	br := "0"

	if bedroom {
		br = "1"
	}

	st := "0"

	if study {
		st = "1"
	}

	zoneOnff := fmt.Sprintf("%v;%v;0;0;0;0;0;0", st, br)

	zoneOnffEncoded := url.QueryEscape(zoneOnff)

	qs := fmt.Sprintf("&%v&zone_onoff=%v", zones, zoneOnffEncoded)

	err = s.set("set_zone_setting", qs)

	return err

}

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
