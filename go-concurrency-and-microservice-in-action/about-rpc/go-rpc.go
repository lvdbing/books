package aboutrpc

import "strings"

type StringRequest struct {
	A string
	B string
}

type Service interface {
	Concat(req StringRequest, ret *string) error
	Diff(req StringRequest, ret *string) error
}

type StringService struct{}

func (s StringService) Concat(req StringRequest, ret *string) error {
	*ret = req.A + req.B
	return nil
}

func (s StringService) Diff(req StringRequest, ret *string) error {
	if len(req.A) == 0 || len(req.B) == 0 {
		*ret = ""
		return nil
	}
	res := ""
	longStr, shortStr := req.A, req.B
	if len(req.A) < len(req.B) {
		longStr, shortStr = req.B, req.A
	}
	for _, char := range shortStr {
		if strings.Contains(longStr, string(char)) {
			res += string(char)
		}
	}
	// if len(req.A) >= len(req.B) {
	// 	for _, char := range req.B {
	// 		if strings.Contains(req.A, string(char)) {
	// 			res += string(char)
	// 		}
	// 	}
	// } else {
	// 	for _, char := range req.A {
	// 		if strings.Contains(req.B, string(char)) {
	// 			res += string(char)
	// 		}
	// 	}
	// }
	*ret = res
	return nil
}
