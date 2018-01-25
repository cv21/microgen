// This file was automatically generated by "microgen 0.7.0b" utility.
// Please, do not change functions names!
package httpconv

import (
	bytes "bytes"
	context "context"
	json "encoding/json"
	errors "errors"
	generated "github.com/devimteam/microgen/example/generated"
	mux "github.com/gorilla/mux"
	ioutil "io/ioutil"
	http "net/http"
	path "path"
	strconv "strconv"
)

func CommonHTTPRequestEncoder(_ context.Context, r *http.Request, request interface{}) error {
	var buf bytes.Buffer
	if err := json.NewEncoder(&buf).Encode(request); err != nil {
		return err
	}
	r.Body = ioutil.NopCloser(&buf)
	return nil
}

func CommonHTTPResponseEncoder(_ context.Context, w http.ResponseWriter, response interface{}) error {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	return json.NewEncoder(w).Encode(response)
}

func DecodeHTTPUppercaseRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var req generated.UppercaseRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	return &req, err
}

func DecodeHTTPCountRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var _param string
	var ok bool
	_vars := mux.Vars(r)
	_param, ok = _vars["text"]
	if !ok {
		return nil, errors.New("param text not found")
	}
	text := _param
	_param, ok = _vars["symbol"]
	if !ok {
		return nil, errors.New("param symbol not found")
	}
	symbol, err := strconv.ParseInt(_param, 10, 64)
	if err != nil {
		return nil, err
	}
	return &generated.CountRequest{
		Symbol: int(symbol),
		Text:   string(text),
	}, nil
}

func DecodeHTTPTestCaseRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var req generated.TestCaseRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	return &req, err
}

func DecodeHTTPUppercaseResponse(_ context.Context, r *http.Response) (interface{}, error) {
	var resp generated.UppercaseResponse
	err := json.NewDecoder(r.Body).Decode(&resp)
	return &resp, err
}

func DecodeHTTPCountResponse(_ context.Context, r *http.Response) (interface{}, error) {
	var resp generated.CountResponse
	err := json.NewDecoder(r.Body).Decode(&resp)
	return &resp, err
}

func DecodeHTTPTestCaseResponse(_ context.Context, r *http.Response) (interface{}, error) {
	var resp generated.TestCaseResponse
	err := json.NewDecoder(r.Body).Decode(&resp)
	return &resp, err
}

func EncodeHTTPUppercaseRequest(ctx context.Context, r *http.Request, request interface{}) error {
	r.URL.Path = path.Join(r.URL.Path, "uppercase")
	return CommonHTTPRequestEncoder(ctx, r, request)
}

func EncodeHTTPCountRequest(ctx context.Context, r *http.Request, request interface{}) error {
	req := request.(*generated.CountRequest)
	r.URL.Path = path.Join(r.URL.Path, "count",
		req.Text,
		strconv.FormatInt(int64(req.Symbol), 10),
	)
	return nil
}

func EncodeHTTPTestCaseRequest(ctx context.Context, r *http.Request, request interface{}) error {
	r.URL.Path = path.Join(r.URL.Path, "test-case")
	return CommonHTTPRequestEncoder(ctx, r, request)
}

func EncodeHTTPUppercaseResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	return CommonHTTPResponseEncoder(ctx, w, response)
}

func EncodeHTTPCountResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	return CommonHTTPResponseEncoder(ctx, w, response)
}

func EncodeHTTPTestCaseResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	return CommonHTTPResponseEncoder(ctx, w, response)
}
