package parser

import (
	"github.com/cx-antonio-lopes/manifest-parser/pkg/parser/models"
)

type Parser interface {
	Parse(manifestFile string) ([]models.Package, error)
}
