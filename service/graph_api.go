package service

import (
	"IkezawaYuki/a-root-backend/config"
	"IkezawaYuki/a-root-backend/domain/arootErr"
	"IkezawaYuki/a-root-backend/infrastructure"
	"IkezawaYuki/a-root-backend/interface/dto/external"
	"context"
	"encoding/json"
	"fmt"
	"net/url"
)

type graphAPI struct {
	httpClient infrastructure.HttpClient
	baseURL    string
}

type GraphAPI interface {
	GetInstagramBusinessAccountID(ctx context.Context, facebookToken string) (string, error)
	GetInstagramPosts(ctx context.Context, facebookToken string, instagramID string) (*external.InstagramPosts, error)
	GetOAuthAccessToken(ctx context.Context, facebookToken string) (*external.OAuthAccessToken, error)
}

func NewGraph(httpClient infrastructure.HttpClient) GraphAPI {
	return &graphAPI{
		httpClient: httpClient,
		baseURL:    config.Env.GraphApiURL,
	}
}

const getInstagramBusinessAccountURL = "/me?fields=id,name,accounts{instagram_business_account}"

type GraphApiMeResponse struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	Accounts struct {
		Data []struct {
			InstagramBusinessAccount struct {
				ID string `json:"id"`
			} `json:"instagram_business_account"`
			ID string `json:"id"`
		} `json:"data"`
	} `json:"accounts"`
}

func (i *graphAPI) GetInstagramBusinessAccountID(ctx context.Context, facebookToken string) (string, error) {
	resp, err := i.httpClient.GetRequest(ctx,
		i.baseURL+getInstagramBusinessAccountURL,
		fmt.Sprintf("Bearer %s", facebookToken))
	if err != nil {
		return "", err
	}
	var instagram GraphApiMeResponse
	err = json.Unmarshal(resp, &instagram)
	if err != nil {
		return "", err
	}
	if len(instagram.Accounts.Data) == 0 {
		return "", arootErr.ErrNotFound
	}
	return instagram.Accounts.Data[0].InstagramBusinessAccount.ID, nil
}

const getInstagramPosts = "%s?fields=media{id,permalink,caption,timestamp,media_type,media_url,children{media_type,media_url}}"

func (i *graphAPI) GetInstagramPosts(ctx context.Context, facebookToken string, instagramID string) (*external.InstagramPosts, error) {
	resp, err := i.httpClient.GetRequest(ctx,
		i.baseURL+fmt.Sprintf(getInstagramPosts, instagramID),
		fmt.Sprintf("Bearer %s", facebookToken),
	)
	if err != nil {
		return nil, err
	}
	var posts external.InstagramPosts
	if err := json.Unmarshal(resp, &posts); err != nil {
		return nil, err
	}
	return &posts, nil
}

const postOAuthAccessToken = "/oauth/access_token"

func (i *graphAPI) GetOAuthAccessToken(ctx context.Context, facebookToken string) (*external.OAuthAccessToken, error) {
	params := url.Values{}
	params.Add("grant_type", "fb_exchange_token")
	params.Add("fb_exchange_token", facebookToken)
	params.Add("client_id", config.Env.MetaClientID)
	params.Add("client_secret", config.Env.MetaClientSecret)
	u := i.baseURL + postOAuthAccessToken + "?" + params.Encode()
	resp, err := i.httpClient.GetRequest(ctx, u, "")
	if err != nil {
		return nil, err
	}
	var oauthAccessToken external.OAuthAccessToken
	err = json.Unmarshal(resp, &oauthAccessToken)
	if err != nil {
		return nil, err
	}
	return &oauthAccessToken, nil
}
