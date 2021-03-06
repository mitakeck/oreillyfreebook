# oreillyfreebook

[![Circle CI](https://circleci.com/gh/mitakeck/oreillyfreebook/tree/master.svg?style=shield)](https://circleci.com/gh/mitakeck/oreillyfreebook/tree/master) [![Go Report Card](https://goreportcard.com/badge/github.com/mitakeck/oreillyfreebook)](https://goreportcard.com/report/github.com/mitakeck/oreillyfreebook) [![MIT License](https://img.shields.io/badge/license-MIT-blue.svg)](https://github.com/mitakeck/oreillyfreebook/blob/master/LICENSE)

O'Reilly Free Ebook Downloader


## Install

### From Homebrew

```
$ brew tap mitakeck/oreillyfreebook
$ brew install oreillyfreebook
```

### From binary

- Download from [here](https://github.com/mitakeck/oreillyfreebook/releases/latest).

```
$ chmod +x oreillyfreebook_*
$ mv oreillyfreebook_* /usr/loca/bin/oreillyfreebook
```

### From source code

```
$ go get github.com/mitakeck/oreillyfreebook
$ cd $GOPATH/src/github.com/mitakeck/oreillyfreebook
$ go install
```

## Usage

[![asciicast](https://asciinema.org/a/4y4udfbh4lf26jbfxzown0zkz.png)](https://asciinema.org/a/4y4udfbh4lf26jbfxzown0zkz)


```
$ oreillyfreebook -d directory [-f format] [-c category] [-s search-word]
```

```
options
  -d directory    Specify the directory to save
  -f format       Specify the ebook format to download; the default is "pdf"
                  "pdf", "mobi", "epub"
  -c category     Specify the ebook category to download;
                  if not specified, all categories will be download
                  "business", "design", "iot", "data", "programming",   "security", "web-platform", "webops"
  -s search-word  Specify the search word for ebook
```

```
$ # ex) download all free ebook
$ oreillyfreebook -d files
Download : files/2016-data-science-salary-survey.pdf
Download : files/business-models-for-the-data-economy.pdf
Download : files/embedding-analytics-in-modern-applications.pdf
Download : files/mapping-big-data.pdf
Download : files/data-science-microsoft-azure-ml-python.pdf
...
```

```
$ # ex) download iot category
$ oreillyfreebook -d files -c iot
Download : files/ambient-computing.pdf
Download : files/governing-the-iot.pdf
Download : files/creating-functional-teams-for-iot.pdf
Download : files/iot-opportunities-challenges.pdf
...
```
