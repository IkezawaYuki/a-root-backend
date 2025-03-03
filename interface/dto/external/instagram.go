package external

import "strings"

type InstagramPosts struct {
	ID    string `json:"id"`
	Media struct {
		Data []InstagramPost `json:"data"`
	} `json:"media"`
}

type InstagramPost struct {
	ID        string                `json:"id"`
	Permalink string                `json:"permalink"`
	Caption   string                `json:"caption,omitempty"`
	Timestamp string                `json:"timestamp"`
	MediaType string                `json:"media_type"`
	MediaURL  string                `json:"media_url"`
	Children  InstagramPostChildren `json:"children"`
}

func (i *InstagramPost) Title() string {
	return strings.Split(i.Caption, " ")[0]
}

type InstagramPostChildren struct {
	Data []struct {
		MediaType string `json:"media_type"`
		MediaURL  string `json:"media_url"`
		ID        string `json:"id"`
	} `json:"data"`
}
