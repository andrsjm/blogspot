package handler

import (
	"blogspot/flow"
	"blogspot/parser"
	"blogspot/util"
	"net/http"
)

type blogHandler struct {
	parser    parser.IBlogParser
	presenter util.IPresenterJSON
	flow      flow.IBlogFlow
}

func NewBlogHandler(parser parser.IBlogParser, presenter util.IPresenterJSON, flow flow.IBlogFlow) *blogHandler {
	return &blogHandler{
		parser:    parser,
		presenter: presenter,
		flow:      flow,
	}
}

func (h *blogHandler) Insert(w http.ResponseWriter, r *http.Request) {
	blog, err := h.parser.ParseBlogEntity(r)
	if err != nil {
		h.presenter.SendError(w, "Error Parsing")
		return
	}

	err = h.flow.Insert(blog)
	if err != nil {
		h.presenter.SendError(w, "Error Insert")
		return
	}

	h.presenter.SendSuccess(w)
}

func (h *blogHandler) Update(w http.ResponseWriter, r *http.Request) {
	blog, err := h.parser.ParseBlogEntity(r)
	if err != nil {
		h.presenter.SendError(w, "Error Parsing")
		return
	}

	err = h.flow.Update(blog)
	if err != nil {
		h.presenter.SendError(w, "Error Update")
		return
	}

	h.presenter.SendSuccess(w)
}

func (h *blogHandler) GetBlog(w http.ResponseWriter, r *http.Request) {
	blogFilter := h.parser.ParseBlogFiler(r)

	blog, err := h.flow.GetBlog(blogFilter)
	if err != nil {
		h.presenter.SendError(w, "Error Get")
		return
	}

	h.presenter.SendSuccess(w, blog)
}

func (h *blogHandler) GetByUser(w http.ResponseWriter, r *http.Request) {
	blogFilter := h.parser.ParseBlogFiler(r)

	blogs, err := h.flow.GetByUser(blogFilter)
	if err != nil {
		h.presenter.SendError(w, "Error Get")
		return
	}

	h.presenter.SendSuccess(w, blogs)
}

func (h *blogHandler) GetHidden(w http.ResponseWriter, r *http.Request) {
	blogFilter := h.parser.ParseBlogFiler(r)

	blogs, err := h.flow.GetHidden(blogFilter)
	if err != nil {
		h.presenter.SendError(w, "Error Get")
		return
	}

	h.presenter.SendSuccess(w, blogs)
}

func (h *blogHandler) Hidden(w http.ResponseWriter, r *http.Request) {
	blogFilter := h.parser.ParseBlogFiler(r)

	err := h.flow.Hidden(blogFilter)
	if err != nil {
		h.presenter.SendError(w, "Error Hidden")
		return
	}

	h.presenter.SendSuccess(w)
}

func (h *blogHandler) Delete(w http.ResponseWriter, r *http.Request) {
	blogID := h.parser.ParseBlogID(r)

	err := h.flow.Delete(blogID)
	if err != nil {
		h.presenter.SendError(w, "Error Delete")
		return
	}

	h.presenter.SendSuccess(w)
}
