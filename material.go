package gci

import (
	"errors"
	"fmt"
	"strings"
)

// Material 手工构造物料
type Material struct {
	Name         string
	PreviewImage string
	Texts        []string
	Images       []string
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

		label, value, ok := strings.Cut(lineStr, ":")
		if !ok {
			return fmt.Errorf("gci: invalid text: %s", lineStr)
		}

		label = strings.TrimSpace(label)
		value = strings.TrimSpace(value)
		if label == "" {
			return errors.New("gci: invalid label")
		}
	}
	for _, image := range images {
		image = strings.TrimSpace(image)
		if image == "" {
			return errors.New("gci: invalid image")
		}
	}

	return nil
}
