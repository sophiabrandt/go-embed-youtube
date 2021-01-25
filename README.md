## go-embed-youtube

Embedding a YouTube video in Markdown format is cumbersome. This project aims to help with that.

## Motivation

If you want to embed a video using _standard markdown_, you need to do [this](https://stackoverflow.com/questions/11804820/how-can-i-embed-a-youtube-video-on-github-wiki-pages):

```markdown
[![IMAGE ALT TEXT](http://img.youtube.com/vi/YOUTUBE_VIDEO_ID_HERE/0.jpg)](http://www.youtube.com/watch?v=YOUTUBE_VIDEO_ID_HERE "Video Title")
```

Trying to remember the correct syntax, plus copy and pasting the needed values is a bit of a hassle.

Ideally, I want to use a link to a YouTube video and let a program figure out the rest.

## Tech/framework used

<b>Built with</b>

- [Go](https://golang.org)

## Features

The project is a command line program that uses the YouTube API and returns markdown-formatted text as [`stdout`](https://www.howtogeek.com/435903/what-are-stdin-stdout-and-stderr-on-linux/).

You can use this tool in conjunction with other utilities, like [xclip](https://linux.die.net/man/1/xclip) or [pbpaste](https://osxdaily.com/2007/03/05/manipulating-the-clipboard-from-the-command-line/).

## Example Usage

```sh
go-embed-youtube -y 'https://youtube.com/watch?v=ScMzIvxBSi4' -k <google developer API key>
```

Output:

```sh
[![Placeholder Video](https://i.ytimg.com/vi/ScMzIvxBSi4/sddefault.jpg)](https://youtube.com/watch?v=ScMzIvxBSi4 "Placeholder Video")
```

Rendered Markdown:

[![Placeholder Video](https://i.ytimg.com/vi/ScMzIvxBSi4/sddefault.jpg)](https://youtube.com/watch?v=ScMzIvxBSi4 "Placeholder Video")

## Installation

1. Get a [Google Developer API key](https://elfsight.com/blog/2016/12/how-to-get-youtube-api-key-tutorial/).

2. You need [Go](https://golang.org/dl/) on your machine. Then clone the repository:

   ```sh
   git clone https://github.com/sophiabrandt/go-embed-youtube.git
   ```

3. Build binary:

   ```sh
   cd go-embed-youtube
   go build -o go-embed-youtube
   ```

## How to use?

```bash
go-embed-youtube -h

Usage of embed-youtube:
  -k string
        Google Developers API Key (required)
  -t duration
        Client timeout (default 30s)
  -y string
        Youtube Video to embed as Markdown (default "https://youtube.com/watch?v=ScMzIvxBSi4")
```

## Advanced Example

Use [pass](https://www.passwordstore.org/) to store your Google API Developer Key:

```sh
pass add google-dev/youtube-api-key
```

Use `xargs` to pipe the API key to `go-embed-youtube`:

```sh
pass show google-dev/youtube-api-key | xargs -I % ./go-embed-youtube -k % -y 'https://youtube.com/watch?v=ScMzIvxBSi4'
```

**Bonus tip**: If you use the [fish shell](https://fishshell.com), you can create a wrapper function:

```sh
function embedyoutube
   pass show google-dev/youtube-api-key | xargs -I % go-embed-youtube -k % -y $argv
end
```

(The function assumes that you have `go-embed-youtube` in your [`$PATH`](https://astrobiomike.github.io/unix/modifying_your_path).)

Don't forget to save it:

```sh
funcsave embedyoutube
```

Use it:

```sh
embedyoutube 'https://youtube.com/watch?v=ScMzIvxBSi4'
```

## Contribute

Pull requests are welcome. For major changes, please open an issue first to discuss what you would like to change.

Please make sure to update tests as appropriate.

## Credits

The project owns its thanks to Carl M. Johnson's blog **[Writing Go CLIs With Just Enough Architecture](https://blog.carlmjohnson.net/post/2020/go-cli-how-to-and-advice/)** and accompanying [code](https://github.com/carlmjohnson/go-grab-xkcd).

## License

Distributed under the MIT License. See [LICENSE](LICENSE) for details.
