package handlers

import (
	"net/http"
	"github.com/flosch/pongo2"
	"io"
	"github.com/labstack/echo"
	"errors"
	
)

type Renderer struct {
	Debug bool
}

func (r Renderer) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	
		var ctx pongo2.Context
	
		if data != nil {
			var ok bool
			ctx, ok = data.(pongo2.Context)
	
			if !ok {
				return errors.New("no pongo2.Context data was passed...")
			}
		}
	
		var t *pongo2.Template
		var err error
	
		if r.Debug {
			t, err = pongo2.FromFile(name)
		} else {
			t, err = pongo2.FromCache(name)
		}
	
		// Add some static values
		ctx["version_number"] = "v0.0.1-beta"
	
		if err != nil {
			return err
		}
	
		return t.ExecuteWriter(ctx, w)
	}


func HtmlIndex(c echo.Context) error {
	data := pongo2.Context{
		"name": "Machiel",
		"answer": func() int {
			return 42
		},
		"items": []string{
			"Apple", "Pear", "Banana", "Pineapple",
		},
	}

	return c.Render(http.StatusOK, "../../public/views/index.html", data)
}