package websocket

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/require"
)

type authMiddlewareTestCase struct {
	t *testing.T
	// input
	tokenLookup    TokenLookup
	requestHeaders map[string]string
	requestCookies map[string]string
	// output
	expectedToken string
}

func (th authMiddlewareTestCase) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	actualToken, _ := ctx.Value(tokenContextKey).(string)
	require.Equal(th.t, actualToken, th.expectedToken)
}

func TestAuthMiddleware(t *testing.T) {
	tt := []authMiddlewareTestCase{
		{
			tokenLookup: TokenLookup{
				Header:       "Authorization",
				Cookie:       "SSID",
				HeaderPrefix: "Bearer",
			},

			requestHeaders: map[string]string{
				"Authorization": "Bearer token from header",
			},

			requestCookies: map[string]string{
				"SSID": "Bearer token from cookie",
			},

			expectedToken: "token from header",
		},
		{
			tokenLookup: TokenLookup{
				Header:       "Authorization",
				Cookie:       "SSID",
				HeaderPrefix: "Bearer",
			},

			requestHeaders: map[string]string{},

			requestCookies: map[string]string{
				"SSID": "Bearer token from cookie",
			},

			expectedToken: "Bearer token from cookie",
		},
		{
			tokenLookup: TokenLookup{
				Header: "Authorization",
				Cookie: "SSID",
			},

			requestHeaders: map[string]string{},

			requestCookies: map[string]string{
				"SSID": "token from cookie",
			},

			expectedToken: "token from cookie",
		},
	}
	for _, testCase := range tt {

		testCase.t = t

		handler := authMiddleware(testCase, testCase.tokenLookup)

		req := &http.Request{}

		req.Header = make(http.Header)

		for name, value := range testCase.requestHeaders {
			req.Header.Add(name, value)
		}

		for name, value := range testCase.requestCookies {
			cookie := &http.Cookie{
				Name:  name,
				Value: value,
			}
			req.AddCookie(cookie)
		}
		w := httptest.NewRecorder()
		handler.ServeHTTP(w, req)
	}
}
