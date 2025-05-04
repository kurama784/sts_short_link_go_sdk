package tests

import (
	"os"
	"testing"

	"short_link_library/sdk"
)

func TestLinkCreate(t *testing.T) {

	api := sdk.CreateRequest(os.Getenv("SHORT_LINK_TOKEN"), nil)

	dto := sdk.CreateDto{
		RedirectURL: "https://google.com",
	}

	resp, err := api.SendRequest(dto)

	if err != nil {
		t.Fatal(err)
	}

	t.Logf("Response: %+v", resp.URL)
}
