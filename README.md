## go-github-streak

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
