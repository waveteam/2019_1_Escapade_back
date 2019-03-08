package api

import (
	"encoding/json"
	"errors"
	database "escapade/internal/database"
	"escapade/internal/misc"
	"fmt"
	"io"
	"net/http"
	"os"

	//"reflect"

	"github.com/gorilla/mux"
	"github.com/segmentio/ksuid"
)

// Handler is struct
type Handler struct {
	DB database.DataBase
}

// Init creates Handler
func Init(DB *database.DataBase) (handler *Handler) {
	handler = &Handler{
		DB: *DB,
	}
	return
}

// UploadAvatar uploads avatar
func (h *Handler) UploadAvatar(r *http.Request) (created bool, path string) {
	file, _, err := r.FormFile("avatar")

	if err != nil || file == nil {
		return true, "img/avatars/default"
	}

	defer file.Close()

	prefix := "img/avatars/"
	hash := ksuid.New()
	fileName := hash.String()

	createPath := "./" + prefix + fileName
	path = prefix + fileName

	out, err := os.Create(createPath)
	defer out.Close()

	if err != nil {

		return false, ""
	}

	_, err = io.Copy(out, file)
	if err != nil {
		return false, ""
	}

	return true, path
}

// Ok always returns StatusOk
func (h *Handler) Ok(rw http.ResponseWriter, r *http.Request) {
	rw.WriteHeader(http.StatusOK)
	sendSuccessJSON(rw, "Ok")

	fmt.Println("api/ok - ok")
	return
}

// Register handle registration
func (h *Handler) Register(rw http.ResponseWriter, r *http.Request) {

	const place = "Register"
	user := getUser(r)
	sessionID, err := h.DB.Register(&user)

	if err != nil {
		rw.WriteHeader(http.StatusForbidden)
		sendErrorJSON(rw, err, place)

		fmt.Println("api/register failed")
		return
	}

	sessionCookie := misc.CreateCookie(sessionID)
	http.SetCookie(rw, sessionCookie)
	rw.WriteHeader(http.StatusCreated)
	sendSuccessJSON(rw, place)

	fmt.Println("api/register ok")

	return
}

// Login handle login
func (h *Handler) Login(rw http.ResponseWriter, r *http.Request) {
	const place = "Login"
	user := getUser(r)
	sessionID, err := h.DB.Login(&user)

	if err != nil {
		rw.WriteHeader(http.StatusForbidden)
		sendErrorJSON(rw, err, place)
		return
	}

	sessionCookie := misc.CreateCookie(sessionID)
	http.SetCookie(rw, sessionCookie)
	rw.WriteHeader(http.StatusOK)
	sendSuccessJSON(rw, place)

	return
}

// Login handle logout
func (h *Handler) Logout(rw http.ResponseWriter, r *http.Request) {
	const place = "Logout"
	sessionID := misc.GetSessionCookie(r)
	err := h.DB.Logout(sessionID)

	if err != nil {
		rw.WriteHeader(http.StatusInternalServerError)
		sendErrorJSON(rw, err, place)

		return
	}

	http.SetCookie(rw, misc.CreateCookie(""))
	rw.WriteHeader(http.StatusOK)

	fmt.Println("api/logout ok")

	return
}

// DeleteAccount handle registration
func (h *Handler) DeleteAccount(rw http.ResponseWriter, r *http.Request) {

	const place = "DeleteAccount"
	user := getUser(r)
	sessionID := misc.GetSessionCookie(r)
	sessionID, err := h.DB.DeleteAccount(&user, sessionID)

	if err != nil {
		rw.WriteHeader(http.StatusForbidden)
		sendErrorJSON(rw, err, place)

		fmt.Println("api/DeleteAccount failed")
		return
	}

	http.SetCookie(rw, misc.CreateCookie(""))
	rw.WriteHeader(http.StatusOK)

	fmt.Println("api/DeleteAccount ok")
	return
}

// DeleteAccountOptions handle preCORS request
func (h *Handler) DeleteAccountOptions(rw http.ResponseWriter, r *http.Request) {
	fmt.Println("api/DeleteAccountOptions ok")
	rw.WriteHeader(http.StatusOK)
}

// GetPlayerGames handle get games list
func (h *Handler) GetPlayerGames(rw http.ResponseWriter, r *http.Request) {
	const place = "GetPlayerGames"

	vars := mux.Vars(r)
	username := vars["name"]

	if username == "" {
		fmt.Println("No username found")

		rw.WriteHeader(http.StatusInternalServerError)
		sendErrorJSON(rw, errors.New("No username found"), place)
		return
	}

	games, err := h.DB.GetGames(username)

	if err != nil {
		rw.WriteHeader(http.StatusInternalServerError)
		sendErrorJSON(rw, err, place)
		return
	}

	bytes, errJSON := json.Marshal(games)
	if errJSON == nil {
		rw.WriteHeader(http.StatusOK)
		fmt.Fprintln(rw, string(bytes))

		fmt.Println("api/GetPlayerGames ok")
	} else {
		rw.WriteHeader(http.StatusInternalServerError)
		sendErrorJSON(rw, err, place)

		fmt.Println("api/GetPlayerGames cant create json")
	}

}