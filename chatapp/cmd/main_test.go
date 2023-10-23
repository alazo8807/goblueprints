package main

import (
	"net/http"
	"sync"
	"testing"
	"text/template"
)

func Test_templateHandler_ServeHTTP(t *testing.T) {
	type fields struct {
		once     sync.Once
		filename string
		templ    *template.Template
	}
	type args struct {
		w http.ResponseWriter
		r *http.Request
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tr := &templateHandler{
				once:     tt.fields.once,
				filename: tt.fields.filename,
				templ:    tt.fields.templ,
			}
			tr.ServeHTTP(tt.args.w, tt.args.r)
		})
	}
}
