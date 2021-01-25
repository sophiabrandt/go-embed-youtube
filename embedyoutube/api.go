package embedyoutube

import (
	"fmt"
	"regexp"
	"strings"
	"time"
)

// YoutubePlaceholder is a placeholder video URL.
const YoutubePlaceholder = "https://youtube.com/watch?v=ScMzIvxBSi4"

// extractYoutubeID checks if an input matches the Youtube URL pattern
// and returns the Youtube video ID or an error.
func extractYoutubeID(url string) (string, error) {
	re := regexp.MustCompile(`(?:youtube\.com\/\S*(?:(?:\/e(?:mbed))?\/|watch\/?\?(?:\S*?&?v\=))|youtu\.be\/)([a-zA-Z0-9_-]{6,11})`)
	found := re.FindAllString(url, -1)
	if len(found) == 0 {
		return "", fmt.Errorf("invalid Youtube URL")
	}
	verifiedYoutubeURL := strings.Split(found[0], "=")
	id := verifiedYoutubeURL[len(verifiedYoutubeURL)-1]
	return id, nil
}

// BuildURL creates a Google API URL from an input URL.
func BuildURL(youtubeURL, apiKey string) (string, error) {
	id, err := extractYoutubeID(youtubeURL)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("https://www.googleapis.com/youtube/v3/videos?id=%s&part=snippet&key=%s", id, apiKey), nil
}

// APIResponse returned by the Youtube API:
// https://developers.google.com/youtube/v3/docs/videos
// Generated with https://mholt.github.io/json-to-go/
type APIResponse struct {
	Kind  string `json:"kind"`
	Etag  string `json:"etag"`
	Items []struct {
		Kind    string `json:"kind"`
		Etag    string `json:"etag"`
		ID      string `json:"id"`
		Snippet struct {
			PublishedAt time.Time `json:"publishedAt"`
			ChannelID   string    `json:"channelId"`
			Title       string    `json:"title"`
			Description string    `json:"description"`
			Thumbnails  struct {
				Default struct {
					URL    string `json:"url"`
					Width  int    `json:"width"`
					Height int    `json:"height"`
				} `json:"default"`
				Medium struct {
					URL    string `json:"url"`
					Width  int    `json:"width"`
					Height int    `json:"height"`
				} `json:"medium"`
				High struct {
					URL    string `json:"url"`
					Width  int    `json:"width"`
					Height int    `json:"height"`
				} `json:"high"`
				Standard struct {
					URL    string `json:"url"`
					Width  int    `json:"width"`
					Height int    `json:"height"`
				} `json:"standard"`
				Maxres struct {
					URL    string `json:"url"`
					Width  int    `json:"width"`
					Height int    `json:"height"`
				} `json:"maxres"`
			} `json:"thumbnails"`
			ChannelTitle         string   `json:"channelTitle"`
			Tags                 []string `json:"tags"`
			CategoryID           string   `json:"categoryId"`
			LiveBroadcastContent string   `json:"liveBroadcastContent"`
			Localized            struct {
				Title       string `json:"title"`
				Description string `json:"description"`
			} `json:"localized"`
		} `json:"snippet"`
	} `json:"items"`
	PageInfo struct {
		TotalResults   int `json:"totalResults"`
		ResultsPerPage int `json:"resultsPerPage"`
	} `json:"pageInfo"`
}
