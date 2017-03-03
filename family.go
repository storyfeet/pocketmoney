package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/coderconvoy/dbase2"
	"github.com/pkg/errors"
)

func LoggedInFamily(w http.ResponseWriter, r *http.Request) (*Family, string, error) {
	ld, iok := loginControl.GetLogin(w, r)
	if iok != dbase2.OK {
		return nil, "", errors.New("No login")
	}
	fam, err := LoadFamily(ld.Familyname)
	return fam, ld.Username, err

}

// LoadFamily Reads and Unmarshals the Family File
// Params family name,
// Returns loaded family or nil, followed by error
func LoadFamily(fname string) (*Family, error) {
	d, err := FamDB.ReadMap(fname)
	if err != nil {
		return nil, errors.Wrap(err, "No Family Exists")
	}

	var f Family
	err = json.Unmarshal(d, &f)
	if err != nil {
		return nil, errors.Wrap(err, "Corrupted Family File")
	}
	return &f, nil
}

func SaveFamily(f *Family) error {
	mar, err := json.MarshalIndent(&f, "", " ")
	if err != nil {
		return err
	}

	ok := FamDB.WriteMap(f.FamilyName, mar)
	if !ok {
		return errors.New("Could not save Family")
	}
	return nil
}

func HandleLogin(w http.ResponseWriter, r *http.Request) {
	//Load Family File
	famname := r.FormValue("family")
	fam, err := LoadFamily(famname)
	if err != nil {
		ExTemplate(GT, w, "index.html", err.Error())
		return
	}

	//Check Password
	uname := r.FormValue("username")
	pw := r.FormValue("pwd")
	if uname == "" {
		ExTemplate(GT, w, "index.html", "Username Blank")
	}

	sel := -1
	for i, v := range fam.Members {
		if v.Username == uname || (v.Email == uname) {
			if v.Password.Check(pw) {
				sel = i
				break
			}
		}
	}
	if sel == -1 {
		ExTemplate(GT, w, "index.html", "No Username-Password match")
		return
	}

	loginControl.Login(w, fam.FamilyName, uname)
	ExTemplate(GT, w, "familypage.html", PageData{"", uname, fam})

}
func HandleNewFamily(w http.ResponseWriter, r *http.Request) {
	var f Family
	f.FamilyName = r.FormValue("familyname")

	//Check if family already exists
	_, err := FamDB.ReadMap(f.FamilyName)
	if err == nil {
		ExTemplate(GT, w, "index.html", "Family Name Already Exists")
		return
	}

	p1 := r.FormValue("pwd1")
	p2 := r.FormValue("pwd2")

	if p1 != p2 {
		ExTemplate(GT, w, "index.html", "Passwords don't match")
		return
	}

	pw, err := dbase2.NewPassword(p1)
	if err != nil {
		ExTemplate(GT, w, "index.html", "Passwords error: "+err.Error())
		return
	}

	uname := r.FormValue("username")
	email := r.FormValue("email")
	f.Members = append(f.Members, User{
		Username: uname,
		Email:    email,
		Password: pw,
		Parent:   true,
	})

	err = SaveFamily(&f)
	if err != nil {
		//TODO
	}
	loginControl.Login(w, f.FamilyName, uname)
	ExTemplate(GT, w, "familypage.html", f)

}

func HandleAddMember(w http.ResponseWriter, r *http.Request) {
	fam, fmem, err := LoggedInFamily(w, r)
	if err != nil {
		ExTemplate(GT, w, "index.html", "Not Logged In")
		return
	}

	uname := r.FormValue("username")
	parent := r.FormValue("parent")
	fmt.Println("Parent == " + parent)
	pwd1 := r.FormValue("pwd1")
	pwd2 := r.FormValue("pwd2")

	if uname == "" {
		ExTemplate(GT, w, "familypage.html", PageData{"No Username", fmem, fam})
		return
	}
	if pwd1 != pwd2 || pwd1 == "" {
		ExTemplate(GT, w, "familypage.html", PageData{"Passwords not matching", fmem, fam})
		return
	}

	pw, err := dbase2.NewPassword(pwd1)
	if err != nil {
		fmt.Println("Could not Password: ", err)
		ExTemplate(GT, w, "index.html", PageData{"Password Failed", fmem, fam})
	}
	fam.Members = append(fam.Members, User{
		Username: uname,
		Parent:   parent == "on",
		Password: pw,
	})
	fam.Accounts = append(fam.Accounts, &Account{
		Username:  uname,
		Name:      "Checking",
		Current:   0,
		StartDate: time.Now(),
	})

	err = SaveFamily(fam)
	if err != nil {
		//TODO
	}
	ExTemplate(GT, w, "familypage.html", PageData{"", fmem, fam})
}