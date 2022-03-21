# teams-webhook

Send a message to Microsoft Teams® via Webhook.


## Getting started

1. Download a pre-compiled binary from the release [page](https://github.com/dassump/teams-webhook/releases).
2. Run `teams-webhook`

```shell
$ teams-webhook
NAME:
   teams-webhook - Send a message to Microsoft Teams® via Webhook

USAGE:
   teams-webhook [global options] [arguments...]

VERSION:
   v1.0.0

AUTHOR:
   Daniel Dias de Assumpção <dassump@gmail.com>

GLOBAL OPTIONS:
   --color value, -c value    card theme color (default: "#DF813D")
   --message value, -m value  card text or file path to get content
   --title value, -t value    card title
   --url value, -u value      teams webhook url
   --version, -v              print the version (default: false)

COPYRIGHT:
   GNU General Public License

Required flags "message, title, url" not set
```


## Examples

### Inline message

```shell
$ teams-webhook --title "Test title" --message "Test message" --url "https://xxx.webhook.office.com/webhookb2/..."
```


### File content

```shell
$ echo "File content text" > file.txt
$ teams-webhook --title "Test title" --message ./file.txt --url "https://xxx.webhook.office.com/webhookb2/..."
```


### Shell stdout

```shell
$ ls | read -d "" content; teams-webhook --title "Test title" --message "$content" --url "https://xxx.webhook.office.com/webhookb2/..."
```


## Contributing

Bug reports and pull requests are welcome on GitHub at https://github.com/dassump/teams-webhook.