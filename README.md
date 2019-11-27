# go-record-jar

[![Build Status](https://travis-ci.com/Xuanwo/go-record-jar.svg?branch=master)](https://travis-ci.com/Xuanwo/go-record-jar)
[![GoDoc](https://godoc.org/github.com/Xuanwo/go-record-jar?status.svg)](https://godoc.org/github.com/Xuanwo/go-record-jar)
[![Go Report Card](https://goreportcard.com/badge/github.com/Xuanwo/go-record-jar)](https://goreportcard.com/report/github.com/Xuanwo/go-record-jar)
[![codecov](https://codecov.io/gh/Xuanwo/go-record-jar/branch/master/graph/badge.svg)](https://codecov.io/gh/Xuanwo/go-record-jar)
[![License](https://img.shields.io/badge/license-apache%20v2-blue.svg)](https://github.com/Xuanwo/go-record-jar/blob/master/LICENSE)

go-record-jar provide support for [Record jar](https://tools.ietf.org/html/draft-phillips-record-jar-01)

## Quickstart

Let's parse the Language Subtag Registry spec

```go
url := "https://www.iana.org/assignments/language-subtag-registry/language-subtag-registry"
resp, err := http.Get(url)
if err != nil {
    log.Fatalf("downloadLanguageSubtagRegistry http.Get failed for %v", err)
}
defer resp.Body.Close()

content, err := ioutil.ReadAll(resp.Body)
if err != nil {
    log.Fatalf("downloadLanguageSubtagRegistry ioutil.ReadAll failed for %v", err)
}

data := recordjar.Parse(content)
```