package main

import (
	"crypto/sha1"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"text/template"

	uuid "github.com/satori/go.uuid"
)

var users = map[string]UserData{}

// UserData contains the UUID & a the names of image files.
type UserData struct {
	UUID   string
	Images []Photo
}

// Photo represents an image.
type Photo struct {
	Name string
	Img  []byte
}

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("templates/*.html"))
}

func main() {
	http.Handle("/favicon.ico", http.NotFoundHandler())
	serveFile("public")
	http.HandleFunc("/", index)
	http.ListenAndServe("localhost:8080", nil)
}

func serveFile(fName string) {
	fs := http.FileServer(http.Dir("./" + fName))
	http.Handle("/"+fName+"/", http.StripPrefix("/"+fName, fs))
}

func index(w http.ResponseWriter, req *http.Request) {
	c := getCookie(w, req)
	user := getUser(c.Value)
	fmt.Println("userData on /index", user)
	fmt.Println(req.Method)
	if req.Method == http.MethodGet {
		fmt.Println("userData on GET", user)
	}

	if req.Method == http.MethodPost {
		fmt.Println("userData on POST", user)

		// Get the file & the file header; handle error
		img, fh, err := req.FormFile("name")
		check(err, "req.FormFile")
		defer img.Close()

		// create sha for file
		imgName := fh.Filename
		h := sha1.New()
		io.Copy(h, img)
		imgSha := h.Sum(nil)

		// create new file
		wd, err := os.Getwd()
		check(err, "os.Getwd")
		path := filepath.Join(wd, "public", "img", imgName)
		nf, err := os.Create(path)
		check(err, "os.Create")
		defer nf.Close()

		// copy
		img.Seek(0, 0)
		io.Copy(nf, img)

		// add file info to this user's data
		storeImg(imgName, imgSha, &user)
		saveUser(user)
		http.Redirect(w, req, "/", http.StatusOK)
	}

	tpl.ExecuteTemplate(w, "index.html", user)
}

func getCookie(w http.ResponseWriter, req *http.Request) *http.Cookie {
	c, err := req.Cookie("photosession")

	if err != nil {
		sessionID, err := uuid.NewV4()
		check(err, "uuid.NewV4")

		c = &http.Cookie{
			Name:  "photosession",
			Value: sessionID.String(),
		}
	}

	http.SetCookie(w, c)

	return c
}

func getUser(id string) UserData {
	fmt.Println("getUser", users)

	user, ok := users[id]
	if !ok {
		fmt.Println("User not found.")
		user.UUID = id
		saveUser(user)
		return user
	}
	return user
}

func saveUser(user UserData) {
	fmt.Println("Before saveUser", users)
	users[user.UUID] = user
	fmt.Println("After saveUser", users)

}

func storeImg(imgName string, imgSha []byte, user *UserData) {
	fmt.Println("Before storeImg", user)
	images := user.Images
	images = append(images, Photo{imgName, imgSha})
	user.Images = images
	fmt.Println("After storeImage", user)

}
