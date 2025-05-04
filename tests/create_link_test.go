package tests

import (
	"os"
	"testing"

	"github.com/kurama784/sts_short_link_go_sdk/sdk"
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
