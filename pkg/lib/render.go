package lib

import (
	"net/http"

	"github.com/a-h/templ"
	"github.com/labstack/echo/v4"
)

// Render renders zero to many templ components to the response with the given status code.
func Render(c echo.Context, status int, components ...templ.Component) error {
	buf := templ.GetBuffer()
	defer templ.ReleaseBuffer(buf)

	for _, cmp := range components {
		if err := cmp.Render(c.Request().Context(), buf); err != nil {
			return err
		}
	}

	return c.HTML(status, buf.String())
}

// IsHTMX returns true if the HTTP request is from HTMX, otherwise returns false.
func IsHTMX(c echo.Context) bool {
	return c.Request().Header.Get("HX-Request") == "true"
}

// HxRender conditionally renders the component if the request is from HTMX, otherwise page is rendered.
func HxRender(c echo.Context, status int, page, component templ.Component) error {
	if IsHTMX(c) {
		return Render(c, status, component)
	}
	return Render(c, status, page)
}

// Redirect performs a HTTP redirect with special handling for HTMX requests.
//
// If the request was made by HTMX, the "HX-Redirect" header is set with the route and is returned with a 200 OK.
// HTMX does not handle 3xx redirects, it only sees the final 2xx, 4xx, or 5xx response.
// For non-HTMX requests, a standard HTTP redirect is performed with the provided status.
//
// See github.com/bigskysoftware/htmx/issues/2052#issuecomment-1979805051 for more details.
func Redirect(c echo.Context, status int, route string) error {
	if IsHTMX(c) {
		c.Response().Header().Add("HX-Redirect", route)
		status = http.StatusOK
	}
	return c.Redirect(status, route)
}
