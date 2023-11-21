package oauth2wrap

import (
	"log"
	"testing"

	"golang.org/x/oauth2/endpoints"
)

func TestEp(t *testing.T) {
	t.Skip()
	g := endpoints.GitHub
	my := EndpointM(g)
	log.Println(my)
}
