# Last Week

Generate an automatic report based on your git commits in a time frame.

## Installation

```bash
go install github.com/newtoallofthis123/last_week@latest
```

## Usage

```bash
last_week -time "last week" -submodules
```

You can describe the time frame in a few different ways:

- last week
- last month
- last 3 days
- 6th of May etc

The submodules flag is used to include the submodules in the report.
There is also a `local` flag to signal where the .env file is located and
a `pipe` flag to signal that the output should be piped to the terminal.

By default, `last_week` will use the `GOOGLE_GEMINI_KEY` environment variable to authenticate with Google Gemini.
This is searched for in the `~/.gcommit` file as well as the standard environment variables.

## Recommended Usage

```bash
last_week -time "last week" -submodules | slides
```

Install the [slides](https://github.com/maaslalani/slides) tool and pipe the output to it. This will automatically create a presentation from the report.

## Todo

This is a work in progress and there are a few things that could be improved:

- [ ] Custom system prompts
- [ ] Local LLM support
- [ ] More code related options
- [ ] Presentation Customization
- [ ] Better git options
- [ ] Coauthors and custom author

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.