package main

import (
	export "export/data"
)

import (
	"net/http"
)

var i int

func main() {
	http.HandleFunc("/", myfunc)
	http.ListenAndServe(":9010", nil)
}
func myfunc(w http.ResponseWriter, r *http.Request) {
	new(export.Export).ExportStart(r)
}
