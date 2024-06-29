package book

import (
	"context"
	"time"

	"github.com/sirupsen/logrus"
	"golang.org/x/sync/errgroup"

	"github.com/bxcodec/go-clean-arch/domain"
)


type Service struct {
}

// NewService will create a new article service object
func NewService() *Service {
	return &Service{
	}
}

func (a *Service) Upload(ctx context.Context, m *domain.Article) (err error) {
	
}

func (a *Service) Summary(ctx context.Context, m *domain.Article) (err error) {
	
}