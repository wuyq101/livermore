package workflow

import "qxf-backend/logger"

func (w *WorkFlow) ShowByMoneyFlow(args []string) error {
	stocks, err := w.m.GetAllStocks()
	if err != nil {
		return err
	}
	for _, stock := range stocks {
		logger.Info("start to check stock %+v", stock)
		//读取该股票的所有money flow
		mfs, err := w.m.GetMoneyFlows(stock.Code)
		if err != nil {
			logger.Error("Failed to get money flow records, err: %+v", err)
			continue
		}

		//统计所有流入资金的和，超大单和大单 净流入
		sum := int64(0)
		for _, mf := range mfs {
			sum += mf.R0Net + mf.R1Net
		}
		//一共流入资金为
		logger.Info("一共有%d条记录, 净流入资金为 %d万", len(mfs), int64(sum/1000000.0))
		daily, err := w.m.GetLastStockDaily(stock.Code)
		if err != nil {
			logger.Error("Failed to get last stock daily records, err: %+v", err)
			continue
		}
		logger.Info("Last stock daily %+v", daily)
		//总流通股本
		total := daily.CurPrice * stock.CurrCapital
		if total != 0 {
			ratio := float64(sum * 1.0 / total)
			logger.Info("total %d, sum %d , ratio %.f", total, sum, ratio)
		} else {
			logger.Error("错误 %+v", stock)
		}
	}
	return nil
}
