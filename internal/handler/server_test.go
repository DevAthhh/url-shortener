package handler

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/DevAthhh/url-shortener/internal/database"
	"github.com/DevAthhh/url-shortener/internal/initializers"
	"github.com/DevAthhh/url-shortener/internal/lib/logger"
	"github.com/DevAthhh/url-shortener/internal/lib/transport"
)

func TestURLServer(t *testing.T) {
	var testCases = []struct {
		Name  string
		Input transport.RequestToSave
		Want  int
	}{
		{
			Name:  "zero size",
			Input: transport.RequestToSave{Root: "аа", Size: 0},
			Want:  0,
		},
		{
			Name:  "5 size",
			Input: transport.RequestToSave{Root: "ыва", Size: 5},
			Want:  5,
		},
		{
			Name:  "99 size",
			Input: transport.RequestToSave{Root: "ывап", Size: 99},
			Want:  99,
		},
		{
			Name:  "6 size",
			Input: transport.RequestToSave{Root: "аыпрвапр", Size: 6},
			Want:  6,
		},
	}

	initializers.LoadEnv()
	cfg := initializers.LoadConfig()
	logger.LoadLogger(cfg)
	db := database.LoadDatabase()

	srv := NewServer(cfg, db)

	for _, tc := range testCases {
		t.Run(tc.Name, func(t *testing.T) {
			rec := httptest.NewRecorder()

			jsonReq, _ := json.Marshal(tc.Input)

			req, _ := http.NewRequest("POST", "/", bytes.NewBuffer(jsonReq))
			srv.server.Handler.ServeHTTP(rec, req)

			var resp transport.ResponseFromSave
			json.Unmarshal(rec.Body.Bytes(), &resp)

			if len(resp.Alias) != tc.Want {
				t.Errorf("got %v, want %v", len(resp.Alias), tc.Want)
			}
		})
	}
}
