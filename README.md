## go-github-streak

[![Circle CI](https://circleci.com/gh/hico-horiuchi/go-github-streak.svg?style=shield)](https://circleci.com/gh/hico-horiuchi/go-github-streak)

`go-github-streak` is pacakge to get public contributions from GitHub profile.  
Contributions include "Contributions in the last year", "Longest streak" and "Current streak."

#### Documents

  - [github.com/hico-horiuchi/go-github-streak](http://godoc.org/github.com/hico-horiuchi/go-github-streak)

#### Installation

    $ go get github.com/hico-horiuchi/go-github-streak

#### Example

```go
package main

import (
	"fmt"
	streak "github.com/hico-horiuchi/go-github-streak"
)

func main() {
	contributions, _ := streak.GetContributions("hico-horiuchi")
	fmt.Printf("%+v\n", contributions)
    // &{
    //     Contributions: 1185,
    //     LongestStreak: 13,
    //     CurrentStreak: 0,
    // }
}
```

#### License

go-github-streak is released under the [MIT license](https://raw.githubusercontent.com/hico-horiuchi/go-github-streak/master/LICENSE).
