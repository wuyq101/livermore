package workflow

import (
	"github.com/wuyq101/livermore/model"
)

type WorkFlow struct {
	m *model.Model
}

func NewWorkFlow(m *model.Model) *WorkFlow {
	return &WorkFlow{m: m}
}

func (w *WorkFlow) Model() *model.Model {
	return w.m
}

func (w *WorkFlow) UpdateStockMeta() error {
	w.m.InsertStock(&model.Stock{})
	return nil
}
