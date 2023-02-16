package parser

import (
	"github.com/emptywe/trading_sim/internal/storage/postgres/parser_repo"
	"github.com/emptywe/trading_sim/pkg/binance/binancews"
	"go.uber.org/zap"
	"strconv"
	"strings"
)

type Parser struct {
	repo         *parser_repo.Repository
	poolSize     int
	currencyList []string
}

func NewParser(repo *parser_repo.Repository, poolSize int, currencyList []string) *Parser {
	return &Parser{repo: repo, poolSize: poolSize, currencyList: currencyList}
}

func (a *Parser) parseBinanceData(Data chan binancews.DataPrice) {
	for i := range Data {
		price, _ := strconv.ParseFloat(i.Price, 64)
		if err := a.repo.UpdateCurrency(strings.ToLower(i.Symbol), price); err != nil {
			zap.S().Errorf("can't update currency: %v", err)
		}
	}
}

func (a *Parser) createWorker(list []string) {
	wsClient := binancews.NewBinanceWSClient()
	go wsClient.WSHandlerBinance(list)
	a.parseBinanceData(wsClient.Data)
}

func (a *Parser) currencyUpdater() {
	var list []string
	for _, v := range a.currencyList {
		list = append(list, v+binancews.Trade)
		if len(list)%a.poolSize == 0 {
			zap.S().Infof("Creating worker on %v", list)
			go a.createWorker(list)
			list = []string{}
			continue
		}
	}
	if len(list) > 0 {
		go a.createWorker(list)
	}
}

func (a *Parser) createCurrencies() {
	for _, cur := range a.currencyList {
		err := a.repo.CreateNewCurrency(cur)
		if err != nil {
			zap.S().Errorf("can't create currency storage: %v", err)
		}
	}
}

func (a *Parser) InitParser() {
	a.createCurrencies()
	a.currencyUpdater()
}
