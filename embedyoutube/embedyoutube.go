package embedyoutube

import (
	"encoding/json"
	"flag"
	"fmt"
	"net/http"
	"os"
	"time"
)

// CLI runs the go-embed-youtube command line cli and returns its exit status.
func CLI(args []string) int {
	var cli cliEnv
	err := cli.fromArgs(args)
	if err != nil {
		return 2
	}
	if err = cli.run(); err != nil {
		fmt.Fprintf(os.Stderr, "Runtime error: %v\n", err)
		return 1
	}
	return 0
}

type cliEnv struct {
	hc         http.Client
	youtubeURL string
	apiKey     string
}

func (cli *cliEnv) fromArgs(args []string) error {
	// Shallow copy of default client
	cli.hc = *http.DefaultClient
	fl := flag.NewFlagSet("embed-youtube", flag.ContinueOnError)
	fl.StringVar(
		&cli.youtubeURL, "y", YoutubePlaceholder, "Youtube Video to embed as Markdown",
	)
	fl.StringVar(
		&cli.apiKey, "k", "", "Google Developers API Key (required)",
	)
	fl.DurationVar(&cli.hc.Timeout, "t", 30*time.Second, "Client timeout")

	if err := fl.Parse(args); err != nil {
		return err
	}

	if len(cli.apiKey) == 0 {
		fmt.Fprintf(os.Stderr, "Missing or incorrect API Key: %q\n", cli.apiKey)
		fl.Usage()
		return flag.ErrHelp
	}

	return nil
}

func (cli *cliEnv) run() error {
	u, err := BuildURL(cli.youtubeURL, cli.apiKey)
	if err != nil {
		return err
	}

	var resp APIResponse
	if err = cli.fetchJSON(u, &resp); err != nil {
		return err
	}
	return markdownPrint(resp)
}

func (cli *cliEnv) fetchJSON(url string, data interface{}) error {
	resp, err := cli.hc.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	return json.NewDecoder(resp.Body).Decode(data)
}

func markdownPrint(ar APIResponse) error {
	_, err := fmt.Printf(
		`[![%s](%s)](https://youtube.com/watch?v=%s "%s")`,
		ar.Items[0].Snippet.Title,
		ar.Items[0].Snippet.Thumbnails.Standard.URL,
		ar.Items[0].ID,
		ar.Items[0].Snippet.Title,
	)
	return err
}
