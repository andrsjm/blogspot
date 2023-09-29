package handler

import (
	"blogspot/flow"
	"blogspot/parser"
	"blogspot/util"
	"net/http"
)

type tagHandler struct {
	parser    parser.ITagParser
	presenter util.IPresenterJSON
	flow      flow.ITagFlow
}

func NewTagHandler(parser parser.ITagParser, presenter util.IPresenterJSON, flow flow.ITagFlow) *tagHandler {
	return &tagHandler{
		parser:    parser,
		presenter: presenter,
		flow:      flow,
	}
}

func (h *tagHandler) Insert(w http.ResponseWriter, r *http.Request) {
	tag, err := h.parser.ParseTagEntity(r)
	if err != nil {
		h.presenter.SendError(w, "Error Parsing")
		return
	}

	err = h.flow.Insert(tag)
	if err != nil {
		h.presenter.SendError(w, "Error Insert")
		return
	}

	h.presenter.SendSuccess(w)
}
