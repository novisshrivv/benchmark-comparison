package routes

import (
	"bytes"
	"io"
	"net/http"
)

var (
	GetPosts   = (&JsonPlaceholder{}).GetPosts
	CreatePost = (&JsonPlaceholder{}).CreatePost
)

type JsonPlaceholder struct{}

func (*JsonPlaceholder) GetPosts() (*http.Response, error) {
	resp, err := http.Get("https://jsonplaceholder.typicode.com/posts")

	if err != nil {
		return &http.Response{
			Body: io.NopCloser(bytes.NewBuffer(nil)),
		}, err
	}

	return resp, nil
}

func (*JsonPlaceholder) CreatePost(body io.ReadCloser) (*http.Response, error) {
	resp, err := http.Post("https://jsonplaceholder.typicode.com/posts", "application/json", body)

	if err != nil {
		return &http.Response{
			Body: io.NopCloser(bytes.NewBuffer(nil)),
		}, err
	}

	return resp, err
}

type PostsResource struct{}

func (rs PostsResource) List(w http.ResponseWriter, r *http.Request) {
	resp, err := GetPosts()

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	defer resp.Body.Close()

	w.Header().Set("Content-Type", "application/json")

	if _, err := io.Copy(w, resp.Body); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (rs PostsResource) Create(w http.ResponseWriter, r *http.Request) {
	resp, err := CreatePost(r.Body)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	defer resp.Body.Close()

	w.Header().Set("Content-Type", "application/json")

	if _, err := io.Copy(w, resp.Body); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
