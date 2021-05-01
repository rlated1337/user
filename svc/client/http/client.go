// Code generated by truss. DO NOT EDIT.
// Rerunning truss will overwrite this file.
// Version: 72999ebd2f
// Version Date: Wed Mar 17 08:36:51 UTC 2021

// Package http provides an HTTP client for the User service.
package http

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"

	"github.com/gogo/protobuf/jsonpb"

	"github.com/go-kit/kit/endpoint"
	httptransport "github.com/go-kit/kit/transport/http"
	"github.com/pkg/errors"

	// This Service
	pb "uas/user"
	"uas/user/svc"
)

var (
	_ = endpoint.Chain
	_ = httptransport.NewClient
	_ = fmt.Sprint
	_ = bytes.Compare
	_ = ioutil.NopCloser
)

// New returns a service backed by an HTTP server living at the remote
// instance. We expect instance to come from a service discovery system, so
// likely of the form "host:port".
func New(instance string, options ...httptransport.ClientOption) (pb.UserServer, error) {

	if !strings.HasPrefix(instance, "http") {
		instance = "http://" + instance
	}
	u, err := url.Parse(instance)
	if err != nil {
		return nil, err
	}
	_ = u

	var InformationZeroEndpoint endpoint.Endpoint
	{
		InformationZeroEndpoint = httptransport.NewClient(
			"GET",
			copyURL(u, "/user/"),
			EncodeHTTPInformationZeroRequest,
			DecodeHTTPInformationResponse,
			options...,
		).Endpoint()
	}

	endpoints := svc.NewEndpoints()
	endpoints.InformationEndpoint = InformationZeroEndpoint

	return endpoints, nil
}

func copyURL(base *url.URL, path string) *url.URL {
	next := *base
	next.Path = path
	return &next
}

// CtxValuesToSend configures the http client to pull the specified keys out of
// the context and add them to the http request as headers.  Note that keys
// will have net/http.CanonicalHeaderKey called on them before being send over
// the wire and that is the form they will be available in the server context.
func CtxValuesToSend(keys ...string) httptransport.ClientOption {
	return httptransport.ClientBefore(func(ctx context.Context, r *http.Request) context.Context {
		for _, k := range keys {
			if v, ok := ctx.Value(k).(string); ok {
				r.Header.Set(k, v)
			}
		}
		return ctx
	})
}

// HTTP Client Decode

// DecodeHTTPInformationResponse is a transport/http.DecodeResponseFunc that decodes
// a JSON-encoded InformationResponse response from the HTTP response body.
// If the response has a non-200 status code, we will interpret that as an
// error and attempt to decode the specific error message from the response
// body. Primarily useful in a client.
func DecodeHTTPInformationResponse(_ context.Context, r *http.Response) (interface{}, error) {
	defer r.Body.Close()
	buf, err := ioutil.ReadAll(r.Body)
	if err == io.EOF {
		return nil, errors.New("response http body empty")
	}
	if err != nil {
		return nil, errors.Wrap(err, "cannot read http body")
	}

	if r.StatusCode != http.StatusOK {
		return nil, errors.Wrapf(errorDecoder(buf), "status code: '%d'", r.StatusCode)
	}

	var resp pb.InformationResponse
	if err = jsonpb.UnmarshalString(string(buf), &resp); err != nil {
		return nil, errorDecoder(buf)
	}

	return &resp, nil
}

// HTTP Client Encode

// EncodeHTTPInformationZeroRequest is a transport/http.EncodeRequestFunc
// that encodes a information request into the various portions of
// the http request (path, query, and body).
func EncodeHTTPInformationZeroRequest(_ context.Context, r *http.Request, request interface{}) error {
	strval := ""
	_ = strval
	req := request.(*pb.InformationRequest)
	_ = req

	r.Header.Set("transport", "HTTPJSON")
	r.Header.Set("request-url", r.URL.Path)

	// Set the path parameters
	path := strings.Join([]string{
		"",
		"user",
		fmt.Sprint(req.Id),
	}, "/")
	u, err := url.Parse(path)
	if err != nil {
		return errors.Wrapf(err, "couldn't unmarshal path %q", path)
	}
	r.URL.RawPath = u.RawPath
	r.URL.Path = u.Path

	// Set the query parameters
	values := r.URL.Query()
	var tmp []byte
	_ = tmp

	r.URL.RawQuery = values.Encode()
	return nil
}

func errorDecoder(buf []byte) error {
	var w errorWrapper
	if err := json.Unmarshal(buf, &w); err != nil {
		const size = 8196
		if len(buf) > size {
			buf = buf[:size]
		}
		return fmt.Errorf("response body '%s': cannot parse non-json request body", buf)
	}

	return errors.New(w.Error)
}

type errorWrapper struct {
	Error string `json:"error"`
}
