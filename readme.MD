# Kraken Advanced Proxy Checker

![Kraken Logo](https://i.imgur.com/RTUdcfp.png)

Kraken is a powerful, cross-platform proxy checker designed to offer blazing fast speed and support for both HTTP and SOCKS proxies. It's built with simplicity and efficiency in mind, ensuring reliable performance across various environments.

## Features

- **HTTP/SOCKS Support**: Compatible with both HTTP and SOCKS proxies.
- **Blazing Fast**: Optimized for speed, ensuring quick proxy checking.
- **Cross-Platform**: Runs smoothly on various operating systems.
- **Configurable**: Easy to customize settings via `config.json`.

## Tool Speed

![Tool Speed](https://media.discordapp.net/attachments/720287329859797024/1174015280184303666/image.png)
#### kraken can realistically check between 1-2 million proxies per minute on the average system. 

## Configuration (`config.json`)

```json
{
  "threads": 400,
  "prefix": "http",
  "checkURL": "http://ip-api.com/json?fields=8194",
  "successKey": "countryCode",
  "proxy-filepath": "data/http_proxies.txt",
  "debug": false,
  "timeout": 10
}
```
You can modify the checkURL and successKey to your liking, but the default settings should work fine for most users. Proxies are read from the proxy-filepath line by line, and the output is saved to the Results folder along with date and proxy type.
## Getting started
Head over to https://github.com/xvertile/Kraken/releases/latest and download the latest release for your operating system.

## Building from source
1. Install GoLang: Ensure that GoLang is installed on your system.
2. Clone Repository: ``git clone https://github.com/xvertile/Kraken.git``  to clone the Kraken repository.
2. Modify the ``config.json`` for your needs.
3. Run: Navigate to the cloned directory and run ```go run main.go``` to start the proxy checker.
