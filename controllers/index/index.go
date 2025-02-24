package index

import (
	"net/http"
	"tribegroups/configuration"
)

var ()

func IndexFunc(w http.ResponseWriter, r *http.Request) {
	varmap := map[string]interface{}{
		"error": false,
	}

	configuration.Templates.ExecuteTemplate(w, "index.html", varmap)
}
