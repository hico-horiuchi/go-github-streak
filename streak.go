/*
go-github-streak is pacakge to get public cotributions from GitHub profile.
Conrtibutions include "Contributions in the last year", "Longest streak" and "Current streak."

https://github.com/hico-horiuchi/go-github-streak

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

*/
package streak

import (
	"io/ioutil"
	"net/http"
	"regexp"
	"strconv"
	"strings"
)

const (
	GITHUB_URL   = "https://github.com"
	SPAN_PATTERN = `<span class="contrib-number">([0-9,]+) [a-z]+</span>`
)

type Contributions struct {
	Contributions int `Contributions in the last year`
	LongestStreak int `Longest streak`
	CurrentStreak int `Current streak`
}

// Get public contributions from GitHub profile.
func GetContributions(user string) (*Contributions, error) {
	url := GITHUB_URL + "/" + user
	request, err := http.NewRequest("GET", url, nil)

	response, err := http.DefaultClient.Do(request)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	re := regexp.MustCompile(SPAN_PATTERN)
	matches := re.FindAllStringSubmatch(string(body), -1)
	return &Contributions{
		Contributions: atoi(removeComma(matches[0][1])),
		LongestStreak: atoi(removeComma(matches[1][1])),
		CurrentStreak: atoi(removeComma(matches[2][1])),
	}, nil
}

func atoi(s string) int {
	v, _ := strconv.Atoi(s)
	return v
}

func removeComma(s string) string {
	return strings.Replace(s, ",", "", -1)
}
