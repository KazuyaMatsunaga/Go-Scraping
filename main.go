package main

import (
	"fmt"
	"time"
	"strings"
    "github.com/PuerkitoBio/goquery"
)

type Game struct {
	Title string
	Price int
	ReleaseDate string
	Genre string
	Link string
	ImageLink string
}

func extraReplacer(str string) string {
	var afterStr string

	afterStr = strings.Replace(str," ","",-1)
	afterStr = strings.Replace(str,"\n","",-1)

	return afterStr
}

func main() {
	baseURL := "https://w.atwiki.jp/gcmatome/"

	doc, _ := goquery.NewDocument(baseURL)

	var gameHardTitle string
	var gameHardLink string

	// var gameLink string
	// var gameGenre string

	var gameURL string

	var gameList []Game

	mapGames := make(map[string]map[string]string)

	doc.Find("ul.pl-l-level-2 a[title]").Each(func(_ int, s * goquery.Selection) {
		titleAttr, _ := s.Attr("title")
		if c := strings.Contains(titleAttr,"ゲーム記事一覧 (PS4)"); c {
			hrefsStr , _ := s.Attr("href")
			gameHardTitle = titleAttr
			gameHardLink = hrefsStr
		}
	})

	fmt.Printf("%s: %s\n", gameHardTitle, gameHardLink)

	gameListURL := "https:" + gameHardLink

	docGameList, _ := goquery.NewDocument(gameListURL)

	docGameList.Find("table[cellspacing] a[title]").Each(func(_ int, s * goquery.Selection) {
		titleAttr, _ := s.Attr("title")
		if c := strings.Contains(titleAttr,"ペルソナ5"); c {
			genreStr := s.ParentFiltered("td").Next().Text()
			genreStr = extraReplacer(genreStr)
			titleStr := s.Text()
			titleStr = extraReplacer(titleStr)
			hrefAttr , _ := s.Attr("href")
			hrefStr := "https:" + hrefAttr
			mapGames[titleStr] = make(map[string]string)
			mapGames[titleStr]["link"] = hrefStr
			mapGames[titleStr]["genre"] = genreStr
		}
	})

	for key, value := range mapGames {
		fmt.Printf("%s: %s , %s\n",key, value["genre"], value["link"])
	}

	for key, value := range mapGames {
		gameURL = value["link"]
		docGame, _ := goquery.NewDocument(gameURL)

		docGame.Find("h2[id]").Each(func(_ int, s * goquery.Selection) {
			if extraReplacer(s.Text()) == key {
				// fmt.Println("if true")
				sTable := s.Next().Next()
				sTable.Find("td").Each(func(_ int, sd * goquery.Selection) {
					if extraReplacer(sd.Text()) == "定価" {
						price := sd.Next().Text()
						price = extraReplacer(price)
						value["price"] = price
					}
					if extraReplacer(sd.Text()) == "発売日" {
						releaseDate := sd.Next().Text()
						releaseDate = extraReplacer(releaseDate)
						value["release_date"] = releaseDate
					}
				})
				value["image_link"], _ = sTable.Find("img").Attr("src")
			}
		})
	}

	for key, value := range mapGames {
		var boxGame Game

		
	}

	for key, value := range mapGames {
		fmt.Printf("%s: %s, %s, %s, %s, %s\n",key, value["price"], value["release_date"], value["genre"], value["link"], value["image_link"])
	}
}

	