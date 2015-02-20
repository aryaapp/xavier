package routes

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/zenazn/goji/web"
)

const (
	ThemeIDValid   = "5911e506-d0dc-4c87-b78e-f64f79790180"
	ThemeIDInvalid = "1911e506-d0dc-4c87-b78e-f64f79790180"
)

func TestThemesIndexValid(t *testing.T) {
	w := httptest.NewRecorder()
	r, _ := http.NewRequest("GET", "http://example.com/themes", nil)

	handler{engine, ThemesIndex}.ServeHTTPC(web.C{}, w, r)
	assert.Equal(t, 200, w.Code, "Response code should be equal")
}

func TestThemesShow(t *testing.T) {
	w := httptest.NewRecorder()
	r, _ := http.NewRequest("GET", "http://example.com/themes/:uuid", nil)

	handler{engine, ThemesShow}.ServeHTTPC(web.C{URLParams: map[string]string{"uuid": ThemeIDValid}}, w, r)
	assert.Equal(t, 200, w.Code, "Response code should be equal")
}

func TestThemesShowInvalidUUID(t *testing.T) {
	w := httptest.NewRecorder()
	r, _ := http.NewRequest("GET", "http://example.com/themes/:uuid", nil)

	handler{engine, ThemesShow}.ServeHTTPC(web.C{URLParams: map[string]string{"uuid": ThemeIDInvalid}}, w, r)
	assert.Equal(t, 404, w.Code, "Response code should be equal")
}

func TestThemesShowNoUUID(t *testing.T) {
	w := httptest.NewRecorder()
	r, _ := http.NewRequest("GET", "http://example.com/themes/:uuid", nil)

	handler{engine, ThemesShow}.ServeHTTPC(web.C{URLParams: map[string]string{}}, w, r)
	assert.Equal(t, 404, w.Code, "Response code should be equal")
}
