package control_test

import (
	"daikincli/internal/dclilog"
	"daikincli/pkg/control"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSet(t *testing.T) {
	logger := dclilog.GetInstance()
	manager := control.NewManager(logger)

	before, after, err := manager.SetState("24", "", "", "")

	assert.NoError(t, err)

	assert.NotEmpty(t, before.Temp)
	assert.NotEmpty(t, after.Temp)
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
