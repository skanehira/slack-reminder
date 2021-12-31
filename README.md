![GitHub Repo stars](https://img.shields.io/github/stars/skanehira/slack-reminder?style=social)
![GitHub](https://img.shields.io/github/license/skanehira/slack-reminder)
![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/skanehira/slack-reminder)
![GitHub all releases](https://img.shields.io/github/downloads/skanehira/slack-reminder/total)
![GitHub CI Status](https://img.shields.io/github/workflow/status/skanehira/slack-reminder/ci?label=CI)
![GitHub Release Status](https://img.shields.io/github/workflow/status/skanehira/slack-reminder/Release?label=release)

# slack-reminder
Slack remind generator

![](https://i.gyazo.com/1c8d4750ab8ae601029a00cf58faa89b.gif)

## Why
Slack's `/reminder` is very hard to use.

## Required
- Go 1.16 ~

## Installation

```go
$ go install github.com/skanehira/slack-reminder@latest
```

## Usage
### onetime reminder
```sh
MacbookPro13% slack-reminder
? Kind of remind onetime
? Date(YYYY-MM-DD) 2020-05-13
? Hour(HH:MM) 10:00
? @someone or #channel or me @gorilla
? Message hello, I'm here.
/remind @gorilla "hello, I'm here." at 10:00 on 2020-05-13
```
### repetition reminder
```sh
MacbookPro13% slack-reminder
? Kind of remind repetition
? What kind of repetition every week
? What day of week choice
? Choice days Tuesday, Thursday, Sunday
? Hour(HH:MM) 10:00
? @someone or #channel or me me
? Message some remind
/remind me "some remind" at 10:00 on every Tuesday ,Thursday ,Sunday
```

## Author
skanehira
