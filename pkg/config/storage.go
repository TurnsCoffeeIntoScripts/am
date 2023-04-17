package config

import (
	"turnscoffeeintoscripts/am/pkg/model"
	"turnscoffeeintoscripts/am/pkg/terminal"

	"gopkg.in/yaml.v3"
)

type Storage struct {
	Aliases []model.Alias
}

func (s *Storage) Marshal() []byte {
	if yd, err := yaml.Marshal(&s); err != nil {
		terminal.ErrorMessage("Failed to marshal 'Storage' struct")
		return nil
	} else {
		return yd
	}
}
