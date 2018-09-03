# Void's finder

__Find and report dead links in HTML files.__

## Why ?
Because it's annoying to have some dead links in our web pages.

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

## Author

Alexandre Bouthinon (@alexandrebouthinon)

_Special thanks to @scottinet for performances improvements_
