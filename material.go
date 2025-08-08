package gci

import (
	"errors"
	"fmt"
	"strings"
)

// Material 手工构造物料
type Material struct {
	Name         string   // Surface name
	PreviewImage string   // Preview image
	Texts        []string // Customization texts
	Images       []string // Customization images
}

func (m Material) validate() error {
	texts := m.Texts
	images := m.Images
	if len(texts) == 0 && len(images) == 0 {
		return errors.New("gci: Either texts or images must be filled in, but not both can be empty")
	}

	for _, lineStr := range texts {
		lineStr = strings.TrimSpace(lineStr)
		if lineStr == "" {
			return errors.New("gci: text cannot be empty")
		}

		label, _, ok := strings.Cut(lineStr, ":")
		if !ok {
			return fmt.Errorf("gci: invalid text: %s", lineStr)
		}
		if strings.TrimSpace(label) == "" {
			return errors.New("gci: invalid label")
		}
	}
	for _, image := range images {
		if strings.TrimSpace(image) == "" {
			return errors.New("gci: invalid image")
		}
	}

	return nil
}
