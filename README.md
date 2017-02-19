# O'Reilly Free Ebook Downloader


## Install

```
$ wget  https://github.com/mitakeck/oreillyfreebook/releases/download/v0.0.1/oreillyfreebook
$ mv oreillyfreebook /usr/loca/bin/
```

## Usage


```
$ oreillyfreebook -d directory [-f format] [-c category]
```

```
option
  -d directory  Specify the directory to save
  -f format     Specify the ebook format to download; the default is "pdf"
                "pdf", "mobi", "epub"
  -c category   Specify the ebook category to download;
                if not specified, all categories will be download
                "business", "design", "iot", "data", "programming", "security", "web-platform", "webops"
```

```
$ oreillyfreebook -d files
Download : file/2016-data-science-salary-survey.pdf
Download : file/business-models-for-the-data-economy.pdf
Download : file/embedding-analytics-in-modern-applications.pdf
Download : file/mapping-big-data.pdf
Download : file/data-science-microsoft-azure-ml-python.pdf
...
```
