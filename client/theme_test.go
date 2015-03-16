package client

import "testing"

var client = NewClient("668ac08d-82b4-42a6-943b-2f6ca2c38258", "clientSecret", nil)

func TestAllThemes(t *testing.T) {
	themes, _, err := client.Themes.All()
	if err != nil {
		t.Errorf("Themes.All returned error: %v", err)
	}

	if len(themes) != 6 {
		t.Errorf("Themes.All did not return 6 themes, it returned: %d", len(themes))
	}
}

func TestFindTheme(t *testing.T) {
	uuid := "de7c6ae3-712c-4452-ae0c-5a4e5aea5359"
	theme, _, err := client.Themes.Find(uuid)
	if err != nil {
		t.Errorf("Themes.Find returned error: %v", err)
	}

	if theme.UUID != uuid {
		t.Errorf("Themes.Find did not the same theme, it returned: %s", theme.UUID)
	}
}
