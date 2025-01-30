package middleware

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/fydhfzh/ecommerce-go-application/src/broker-service/common"
	"github.com/fydhfzh/ecommerce-go-application/src/broker-service/dto"
	"github.com/labstack/echo/v4"
)

func LogRequest(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		log.Printf("Request accepted with method: %s and url: %s", c.Request().Method, c.Request().RequestURI)
		tz, err := time.LoadLocation("Asia/Jakarta")
		if err != nil {
			log.Println(err)
			return err
		}

		time.Local = tz
		content := fmt.Sprintf("[INFO] %s %s - %v", c.Request().Method, c.Request().RequestURI, time.Now().Local())
		logRequest := dto.LogRequest{
			Content: content,
		}

		json, err := json.Marshal(logRequest)
		if err != nil {
			log.Println(err)
			return err
		}

		loggerServiceConfig := common.Config.ServicesConfig.LoggerServiceConfig

		url := fmt.Sprintf("http://%s:%d/api/v1", loggerServiceConfig.Host, loggerServiceConfig.Port)
		req, err := http.NewRequest("POST", url, bytes.NewReader(json))
		if err != nil {
			return err
		}
		req.Header.Set("Content-Type", "application/json")

		client := &http.Client{}
		resp, err := client.Do(req)
		if err != nil {
			return err
		}
		defer resp.Body.Close()

		return next(c)
	}
}
