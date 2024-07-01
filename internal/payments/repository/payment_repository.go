package repo

import (
	"bytes"
	"encoding/json"
	"net/http"
	"payd/internal/payments/model"
	"payd/pkg/db"
	"payd/pkg/utils"

	"github.com/gin-gonic/gin"
)

func MakePayment() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		var payment model.Payment
		if err := ctx.ShouldBindJSON(&payment); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalide"})
			return
		}

		paymentRequest := model.Payment{
			AccountId:   payment.AccountId,
			PhoneNumber: payment.PhoneNumber,
			Amount:      payment.Amount,
			Narration:   payment.Narration,
			CallbackURL: payment.CallbackURL,
			Channel:     payment.Channel,
		}
		db.DB.Create(&paymentRequest)
		requestBody, err := json.Marshal(payment)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "error occurred"})
			return
		}
		if err := utils.PublishMessage(requestBody); err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "failed to publish message"})
			return
		}
		db.DB.Create(&payment)

		paydUrl := "https://api.mypayd.app/api/v2/withdrawal"
		client := &http.Client{}
		req, err := http.NewRequest("POST", paydUrl, bytes.NewBuffer(requestBody))
		if err != nil {
			panic(err)
		}
		req.SetBasicAuth("NqRu98iDE7VYnObu8mTK", "fhOf4vHqvLhrrRiFltGK5anrtzP1BmoDqjqMSsMf")
		req.Header.Set("Content-Type", "application/json")
		response, err := client.Do(req)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "failed to send request"})
			return
		}
		//fmt.Println(response)
		defer response.Body.Close()
		//var result map[string]interface{}
		var paymentResponse model.PaymentResponse
		if err := json.NewDecoder(req.Body).Decode(&paymentResponse); err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "failed to send api response"})
			return
		}

		ctx.JSON(http.StatusOK, gin.H{"payment": payment})
	}
}

func CheckPaymentStatus() gin.HandlerFunc {
	return func(ctx *gin.Context) {

	}
}
