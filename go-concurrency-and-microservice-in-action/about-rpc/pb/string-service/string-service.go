package stringservice

import (
	"context"
	"strings"

	"github.com/lvdbing/books/go-concurrency-and-microservice-in-action/about-rpc/pb"
)

type StringService struct {
}

func (s *StringService) Concat(ctx context.Context, req *pb.StringRequst) (*pb.StringResponse, error) {
	resp := pb.StringResponse{Ret: req.A + req.B}
	return &resp, nil
}

func (s *StringService) Diff(ctx context.Context, req *pb.StringRequst) (*pb.StringResponse, error) {
	res := ""
	longText, shortText := req.A, req.B
	if len(req.A) < len(req.B) {
		longText, shortText = req.B, req.A
	}
	for _, char := range shortText {
		if strings.Contains(longText, string(char)) {
			res += string(char)
		}
	}
	resp := pb.StringResponse{Ret: res}
	return &resp, nil
}
