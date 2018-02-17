# leap [![Travis Build Status](https://img.shields.io/travis/gmcabrita/leap.svg)](https://travis-ci.org/gmcabrita/leap) [![GoDoc](https://godoc.org/github.com/gmcabrita/leap?status.svg)](https://godoc.org/github.com/gmcabrita/leap) [![Go Report Card](https://goreportcard.com/badge/gmcabrita/leap)](https://goreportcard.com/report/gmcabrita/leap) [![License](https://img.shields.io/github/license/gmcabrita/leap.svg)](https://github.com/gmcabrita/leap/blob/master/LICENSE)

Leap implements the jump consistent hashing algorithm proposed by John Lamping and Eric Veach.<sup
id="anchor-paper">[1](#footnote-paper)</sup>

## Usage

```go
package main

import (
	"fmt"

	"github.com/gmcabrita/leap"
)

func main() {
	key := uint64(42)
	bucketPlacement := leap.Hash(key, 100)

	fmt.Println(bucketPlacement)
}

```

[^1]: https://arxiv.org/pdf/1406.2294.pdf

<sup id="footnote-paper">1</sup> https://arxiv.org/pdf/1406.2294.pdf [↩](#anchor-paper)
