package handlers

import (
	"encoding/json"
	"github.com/go-http-utils/headers"
	"github.com/willkoerich/rock-challenge/cmd/api/response"
	"github.com/willkoerich/rock-challenge/internal/domain"
	"net/http"
	"time"
)

const AccountDescription = "access_account"

type (
	TransferHandler struct {
		controller domain.TransferController
	}
)

func NewTransferHandler(controller domain.TransferController) TransferHandler {
	return TransferHandler{
		controller: controller,
	}
}

func (handler TransferHandler) Process(writer http.ResponseWriter, request *http.Request) {
	var body domain.Transfer
	if err := json.NewDecoder(request.Body).Decode(&body); err != nil {
		response.CreateHandlerResponse(writer, http.StatusBadRequest,
			" unable to decode submitted transfer information body.", &err)
		return
	}
	body = updateBodyWithContextInformation(body, request)
	createdTransfer, err := handler.controller.Process(request.Context(), body)
	if err != nil {
		response.CreateHandlerResponse(writer, http.StatusInternalServerError,
			" can't process transfer", &err)
		return
	}
	writer.Header().Add(headers.ContentType, "application/json")
	writer.WriteHeader(http.StatusCreated)
	err = json.NewEncoder(writer).Encode(createdTransfer)
	if err != nil {
		return
	}
}

func updateBodyWithContextInformation(transferBody domain.Transfer, request *http.Request) domain.Transfer {
	transferBody.AccountOriginID = request.Context().Value(AccountDescription).(domain.AuthenticationResponse).AccountID
	transferBody.CreatedAt = time.Now()
	return transferBody
}

func (handler TransferHandler) Get(writer http.ResponseWriter, request *http.Request) {
	retrievedTransfer, err := handler.controller.GetAll(request.Context())
	if err != nil {
		response.CreateHandlerResponse(writer, http.StatusInternalServerError,
			" can't retrieve transfer", &err)
		return
	}
	writer.Header().Add(headers.ContentType, "application/json")
	writer.WriteHeader(http.StatusOK)
	err = json.NewEncoder(writer).Encode(retrievedTransfer)
	if err != nil {
		return
	}
}
