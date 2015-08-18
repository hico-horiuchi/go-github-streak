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
	var contributions Contributions

	profile, err := getGitHubProfile(user)
	if err != nil {
		return nil, err
	}

	re := regexp.MustCompile(SPAN_PATTERN)
	matches := re.FindAllStringSubmatch(profile, -1)

	contributions.Contributions, err = strconv.Atoi(removeComma(matches[0][1]))
	if err != nil {
		return nil, err
	}

	contributions.LongestStreak, err = strconv.Atoi(removeComma(matches[1][1]))
	if err != nil {
		return nil, err
	}

	contributions.CurrentStreak, err = strconv.Atoi(removeComma(matches[1][1]))
	if err != nil {
		return nil, err
	}

	return &contributions, nil
}

func getGitHubProfile(user string) (string, error) {
	url := GITHUB_URL + "/" + user
	request, err := http.NewRequest("GET", url, nil)

	response, err := http.DefaultClient.Do(request)
	if err != nil {
		return "", err
	}
	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return "", err
	}

	return string(body), nil
}

func removeComma(s string) string {
	return strings.Replace(s, ",", "", -1)
}
