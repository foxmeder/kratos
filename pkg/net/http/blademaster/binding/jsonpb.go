package binding

import (
	json "github.com/golang/protobuf/jsonpb"
	proto "github.com/golang/protobuf/proto"
	"github.com/pkg/errors"
	"net/http"
)

type jsonPbBinding struct{}

func (jsonPbBinding) Name() string {
	return "jsonpb"
}

func (jsonPbBinding) Bind(req *http.Request, obj interface{}) error {
	if msg, ok := obj.(proto.Message); ok {
		if err := json.Unmarshal(req.Body, msg); err != nil {
			return errors.WithStack(err)
		}
	} else {
		return errors.New("invalid proto.Message")
	}
	return nil
}
