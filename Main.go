package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"strings"
)

type uni struct {
	name string
	id   int
}

func main() {
	url := "https://yokatlas.yok.gov.tr/lisans-anasayfa.php"
	resp, err := http.Get(url)
	if err != nil {
		log.Fatal("Page not found")
		panic(err)

	}
	defer resp.Body.Close()
	content, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)

	}
	strContent := string(content[:])
	scanner := bufio.NewScanner(strings.NewReader(strContent))
	index := 0
	mapOfUni := make(map[int]uni)
	for scanner.Scan() {
		var temp string
		temp = scanner.Text()
		if strings.Contains(temp, "option value") && strings.Contains(temp, "ÜNİVERSİTESİ") {
			temp = strings.Trim(temp, " ")
			temp = strings.TrimLeft(temp, " ")
			temp = strings.TrimLeft(temp, "															")
			temp = strings.ReplaceAll(temp, "<option value=\"", "")
			temp = strings.ReplaceAll(temp, "\">", " ")
			temp = strings.ReplaceAll(temp, "</option>", "")
			temp = strings.ReplaceAll(temp, "<", "")
			num, err := strconv.Atoi(temp[0:4])
			if err != nil {
				fmt.Print("atoi not working")
			}
			var tempUni = new(uni)
			tempUni.id = num
			tempUni.name = temp[5:]
			mapOfUni[index] = *tempUni
		}
	}

}
