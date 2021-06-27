# devgo
> a command-line launcher

<img src="screenshots/devgo1.png" alt="devgo" width="250"/> <img src="screenshots/devgo2.png" alt="devgo" width="250"/><img src="screenshots/devgo3.png" alt="devgo" width="250"/>

## Install

```shell script
curl -o- https://raw.githubusercontent.com/TheWinds/devgo/main/install.sh | bash
```

## QuickStart
```shell script
dg
```

## Hotkeys
* select group `←` `→`
* select item  `↑` `↓`
* search item `type keywords`
* exit `ctrl+c` `ctrl+d` `esc`

## Config
> file path: ~/.devgo
```toml
# random group tab name preifx emoji
tab_emojis="🐶🐱🐭🦊🐻🐼🐮🐷🐸🐵🦉🦄🐟🐳🐖🐂💥🌈🌞"

[[group]]
# group name
name="tools"
[[group.item]]
# group item title
title="hello"
# group item command to exec
exec="echo hello devgo"
[[group.item]]
title="date now"
exec="date"

[[group]]
name="website"
[[group.item]]
title="github"
exec="open https://github.com"

[[group]]
name="devgo"
[[group.item]]
title="edit config"
exec="vim $HOME/.devgo"
```

## Uninstall
```shell script
rm $(which dg)
```