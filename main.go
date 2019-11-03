package main

import (
	"asanaSync/app"
	"io"
	"net/http"
	"os"
)

func main() {
	app.Start()

	http.HandleFunc("/", hello)
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	http.ListenAndServe(":"+port, nil)
}

func hello(res http.ResponseWriter, req *http.Request) {
	res.Header().Set(
		"Content-Type",
		"text/html",
	)

	io.WriteString(
		res,
		`<DOCTYPE html>
 <html>
	 <head>
	 	<title>Asana Sync service</title>
	 </head>
	 <body>
		 <h1>ASANA SYNC SERVICE</h1>
		 <h1>This service is saving data  from https://traction.tools to https://app.asana.com</h1>		 
	 </body>
 </html>
		`,
	)
}
