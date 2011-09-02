/**
 *
 */

package main

import (
    "flag"
    "http"
    "template"
    "json"
)

var (
    listenAddr = flag.String("port", "80", "http listen port")
    hostName   = flag.String("host", "you.rl", "http host name")
    //hostName   = flag.String("host", "localhost", "http host name")
)

type Title struct {
    Logo string
}

var store Store

func Redirect(w http.ResponseWriter, r *http.Request) {
    key := r.URL.Path[1:]

    if key == "favicon.ico" {
        http.NotFound(w, r)
    } else if key == "" {
        logo   := &Title{Logo: "You.RL"}
        tpl, _ := template.ParseFile("index.html", nil)
        tpl.Execute(w, logo)
    } else if key == "index.js" {
        http.ServeFile(w, r, "index.js")
    } else {
        var url string
        store.Get(&key, &url)
        http.Redirect(w, r, url, http.StatusFound)
    }
}

func Shorten(w http.ResponseWriter, r *http.Request) {
    var key string
    for key = genKey(); store.Check(&key) == false; { }

    long_url := r.FormValue("url")

    store.Put(&key, &long_url)

    v := make(map[string]interface{})
    v["url"] = "http://" + *hostName + "/" + key

    enc := json.NewEncoder(w)
    enc.Encode(&v)
}

func main() {
    flag.Parse()
    store = new(URLStore)

    http.HandleFunc("/", Redirect)
    http.HandleFunc("/shorten", Shorten)
    http.ListenAndServe(":" + *listenAddr, nil)
}

/* EOF */
