package server

import (
	"context"
	"interceptor/gen/interceptor"
	"interceptor/internal/storage/sqlite"
	"io/ioutil"
	"log/slog"
	"net/http"
	"strings"
)

type handler struct {
	logger *slog.Logger
	db     sqlite.Storage
}

func NewHandler(logger *slog.Logger, db sqlite.Storage) *handler {
	return &handler{
		logger: logger,
		db:     db,
	}
}

func (h *handler) HealthGet(ctx context.Context) (*interceptor.HealthGetOK, error) {
	//req := ctx.Value("req").(*http.Request)
	return &interceptor.HealthGetOK{Status: interceptor.NewOptString("server is running")}, nil
}

func (h *handler) StealPost(ctx context.Context, req interceptor.StealPostReq) (interceptor.StealPostRes, error) {
	ctxReq := ctx.Value("req").(*http.Request)
	h.logger.Info("/StealPost", slog.Any("req", req), slog.Any("ctx", ctxReq))
	h.logger.Debug("Remote Address", slog.Any("ip", ctxReq.RemoteAddr))

	// save to db
	data, _ := ioutil.ReadAll(req.Data)
	err := h.db.InsertInformation(string(data), ctxReq.RemoteAddr)
	if err != nil {
		return nil, err
	}

	h.logger.Info("Data saved successfully")
	return &interceptor.StealPostOK{
		Data: ioutil.NopCloser(strings.NewReader("Data saved successfully")),
	}, nil
}

//HealthGet(ctx context.Context) (*interceptor.HealthGetOK, error)
//StealPost(ctx context.Context, req *interceptor.StealPostReq) (interceptor.StealPostRes, error)
