package service

import (
	"IkezawaYuki/a-root-backend/domain/model"
	"IkezawaYuki/a-root-backend/interface/filter"
	"IkezawaYuki/a-root-backend/interface/repository"
	"context"
)

type postService struct {
	postRepo repository.PostRepository
}

type PostService interface {
	IsLinked(ctx context.Context, instagramMediaID string) (bool, error)
	Create(ctx context.Context, post *model.Post) error
	Update(ctx context.Context, post *model.Post) error
	FindByCustomerID(ctx context.Context, customerID int) ([]*model.Post, error)
	GetLinkedMediaIDs(ctx context.Context, customerID int) ([]string, error)
}

func NewPostService(postRepo repository.PostRepository) PostService {
	return &postService{
		postRepo: postRepo,
	}
}

func (s *postService) IsLinked(ctx context.Context, instagramMediaID string) (bool, error) {
	posts, err := s.postRepo.Get(ctx, &filter.PostFilter{
		InstagramMediaID: &instagramMediaID,
	})
	if err != nil {
		return false, err
	}
	return len(posts) > 0, nil
}

func (s *postService) Create(ctx context.Context, post *model.Post) error {
	return s.postRepo.Save(ctx, post)
}

func (s *postService) Update(ctx context.Context, post *model.Post) error {
	return s.postRepo.Save(ctx, post)
}

func (s *postService) FindByCustomerID(ctx context.Context, customerID int) ([]*model.Post, error) {
	return s.postRepo.Get(ctx, &filter.PostFilter{
		CustomerID: &customerID,
	})
}

func (s *postService) GetLinkedMediaIDs(ctx context.Context, customerID int) ([]string, error) {
	posts, err := s.postRepo.Get(ctx, &filter.PostFilter{
		CustomerID: &customerID,
	})
	if err != nil {
		return nil, err
	}
	linkedMediaIDs := make([]string, len(posts))
	for i, post := range posts {
		linkedMediaIDs[i] = post.InstagramMediaID
	}
	return linkedMediaIDs, nil
}
