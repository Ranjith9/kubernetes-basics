package main

import (
        "html/template"
        "log"
        "net/http"
//        "os"
        "io/ioutil"
//        "fmt"
        "encoding/json"
        "gopkg.in/mgo.v2"
//        "gopkg.in/mgo.v2/bson"
)

type Profile struct {
        FirstName  string
        LastName   string
}

func main() {
        http.HandleFunc("/", foo)
        http.Handle("/favicon.ico", http.NotFoundHandler())
        http.HandleFunc("/users", getUsers)
        http.ListenAndServe(":80", nil)
}

func foo(w http.ResponseWriter, req *http.Request) {

        form := `<!DOCTYPE html>

<html>
<head>
    <meta charset="UTF-8">
    <title>Document</title>
</head>
<body>
<form method="POST">
    <label for="firstName">First Name</label>
    <input type="text" id="firstName" name="first">
    <br>
    <label for="lastName">Last Name</label>
    <input type="text" id="lastName" name="last">
    <br>
    <input type="submit">
</form>
</body>
</html>`

        t, errt := template.New("Title").Parse(form)
        if errt != nil {
               log.Fatalln(errt)
        }

        f := req.FormValue("first")
        l := req.FormValue("last")

        mine := Profile{f, l}
        if f != "" {
                dbInsert(mine)
        }
        err := t.Execute(w, nil)
        if err != nil {
                http.Error(w, err.Error(), 500)
                log.Fatalln(err)
        }
}


func dbInsert(p Profile) {

        session, err := mgo.Dial("mongo:27017")
        if err != nil {
                panic(err)
        }
        defer session.Close()

        session.SetMode(mgo.Monotonic, true)

        c := session.DB("profiles").C("users")
        err = c.Insert(&Profile{p.FirstName,p.LastName})
        if err != nil {
                log.Fatal(err)
        }

}

func getUsers(w http.ResponseWriter, req *http.Request){
       session, err := mgo.Dial("mongo:27017")
        if err != nil {
                panic(err)
        }
        defer session.Close()
        var results []Profile

        c := session.DB("profiles").C("users")

        err1 := c.Find(nil).All(&results)
        if err1 != nil {
            panic(err1)
        }

        b, _ := ioutil.ReadAll(req.Body)
        defer req.Body.Close()
        // Unmarshal

        json.Unmarshal(b, &results)

        output,_ := json.MarshalIndent(results, "", " ")
        w.Header().Set("content-type", "application/json")
        w.Write(output)

}
