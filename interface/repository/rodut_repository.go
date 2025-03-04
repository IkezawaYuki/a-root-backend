package repository

import (
	"IkezawaYuki/a-root-backend/config"
	"IkezawaYuki/a-root-backend/infrastructure"
	"IkezawaYuki/a-root-backend/interface/dto/external"
	"context"
	"encoding/json"
	"fmt"
)

type RodutRepository interface {
	CreatePost(ctx context.Context, wordpressUrl string, post external.WordpressPost) (*external.CreatePostResponse, error)
	UploadMedia(ctx context.Context, wordpressUrl string, filePath string) (*external.UploadMediaResponse, error)
	UploadMedias(ctx context.Context, wordpressUrl string, filePath []string) ([]*external.UploadMediaResponse, error)
}

func NewRodutRepository(httpClient infrastructure.HttpClient) RodutRepository {
	return &rodutRepository{
		httpClient: httpClient,
		ApiKey:     config.Env.RodutKey,
		AdminEmail: config.Env.WordpressAdminEmail,
	}
}

type rodutRepository struct {
	httpClient infrastructure.HttpClient
	ApiKey     string
	AdminEmail string
}

const createPostEndpoint = "create-post"

func (r *rodutRepository) CreatePost(ctx context.Context, wordpressUrl string, post external.WordpressPost) (*external.CreatePostResponse, error) {
	url := fmt.Sprintf("https://%s/rodut/v1/%s", wordpressUrl, createPostEndpoint)
	responseBody, err := r.httpClient.PostRequest(ctx, url, &external.CreatePostRequest{
		ApiKey:        r.ApiKey,
		Email:         r.AdminEmail,
		Title:         post.Title,
		Content:       post.Content,
		FeaturedMedia: post.FeaturedMedia,
	}, "")
	if err != nil {
		return nil, err
	}
	createdPost := external.CreatePostResponse{}
	err = json.Unmarshal(responseBody, &createdPost)
	if err != nil {
		return nil, err
	}
	return &createdPost, nil
}

const uploadMediaEndpoint = "upload-media"

func (r *rodutRepository) UploadMedia(ctx context.Context, wordpressUrl string, filePath string) (*external.UploadMediaResponse, error) {
	url := fmt.Sprintf("https://%s/rodut/v1/%s", wordpressUrl, uploadMediaEndpoint)
	response, err := r.httpClient.UploadFile(ctx, url, filePath, "")
	if err != nil {
		return nil, err
	}
	uploadMedia := external.UploadMediaResponse{}
	err = json.Unmarshal(response, &uploadMedia)
	if err != nil {
		return nil, err
	}
	return &uploadMedia, nil
}

func (r *rodutRepository) UploadMedias(ctx context.Context, wordpressUrl string, filePaths []string) ([]*external.UploadMediaResponse, error) {
	result := make([]*external.UploadMediaResponse, len(filePaths))
	for i, filePath := range filePaths {
		resp, err := r.UploadMedia(ctx, wordpressUrl, filePath)
		if err != nil {
			return nil, err
		}
		result[i] = resp
	}
	return result, nil
}
