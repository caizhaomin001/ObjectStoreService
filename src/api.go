package main

import (
	"encoding/json"
	"net/http"
	"os"

	"github.com/go-martini/martini"
	"github.com/martini-contrib/binding"
)

type apiServer struct {
	server *Server
}

func (s *apiServer) DeleteObject(params martini.Params) (int, string) {
	err := s.server.DeleteObject(params["bucket"], params["object"])
	if err != nil {
		if os.IsNotExist(err) {
			return 404, "The object doesn't exists!"
		}
		return 500, err.Error()
	}
	return 200, ""
}

func (s *apiServer) DeleteBucket(params martini.Params) (int, string) {
	err := s.server.DeleteBucket(params["bucket"])
	if err != nil {
		return 500, err.Error()
	}
	return 200, ""
}

func (s *apiServer) CreateBucket(params martini.Params) (int, string) {
	err := s.server.CreateBucket(params["bucket"])
	if err != nil {
		return 500, err.Error()
	}
	return 200, ""
}

func (s *apiServer) ListBucket() (int, string) {
	data := make(map[string][]string)
	bucketList, err := s.server.ListBucket()
	if err != nil {
		return 500, err.Error()
	}
	data["bucket_list"] = bucketList
	return ApiResponseJson(data)
}

func (s *apiServer) ListObject(params martini.Params) (int, string) {
	data := make(map[string][]string)
	objectList, err := s.server.ListObject(params["bucket"])
	if err != nil {
		if os.IsNotExist(err) {
			return 404, "The bucket doesn't exists!"
		}
		return 500, err.Error()
	}
	if objectList == nil {
		data["object_list"] = []string{}
	} else {
		data["object_list"] = objectList
	}
	return ApiResponseJson(data)
}

func (s *apiServer) GetObject(params martini.Params) (int, string) {
	data, err := s.server.GetObject(params["bucket"], params["object"])
	if err != nil {
		if os.IsNotExist(err) {
			return 404, "The object doesn't exists!"
		}
		return 500, err.Error()
	}
	return ApiResponseJson(data)
}

func (s *apiServer) UploadObject(body uploadObjectRequestBody, params martini.Params, ) (int, string) {
	content := body.Content
	err := s.server.UploadObject(content, params["bucket"], params["object"])
	if err != nil {
		if os.IsNotExist(err) {
			return 404, "The bucket doesn't exists!"
		}
		return 500, err.Error()
	}
	return ApiResponseJson("OK")
}

func NewApiServer(s *Server) http.Handler {
	api := &apiServer{server: s}
	m := martini.Classic()
	m.Get("/:bucket/:object", api.GetObject)
	m.Put("/:bucket/:object", binding.Json(uploadObjectRequestBody{}), api.UploadObject)
	m.Delete("/:bucket/:object", api.DeleteObject)
	m.Post("/:bucket", api.CreateBucket)
	m.Get("/", api.ListBucket)
	m.Get("/:bucket", api.ListObject)
	m.Delete("/:bucket", api.DeleteBucket)
	return m
}

func ApiResponseJson(v interface{}) (int, string) {
	b, err := apiMarshalJson(v)
	if err != nil {
		return 500, err.Error()
	} else {
		return 200, string(b)
	}
}

func apiMarshalJson(v interface{}) ([]byte, error) {
	return json.MarshalIndent(v, "", "    ")
}
