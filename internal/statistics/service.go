package statistics

import (
	"context"

	"github.com/redis/go-redis/v9"
)

type service struct {
	rdb *redis.Client
	ctx context.Context
}

func NewService(c *redis.Client) Service {
	return &service{rdb: c, ctx: context.Background()}
}

func (s *service) IncrementVisitorCount() error {
	return s.rdb.Incr(s.ctx, "site:visits").Err()
}

func (s *service) GetVisitorCount() (int64, error) {
	return s.rdb.Get(s.ctx, "site:visits").Int64()
}

func (s *service) IncrementArticleViews(articleID string) error {
	key := "article:visits:" + articleID
	return s.rdb.Incr(s.ctx, key).Err()
}

func (s *service) GetArticleViews(articleID string) (int64, error) {
	key := "article:visits:" + articleID
	return s.rdb.Get(s.ctx, key).Int64()
}

type Service interface {
	IncrementVisitorCount() error
	GetVisitorCount() (int64, error)
	IncrementArticleViews(string) error
	GetArticleViews(string) (int64, error)
}
