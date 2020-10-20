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

	isDebugEnabled bool
	isTraceEnabled bool
}

func NewHTTPLogger(transport http.RoundTripper, logger GenericLogger) *httpLogger {
	return &httpLogger{transport, logger, logger.IsDebugEnabled(), logger.IsTraceEnabled()}
}

func (this *httpLogger) RoundTrip(req *http.Request) (*http.Response, error) {
	var reqDump []byte

	if this.isDebugEnabled {
		reqDump, _ = httputil.DumpRequestOut(req, this.isTraceEnabled)
	}

	resp, err := this.transport.RoundTrip(req)

	if this.isDebugEnabled {
		respDump, _ := httputil.DumpResponse(resp, this.isTraceEnabled)

		this.logger.WithFields(Fields{
			"obj":  Fields{"request": string(reqDump), "response": string(respDump)},
			"tags": []string{"httplogger"},
		}).Debugf("send request %s %s", resp.Request.Method, resp.Request.URL.String())
	}

	if err != nil {
		return resp, err
	}

	return resp, err
}
