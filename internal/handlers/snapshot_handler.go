package handlers

import (
	"encoding/json"
	"net/http"

	"go-typesense-app/internal/search"
)

type SnapshotHandlerInterface interface {
	CreateSnapshot(w http.ResponseWriter, r *http.Request)
}

type SnapshotHandlerImpl struct {
	snapshot search.SnapshotInterface
}

func NewSnapshotHandler(snp search.SnapshotInterface) *SnapshotHandlerImpl {
	return &SnapshotHandlerImpl{
		snapshot: snp,
	}
}

func (h *SnapshotHandlerImpl) CreateSnapshot(w http.ResponseWriter, _ *http.Request) {
	snapshot, err := h.snapshot.CreateSnapshot()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(map[string]interface{}{
		"message":  "SnapshotImpl created successfully",
		"snapshot": snapshot,
	})
}
