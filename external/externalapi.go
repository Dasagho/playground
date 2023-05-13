package external

import (
	"fmt"

	"github.com/PuerkitoBio/goquery"
	"github.com/dasagho/playground/api/client"
	"github.com/dasagho/playground/api/model"
)

func GetPoliformatMainPageData(doc goquery.Document, client client.Client) []model.Subject {
	var SubjectList []model.Subject
	// Search Subjects Link
	doc.Find("#otherSitesCategorWrap > div.moresites-right-col > div:nth-child(1)").First().Find("li > div > a").Each(
		func(i int, s *goquery.Selection) {
			href, exHref := s.Attr("href")
			name, exTitle := s.Attr("title")
			if exHref && exTitle {
				SubjectList = append(SubjectList, GetSubjectData(name, href, client))
			}
		},
	)
	return SubjectList
}

func GetSubjectData(name string, link string, client client.Client) model.Subject {
	res, err := client.Get(link)
	if err != nil {
		panic("Error subject request " + err.Error())
	}
	defer res.Body.Close()

	if res.StatusCode != 200 {
		panic(fmt.Sprintf("Error on subject request. Status Code: %d", res.StatusCode))
	}
	var Subject = model.Subject{Name: name, Link: link}
	document, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		panic("Error parsing HTML response" + err.Error())
	}

	recursosHref, exRecursos := document.Find("#toolMenu > ul > li:nth-child(3) > a").First().Attr("href")
	if exRecursos {
		Subject.Recursos = recursosHref
	}

	return Subject
}
