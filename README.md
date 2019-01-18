# color-id

[![GoDoc](https://godoc.org/github.com/alex-ant/color-id?status.svg)](https://godoc.org/github.com/alex-ant/color-id)

A library generating random colors for each supplied ID.

### Sample use case

An application where each user upon the registration process is being assigned a unique color - random, but distinct enough, so that another user won't have the same or a similar one.

Another use case could be any type of list where each item has to be visually unique.

### Logic and examples

Let's create 5 pseudo-random colors:

```
package main

import (
	"fmt"
	"strconv"

	"github.com/alex-ant/color-id"
)

func main() {
	s := color.NewSet()

	for i := 0; i < 5; i++ {
		id := strconv.Itoa(i)
		hexColor := s.GetColor(id).Hex()
		fmt.Println(hexColor)
	}
}

// OUTPUT:
// #00B1FF
// #FFD200
// #9F00FF
// #00FF00
// #FF1600

```

A single ID will have the same color on each call in the same `*Set`, useful for API integration.

The package uses a flat RGB palette as a scale from 0 to 1. It generates the first color in `*Set` randomly, and then looks for the widest "gap" on the scale to get the next requested color from the middle of it. The picture below describes the distribution of colors across the palette in the mentioned example.

![alt text](https://raw.githubusercontent.com/alex-ant/color-id/master/distribution.png "Palette color distribution")
