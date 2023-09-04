package domain

import (
	"context"
	"io"
)

const (
	CycloneDX     = "cyclonedx"
	CycloneDXJSON = "cyclonedx-json"
	SPDX          = "spdx"
	SPDXJSON      = "spdx-json"
)

var BOMOutputs = []string{CycloneDXJSON}

type BomService interface {
	Generate(ctx context.Context, in io.ReadSeekCloser, out io.WriteCloser) error
}
