# Kobo Remote

**This is proof-of-concept network remote controller tested on my Kobo Glo HD**

## Build
Cross-compile a compatible ARM binary.
```
$ docker run --rm -v "$PWD":/usr/src/myapp -w /usr/src/myapp -e GOOS=linux -e GOARCH=arm golang:1.15 go build -v
```

## Installation
SCP the compiled binary to the Kobo reader and run it in the background.
```
$ ./kobo-remote &
```

## Usage examples
To change page, send a left or right touch event.

```
$ curl http://192.168.1.54:10000/left
$ curl http://192.168.1.54:10000/right
```