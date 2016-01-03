package workflow

import (
	ui "github.com/gizak/termui"
	"github.com/wuyq101/livermore/logger"
	"github.com/wuyq101/livermore/model"
	"time"
)

func (w *WorkFlow) DisplayStockInfo(codes []string) error {
	code := codes[0]
	stock, err := w.m.GetStockByCode(code)
	if err != nil {
		logger.Error("Failed to get stock by code %s, err %+v", code, err)
		return err
	}
	mfs, err := w.m.GetMoneyFlows(code)
	if err != nil {
		logger.Error("Failed to get stock money flow record for %s, err %+v", code, err)
		return err
	}
	logger.Info("mfs len %d", len(mfs))
	err = ui.Init()
	if err != nil {
		return err
	}
	defer ui.Close()
	bc := ui.NewBarChart()
	bc.BorderLabel = stock.Name + " 净流入"
	startDay := time.Now()
	data, labels := w.getRenderData(startDay, mfs)
	bc.Width = 120
	bc.Height = 40
	bc.Data = data
	bc.DataLabels = labels
	bc.TextColor = ui.ColorGreen
	bc.BarColor = ui.ColorRed
	bc.NumColor = ui.ColorYellow
	bc.BarWidth = 6
	ui.Render(bc)
	ui.Handle("/sys/kbd/q", func(ui.Event) {
		ui.StopLoop()
	})
	ui.Handle("/sys/kbd/h", func(ui.Event) {
		startDay = startDay.AddDate(0, 0, -1)
		bc.Data, bc.DataLabels = w.getRenderData(startDay, mfs)
		ui.Render(bc)
	})
	ui.Handle("/sys/kbd/l", func(ui.Event) {
		startDay = startDay.AddDate(0, 0, 1)
		bc.Data, bc.DataLabels = w.getRenderData(startDay, mfs)
		ui.Render(bc)
	})
	ui.Loop()
	return nil
}

func (w *WorkFlow) getRenderData(day time.Time, mfs []*model.MoneyFlow) ([]int, []string) {
	firstDay := day.AddDate(0, -1, 0)
	data := make([]int, 0)
	labels := make([]string, 0)
	for _, wf := range mfs {
		t, _ := time.ParseInLocation("20060102", wf.MarketDay, time.Local)
		if t.After(firstDay) && len(data) < 30 {
			data = append(data, int(wf.ClosingPrice))
			labels = append(labels, t.Format("01-02"))
		} else {
			break
		}
	}
	for i, j := 0, len(data)-1; i < j; i, j = i+1, j-1 {
		data[i], data[j] = data[j], data[i]
	}
	for i, j := 0, len(labels)-1; i < j; i, j = i+1, j-1 {
		labels[i], labels[j] = labels[j], labels[i]
	}
	return data, labels
}
