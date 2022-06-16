package main

import (
	"github.com/IT-Nick/go-project-cms/cmd/rcms/config"
	"github.com/stretchr/testify/assert"
	"testing"
)

// Пример теста, используем команду `make test` (прям тут, в любом свободном поле), чтобы создать тест
func TestDummy(t *testing.T) {
	assert.Equal(t, 1, 1)
}

func TestLoadConfig(t *testing.T) {
	if err := config.LoadConfig("./config"); err != nil {
		assert.Error(t, err)
	}
	assert.Equal(t, "Castro washes socks at the university", config.Config.ConfigVar)
}
