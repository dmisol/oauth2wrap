package oauth2wrap

import (
	"encoding/json"
	"log"
	"os"
	"testing"

	"golang.org/x/oauth2/endpoints"
)

func TestJson(t *testing.T) {
	m := make(map[string]*Conf)

	ep := EndpointM(endpoints.GitHub)
	c := &Conf{
		Endpoint: &ep,
		Api:      "https://api.github.com/user",
	}
	m["github"] = c

	ep = EndpointM(endpoints.Google)
	c = &Conf{
		Endpoint: &ep,
		Api:      "https://www.googleapis.com/oauth2/v2/userinfo",
	}
	m["google"] = c

	b, err := json.Marshal(m)
	if err != nil {
		t.Fatalf(err.Error())
	}

	log.Println(string(b))

	err = os.WriteFile("auth.json", b, 0644)
	if err != nil {
		t.Fatalf(err.Error())
	}
}
