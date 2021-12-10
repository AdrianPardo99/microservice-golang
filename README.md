# Microservices with golang

You can install golang from source in the main page of golang

Golang is speally development for the web applications, and sometimes we development an endpoints for exchange information.

## Hello world (endpoint)

```go
package main

import (
	"fmt"
	"net/http"
)

func main() {
    http.HandleFunc("/hello-world", helloWorld)
    /* 0.0.0.0 for listen in all interfaces */
    http.ListenAndServe("0.0.0.0:8081", nil)
}

func helloWorld(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello World!!")
}
```
### Creating an endpoint that uses json encode
Sometimes we use the standard JSON for sending or receiving data, the next example is for sending data:
```go
package main

import (
	"encoding/json"
	"net/http"
)

type Product struct {
	Name  string `json:"name"`
	Brand string `json:"brand"`
	Sku   string `json:"sku"`
}

func main() {
	http.HandleFunc("/get-products", getAllProducts)
	http.ListenAndServe("localhost:8081", nil)
}

func getAllProducts(w http.ResponseWriter, r *http.Request) {
	products := []Product{
		{"Papitas", "Barcel", "12345"},
		{"Takis", "Barcel", "12346"},
	}
	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(products)
}
```

### Creating an endpoint that uses xml encode
```go
package main

import (
	"encoding/xml"
	"net/http"
)

type Product struct {
	Name  string `xml:"name"`
	Brand string `xml:"brand"`
	Sku   string `xml:"sku"`
}

func main() {
	http.HandleFunc("/get-products", getAllProducts)
	http.ListenAndServe("localhost:8081", nil)
}

func getAllProducts(w http.ResponseWriter, r *http.Request) {
	products := []Product{
		{"Papitas", "Barcel", "12345"},
		{"Takis", "Barcel", "12346"},
	}
	w.Header().Add("Content-Type", "application/xml")
	xml.NewEncoder(w).Encode(products)
}
```

#### Combining xml and json enconde

If we want to development and use the two encodes methods we only need to check the type of petition that we receive from our client, i.e:

```go
/**
*	In out function where we return the data:
*	Check if the content petition type is XML
**/
if r.Header.Get("Content-Type") == "application/xml" {
	w.Header().Add("Content-Type", "application/xml")
	xml.NewEncoder(w).Encode(products)
} else {
	/**
	*	In other cases return information like json
	**/
	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(products)
}
```

## Creating our first microservice

The first microservice will be separate in 3 simple files in the same work folder:

* microservice -> folder
	* app.go -> contain the init server function
	* handlers.go -> contains routes and functions
	* main.go -> contains the entry point for create all the application

When we try to run the microservice we need to create the go module, so we need to execute like:

```bash
#	For create the module
go mod init name-module
#	For execute
go run .
```

### Creating the same microservice but with different work folder

Our microservice seems like:

* microservice -> folder
	* app -> second folder
		* app.go
		* handler.go
	* main.go

For working with these idea, we need to create a module and import the module in main.go file

```bash
#	Creating the module
go mod init name-module
```
For importing in main
```go
import "name-module/app"
```
For executing
```bash
go run .
```
For more example check the folders first_microservice and second_microservice
