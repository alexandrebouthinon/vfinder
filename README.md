# Void's finder

__Find and report dead links in HTML files.__

[![GitHub license](https://img.shields.io/github/license/alexandrebouthinon/vfinder.svg)](https://github.com/alexandrebouthinon/vfinder/blob/master/LICENSE)
[![Go Report Card](https://goreportcard.com/badge/github.com/alexandrebouthinon/vfinder)](https://goreportcard.com/report/github.com/alexandrebouthinon/vfinder)

## Why ?
Because it's annoying to have some dead links in our web pages.


## Installation

### Using sources

Install Golang, clone the repository and build:
```
$ git clone https://github.com/alexandrebouthinon/vfinder
$ cd vfinder
$ go build
```

### Using __go get__

If you have a fully configured Golang environment, set up and add your `GOBIN` to your `PATH`. Finally, just run:
```
$ go get github.com/alexandrebouthinon/vfinder
$ go install github.com/alexandrebouthinon/vfinder
```

## Usage

You can check the documentation using the `-h` flag:

```
$ vfinder -h

         _    _________           __
        | |  / / ____(_)___  ____/ /__  _____
        | | / / /_  / / __ \/ __  / _ \/ ___/
        | |/ / __/ / / / / / /_/ /  __/ /
        |___/_/   /_/_/ /_/\__,_/\___/_/

Usage of vfinder:
-d string
A directory location as a string, this directory or sub-directories should contain HTML files to analize.
-f string
A file path as a string, This file should contain HTML code.
-x string
An exception filename as a string, this file sould contains prefix that need to be ignored in parsing.

```

### Check links in HTML file

You can check URLs in only one file using `-f` flag:

```
$ vfinder -f index.html
```

Or recursively, in files tree using `-d` flag:

```
$ vfinder -d mywebsite/build
```

### Add ignored URLs

Sometimes you need to ignore some of the errored URLs like `http://localhost/a/super/url` this can be performed using the `-x` flag and a custom file containing your ignored links or prefixes:


```
http://localhost/a/super/url
http://localhost
```

and use it in your command line for a single file:

```
$ vfinder -f index.html -x exceptions.txt
```

or a complete directory:

```
$ vfinder -d mywebsite/build -x exceptions.txt
```

## Docker image

If you don't want to install Golang, you can use the Docker image.
It's the recommended way to use Kuttlefish when you want to use it in your CI system

* Pull the image
```
$ docker pull alexandrebouthinon/vfinder
```

* Use it:
```
docker run --rm -it -v /absolute/path/to/directory:/mnt alexandrebouthinon/vfinder vfinder \
        -d /mnt/build \
        -x /mnt/exceptions.txt

```

## Author

Alexandre BOUTHINON [(@alexandrebouthinon)](https://github.com/alexandrebouthinon)

_Special thanks to_ [@scottinet](https://github.com/scottinet) _for performances improvements_

##  [MIT LICENSE](https://github.com/alexandrebouthinon/vfinder/raw/master/LICENSE)
