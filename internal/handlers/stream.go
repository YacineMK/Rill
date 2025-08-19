package handlers

import (
	"net/http"

	"github.com/YacineMK/Rill/pkg/jwt"
	"github.com/YacineMK/Rill/pkg/utils"
)

type StreamHandler struct{}

func (h *StreamHandler) StreamKeyHandler(w http.ResponseWriter, r *http.Request) {
	streamID := utils.GenerateStreamKey()

	token, err := jwt.GenerateJWT(streamID)
	if err != nil {
		utils.WriteJSONError(w, http.StatusInternalServerError, err)
		return
	}

	utils.WriteJSONResponse(w, http.StatusOK, utils.Map{
		"success": true,
		"data": utils.Map{
			"stream_key": token,
			"stream_id":  streamID,
		},
	})
}
