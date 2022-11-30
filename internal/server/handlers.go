package server

import (
	"Uploader/internal/models"
	"encoding/json"
	"io"
	"log"
	"net/http"
)

func (s *Server) Registration(w http.ResponseWriter, r *http.Request) {

	bytes, err := io.ReadAll(r.Body)
	if err != nil {
		log.Println(err)
	}

	var userInfo models.AuthInfo

	err = json.Unmarshal(bytes, &userInfo)
	if err != nil {
		log.Println(err)
		return
	}

	err = s.Services.Register(&userInfo)
	if err != nil {
		log.Println(err)
		return
	}

	w.WriteHeader(200)

}

func (s *Server) Login(w http.ResponseWriter, r *http.Request) {

	body, err := io.ReadAll(r.Body)
	if err != nil {
		log.Println(err)
	}

	var userInfo models.AuthInfo

	err = json.Unmarshal(body, &userInfo)
	if err != nil {
		log.Println(err)
		return
	}

	token, err := s.Services.Login(&userInfo)
	if err != nil {
		log.Println(err)
		return
	}
	log.Println(token)
	w.Write([]byte(token))
	w.WriteHeader(200)
}

func (s *Server) FolderCreator(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	value := ctx.Value(userID)
	userId := value.(string)

	body, err := io.ReadAll(r.Body)
	if err != nil {
		log.Println(err)
	}
	var FolderInfo models.Folder

	FolderInfo.UserID = userId

	err = json.Unmarshal(body, &FolderInfo)
	if err != nil {
		log.Println(err)
		return
	}

	s.Services.FolderCreation(&FolderInfo)

	log.Println("Folder created successful")
}