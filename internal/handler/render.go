package handler

import (
	"bytes"
	"log"
	"net/http"
	"text/template"
)

var funcs = template.FuncMap{}

type Renderer struct {
	template map[string]*template.Template
}

func NewRenderer() *Renderer {
	files := []string{"home.html", "product-page.html"}

	renderer := &Renderer{
		template: make(map[string]*template.Template),
	}

	for _, file := range files {
		renderer.template[file] = template.Must(template.New("layout").Funcs(funcs).ParseFiles(
			"templates/layout.html",
			// "templates/header.html",
			// "templates/sidebar.html",
			// "templates/footer.html",
			"templates/"+file,
		))
	}

	return renderer
}

func (r *Renderer) Render(w http.ResponseWriter, file string, data any) {
	var buffer bytes.Buffer

	if err := r.template[file].ExecuteTemplate(&buffer, "layout", data); err != nil {
		log.Printf("render %s: %v", file, err)
		http.Error(w, "Terjadi kesalahan pada server", http.StatusInternalServerError)
		return
	}

	buffer.WriteTo(w)
}
