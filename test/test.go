package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type riotAccount struct {
	ID            string `json:"id"`
	AccountID     string `json:"accountId"`
	Puuid         string `json:"puuid"`
	Name          string `json:"name"`
	ProfileIconID int    `json:"profileIconId"`
	RevisionDate  int64  `json:"revisionDate"`
	SummonerLevel int    `json:"summonerLevel"`
}

func main() {
	//name := ""
	var sName string
	fmt.Println("Summoner name:")
	fmt.Scanln(&sName)

	var account riotAccount
	json.Unmarshal(getSummoner(sName), &account)

	fmt.Println(account.Name, account.SummonerLevel)

	//json.Unmarshal(bodyBytes, &account)

	//sLevel := getSummoner(sName)
	//getSummoner(sName)
	//fmt.Println(sLevel)

}

func getSummoner(name string) []byte {
	client := &http.Client{}
	req, _ := http.NewRequest("GET", "https://na1.api.riotgames.com/lol/summoner/v4/summoners/by-name/"+name, nil)
	req.Header.Set("X-Riot-Token", "RGAPI-2e6fecae-49ee-48eb-9aba-98fc67cc0d99")
	res, err := client.Do(req)
	if err != nil {
		log.Fatalln(err)
	}
	bodyBytes, _ := ioutil.ReadAll(res.Body)
	//fmt.Println(reflect.TypeOf(bodyBytes))

	//fmt.Println(bodyBytes)
	// bodyString := string(bodyBytes)
	// fmt.Println("API Response as String:\n" + bodyString)

	//fmt.Println(reflect.TypeOf(bodyString))

	//var account riotAccount
	//json.Unmarshal(bodyBytes, &account)

	//fmt.Println(account)
	//fmt.Printf("%+v\n", sAccount)
	//fmt.Println(sAccount.SummonerLevel)

	defer res.Body.Close()

	//return account.SummonerLevel
	return bodyBytes
}

func showRank(sa struct{}) {

	fmt.Print()

}
