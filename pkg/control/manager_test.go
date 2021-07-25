package control_test

import (
	"daikincli/internal/dclilog"
	"daikincli/pkg/control"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGet(t *testing.T) {
	logger := dclilog.GetInstance()
	manager := control.NewManager(logger)

	state, err := manager.GetState()

	assert.NoError(t, err)

	assert.NotEmpty(t, state)

	logger.Infof(state)
}
