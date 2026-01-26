package api

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"regexp"

	"github.com/NikkiAung/go-fundmentals/internal/store"
	"github.com/NikkiAung/go-fundmentals/internal/utils"
)


type UserHandler struct {
	userStore store.UserStore
	logger *log.Logger
}

func NewUserHandler(userStore store.UserStore, logger *log.Logger) *UserHandler {
	return &UserHandler{
		userStore: userStore,
		logger: logger,
	}
}


type registerUserRequest struct {
	Username string `json:"username"`
	Email string `json:"email"`
	Password string `json:"password"`
}

func (h *UserHandler) validateRegsiterRequest(req *registerUserRequest) error {
	if req.Username == "" {
		return errors.New("username is required")
	}

	if len(req.Username) > 20 {
		return errors.New("username cannot be greater than 20 characters")
	}

	if req.Email == "" {
		return errors.New("email is required")
	}

	emailRegex := regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)

	if !emailRegex.MatchString(req.Email) {
		return errors.New("Invalid email")
	}

	if req.Password == "" {
		return errors.New("password is required") 
	}

	if len(req.Password) < 3 {
		return errors.New("password must be greater than 3 characters")
	}

	return nil
}

func (h *UserHandler) HandleRegister (w http.ResponseWriter, r *http.Request) {
	var req registerUserRequest
	err := json.NewDecoder(r.Body).Decode(&req)

	if err != nil {
		h.logger.Printf("ERROR: decoding request: %v", err)
		utils.WriteJSON(w, http.StatusBadRequest, utils.Payload{"error" : "Invalid request payload"})
	}

	err = h.validateRegsiterRequest(&req)
	if err != nil {
		utils.WriteJSON(w, http.StatusBadRequest, utils.Payload{"error" : err.Error()})
	}

	user := &store.User{
		Username: req.Username,
		Email: req.Email,
	}

	err = user.PasswordHash.Hash(req.Password)
	if err != nil {
		h.logger.Printf("ERROR: hashing password: %v", err)
		utils.WriteJSON(w, http.StatusBadRequest, utils.Payload{"error": "internal server error"})
		return
	}

	err = h.userStore.CreateUser(user)
	if err != nil {
		h.logger.Printf("ERROR: registering user: %v", err)
		utils.WriteJSON(w, http.StatusBadRequest, utils.Payload{"error": "internal server error"})
		return
	}

	utils.WriteJSON(w, http.StatusBadRequest, utils.Payload{"user": user})
}