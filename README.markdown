# podlove-transfilter

This tool transforms the JSON transcript produced by [Whisper](https://github.com/openai/whisper) to the format used by the [Podlove Web Player](https://docs.podlove.org/podlove-web-player/).

TODO It also allows for shifting all timestamps by some duration (useful for correcting skew).

# Transcription Workflow

1. Install whisper:

    ```command
    $ brew install openai-whisper
    ```

1. Transcribe a podcast episode:

	```command
	$ whisper episode.mp3 --model medium --language en --output_format json
	```

1. Convert the transcript to the Podlove format:

	```command
	$ podlove-transfilter < episode.json > transcript.json
	```

1. [Configure the web player](https://docs.podlove.org/podlove-web-player/v5/configuration#transcripts) to use the transfiltered `transcript.json`.

# Development

Contributions are welcome! I'd appreciate if you get in touch via [a new issue](https://github.com/suhlig/podlove-transfilter/issues/new) before investing into a PR. We'll work out things from there.
