package handlers

import (
	"net/http"

	"github.com/YacineMK/Rill/pkg/jwt"
	"github.com/YacineMK/Rill/pkg/utils"
)

type StreamHandler struct{}

func (h *StreamHandler) StreamKeyHandler(w http.ResponseWriter, r *http.Request) {
	key := utils.GenerateStreamKey()

	claims := jwt.Claims{
		StreamID: key, 
	}

	token, err := jwt.GenerateJWT(claims.StreamID)
	if err != nil {
		utils.WriteJSONError(w, http.StatusInternalServerError, err)
		return
	}

	utils.WriteJSONResponse(w, http.StatusOK, utils.Map{
		"message": "Stream key generated successfully",
		"token": token,
	})
}
