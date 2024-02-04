package main

import (
    "fmt"
    "log"
    "net/http"
    "strconv"

    "github.com/julienschmidt/httprouter"
)

type Item struct {
    Name string
    Value int
}

var db = []Item{}

func main() {
    router := httprouter.New()
    router.GET("/", Index)
    router.GET("/hello/:name", Hello)
    router.GET("/add/:name/:value", AddItem)
    router.GET("/list/", ListItems)

    log.Print("Listening for requests...")

    log.Fatal(http.ListenAndServe(":8000", router))
}

func Index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
    fmt.Fprint(w, "Welcome!\n")
}

func Hello(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
    fmt.Fprintf(w, "Hello, %s!\n", ps.ByName("name"))
}

func AddItem(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
    name := ps.ByName("name")
    value, err := strconv.Atoi(ps.ByName("value"))
    if err != nil {
        fmt.Fprintf(w, "Error, `%s` is not a valid integer", ps.ByName("value"))
        return
    }
    fmt.Fprintf(w,
        "Adding item:\n\tName: %s\n\tValue: %d",
        name,
        value,
    )

    item := Item{ Name: name, Value: value }

    db = append(db, item)
    log.Printf("Adding item: %v", item)
}

func ListItems(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
    if len(db) == 0 {
        log.Print("Not listing items")
        fmt.Fprint(w, "No items, add them by going to `/add/<name>/<value>`")
        return
    }
    log.Print("Listing items")
    for i, item := range db {
        fmt.Fprintf(w,
            "Item %d:\n\tName: %s\n\tValue: %d\n\n",
            i,
            item.Name,
            item.Value,
        )
    }
}

