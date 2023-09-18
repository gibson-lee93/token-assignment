package service

import (
	"encoding/hex"
	"encoding/json"
	"log"
	"math/big"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"

	"token-assignment/internal/project/dto"
	"token-assignment/internal/project/entity"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/rpc"
)

var bitfinetSymbolMap = map[string]string{
	"USDT": "USTUSD",
	"USDC": "UDCUSD",
	"ETH":  "ETHUSD",
}

var chainlinkSymbolMap = map[string]string{
	"USDT": "0x22f15736aB9A5b944B6ec41bAeE2782e0f1B7724",
	"USDC": "0x053Fc65249dF91a02Ddb294A081f774615aB45F4",
	"ETH":  "0x143db3CEEfbdfe5631aDD3E50f7614B6ba708BA7",
}

func (svc *ProjectSvc) GetTokenInfo(dtoReq dto.GetTokenInfoReq) (dtoResp dto.GetTokenInfoResp, err error) {
	entReq := entity.GetTokenInfoReq{
		Symbol:    dtoReq.Symbol,
		StartTime: dtoReq.StartTime,
		EndTime:   dtoReq.EndTime,
	}

	var tokenInfoList []entity.TokenInfo
	if dtoReq.Source != "" {
		tokenInfoList = make([]entity.TokenInfo, 0, 1)
		entReq.Source = dtoReq.Source
		entResp, err := svc.repository.GetTokenInfo(entReq)
		if err != nil {
			return dtoResp, err
		}
		if entResp.Price > 0 {
			tokenInfoList = append(tokenInfoList, entResp)
		}

	} else {
		tokenInfoList = make([]entity.TokenInfo, 0, 2)
		entReq.Source = "bitfinex"
		entResp, err := svc.repository.GetTokenInfo(entReq)
		if err != nil {
			return dtoResp, err
		}
		if entResp.Price > 0 {
			tokenInfoList = append(tokenInfoList, entResp)
		}

		entReq.Source = "chainlink"
		entResp, err = svc.repository.GetTokenInfo(entReq)
		if err != nil {
			return dtoResp, err
		}
		if entResp.Price > 0 {
			tokenInfoList = append(tokenInfoList, entResp)
		}
	}

	dtoResp.ToDTO(tokenInfoList)
	return dtoResp, nil
}

func (svc *ProjectSvc) GetBitfinexTokenInfo(tokenSymbol string) error {
	resp, err := http.Get("https://api.bitfinex.com/v1/pubticker/" + bitfinetSymbolMap[tokenSymbol])
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

	// convert price's string value to float
	price, err := strconv.ParseFloat(bitfinexResp.Price, 64)
	if err != nil {
		log.Println(err)
		return err
	}

	tokenInfoEntity := entity.TokenInfo{
		Symbol:    tokenSymbol,
		Price:     float32(price),
		Source:    "bitfinex",
		Timestamp: time.Now(),
	}

	if err = svc.repository.CreateTokenInfo(tokenInfoEntity); err != nil {
		return err
	}

	return nil
}

func (svc *ProjectSvc) GetChainlinkTokenInfo(tokenSymbol, contractABI, rpcURL string) error {
	client, err := rpc.Dial(rpcURL)
	if err != nil {
		log.Println("Failed to connect to the Ethereum client: ", err)
		return err
	}

	contractAbi, err := abi.JSON(strings.NewReader(contractABI))
	if err != nil {
		log.Println("Failed to parse contract ABI: ", err)
		return err
	}

	contractAddress := common.HexToAddress(chainlinkSymbolMap[tokenSymbol])
	// call latestRoundData function
	callMsg := map[string]interface{}{
		"to":   contractAddress.Hex(),
		"data": "0x" + hex.EncodeToString(contractAbi.Methods["latestRoundData"].ID),
	}

	var result string
	err = client.Call(&result, "eth_call", callMsg, "latest")
	if err != nil {
		log.Println("Failed to call contract method: ", err)
		return err
	}

	// convert result (hex string) to byte slice
	resultBytes, err := hex.DecodeString(strings.TrimPrefix(result, "0x"))
	if err != nil {
		log.Println("Failed to decode hex: ", err)
		return err
	}

	var chainlinkResp dto.ChainLinkTokenResp
	err = contractAbi.UnpackIntoInterface(&chainlinkResp, "latestRoundData", resultBytes)
	if err != nil {
		log.Println("Failed to unpack data: ", err)
		return err
	}

	// convert price
	divisor := big.NewFloat(1e8)
	chainlinkPriceFloat := new(big.Float).SetInt(chainlinkResp.Answer)
	priceBigFloat := new(big.Float).Quo(chainlinkPriceFloat, divisor)
	price, _ := priceBigFloat.Float32()

	tokenInfoEntity := entity.TokenInfo{
		Symbol:    tokenSymbol,
		Price:     price,
		Source:    "chainlink",
		Timestamp: time.Now(),
	}

	if err = svc.repository.CreateTokenInfo(tokenInfoEntity); err != nil {
		return err
	}

	return nil
}

func (svc *ProjectSvc) GetTokenInfoScheduler() {
	ticker := time.NewTicker(30 * time.Second)

	chainlinkABI := ConvertAbiToString()
	// BSC Testnet
	rpcURL := "https://data-seed-prebsc-1-s1.binance.org:8545/"

	log.Println("Running Scheduler")

	// Loop to handle ticks
	for {
		select {
		case <-ticker.C:
			log.Println("Fetching Token Information")

			// get bitfinex token info
			svc.GetBitfinexTokenInfo("USDT")
			svc.GetBitfinexTokenInfo("USDC")
			svc.GetBitfinexTokenInfo("ETH")

			// get chainlink token info
			svc.GetChainlinkTokenInfo("USDT", chainlinkABI, rpcURL)
			svc.GetChainlinkTokenInfo("USDC", chainlinkABI, rpcURL)
			svc.GetChainlinkTokenInfo("ETH", chainlinkABI, rpcURL)
		}
	}
}

func ConvertAbiToString() string {
	abiFile, err := os.ReadFile("./internal/project/abi/ChainLink.abi")
	if err != nil {
		log.Println("Error reading ABI file:", err)
	}

	var abiJSON interface{}
	err = json.Unmarshal(abiFile, &abiJSON)
	if err != nil {
		log.Println("Error unmarshalling JSON:", err)
	}

	abiString, err := json.Marshal(abiJSON)
	if err != nil {
		log.Println("Error marshalling back to JSON:", err)
	}
	return string(abiString)
}
