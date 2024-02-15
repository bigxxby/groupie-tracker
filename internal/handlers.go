package bigxxby

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"regexp"
	"strconv"
	"strings"
)

var artists []Artist = GetContent()

func StaticHandler(w http.ResponseWriter, r *http.Request) {
	filePath := r.URL.Path
	if r.Method != "GET" {
		ErrorHandler(w, r, http.StatusMethodNotAllowed)
		return
	}
	if r.Header.Get("Sec-Fetch-Dest") == "document" { // если пользовтель просит получить ответ в виде документа выдаем ошибку , в других случаях выдаем 200
		fmt.Println("Permission denied")
		ErrorHandler(w, r, 404)
		return
	}

	// Предотвращение навигации вверх по дереву каталогов
	if strings.Contains(filePath, "..") {
		fmt.Println("ERROR:" + " PATH CONTAINS ..")
		ErrorHandler(w, r, 404)
		return
	}

	// Открываем файл
	file, err := http.Dir(".").Open(filePath)
	if err != nil {
		fmt.Println("ERROR: ", err)
		ErrorHandler(w, r, 404)
		return
	}
	defer file.Close()

	// Получаем информацию о файле
	fileInfo, err := file.Stat()
	if err != nil {
		fmt.Println("ERROR: ", err)
		ErrorHandler(w, r, 404)
		return
	}

	// Проверяем, является ли файл каталогом
	if fileInfo.IsDir() {
		// Если это каталог, добавляем индексный файл
		fmt.Println("INDEX")
		filePath += "/index.html"
	}

	// Открываем файл
	file, err = http.Dir(".").Open(filePath)
	if err != nil {
		fmt.Println("ERROR: ", err)
		ErrorHandler(w, r, 404)
		return
	}
	defer file.Close()
	// Копируем содержимое файла в ответ
	http.ServeContent(w, r, filePath, fileInfo.ModTime(), file)
}

func MainHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		log.Printf("Invalid URL: %s", r.URL.Path)

		ErrorHandler(w, r, 404)
		return
	}
	if r.Method != "GET" {
		ErrorHandler(w, r, http.StatusMethodNotAllowed)
		return
	}

	data, err := template.ParseFiles("./templates/index.html")
	if err != nil {
		log.Printf("Error parsing template: %v", err)
		ErrorHandler(w, r, 500)
		return
	}

	err = data.Execute(w, artists)
	if err != nil {
		log.Printf("Error executing template: %v", err)
		ErrorHandler(w, r, 500)
		return
	}
}

func ArtistIdHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		ErrorHandler(w, r, http.StatusMethodNotAllowed)
		return
	}

	route := regexp.MustCompile(`/artists/(\d+)$`) // проверка на /artist/любая цифра
	if !route.MatchString(r.URL.Path) {
		ErrorHandler(w, r, 404)
		return
	}
	matches := route.FindStringSubmatch(r.URL.Path)
	if matches[1][0] == '0' {

		ErrorHandler(w, r, 404)
		return
	}
	html, err := template.ParseFiles("templates/artist.html")
	if err != nil {
		fmt.Println("Internal server error", err.Error())
		ErrorHandler(w, r, 500)
		return
	}
	if len(matches) < 2 {
		ErrorHandler(w, r, 404)
		return
	}
	id, err := strconv.Atoi(matches[1])
	if err != nil {
		ErrorHandler(w, r, 404)
		return
	}

	if id <= 0 || id > len(artists) {
		ErrorHandler(w, r, 404)
		return
	}

	artist := artists[id-1]

	err = html.Execute(w, artist)
	if err != nil {
		log.Println(err.Error())
		ErrorHandler(w, r, 500)
		return
	}
}

type Error struct {
	Status    int
	ErrString string
}

func ErrorHandler(w http.ResponseWriter, r *http.Request, status int) {
	var err Error
	switch status {
	case 404:
		err.Status = http.StatusNotFound
		err.ErrString = "Not Found"
	case 500:
		err.Status = http.StatusInternalServerError
		err.ErrString = "Internal Server Error"
	case 502:
		err.Status = http.StatusBadGateway
		err.ErrString = "Bad Gateway"
	case http.StatusMethodNotAllowed:
		err.Status = http.StatusMethodNotAllowed

		err.ErrString = "Method not allowed"
	}

	html, errParsing := template.ParseFiles("templates/error.html")
	if errParsing != nil {
		fmt.Fprint(w, "Internal server error")
		return
	}

	w.WriteHeader(status)
	execErr := html.Execute(w, err)
	if execErr != nil {
		log.Println("Error Executing template ", execErr.Error())

		fmt.Fprint(w, "Internal server error")
		return
	}
}
