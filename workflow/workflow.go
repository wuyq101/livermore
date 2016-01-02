package workflow

import (
	"github.com/wuyq101/livermore/model"
)

type WorkFlow struct {
	m *model.Model
}

var w *WorkFlow

func Instance() *WorkFlow {
	if w == nil {
		w = &WorkFlow{
			m: model.NewModel(),
		}
	}
	return w
}

func (w *WorkFlow) Model() *model.Model {
	return w.m
}
