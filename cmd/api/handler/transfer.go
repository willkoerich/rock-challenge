package handlers

import (
	"encoding/json"
	"fmt"
	"github.com/willkoerich/rock-challenge/internal/domain"
	"net/http"
)

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
	var bodyContent domain.Transfer
	if err := json.NewDecoder(request.Body).Decode(&bodyContent); err != nil {
		writer.WriteHeader(http.StatusBadRequest)
		_, err := writer.Write([]byte(" unable to decode submitted transfer information body. Error: " + err.Error()))
		if err != nil {
			return
		}
		return
	}
	//updatedTransferBody := transfer.Transfer{AccountOriginID: request.Context().Value("accountID")}
	createdTransfer, err := handler.controller.Save(request.Context(), bodyContent)
	if err != nil {
		fmt.Println("Can't process Transfer", err)
		writer.WriteHeader(http.StatusInternalServerError)
		_, err := writer.Write([]byte(err.Error()))
		if err != nil {
			return
		}
		return
	}
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusCreated)
	err = json.NewEncoder(writer).Encode(createdTransfer)
	if err != nil {
		return
	}
}

func (handler TransferHandler) Get(writer http.ResponseWriter, request *http.Request) {
	retrievedTransfer, err := handler.controller.GetAll(request.Context())
	if err != nil {
		fmt.Println("Can't retrieve Transfer", err)
		// Aqui deveria adicionar uma maneira de distinguir entre um 404 e demais erros.
		// c.JSON(código_definido, mensagem_erro)
		writer.WriteHeader(http.StatusInternalServerError)
		_, err := writer.Write([]byte(err.Error()))
		if err != nil {
			return
		}
		return
	}
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusOK)
	err = json.NewEncoder(writer).Encode(retrievedTransfer)
	if err != nil {
		return
	}
}
