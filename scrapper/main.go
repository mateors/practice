package main

import (
	"fmt"
	"strings"

	"github.com/anaskhan96/soup"
	"github.com/xuri/excelize/v2"
)

type Participant struct {
	Name    string
	Address string
	Email   string
	Website string
	Sector  string
	Country string
}

type keyval struct {
	Key string
	Val string
}

func main() {

	//companyList := make([]*Participant, 0)
	//3956~4508
	//Total participant~553
	//4510
	// Create a new sheet.
	f := excelize.NewFile()
	index := f.NewSheet("Sheet1")
	f.SetCellValue("Sheet1", "A1", "Name")
	f.SetCellValue("Sheet1", "B1", "Address")
	f.SetCellValue("Sheet1", "C1", "Email")
	f.SetCellValue("Sheet1", "D1", "Website")
	f.SetCellValue("Sheet1", "E1", "Sector")
	f.SetCellValue("Sheet1", "F1", "Country")
	count := 0

	for i := 3956; i < 4510; i++ {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("recovered from ", r)
			}
		}()

		fetchUrl := fmt.Sprintf(`https://.com/expo_zone/%d`, i)
		list := fetchData(fetchUrl)

		if len(list) == 6 {

			count++
			fmt.Println(i, list[0].Key, list[0].Val)
			// companyList = append(companyList, &Participant{
			// 	Name:    list[0].Val,
			// 	Address: list[1].Val,
			// 	Email:   list[2].Val,
			// 	Website: list[3].Val,
			// 	Sector:  list[4].Val,
			// 	Country: list[5].Val,
			// })
			//break
			f.SetCellValue("Sheet1", fmt.Sprintf(`A%d`, count+1), list[0].Val)
			f.SetCellValue("Sheet1", fmt.Sprintf(`B%d`, count+1), list[1].Val)
			f.SetCellValue("Sheet1", fmt.Sprintf(`C%d`, count+1), list[2].Val)
			f.SetCellValue("Sheet1", fmt.Sprintf(`D%d`, count+1), list[3].Val)
			f.SetCellValue("Sheet1", fmt.Sprintf(`E%d`, count+1), list[4].Val)
			f.SetCellValue("Sheet1", fmt.Sprintf(`F%d`, count+1), list[5].Val)

			f.SetActiveSheet(index)
			if err := f.SaveAs("participants.xlsx"); err != nil {
				fmt.Println(err)
			}
		}

	}

	// // Create a new sheet.
	// f := excelize.NewFile()
	// index := f.NewSheet("Sheet1")
	// f.SetCellValue("Sheet1", "A1", "Name")
	// f.SetCellValue("Sheet1", "B1", "Address")
	// f.SetCellValue("Sheet1", "C1", "Email")
	// f.SetCellValue("Sheet1", "D1", "Website")
	// f.SetCellValue("Sheet1", "E1", "Sector")
	// f.SetCellValue("Sheet1", "F1", "Country")

	// for i, company := range companyList {

	// 	//fmt.Println(i, company.Country, company.Name)
	// 	//cellPos := fmt.Sprintf(`A%d`, i+1)
	// 	f.SetCellValue("Sheet1", fmt.Sprintf(`A%d`, i+2), company.Name)
	// 	f.SetCellValue("Sheet1", fmt.Sprintf(`B%d`, i+2), company.Address)
	// 	f.SetCellValue("Sheet1", fmt.Sprintf(`C%d`, i+2), company.Email)
	// 	f.SetCellValue("Sheet1", fmt.Sprintf(`D%d`, i+2), company.Website)
	// 	f.SetCellValue("Sheet1", fmt.Sprintf(`E%d`, i+2), company.Sector)
	// 	f.SetCellValue("Sheet1", fmt.Sprintf(`F%d`, i+2), company.Country)
	// }

	// f.SetActiveSheet(index)
	// if err := f.SaveAs("participants.xlsx"); err != nil {
	// 	fmt.Println(err)
	// }

}

func fetchData(fetchUrl string) []*keyval {

	defer func() {
		if r := recover(); r != nil {
			fmt.Println("fetchData:: recovered from ", r)
		}
	}()

	resp, err := soup.Get(fetchUrl) //4442
	if err != nil {
		//os.Exit(1)
		return nil
	}
	doc := soup.HTMLParse(resp)
	// /#profileModal > div
	links := doc.Find("div", "id", "profileModal").FindAll("tr")

	list := make([]*keyval, 0)

	for _, link := range links {
		//fmt.Println(link.Text(), "| Link :", link.Attrs()["href"])
		cells := link.FindAll("td")
		//fmt.Println(link.HTML(), cells[0])

		if len(cells) == 2 {
			key := strings.TrimSpace(cells[0].Text())
			val := strings.TrimSpace(cells[1].Text())
			//fmt.Println(key, val)
			list = append(list, &keyval{key, val})
		}
	}
	return list
}
