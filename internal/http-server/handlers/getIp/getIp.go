package getIp

import (
	"cyberok/internal/lib/api/logger/sl"
	"cyberok/internal/lib/api/response"
	"cyberok/internal/model"
	"cyberok/internal/service"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/render"
	"golang.org/x/exp/slog"
	"net/http"
)

type Request struct {
	FqdnData []string `json:"fqdn_data" validate:"required,dive"`
}

type Response struct {
	response.Response
	FqdnData []model.Fqdn `json:"fqdn_data" validate:"required,dive"`
}

func New(log *slog.Logger, service *service.Service) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		const op = "handlers.getIp.New"

		log = log.With(
			slog.String("op", op),
			slog.String("request_id", middleware.GetReqID(r.Context())),
		)

		var req Request

		if err := render.DecodeJSON(r.Body, &req); err != nil {
			log.Error("failed to decode request body", sl.Err(err))
			render.JSON(w, r, response.Error("failed to decode request"))
			return
		}

		log.Info("request body decoded", slog.Any("request", req))

		fqdns, err := service.GetByFQDNs(req.FqdnData)
		if err != nil {
			log.Error("failed to get fqdns by ips", sl.Err(err))
			return
		}

		render.JSON(w, r, Response{
			Response: response.OK(),
			FqdnData: fqdns,
		})
	}
}
