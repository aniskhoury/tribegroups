package main

import (
	"fmt"
	"net/http"
	"tribegroups/configuration"
	"tribegroups/controllers/handleFunctions"
)

func main() {
	configuration.LoadOAuthData()
	fmt.Printf("clientID %s", configuration.ClientIDOAuth)
	handleFunctions.CallAllHandleFunctions()
	http.ListenAndServe(":"+configuration.PortHTTPServer, nil)

}
