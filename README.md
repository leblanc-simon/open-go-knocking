# OpenGoKnocking

Simple and cross-platform port knocking client written in GoLang.

<p align="center">
<img src="https://raw.githubusercontent.com/leblanc-simon/open-go-knocking/main/assets/open-go-knocking.png">
</p>

## Installation

To install OpenGoKnocking in your $GOPATH:

```bash
git clone https://github.com/leblanc-simon/open-go-knocking.git
cd open-go-knocking
go install
```

alternatively, you can download binary from [releases](https://github.com/leblanc-simon/open-go-knocking/releases). 
Binaries are availables for GNU/Linux, MacOS and Windows

## Usage

```bash
# Knock to the server example.com at port 1337 (tcp), then 1338 (UDP) and finally 1339 (TCP)
open-go-knocking example.com 1337:tcp 1338:udp 1339:tcp

# If the majority of ports are TCP ports, you can use the following syntax (ports without protocol will be in TCP)
open-go-knocking --tcp example.com 1337 1338:udp 1339

# If the majority of ports are UDP ports, you can use the following syntax (ports without protocol will be in UDP)
open-go-knocking --udp example.com 1337 1338:tcp 1339
```

## Options

* `--tcp`: 
* `--udp`: 
* `--delay` (in ms): 
* `--version`: Show the current version
* `--help`: Show the help

## Errors

Possible exit codes :

* 0: all is fine :)
* 1: Not enough arguments (host and port are required)
* 2: Host is not a valid hostname or IP (v4 or v6)
* 3: Problems with ports or protocols (check if you use `--tcp` or `--udp` that the option is located before the host)

## Author

* Simon Leblanc <contact@leblanc-simon.eu>

## License

[WTFPL](http://www.wtfpl.net/)

## Logo

Original Logo from [Takuya Ueda](https://twitter.com/tenntenn), you can find source [in the repository](https://github.com/golang-samples/gopher-vector). Licensed under the Creative Commons 3.0 Attributions license.

Modified by Simon Leblanc.