package stat

import (
	"go_pro_api/configs"
	"go_pro_api/pkg/middleware"
	"go_pro_api/pkg/response"
	"net/http"
	"time"
)

const (
	GroupByDay   = "day"
	GroupByMonth = "month"
)

type StatHandlerDeps struct {
	StatRepository *StatRepository
	Config         *configs.Config
}

type StatHandler struct {
	StatRepository *StatRepository
}

func NewStatHandler(router *http.ServeMux, deps StatHandlerDeps) {
	handler := StatHandler{
		StatRepository: deps.StatRepository,
	}
	router.Handle("GET /stat", middleware.IsAuthed(handler.GetStat(), deps.Config))
}

func (h *StatHandler) GetStat() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fromStr := r.URL.Query().Get("from")
		toStr := r.URL.Query().Get("to")
		by := r.URL.Query().Get("by")
		if by != GroupByDay && by != GroupByMonth {
			http.Error(w, "Invalid from param", http.StatusBadRequest)
			return
		}

		from, err := time.Parse(time.DateOnly, fromStr)
		if err != nil {
			http.Error(w, "Invalid from param", http.StatusBadRequest)
			return
		}

		to, err := time.Parse(time.DateOnly, toStr)
		if err != nil {
			http.Error(w, "Invalid from param", http.StatusBadRequest)
			return
		}

		stats := h.StatRepository.GetStats(by, from, to)
		response.JSON(w, stats, http.StatusOK)
	}
}
