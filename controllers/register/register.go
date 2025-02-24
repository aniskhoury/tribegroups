package register

import (
	"net/http"
	"tribegroups/configuration"
)

var ()

func RegisterFunc(w http.ResponseWriter, r *http.Request) {
	varmap := map[string]interface{}{
		"error": false,
	}

	configuration.Templates.ExecuteTemplate(w, "register.html", varmap)
}
