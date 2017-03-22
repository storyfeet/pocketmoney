package main

import (
	"flag"
	"fmt"
	"net/http"
	"path"
	"strings"
	"text/template"

	"github.com/coderconvoy/dbase2"
	"github.com/coderconvoy/templater/tempower"
)

var GT *template.Template
var FamDB = dbase2.DBase{"data/families"}

type IndexData struct {
	Mes, Family, Username string
}

func GoIndex(w http.ResponseWriter, r *http.Request, m string) {
	c, err := r.Cookie("LastLog")
	if err != nil {
		ExTemplate(GT, w, "index.html", IndexData{m, "", ""})
		return
	}
	s := strings.Split(c.Value, ":")
	if len(s) != 2 {
		ExTemplate(GT, w, "index.html", IndexData{m, c.Value, ""})
		return
	}
	ExTemplate(GT, w, "index.html", IndexData{m, s[0], s[1]})

}

func ExTemplate(t *template.Template, w http.ResponseWriter, name string, data interface{}) {
	err := t.ExecuteTemplate(w, name, data)
	if err != nil {
		fmt.Println(err)
	}
}

func Handle(w http.ResponseWriter, r *http.Request) {
	dbase2.QLog(r.Host + "--" + r.URL.Path)
	GoIndex(w, r, "")
}

func HandleLogout(w http.ResponseWriter, r *http.Request) {
	loginControl.Logout(w, r)
	GoIndex(w, r, "You are now Logged out")
}

func HandleStatic(w http.ResponseWriter, r *http.Request) {
	p := r.URL.Path
	ass, err := Asset(path.Join("assets", p))
	if err != nil {
		dbase2.QLog("Could not serve static, " + p + ":" + err.Error())
		return
	}
	switch path.Ext(p) {
	case ".css":
		w.Header().Set("Content-Type", "text/css")
	case ".js":

	}
	w.Write(ass)
}

func main() {
	insecure := flag.Bool("i", false, "Run without https")
	debug := flag.Bool("d", false, "Debug, log to fmt.Println")
	flag.Parse()
	if *debug {
		dbase2.SetQLogger(dbase2.FmtLog{})
	}

	GT = template.New("index").Funcs(tempower.FMap()).Funcs(TemplateFuncs())
	ad, err := AssetDir("assets/templates")
	for _, n := range ad {
		if path.Ext(n) == ".swp" {
			continue
		}
		t, err := Asset("assets/templates/" + n)
		dbase2.QLog("Parsing :" + n)
		GT = GT.New(n)
		_, err = GT.Parse(string(t))
		if err != nil {
			dbase2.QLog(err.Error())
			fmt.Println(err)
			return
		}
	}

	http.HandleFunc("/", Handle)
	http.HandleFunc("/s/", HandleStatic)
	http.HandleFunc("/newfamily", HandleNewFamily)
	http.HandleFunc("/login", HandleLogin)
	http.HandleFunc("/logout", HandleLogout)

	//Edits
	http.HandleFunc("/addmember", LoggedInFunc(HandleAddMember, true))
	http.HandleFunc("/addaccount", LoggedInFunc(HandleAddAccount, true))
	http.HandleFunc("/pay", LoggedInFunc(HandlePay, true))
	http.HandleFunc("/cancelstanding", LoggedInFunc(HandleCancelStanding, true))
	http.HandleFunc("/makerequest", LoggedInFunc(HandleMakeRequest, true))
	http.HandleFunc("/respondrequest", LoggedInFunc(HandleRespondRequest, true))
	http.HandleFunc("/addstanding", LoggedInFunc(HandleAddStanding, true))
	http.HandleFunc("/chpass", LoggedInFunc(HandlePasswordChange, true))

	//Views
	http.HandleFunc("/transactions", LoggedInVTemp("transactions.html"))
	http.HandleFunc("/personal", LoggedInVTemp("userhome.html"))
	http.HandleFunc("/family", LoggedInVTemp("familypage.html"))
	http.HandleFunc("/view", LoggedInView("viewac.html"))

	if *insecure {
		err = http.ListenAndServe(":8080", nil)
		return

	}
	err = http.ListenAndServeTLS(":8081", "data/server.pub", "data/server.key", nil)
	if err != nil {
		fmt.Println(err)
	}
}
