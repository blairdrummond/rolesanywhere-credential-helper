// Code generated by smithy-go-codegen DO NOT EDIT.

package rolesanywhere

import (
	"bytes"
	"context"
	"fmt"
	smithy "github.com/aws/smithy-go"
	"github.com/aws/smithy-go/encoding/httpbinding"
	smithyjson "github.com/aws/smithy-go/encoding/json"
	"github.com/aws/smithy-go/middleware"
	"github.com/aws/smithy-go/tracing"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

type awsRestjson1_serializeOpCreateSession struct {
}

func (*awsRestjson1_serializeOpCreateSession) ID() string {
	return "OperationSerializer"
}

func (m *awsRestjson1_serializeOpCreateSession) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
) {
	_, span := tracing.StartSpan(ctx, "OperationSerializer")
	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
	defer endTimer()
	defer span.End()
	request, ok := in.Request.(*smithyhttp.Request)
	if !ok {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
	}

	input, ok := in.Parameters.(*CreateSessionInput)
	_ = input
	if !ok {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown input parameters type %T", in.Parameters)}
	}

	opPath, opQuery := httpbinding.SplitURI("/sessions")
	request.URL.Path = smithyhttp.JoinPath(request.URL.Path, opPath)
	request.URL.RawQuery = smithyhttp.JoinRawQuery(request.URL.RawQuery, opQuery)
	request.Method = "POST"
	var restEncoder *httpbinding.Encoder
	if request.URL.RawPath == "" {
		restEncoder, err = httpbinding.NewEncoder(request.URL.Path, request.URL.RawQuery, request.Header)
	} else {
		request.URL.RawPath = smithyhttp.JoinPath(request.URL.RawPath, opPath)
		restEncoder, err = httpbinding.NewEncoderWithRawPath(request.URL.Path, request.URL.RawPath, request.URL.RawQuery, request.Header)
	}

	if err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}

	if err := awsRestjson1_serializeOpHttpBindingsCreateSessionInput(input, restEncoder); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}

	restEncoder.SetHeader("Content-Type").String("application/json")

	jsonEncoder := smithyjson.NewEncoder()
	if err := awsRestjson1_serializeOpDocumentCreateSessionInput(input, jsonEncoder.Value); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}

	if request, err = request.SetStream(bytes.NewReader(jsonEncoder.Bytes())); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}

	if request.Request, err = restEncoder.Encode(request.Request); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}
	in.Request = request

	endTimer()
	span.End()
	return next.HandleSerialize(ctx, in)
}
func awsRestjson1_serializeOpHttpBindingsCreateSessionInput(v *CreateSessionInput, encoder *httpbinding.Encoder) error {
	if v == nil {
		return fmt.Errorf("unsupported serialization of nil %T", v)
	}

	if v.Cert != nil {
		locationName := "X-Amz-X509"
		encoder.SetHeader(locationName).String(*v.Cert)
	}

	if v.ProfileArn != nil {
		encoder.SetQuery("profileArn").String(*v.ProfileArn)
	}

	if v.RoleArn != nil {
		encoder.SetQuery("roleArn").String(*v.RoleArn)
	}

	if v.TrustAnchorArn != nil {
		encoder.SetQuery("trustAnchorArn").String(*v.TrustAnchorArn)
	}

	return nil
}

func awsRestjson1_serializeOpDocumentCreateSessionInput(v *CreateSessionInput, value smithyjson.Value) error {
	object := value.Object()
	defer object.Close()

	if v.DurationSeconds != nil {
		ok := object.Key("durationSeconds")
		ok.Integer(*v.DurationSeconds)
	}

	if v.InstanceProperties != nil {
		ok := object.Key("instanceProperties")
		if err := awsRestjson1_serializeDocumentInstancePropertyMap(v.InstanceProperties, ok); err != nil {
			return err
		}
	}

	if v.RoleSessionName != nil {
		ok := object.Key("roleSessionName")
		ok.String(*v.RoleSessionName)
	}

	if v.SessionName != nil {
		ok := object.Key("sessionName")
		ok.String(*v.SessionName)
	}

	return nil
}

func awsRestjson1_serializeDocumentInstancePropertyMap(v map[string]string, value smithyjson.Value) error {
	object := value.Object()
	defer object.Close()

	for key := range v {
		om := object.Key(key)
		om.String(v[key])
	}
	return nil
}
