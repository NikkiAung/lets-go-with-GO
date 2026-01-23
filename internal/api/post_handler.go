package api

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/NikkiAung/go-fundmentals/internal/store"
	"github.com/NikkiAung/go-fundmentals/internal/utils"
)

type PostHandler struct {
	postStore store.PostStore
	logger *log.Logger
}

func NewPostHandler (ps store.PostStore,logger *log.Logger) *PostHandler {
	return &PostHandler{
		postStore: ps,
		logger: logger,
	}
}

func (ph *PostHandler)HandleCreatePost(w http.ResponseWriter, r *http.Request) {
	var post store.Post

	// JSON -> Go struct
	err := json.NewDecoder(r.Body).Decode(&post)

	if err != nil {
		http.Error(w, "Invalid Request data", http.StatusBadRequest)
		return
	}

	createdPost, err := ph.postStore.CreatePost(&post)
	if err != nil {
		ph.logger.Printf("ERROR: request body error %v", err)
		utils.WriteJSON(w,http.StatusInternalServerError, utils.Payload{"error" : "Can't create post"})
		return
	}

	// w.Header().Set("Content-Type", "application/json")

	// w.WriteHeader(http.StatusCreated) // 201
	// json.NewEncoder(w).Encode(createdPost)
	utils.WriteJSON(w, http.StatusCreated, utils.Payload{"createdPost" : createdPost})
}

func (ph *PostHandler)HandleGetPosts(w http.ResponseWriter, r *http.Request) {
	posts, err := ph.postStore.GetPosts()
	if err != nil {
		ph.logger.Printf("ERROR: handle get post: %v", err)
		utils.WriteJSON(w,http.StatusInternalServerError, utils.Payload{"error" : "Can't get posts"})
		return
	}

	utils.WriteJSON(w, http.StatusOK, utils.Payload{"posts" : posts})
}

func (ph *PostHandler)HandleGetPostById(w http.ResponseWriter, r *http.Request) {
	postId, err := utils.ReadIDFromParams(r)
	if err != nil {
		http.NotFound(w,r)
		return
	}

	post, err := ph.postStore.GetPostById(postId)
	if err != nil {
		ph.logger.Printf("ERROR: handle get post: %v", err)
		utils.WriteJSON(w,http.StatusNotFound, utils.Payload{"error" : "Can't get post"})
		return
	}

	if post == nil {
		http.NotFound(w, r)
		return
	}

	utils.WriteJSON(w, http.StatusOK, utils.Payload{"post" : post})
}

func (ph *PostHandler) HandleUpdatePost(w http.ResponseWriter, r *http.Request) {
	postId, err := utils.ReadIDFromParams(r)
	if err != nil {
		http.NotFound(w,r)
		return
	}

	var post store.Post 
	err = json.NewDecoder(r.Body).Decode(&post)

	if err != nil {
		http.Error(w, "Invalid Request data", http.StatusBadRequest)
		return
	}

	post.ID = postId
	// Go struct 
	updatedPost, err := ph.postStore.UpdatePost(&post)

	if err != nil {
		ph.logger.Printf("ERROR: handle update post: %v", err)
		utils.WriteJSON(w,http.StatusNotFound, utils.Payload{"error" : "Can't update post"})
		return
	}

	if updatedPost == nil {
		http.NotFound(w, r)
		return
	}

	utils.WriteJSON(w, http.StatusOK, utils.Payload{"post" : updatedPost})
}

func (ph *PostHandler) HandleDeletePost(w http.ResponseWriter, r *http.Request) {
	postId, err := utils.ReadIDFromParams(r)
	if err != nil {
		http.NotFound(w,r)
		return
	}

	err = ph.postStore.DeletePost(postId)

	if err != nil {
		ph.logger.Printf("ERROR: handle delete post: %v", err)
		utils.WriteJSON(w,http.StatusNotFound, utils.Payload{"error" : "Can't delete post"})
		return
	}

	w.WriteHeader(http.StatusNoContent)

}