<div align="center">
  <a href="https://github.com/lumaaaaaa/bask">
    🔍
  </a>

<h3 align="center">bask</h3>

  <p align="center">
    Reverse-engineered Bing Chat client for the CLI
    <br />
    <a href="https://github.com/lumaaaaaa/bask/issues">Report Bug</a>
  </p>
</div>

---
## Demo

<img alt="Enola demo" src="https://misato.pw/bask/examples/kanzi.gif" width="600" />

---

## Why?

Microsoft recently released a new feature for the Bing search engine, called Bing Chat. In its current state, it requires users to use the Microsoft Edge browser, and to visit the Bing site to use the Chat functionality.

Personally, I'm not too fond of the Edge browser, and wanted a way to avoid that if possible. While a User-Agent change might be able to fool the Bing site into believing I was using Edge, I realized I wanted to be able to use this wonderful AI without opening a browser. My workflow was becoming less dependent on the browser each day for things like programming, and I hadn't found any CLI software that leveraged these new tools.

As such, bAsk was born.

## Setup

### Compatibility

This program should run wherever Go runs. It's been tested on Linux and macOS, but it should work just fine on Windows.

### Requirements

 - [Go 1.20 or later](https://go.dev/dl/)
 - Microsoft Edge (for grabbing cookies)
 - Microsoft Account with Bing Chat access

### Installation

To install, simply execute the following:
```go
go install github.com/lumaaaaaa/bask@latest
```

That should be all!

## Usage

To print the help message, execute:
```bash
bask -h
```

In general, to set up bask you should get your bing.com cookie from Microsoft Edge, then run the following:
```bash
bask -c {cookie}
```

After this, you should be able to chat straight from you terminal using:
```bash
bask -q {query}
```

Remember to put your cookie and query in double quotes or single quotes! Otherwise, each space will be treated as a different argument.

## Contributing

This program is very much under development, and should be expected to undergo constant changes. If you find a bug, be sure to open a pull request or and issue, and I will take a look at it. I'm aware the codebase is not the greatest at the moment, but expect it to improve in the coming future.

As always, any help is appreciated and thank you in advance! 