package main

import (
	"handmadehttp/pkg/handmadehttp"
	"time"

	"golang.org/x/exp/slog"
)

func main() {
	slog.Info("start server")
	s := handmadehttp.NewServer("tcp", ":8080", 5*time.Second)
	s.UpdateHandler("/echo",
		func(req *handmadehttp.Request, res *handmadehttp.Response) error {
			res.SetContent([]byte(req.URI))
			return nil
		})
	err := s.ListenAndServe()
	if err != nil {
		slog.Error("fail to start server %s", err)
	}
}