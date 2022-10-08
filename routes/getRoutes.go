package routes

import (
	"io/ioutil"

	hbs "github.com/aymerick/raymond"
	"github.com/gin-gonic/gin"
)

func renderHbs(filePath string, ctx map[any]any) string {
	bytes, _ := ioutil.ReadFile(filePath)
	res, _ := hbs.Render(string(bytes), ctx)
	return res
}

func Home(c *gin.Context) {
	res := renderHbs("./templates/index.hbs", map[any]any{
		"title":          "go supabase",
		"supabaseScript": "https://cdn.jsdelivr.net/npm/@supabase/supabase-js",
	})
	c.Header("Content-type", "text/html")
	c.String(200, string(res))
}

func JsRender(c *gin.Context) {
	bytes, _ := ioutil.ReadFile("./templates/app.js")
	c.String(200, string(bytes))
}
