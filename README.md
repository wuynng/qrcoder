# qrcoder #

A simple GUI application written in golang for generating QR code.  

Powered by [fyne](https://github.com/fyne-io/fyne), [qo-qrcode](https://github.com/skip2/go-qrcode).

## Usage
*Running*
```
$ git clone https://github.com/wuynng/qrcoder
$ cd qrcoder
$ go run cmd/qrcoder.go
```

*Packaging*
```
$ go get fyne.io/fyne/cmd/fyne
$ cd cmd
$ fyne package -os darwin -icon icon.png
```
