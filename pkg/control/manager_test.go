package control_test

import (
	"daikincli/internal/dclilog"
	"daikincli/pkg/control"
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestSet(t *testing.T) {
	logger := dclilog.GetInstance()
	manager := control.NewManager(logger)

	err := manager.SetState("24", "", "", "")

	assert.NoError(t, err)

}

func TestGet(t *testing.T) {
	logger := dclilog.GetInstance()
	manager := control.NewManager(logger)

	state, err := manager.GetState()

	assert.NoError(t, err)

	fmt.Println(state)
}

func TestGetZones(t *testing.T) {
	logger := dclilog.GetInstance()
	manager := control.NewManager(logger)

	val, err := manager.GetZones()

	assert.NoError(t, err)
	fmt.Println(val)
}

func TestSetZones(t *testing.T) {
	logger := dclilog.GetInstance()
	manager := control.NewManager(logger)

	err := manager.SetZones(true, true)

	assert.NoError(t, err)

	time.Sleep(2 * time.Second)

	val, err := manager.GetZones()

	assert.NoError(t, err)
	fmt.Println(val)

}
