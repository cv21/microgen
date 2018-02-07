// This file was automatically generated by "microgen 0.8.0alpha" utility.
// Please, do not edit.
package stringsvc

import (
	context "context"
	errors "errors"
	entity "github.com/devimteam/microgen/example/svc/entity"
	endpoint "github.com/go-kit/kit/endpoint"
	opentracing "github.com/go-kit/kit/tracing/opentracing"
	opentracinggo "github.com/opentracing/opentracing-go"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

func AllEndpoints(service StringService, tracer opentracinggo.Tracer) *Endpoints {
	return &Endpoints{
		CountEndpoint:     opentracing.TraceServer(tracer, "Count")(CountEndpoint(service)),
		TestCaseEndpoint:  opentracing.TraceServer(tracer, "TestCase")(TestCaseEndpoint(service)),
		UppercaseEndpoint: opentracing.TraceServer(tracer, "Uppercase")(UppercaseEndpoint(service)),
	}
}

type Endpoints struct {
	UppercaseEndpoint endpoint.Endpoint
	CountEndpoint     endpoint.Endpoint
	TestCaseEndpoint  endpoint.Endpoint
}

func (E *Endpoints) Uppercase(arg0 context.Context, arg1 map[string]string) (res0 string, res1 error) {
	request := UppercaseRequest{StringsMap: arg1}
	response, err := E.UppercaseEndpoint(arg0, &request)
	if err != nil {
		if e, ok := status.FromError(err); ok || e.Code() == codes.Internal || e.Code() == codes.Unknown {
			err = errors.New(e.Message())
		}
		return
	}
	return response.(*UppercaseResponse).Ans, err
}

func (E *Endpoints) Count(arg0 context.Context, arg1 string, arg2 string) (res0 int, res1 []int, res2 error) {
	request := CountRequest{
		Symbol: arg2,
		Text:   arg1,
	}
	response, err := E.CountEndpoint(arg0, &request)
	if err != nil {
		if e, ok := status.FromError(err); ok || e.Code() == codes.Internal || e.Code() == codes.Unknown {
			err = errors.New(e.Message())
		}
		return
	}
	return response.(*CountResponse).Count, response.(*CountResponse).Positions, err
}

func (E *Endpoints) TestCase(arg0 context.Context, arg1 []*entity.Comment) (res0 map[string]int, res1 error) {
	request := TestCaseRequest{Comments: arg1}
	response, err := E.TestCaseEndpoint(arg0, &request)
	if err != nil {
		if e, ok := status.FromError(err); ok || e.Code() == codes.Internal || e.Code() == codes.Unknown {
			err = errors.New(e.Message())
		}
		return
	}
	return response.(*TestCaseResponse).Tree, err
}

func UppercaseEndpoint(svc StringService) endpoint.Endpoint {
	return func(arg0 context.Context, request interface{}) (interface{}, error) {
		req := request.(*UppercaseRequest)
		res0, res1 := svc.Uppercase(arg0, req.StringsMap)
		return &UppercaseResponse{Ans: res0}, res1
	}
}

func CountEndpoint(svc StringService) endpoint.Endpoint {
	return func(arg0 context.Context, request interface{}) (interface{}, error) {
		req := request.(*CountRequest)
		res0, res1, res2 := svc.Count(arg0, req.Text, req.Symbol)
		return &CountResponse{
			Count:     res0,
			Positions: res1,
		}, res2
	}
}

func TestCaseEndpoint(svc StringService) endpoint.Endpoint {
	return func(arg0 context.Context, request interface{}) (interface{}, error) {
		req := request.(*TestCaseRequest)
		res0, res1 := svc.TestCase(arg0, req.Comments)
		return &TestCaseResponse{Tree: res0}, res1
	}
}
