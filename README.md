# Goreview

:octocat: A webhook server to auto-assign (pseudo-randomly) teammates to review a new Github pull request.

[![CircleCI](https://circleci.com/gh/kelvintaywl/goreview.svg?style=svg)](https://circleci.com/gh/kelvintaywl/goreview) [![Go Report Card](https://goreportcard.com/badge/github.com/kelvintaywl/goreview)](https://goreportcard.com/report/github.com/kelvintaywl/goreview)

## Usage

### First things first

You would need to set your personal GitHub access token in the .env file or your bash script.

```shell
EXPORT GITHUB_ACCESS_TOKEN=replaceMe!
```

This is needed to make calls with GitHub API.

Next, add a `goreview.json` in the root directory of your GitHub repository.

This GitHub repository is the repository we want Goreview to assign reviewers.

See a [sample goreview.json file](goreview.json)

The settings are:

| field | remarks |
| --- | --- |
| `num_reviewers` | total number of random reviewers to assign to a new pull request. |
| `reviewers` | list of GitHub handlers of your team mates to pick from. Note that Goreview will not pick the author of the pull request. |
| `webhook_url` | URL to push a success payload from Goreview, if any _(in progress)_ |


### Quick setup
You can download the [latest Docker image at Docker Hub](https://hub.docker.com/r/kelvintaywl/goreview/), and run it.

```shell
$ docker pull kelvintaywl/goreview:latest
# The image exposes port 9999
# in the command below, we bind our localhost:5000 with the container's port 9999
$ docker run --rm -p 127.0.0.1:5000:9999 -e GITHUB_ACCESS_TOKEN=$(GITHUB_ACCESS_TOKEN) kelvintaywl/goreview:latest
```

Alternatively, if you prefer building and running the image yourself,

```shell
$ make init
$ make docker_build
# the command below would use the $PORT environment variable to publish the server (i.e, 5000)
$ make docker_run
```

We use the [`godotenv` library](https://github.com/joho/godotenv) to conveniently load and apply environment variables in the `.env` file onto our make commands.


### Publishing server

Use [ngrok](https://ngrok.com/3) to expose this local server publicly.

```shell
$ ngrok http 5000

ngrok by @inconshreveable                                                                                                                        (Ctrl+C to quit)
                                                                                                                                                                 
Session Status                online                                                                                                                             
Update                        update available (version 2.2.8, Ctrl-U to update)                                                                                 
Version                       2.1.18                                                                                                                             
Region                        United States (us)                                                                                                                 
Web Interface                 http://127.0.0.1:4040                                                                                                              
Forwarding                    http://11111111.ngrok.io -> localhost:500                                                                                        
Forwarding                    https://11111111.ngrok.io -> localhost:5000     
```

### Adding Server to GitHub's webhooks

With the server now discoverable publicly, we can now add our server's URL
onto our desired Github repo's webhook settings.

In the case of this repository, go to https://github.com/pipedpiper/nothotdog/settings/hooks/new

> You should replace `pipedpiper/nothotdog`  with an actual organization/repo :doughnut:  

## Contributing

Feel free to make pull requests! Let me know if this is useful / fun / stupid, and if your team is using this! :beer:
