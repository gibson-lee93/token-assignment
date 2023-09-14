package service

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"time"
	"token-assignment/internal/project/dto"
	"token-assignment/internal/project/entity"
)

func (svc *ProjectSvc) GetBitfinexTokenInfo(tokenSymbol string) error {
	resp, err := http.Get("https://api.bitfinex.com/v1/pubticker/" + tokenSymbol)
	if err != nil {
		log.Println(err)
		return err
	}
	defer resp.Body.Close()

	var bitfinexResp dto.GetBitfinextTokenResp
	if err = json.NewDecoder(resp.Body).Decode(&bitfinexResp); err != nil {
		log.Println(err)
		return err
	}

	// convert unix timestamp to time.Time
	unixTimeStamp, err := strconv.ParseFloat(bitfinexResp.Timestamp, 64)
	if err != nil {
		log.Println(err)
		return err
	}

	seconds := int64(unixTimeStamp)
	nanoseconds := int64((unixTimeStamp - float64(seconds)) * 1e9)
	timestamp := time.Unix(seconds, nanoseconds)

	// convert price's string value to float
	price, err := strconv.ParseFloat(bitfinexResp.Price, 64)
	if err != nil {
		log.Println(err)
		return err
	}

	tokenInfoEntity := entity.TokenInfo{
		Symbol:    tokenSymbol,
		Price:     float32(price),
		Source:    "Bitfinex",
		Timestamp: timestamp,
	}

	if err = svc.repository.CreateTokenInfo(tokenInfoEntity); err != nil {
		return err
	}

	return nil
}
