# `transcode-cli`

ffmpeg-wrapper with a user-friendly CLI

```sh
$ transcode file file.mkv
Select audio track:
  - track 1
  - track 2
Select subtitle track:
  - track 1
  - track 2
Output filename (will use <>.mp4 by default):
$ transcode file -i file.mkv
Select output codec:
  - libx264
  - vp8
  - vp9
...
$ transcode directory # TODO
$ transcode inspect
```

## ToDos

- add directory subcommand
- use more elaborate TUI
- expand and document list of default `ffmpeg` options
