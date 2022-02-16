package transport_test

import (
	"net/http"
	"testing"

	"github.com/stealthmodesoft/gqlgen/graphql/handler/testserver"
	"github.com/stealthmodesoft/gqlgen/graphql/handler/transport"
	"github.com/stretchr/testify/assert"
)

func TestOptions(t *testing.T) {
	h := testserver.New()
	h.AddTransport(transport.Options{})

	t.Run("responds to options requests", func(t *testing.T) {
		resp := doRequest(h, "OPTIONS", "/graphql?query={me{name}}", ``)
		assert.Equal(t, http.StatusOK, resp.Code)
		assert.Equal(t, "OPTIONS, GET, POST", resp.Header().Get("Allow"))
	})

	t.Run("responds to head requests", func(t *testing.T) {
		resp := doRequest(h, "HEAD", "/graphql?query={me{name}}", ``)
		assert.Equal(t, http.StatusMethodNotAllowed, resp.Code)
	})
}
