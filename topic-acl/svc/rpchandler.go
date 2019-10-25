package svc

import (
	"context"
	proto "github.com/tian-yuan/iot-common/iotpb"
	"github.com/sirupsen/logrus"
	"github.com/pkg/errors"
)

type rpchandler struct {

}

func (r *rpchandler) Subscribe(ctx context.Context, in *proto.SubscribeMessageRequest, out *proto.SubscribeMessageResponse) error {
	if in.Guid < 0 ||
		in.Qos < 0 ||
			in.TopicFilter == "" {
				out.Code = 430
				out.Message = "parameter error."
				logrus.Error("parameter error.")
				return errors.New("parameter error.")
	}
	topicId, err := Global.TopicSvc.Subscribe(in.TopicFilter, in.Guid, int(in.Qos))
	out.TopicId = topicId
	if err != nil {
		out.Code = 600
		out.Message = err.Error()
	}
	return nil
}

func (r *rpchandler) UnSubscribe(ctx context.Context, in *proto.UnSubscribeMessageRequest, out *proto.UnSubscribeMessageResponse) error {
	if in.Guid < 0 ||
		in.TopicFilter == "" {
			out.Code = 430
			out.Message = "parameter error."
			logrus.Error("parameter error.")
			return errors.New("parameter error.")
	}
	_, err := Global.TopicSvc.UnSubscribe(in.TopicFilter, in.Guid)
	if err != nil {
		out.Code = 600
		out.Message = err.Error()
	}
	return nil
}