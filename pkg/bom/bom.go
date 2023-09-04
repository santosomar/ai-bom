package bom

import (
	"context"
	"io"

	"github.com/manifest-cyber/ai-bom/pkg/domain"
)

var _ domain.BomService = (*Service)(nil)

type Service struct{}

func NewService(opts ...Option) *Service {
	s := &Service{}

	for _, opt := range opts {
		opt(s)
	}

	return s
}

type Option func(s *Service)

func (s *Service) Generate(_ context.Context, _ io.ReadSeekCloser, _ io.WriteCloser) error {
	return nil
}
