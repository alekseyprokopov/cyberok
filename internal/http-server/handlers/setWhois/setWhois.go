package setWhois

import (
	"cyberok/internal/lib/api/logger/sl"
	"cyberok/internal/lib/api/response"
	resp "cyberok/internal/lib/api/response"
	"cyberok/internal/service"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/render"
	"github.com/go-playground/validator/v10"
	"golang.org/x/exp/slog"
	"net/http"
	"sync"
)

type Request struct {
	DomainData []string `json:"domain_data" validate:"required,dive,fqdn"`
}

type Response struct {
	response.Response
}

func New(log *slog.Logger, service *service.Service) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		const op = "handlers.setWhois.New"

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

		if err := validator.New().Struct(req); err != nil {
			validateErr := err.(validator.ValidationErrors)
			log.Error("invalid request", sl.Err(err))
			render.JSON(w, r, resp.ValidationError(validateErr))
			return
		}

		wg := sync.WaitGroup{}
		for _, domain := range req.DomainData {
			wg.Add(1)
			go func(domain string) {
				defer wg.Done()
				whoisInfo, err := service.LookupWhois(domain)
				if err != nil {
					log.Error("invalid whois", sl.Err(err))
					return
				}
				_, err = service.CreateWhois(domain, whoisInfo)
				if err != nil {
					service.UpdateWhois(domain, whoisInfo)
				}

			}(domain)

		}
		wg.Wait()
		log.Info("request completed", slog.String("op", op))

		render.JSON(w, r, Response{
			Response: response.OK(),
		})
	}
}
