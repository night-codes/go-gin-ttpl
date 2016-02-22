package ttpl

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/render"
	"net/http"
	"text/template"
)

// PageTemplate struct for gin
type PageTemplate struct {
	TemplatePath string
	templates    *template.Template
}

// PageRender struct for gin
type PageRender struct {
	Template *template.Template
	Data     interface{}
	Name     string
}

func (r PageTemplate) Instance(name string, data interface{}) render.Render {
	return PageRender{
		Template: r.templates,
		Name:     name,
		Data:     data,
	}
}

func (r PageRender) Render(w http.ResponseWriter) error {
	header := w.Header()
	if val := header["Content-Type"]; len(val) == 0 {
		header["Content-Type"] = []string{"text/html; charset=utf-8"}
	}

	if len(r.Name) > 0 {
		r.Template.ExecuteTemplate(w, r.Name, r.Data)
	} else {
		r.Template.Execute(w, r.Data)
	}

	return nil
}

// Use ttpl render
func Use(r *gin.Engine, pattern string, funcMap ...template.FuncMap) {
	t := &template.Template{}
	if len(funcMap) > 0 {
		t = template.Must(template.New(pattern).Funcs(funcMap[0]).ParseGlob(pattern))
	} else {
		t = template.Must(template.New(pattern).ParseGlob(pattern))
	}

	r.HTMLRender = PageTemplate{"/", t}
}
