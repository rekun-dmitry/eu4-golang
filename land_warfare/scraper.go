package main

import (
	"fmt"
	"os"
	"regexp"

	"github.com/PuerkitoBio/goquery"
	"github.com/gocolly/colly"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type terrain struct {
	supplyLimit     string `db:"supply_limit"`
	localDef        string `db:"local_defensiveness"`
	movementCost    string `db:"movement_cost"`
	attackerPenalty string `db:"attacker_penalty"`
	localDevCost    string `db:"local_development_cost"`
	terra           string `db:"terrain_name"`
}

func main() {
	crawl()
}

func crawl() {
	re := regexp.MustCompile(`\r?\n`)
	dbPass := os.Getenv("db_pass")
	db, err := sqlx.Connect("postgres", fmt.Sprintf("postgresql://localhost/eu4?user=postgres&password=%s&sslmode=disable", dbPass))
	if err != nil {
		fmt.Println("error:", err)
		return
	}
	c := colly.NewCollector(
		colly.AllowedDomains("eu4.paradoxwikis.com"),
	)
	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting...", r.URL.String())
	})
	c.OnHTML("body", func(body *colly.HTMLElement) {
		tmpTerrain := terrain{}
		body.DOM.Find("table.wikitable").First().Find("tr>td").Each(func(i int, s *goquery.Selection) {
			symbol := s.Find("td span").First().Text()
			k := i % 6
			switch k {
			case 0:
				terraType := re.ReplaceAllString(s.Text(), "")
				tmpTerrain.terra = terraType
			case 1:
				tmpTerrain.supplyLimit = symbol
			case 2:
				tmpTerrain.localDef = symbol
			case 3:
				tmpTerrain.movementCost = symbol
			case 4:
				tmpTerrain.attackerPenalty = symbol
			case 5:
				tmpTerrain.localDevCost = symbol
				db.MustExec("INSERT INTO land_warfare.terrain_war (terrain_name, supply_limit, local_defensiveness, movement_cost, attacker_penalty, local_development_cost) VALUES ($1, $2, $3, $4, $5, $6)", tmpTerrain.terra, tmpTerrain.supplyLimit, tmpTerrain.localDef, tmpTerrain.movementCost, tmpTerrain.attackerPenalty, tmpTerrain.localDevCost)
				tmpTerrain = terrain{}
			}
		})
	})
	c.Visit("https://eu4.paradoxwikis.com/Land_warfare")
}
