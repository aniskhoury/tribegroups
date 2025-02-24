package handleFunctions

import (
	"net/http"
	"tribegroups/controllers/index"
	"tribegroups/controllers/register"
)

func CallAllHandleFunctions() {

	fileServer := http.FileServer(http.Dir("resources"))
	http.Handle("/resources/", http.StripPrefix("/resources", fileServer))
	http.HandleFunc("/", index.IndexFunc)
	http.HandleFunc("/register", register.RegisterFunc)

}
