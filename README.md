# barcomic

![Go](https://img.shields.io/github/go-mod/go-version/TheGrayDot/barcomic_server?filename=go.mod&style=plastic) ![Build Status](https://img.shields.io/github/actions/workflow/status/TheGrayDot/barcomic_server/push_and_pull.yml?branch=main&style=plastic)

![Release](https://img.shields.io/github/v/release/TheGrayDot/barcomic_server?style=plastic) ![Downloads](https://img.shields.io/github/downloads/TheGrayDot/barcomic_server/total?style=plastic)

An HTTP API for receiving comic book barcodes from the Barcomic Android application

## Latest Releases

- [Linux 64bit (`barcomic-linux`)](https://github.com/TheGrayDot/barcomic_server/releases/latest/download/barcomic-linux)
- [Windows 64bit (`barcomic-windows.exe`)](https://github.com/TheGrayDot/barcomic_server/releases/latest/download/barcomic-windows.exe)
- [OS X 64bit (`barcomic-darwin`)](https://github.com/TheGrayDot/barcomic_server/releases/latest/download/barcomic-darwin)

## Quick Start

- Download [latest release](https://github.com/TheGrayDot/barcomic_server/releases/latest/) from GitHub releases page
- Double click to run the program
- This should automatically open and start the server in interactive mode
- Pick an IP address from the list, usually you Ethernet or Wi-Fi adapter, so the Barcomic Android app can connect
- Connect the Barcomic Android app using the QR code

## Command Arguments

```
./barcomic-linux -h
[*] barcomic v1.0.0-6b9b9a5a750be60bc8c8f33a0c3acdbd783406a3
Usage of ./barcomic-linux:
  -a string
    	IP address to listen on (default "0.0.0.0")
  -i	Run interactive configuration (default true)
  -k	Disable keystrokes
  -p string
    	Port to listen on (default "9999")
  -v	Prints verbose information
```

> NOTE: For any of the examples provided below, change `barcomic-linux` to the correct release name you have downloaded. For example, `barcomic-windows.exe` or `barcomic-darwin`.

### Start server with IP and port specified

Use this if you have already configured your Barcomic Android app and want to start the server using known network configuration.

```
./barcomic-linux -a 192.168.1.100 -p 9876
```

### Start server without keystrokes enabled

Use this if you don't want to have the server "type" the barcode out. Good when used in verbose or logging mode.

```
./barcomic-linux -k -v
```
