package httplogger

import (
	"net/http"
	"net/http/httputil"
)

type Fields map[string]interface{}

type GenericLogger interface {
	IsDebugEnabled() bool
	IsTraceEnabled() bool

	WithFields(fields Fields) GenericLogger

	Debugf(format string, args ...interface{})
	Tracef(format string, args ...interface{})
}

type httpLogger struct {
	transport http.RoundTripper
	logger    GenericLogger
}

func NewHTTPLogger(transport http.RoundTripper, logger GenericLogger) *httpLogger {
	return &httpLogger{transport, logger}
}

func (this *httpLogger) RoundTrip(req *http.Request) (*http.Response, error) {
	resp, err := this.transport.RoundTrip(req)

	if this.logger.IsDebugEnabled() {
		isDumpBody := this.logger.IsTraceEnabled()
		reqDump, _ := httputil.DumpRequestOut(resp.Request, isDumpBody)
		respDump, _ := httputil.DumpResponse(resp, isDumpBody)

		this.logger.WithFields(Fields{
			"obj": Fields{"request": string(reqDump), "response": string(respDump)},
		}).Debugf("send request %s %s", resp.Request.Method, resp.Request.URL.String())
	}

	if err != nil {
		return resp, err
	}

	return resp, err
}
