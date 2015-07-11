package main

import (
	"encoding/json"
	"fmt"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

func httpHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	r.ParseForm()
	values, _ := json.Marshal(r.Form)
	fmt.Fprintf(w, "%s\n", values)
}
