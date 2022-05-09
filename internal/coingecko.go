package internal

import (
	"fmt"
	"strconv"

	"encoding/json"
	"net/http"
	"time"

	"github.com/defi-data-analysis/address-analysis/log"
	_ "github.com/ethereum/go-ethereum/common"
	"github.com/rur0/coingecko"
)

var baseurl = "https://api.coingecko.com/api/v3"

type ContractInfo struct {
	ID     string    `json:"id"`
	Symbol string    `json:"symbol"`
	Name   string    `json:"name"`
	Images ImageInfo `json:"image"`
}
type ImageInfo struct {
	Thumb string `json:"thumb"`
	Small string `json:"small"`
	Large string `json:"large"`
}

func GetContractInfo(coinId string, contractAddress string) (info ContractInfo, err error) {
	call := baseurl + "/coins/" + coinId + "/contract/" + contractAddress
	fmt.Println(call)
	//resp, err := http.Get(call)

	req, err := http.NewRequest("GET", call, nil)
	if err != nil {
		return info, err
	}
	req.Header.Add("User-Agent", "Mozilla/5.0")
	req.Header.Add("accept", "application/json")
	client := &http.Client{}
	resp, err := client.Do(req)

	if err != nil {
		fmt.Println(err)
		return info, err
	}
	defer resp.Body.Close()

	err = json.NewDecoder(resp.Body).Decode(&info)
	if err != nil {
		fmt.Println(err)
		return info, err
	}
	return info, nil
}

func GetContractPrice(chainid string, vs_currency string, contractAddress string, from string) (market coingecko.MarketChart, err error) {
	call := baseurl + "/coins/" + chainid + "/contract/" + contractAddress + "/market_chart/range?vs_currency=" + vs_currency + "&from=" + from + "&to=" + strconv.FormatInt(time.Now().Unix(), 10)
	//resp, err := http.Get(call)

	req, err := http.NewRequest("GET", call, nil)
	if err != nil {
		return market, err
	}
	req.Header.Add("User-Agent", "Mozilla/5.0")
	req.Header.Add("accept", "application/json")
	client := &http.Client{}
	resp, err := client.Do(req)

	if err != nil {
		log.Logger().Error(fmt.Sprintf("%v \r\nfailed to GetContractPrice: %v", call, err))
		return market, err
	}
	defer resp.Body.Close()

	err = json.NewDecoder(resp.Body).Decode(&market)
	if err != nil {
		log.Logger().Error(fmt.Sprintf("%v \r\nfailed to NewDecoder: %v", call, err))
		return market, err
	}
	return market, nil
}

func GetContractPriceById(chainid string, vs_currency string, from string) (market coingecko.MarketChart, err error) {
	call := baseurl + "/coins/" + chainid + "/market_chart/range?vs_currency=" + vs_currency + "&from=" + from + "&to=" + strconv.FormatInt(time.Now().Unix(), 10)
	//resp, err := http.Get(call)

	req, err := http.NewRequest("GET", call, nil)
	if err != nil {
		return market, err
	}
	req.Header.Add("User-Agent", "Mozilla/5.0")
	req.Header.Add("accept", "application/json")
	client := &http.Client{}
	resp, err := client.Do(req)

	if err != nil {
		log.Logger().Error(fmt.Sprintf("%v \r\nfailed to GetContractPrice: %v", call, err))
		return market, err
	}
	defer resp.Body.Close()

	err = json.NewDecoder(resp.Body).Decode(&market)
	if err != nil {
		log.Logger().Error(fmt.Sprintf("%v \r\nfailed to NewDecoder: %v", call, err))
		return market, err
	}
	return market, nil
}
