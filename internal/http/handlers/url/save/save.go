package save

import (
	"errors"
	"log/slog"
	"math/rand"
	"net/http"
	res "rest-api-go/internal/lib/api/response"
	"rest-api-go/internal/lib/logger/sl"
	"rest-api-go/internal/storage"

	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/render"
	"github.com/go-playground/validator/v10"
)

type Request struct {
	Url   string `json:"url" validate:"required,uri"`
	Alias string `json:"alias,omitempty"`
}

type Response struct {
	res.Response
	Alias string `json:"alias,omitempty"`
}

type UrlSaver interface {
	SaveUrl(url string, alias string) (int64, error)
}

func New(log *slog.Logger, urlSaver UrlSaver) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		const action = "handlers.url.save.New"

		log = log.With(
			slog.String("action", action),
			slog.String("request_id", middleware.GetReqID(r.Context())),
		)

		var request Request

		err := render.DecodeJSON(r.Body, &request)

		if err != nil {
			log.Error("Failed to json decode request body", sl.Error(err))
			render.JSON(w, r, res.Error("Failed to json decode request"))

			return
		}

		log.Info("Request body decoded", slog.Any("request", request))

		if err := validator.New().Struct(request); err != nil {
			validateErrors := err.(validator.ValidationErrors)
			errMsg := "Invalid request"

			log.Error(errMsg, sl.Error(err))

			render.JSON(w, r, res.ValidaationErrors(validateErrors))

			return
		}

		alias := request.Alias
		if alias == "" {
			alias = generateAlias(10)
		}

		id, err := urlSaver.SaveUrl(request.Url, alias)
		if errors.Is(err, storage.ErrURLExists) {
			log.Info(err.Error(), slog.String("url", request.Url))

			render.JSON(w, r, res.Error(err.Error()))

			return
		}

		if err != nil {
			log.Error("Failed to save url", sl.Error(err))

			render.JSON(w, r, res.Error("Failed to save url"))

			return
		}

		log.Info("Url added", slog.Int64("id", id))

		render.JSON(w, r, Response{
			Response: res.Success(),
			Alias:    alias,
		})
	}
}

func generateAlias(length int) string {
	const charset = "abcdefghijklmnopqrstuvwxyz" + "ABCDEFGHIJKLMNOPQRSTUVWXYZ" + "0123456789"

	if length <= 0 {
		return ""
	}
	result := make([]byte, length)
	for i := range result {
		result[i] = charset[rand.Intn(len(charset))]
	}
	return string(result)
}
