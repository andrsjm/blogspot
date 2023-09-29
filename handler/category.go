package handler

import (
	"blogspot/flow"
	"blogspot/parser"
	"blogspot/util"
	"net/http"
)

type categoryHandler struct {
	parser    parser.ICategoryParser
	presenter util.IPresenterJSON
	flow      flow.ICategoryFlow
}

func NewCategoryHandler(parser parser.ICategoryParser, presenter util.IPresenterJSON, flow flow.ICategoryFlow) *categoryHandler {
	return &categoryHandler{
		parser:    parser,
		presenter: presenter,
		flow:      flow,
	}
}

func (h *categoryHandler) Insert(w http.ResponseWriter, r *http.Request) {
	category, err := h.parser.ParseCategoryEntity(r)
	if err != nil {
		h.presenter.SendError(w, "Error Parsing")
		return
	}

	err = h.flow.Insert(category)
	if err != nil {
		h.presenter.SendError(w, "Error Insert")
		return
	}

	h.presenter.SendSuccess(w)
}
