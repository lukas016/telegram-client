<!-- insert
---
title: Telegram Tee
date: 2021-08-21T16:23:33
gometa: "cj.rs/telegram-tee git https://github.com/cljoly/telegram-tee"
description: "💬 Simple cli tool to create telegram bot behaving like tee"
---
{{< github_badge >}}
end_insert -->
<!-- remove -->
# Telegram Tee
<!-- end_remove -->

**Forked from https://github.com/cljoly/telegram-tee** \
Simple cli tool to send html formatted messages from stdin to any Telegram chat, through a bot. 

## Getting started

### Set up

First, install the tool with
``` bash
go install github.com/lukas016/telegram-client
```

Then, you need to control a bot. Set the environment variable `TLGCLI_TOKEN` to
the token of the bot that will write stdin to a chat for you. You may want to [create a new bot](https://core.telegram.org/bots#3-how-do-i-create-a-bot) or use an existing one.

### Use
Simple input
``` bash
echo "<strong>Hi</strong>" | telegram-client <chatID> ...
```

Multiline input
``` bash
printf "<string>Hi</strong>\nHow are you?" | telegram-client <chatID> ...
```
or
``` bash
telegram-client <chatID> ... << EOF
<string>Hi</strong>
How are you?
EOF
```

You can even send to several chatID at the same time.