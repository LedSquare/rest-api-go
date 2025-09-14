package redirect

import (
	"errors"
	"log/slog"
	"net/http"
	res "rest-api-go/internal/lib/api/response"
	"rest-api-go/internal/lib/logger/sl"
	"rest-api-go/internal/storage"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/render"
)

type UrlGetter interface {
	GetUrl(alias string) (string, error)
}

func New(log *slog.Logger, urlGetter UrlGetter) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		const action = "handlers.url.redirect.New"

		log = log.With(
			slog.String("action", action),
			slog.String("request_id", middleware.GetReqID(r.Context())),
		)

		alias := chi.URLParam(r, "alias")
		if alias == "" {
			log.Info("Alias is empty")

			render.JSON(w, r, res.Error("Alias not found"))

			return
		}

		resUrl, err := urlGetter.GetUrl(alias)
		if errors.Is(err, storage.ErrURLNotFound) {
			log.Info(err.Error(), "alias", alias)

			render.JSON(w, r, res.Error("Url not found"))

			return
		}
		if err != nil {
			log.Error("Failed to get url", sl.Error(err))

			render.JSON(w, r, res.Error("Internal server error"))
		}

		log.Info("Got url", slog.String("url", resUrl))

		http.Redirect(w, r, resUrl, http.StatusFound)
	}
}
