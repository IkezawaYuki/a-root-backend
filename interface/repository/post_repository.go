package repository

import (
	"IkezawaYuki/a-root-backend/domain/model"
	"IkezawaYuki/a-root-backend/infrastructure"
	"IkezawaYuki/a-root-backend/interface/filter"
	"context"
)

type PostRepository interface {
	Get(ctx context.Context, f *filter.PostFilter) ([]*model.Post, error)
	GetTx(ctx context.Context, f *filter.PostFilter, tx infrastructure.Transaction) ([]*model.Post, error)
	Save(ctx context.Context, post *model.Post) error
	SaveTx(ctx context.Context, post *model.Post, tx infrastructure.Transaction) error
	Count(ctx context.Context, f *filter.PostFilter) (int, error)
}

func NewPostRepository(dbDriver infrastructure.DBDriver) PostRepository {
	return &postRepository{
		dbDriver: dbDriver,
	}
}

type postRepository struct {
	dbDriver infrastructure.DBDriver
}

func (p *postRepository) Get(ctx context.Context, f *filter.PostFilter) ([]*model.Post, error) {
	var posts []*model.Post
	err := p.dbDriver.Find(ctx, &posts, f)
	if err != nil {
		return nil, err
	}
	return posts, nil
}

func (p *postRepository) GetTx(ctx context.Context, f *filter.PostFilter, tx infrastructure.Transaction) ([]*model.Post, error) {
	var posts []*model.Post
	err := p.dbDriver.FindTx(ctx, &posts, f, tx)
	if err != nil {
		return nil, err
	}
	return posts, nil
}

func (p *postRepository) Save(ctx context.Context, post *model.Post) error {
	return p.dbDriver.Save(ctx, post)
}

func (p *postRepository) SaveTx(ctx context.Context, post *model.Post, tx infrastructure.Transaction) error {
	return p.dbDriver.SaveTx(ctx, post, tx)
}

func (p *postRepository) Count(ctx context.Context, f *filter.PostFilter) (int, error) {
	count, err := p.dbDriver.Count(ctx, &model.Post{}, f)
	if err != nil {
		return 0, err
	}
	return int(count), nil
}
