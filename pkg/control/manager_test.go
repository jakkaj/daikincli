package control_test

import (
	"daikincli/internal/dclilog"
	"daikincli/pkg/control"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGet(t *testing.T) {
	logger := dclilog.GetInstance()
	manager := control.NewManager(logger)

	state, err := manager.GetState()

	assert.NoError(t, err)

	fmt.Println(state)
}
