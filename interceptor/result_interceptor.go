package interceptor

import (
	"bytes"
	"io/ioutil"
	"strconv"
	"time"

	"example.com/log"
	"example.com/util"

	"github.com/gin-gonic/gin"
)

// https://stackoverflow.com/questions/38501325/how-to-log-response-body-in-gin
type bodyLogWriter struct {
	gin.ResponseWriter
	body *bytes.Buffer
}

func (w bodyLogWriter) Write(b []byte) (int, error) {
	w.body.Write(b)
	return w.ResponseWriter.Write(b)
}

func (w bodyLogWriter) WriteString(s string) (int, error) {
	w.body.WriteString(s)
	return w.ResponseWriter.WriteString(s)
}

// ResultInterceptor ...
func ResultInterceptor() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Before
		start := time.Now()
		method := c.Request.Method
		path := c.Request.URL.Path
		raw := c.Request.URL.RawQuery
		if raw != "" {
			path = path + "?" + raw
		}

		requestBody := ""
		if method == "POST" || method == "PUT" {
			body, err := c.GetRawData()
			if err != nil {
				log.Error("GetRawData failed, err: ", err)
			}
			c.Request.Body = ioutil.NopCloser(bytes.NewBuffer(body)) // 关键点
			requestBody = util.BytesToString(body)
		}

		log.InfoWithFields("request start", map[string]string{
			"api":    path,
			"method": method,
			"body":   requestBody,
		})

		blw := &bodyLogWriter{
			ResponseWriter: c.Writer,
			body:           bytes.NewBufferString(""),
		}
		c.Writer = blw

		c.Next()

		// After
		complete := time.Now()
		duration := complete.Sub(start)

		responseBody := util.BytesToString(blw.body.Bytes())
		log.InfoWithFields("request completed", map[string]string{
			"api":      path,
			"cost":     duration.String(),
			"status":   strconv.Itoa(c.Writer.Status()),
			"response": responseBody,
		})

	}
}
