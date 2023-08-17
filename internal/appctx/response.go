// Package appctx
package appctx

import (
	"strings"
	"sync"

	"github.com/aryayunanta-ralali/shorty/pkg/msg"
)

var rsp *Response
var oneRsp sync.Once

// Response presentation contract object
type Response struct {
	Name        string            `json:"name"`
	Message     interface{}       `json:"message,omitempty"`
	MessageData map[string]string `json:"-"`
	Errors      interface{}       `json:"errors,omitempty"`
	Data        interface{}       `json:"data,omitempty"`
	Lang        string            `json:"-"`
	Meta        interface{}       `json:"meta,omitempty"`
}

// MetaData represent meta data response for multi data
type MetaData struct {
	Page       uint64 `json:"page"`
	Limit      uint64 `json:"limit"`
	TotalPage  uint64 `json:"total_page"`
	TotalCount uint64 `json:"total_count"`
}

// GetCode method to transform response name var to http status
func (r *Response) GetCode() int {
	return msg.GetCode(r.Name)
}

// GetMessage method to transform response name var to message detail
func (r *Response) GetMessage() string {
	return msg.Get(r.Name, r.Lang)
}

// BuildMessage build message
func (r *Response) BuildMessage() {
	if r.Message == nil {
		message := msg.Get(r.Name, r.Lang)

		if len(r.MessageData) > 0 {
			for key, val := range r.MessageData {
				message = strings.ReplaceAll(message, key, val)
			}
		}

		r.Message = message
	}
}

// SetMessage setter message
func (r *Response) SetMessage() {
	r.Message = msg.Get(r.Name, r.Lang)
}

// SetMessageData set message dynamic data
func (r *Response) SetMessageData(data map[string]string) *Response {
	if len(data) > 0 {
		r.MessageData = data
	}

	return r
}

// SetName setter response var name
func (r *Response) SetName(nm string) *Response {
	r.Name = nm
	return r
}

// SetData setter data response
func (r *Response) SetData(v interface{}) *Response {
	r.Data = v
	return r
}

// SetError setter error messages
func (r *Response) SetError(v interface{}) *Response {
	r.Errors = v
	return r
}

// SetMeta setter meta data response
func (r *Response) SetMeta(v interface{}) *Response {
	r.Meta = v
	return r
}

// NewResponse initialize response
func NewResponse() *Response {
	oneRsp.Do(func() {
		rsp = &Response{}
	})

	// clone response
	x := *rsp

	return &x
}
