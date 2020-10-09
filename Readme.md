# xipgen
Generate [xip.io](http://xip.io) links, as well as [nip.io](https://nip.io) and [sslip.io](https://sslip.io).

For example:

```
» xipgen 192.168.0.120
:: xip.io
xao79c.xip.io
192.168.0.120.xip.io

:: nip.io
192.168.0.120.nip.io
192-168-0-120.nip.io
c0a80078.nip.io

:: sslip.io
192.168.0.120.sslip.io
192-168-0-120.sslip.io
```

## Installation
While I work on getting releases on Github, you can currently install `xipgen` if you have [Go](https://golang.org) installed. Simply run the following in your command line:

```
$ go get github.com/codecat/xipgen
```
