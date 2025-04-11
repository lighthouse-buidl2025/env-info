package clients

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

func FetchContractTags(address string) (string, bool, error) {
	url := fmt.Sprintf("https://etherscan.io/address/%s", address)
	res, err := http.Get(url)
	if err != nil {
		return "", false, err
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return "", false, fmt.Errorf("status code %d", res.StatusCode)
	}

	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		return "", false, err
	}

	var isToken bool
	doc.Find("span.badge.text-dark").Each(func(i int, s *goquery.Selection) {
		text := strings.TrimSpace(s.Text())
		if strings.Contains(text, "ERC-") && isToken == false {
			isToken = true
		}
	})

	var nameTag string
	if isToken {
		doc.Find("span.hash-tag").Each(func(i int, s *goquery.Selection) {
			text := s.Text()
			if text != "" && nameTag == "" && i == 0 {
				text = strings.TrimSpace(text)
				text = strings.ToLower(text)
				parts := strings.Split(text, ":")
				nameTag = parts[0]
			}
		})
	} else {
		doc.Find("span.hash-tag").Each(func(i int, s *goquery.Selection) {
			text := s.Text()
			if text != "" && nameTag == "" && i == 1 {
				text = strings.TrimSpace(text)
				text = strings.ToLower(text)
				text = strings.ReplaceAll(text, " ", "")
				nameTag = text
			}
		})
	}

	if nameTag == "" {
		return address, false, fmt.Errorf("couldn't find contract tag")
	}

	return nameTag, isToken, nil
}
