// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/cloud/automl/v1beta1/prediction_service.proto

package automl

import (
	context "context"
	fmt "fmt"
	math "math"

	proto "github.com/golang/protobuf/proto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	longrunning "google.golang.org/genproto/googleapis/longrunning"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

// Request message for [PredictionService.Predict][google.cloud.automl.v1beta1.PredictionService.Predict].
type PredictRequest struct {
	// Name of the model requested to serve the prediction.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// Required.
	// Payload to perform a prediction on. The payload must match the
	// problem type that the model was trained to solve.
	Payload *ExamplePayload `protobuf:"bytes,2,opt,name=payload,proto3" json:"payload,omitempty"`
	// Additional domain-specific parameters, any string must be up to 25000
	// characters long.
	//
	// *  For Image Classification:
	//
	//    `score_threshold` - (float) A value from 0.0 to 1.0. When the model
	//     makes predictions for an image, it will only produce results that have
	//     at least this confidence score. The default is 0.5.
	// *  For Tables:
	//    `feature_importance` - (boolean) Whether
	//
	// [feature_importance][[google.cloud.automl.v1beta1.TablesModelColumnInfo.feature_importance]
	//        should be populated in the returned
	//
	// [TablesAnnotation(-s)][[google.cloud.automl.v1beta1.TablesAnnotation].
	//        The default is false.
	Params               map[string]string `protobuf:"bytes,3,rep,name=params,proto3" json:"params,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *PredictRequest) Reset()         { *m = PredictRequest{} }
func (m *PredictRequest) String() string { return proto.CompactTextString(m) }
func (*PredictRequest) ProtoMessage()    {}
func (*PredictRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_59a9dba5da3c687d, []int{0}
}

func (m *PredictRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PredictRequest.Unmarshal(m, b)
}
func (m *PredictRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PredictRequest.Marshal(b, m, deterministic)
}
func (m *PredictRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PredictRequest.Merge(m, src)
}
func (m *PredictRequest) XXX_Size() int {
	return xxx_messageInfo_PredictRequest.Size(m)
}
func (m *PredictRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_PredictRequest.DiscardUnknown(m)
}

var xxx_messageInfo_PredictRequest proto.InternalMessageInfo

func (m *PredictRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *PredictRequest) GetPayload() *ExamplePayload {
	if m != nil {
		return m.Payload
	}
	return nil
}

func (m *PredictRequest) GetParams() map[string]string {
	if m != nil {
		return m.Params
	}
	return nil
}

// Response message for [PredictionService.Predict][google.cloud.automl.v1beta1.PredictionService.Predict].
type PredictResponse struct {
	// Prediction result.
	// Translation and Text Sentiment will return precisely one payload.
	Payload []*AnnotationPayload `protobuf:"bytes,1,rep,name=payload,proto3" json:"payload,omitempty"`
	// Additional domain-specific prediction response metadata.
	//
	// * For Image Object Detection:
	//  `max_bounding_box_count` - (int64) At most that many bounding boxes per
	//      image could have been returned.
	//
	// * For Text Sentiment:
	//  `sentiment_score` - (float, deprecated) A value between -1 and 1,
	//      -1 maps to least positive sentiment, while 1 maps to the most positive
	//      one and the higher the score, the more positive the sentiment in the
	//      document is. Yet these values are relative to the training data, so
	//      e.g. if all data was positive then -1 will be also positive (though
	//      the least).
	//      The sentiment_score shouldn't be confused with "score" or "magnitude"
	//      from the previous Natural Language Sentiment Analysis API.
	Metadata             map[string]string `protobuf:"bytes,2,rep,name=metadata,proto3" json:"metadata,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *PredictResponse) Reset()         { *m = PredictResponse{} }
func (m *PredictResponse) String() string { return proto.CompactTextString(m) }
func (*PredictResponse) ProtoMessage()    {}
func (*PredictResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_59a9dba5da3c687d, []int{1}
}

func (m *PredictResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PredictResponse.Unmarshal(m, b)
}
func (m *PredictResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PredictResponse.Marshal(b, m, deterministic)
}
func (m *PredictResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PredictResponse.Merge(m, src)
}
func (m *PredictResponse) XXX_Size() int {
	return xxx_messageInfo_PredictResponse.Size(m)
}
func (m *PredictResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_PredictResponse.DiscardUnknown(m)
}

var xxx_messageInfo_PredictResponse proto.InternalMessageInfo

func (m *PredictResponse) GetPayload() []*AnnotationPayload {
	if m != nil {
		return m.Payload
	}
	return nil
}

func (m *PredictResponse) GetMetadata() map[string]string {
	if m != nil {
		return m.Metadata
	}
	return nil
}

// Request message for [PredictionService.BatchPredict][google.cloud.automl.v1beta1.PredictionService.BatchPredict].
type BatchPredictRequest struct {
	// Name of the model requested to serve the batch prediction.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// Required. The input configuration for batch prediction.
	InputConfig *BatchPredictInputConfig `protobuf:"bytes,3,opt,name=input_config,json=inputConfig,proto3" json:"input_config,omitempty"`
	// Required. The Configuration specifying where output predictions should
	// be written.
	OutputConfig *BatchPredictOutputConfig `protobuf:"bytes,4,opt,name=output_config,json=outputConfig,proto3" json:"output_config,omitempty"`
	// Additional domain-specific parameters for the predictions, any string must
	// be up to 25000 characters long.
	//
	// *  For Video Classification :
	//    `score_threshold` - (float) A value from 0.0 to 1.0. When the model
	//        makes predictions for a video, it will only produce results that
	//        have at least this confidence score. The default is 0.5.
	//    `segment_classification` - (boolean) Set to true to request
	//        segment-level classification. AutoML Video Intelligence returns
	//        labels and their confidence scores for the entire segment of the
	//        video that user specified in the request configuration.
	//        The default is "true".
	//    `shot_classification` - (boolean) Set to true to request shot-level
	//        classification. AutoML Video Intelligence determines the boundaries
	//        for each camera shot in the entire segment of the video that user
	//        specified in the request configuration. AutoML Video Intelligence
	//        then returns labels and their confidence scores for each detected
	//        shot, along with the start and end time of the shot.
	//        WARNING: Model evaluation is not done for this classification type,
	//        the quality of it depends on training data, but there are no metrics
	//        provided to describe that quality. The default is "false".
	//    `1s_interval_classification` - (boolean) Set to true to request
	//        classification for a video at one-second intervals. AutoML Video
	//        Intelligence returns labels and their confidence scores for each
	//        second of the entire segment of the video that user specified in the
	//        request configuration.
	//        WARNING: Model evaluation is not done for this classification
	//        type, the quality of it depends on training data, but there are no
	//        metrics provided to describe that quality. The default is
	//        "false".
	//
	// *  For Video Object Tracking:
	//    `score_threshold` - (float) When Model detects objects on video frames,
	//        it will only produce bounding boxes which have at least this
	//        confidence score. Value in 0 to 1 range, default is 0.5.
	//    `max_bounding_box_count` - (int64) No more than this number of bounding
	//        boxes will be returned per frame. Default is 100, the requested
	//        value may be limited by server.
	//    `min_bounding_box_size` - (float) Only bounding boxes with shortest edge
	//      at least that long as a relative value of video frame size will be
	//      returned. Value in 0 to 1 range. Default is 0.
	//
	Params               map[string]string `protobuf:"bytes,5,rep,name=params,proto3" json:"params,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *BatchPredictRequest) Reset()         { *m = BatchPredictRequest{} }
func (m *BatchPredictRequest) String() string { return proto.CompactTextString(m) }
func (*BatchPredictRequest) ProtoMessage()    {}
func (*BatchPredictRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_59a9dba5da3c687d, []int{2}
}

func (m *BatchPredictRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BatchPredictRequest.Unmarshal(m, b)
}
func (m *BatchPredictRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BatchPredictRequest.Marshal(b, m, deterministic)
}
func (m *BatchPredictRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BatchPredictRequest.Merge(m, src)
}
func (m *BatchPredictRequest) XXX_Size() int {
	return xxx_messageInfo_BatchPredictRequest.Size(m)
}
func (m *BatchPredictRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_BatchPredictRequest.DiscardUnknown(m)
}

var xxx_messageInfo_BatchPredictRequest proto.InternalMessageInfo

func (m *BatchPredictRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *BatchPredictRequest) GetInputConfig() *BatchPredictInputConfig {
	if m != nil {
		return m.InputConfig
	}
	return nil
}

func (m *BatchPredictRequest) GetOutputConfig() *BatchPredictOutputConfig {
	if m != nil {
		return m.OutputConfig
	}
	return nil
}

func (m *BatchPredictRequest) GetParams() map[string]string {
	if m != nil {
		return m.Params
	}
	return nil
}

// Result of the Batch Predict. This message is returned in
// [response][google.longrunning.Operation.response] of the operation returned
// by the [PredictionService.BatchPredict][google.cloud.automl.v1beta1.PredictionService.BatchPredict].
type BatchPredictResult struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *BatchPredictResult) Reset()         { *m = BatchPredictResult{} }
func (m *BatchPredictResult) String() string { return proto.CompactTextString(m) }
func (*BatchPredictResult) ProtoMessage()    {}
func (*BatchPredictResult) Descriptor() ([]byte, []int) {
	return fileDescriptor_59a9dba5da3c687d, []int{3}
}

func (m *BatchPredictResult) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BatchPredictResult.Unmarshal(m, b)
}
func (m *BatchPredictResult) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BatchPredictResult.Marshal(b, m, deterministic)
}
func (m *BatchPredictResult) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BatchPredictResult.Merge(m, src)
}
func (m *BatchPredictResult) XXX_Size() int {
	return xxx_messageInfo_BatchPredictResult.Size(m)
}
func (m *BatchPredictResult) XXX_DiscardUnknown() {
	xxx_messageInfo_BatchPredictResult.DiscardUnknown(m)
}

var xxx_messageInfo_BatchPredictResult proto.InternalMessageInfo

func init() {
	proto.RegisterType((*PredictRequest)(nil), "google.cloud.automl.v1beta1.PredictRequest")
	proto.RegisterMapType((map[string]string)(nil), "google.cloud.automl.v1beta1.PredictRequest.ParamsEntry")
	proto.RegisterType((*PredictResponse)(nil), "google.cloud.automl.v1beta1.PredictResponse")
	proto.RegisterMapType((map[string]string)(nil), "google.cloud.automl.v1beta1.PredictResponse.MetadataEntry")
	proto.RegisterType((*BatchPredictRequest)(nil), "google.cloud.automl.v1beta1.BatchPredictRequest")
	proto.RegisterMapType((map[string]string)(nil), "google.cloud.automl.v1beta1.BatchPredictRequest.ParamsEntry")
	proto.RegisterType((*BatchPredictResult)(nil), "google.cloud.automl.v1beta1.BatchPredictResult")
}

func init() {
	proto.RegisterFile("google/cloud/automl/v1beta1/prediction_service.proto", fileDescriptor_59a9dba5da3c687d)
}

var fileDescriptor_59a9dba5da3c687d = []byte{
	// 700 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0x54, 0xcd, 0x6e, 0x13, 0x3d,
	0x14, 0xd5, 0x24, 0xfd, 0xf9, 0xea, 0xb4, 0x1f, 0x60, 0x0a, 0x44, 0x53, 0x7e, 0xaa, 0xc0, 0xa2,
	0x4a, 0xcb, 0x98, 0x96, 0xa2, 0xc2, 0xb4, 0x2c, 0x92, 0xaa, 0x2a, 0x95, 0xa8, 0x1a, 0x05, 0x54,
	0xa4, 0xaa, 0x52, 0xe4, 0x4e, 0xdc, 0xe9, 0x80, 0xc7, 0x36, 0x33, 0x9e, 0x96, 0x0a, 0xb1, 0xe1,
	0x15, 0x78, 0x09, 0x1e, 0x80, 0x2d, 0x2f, 0xd0, 0x25, 0xbc, 0x02, 0x2b, 0x36, 0xb0, 0x63, 0x8b,
	0xc6, 0x76, 0x92, 0x09, 0x45, 0xa3, 0x64, 0xc1, 0x2e, 0x9e, 0x7b, 0xcf, 0xb9, 0xe7, 0xd8, 0x37,
	0x07, 0x2c, 0xfb, 0x9c, 0xfb, 0x94, 0x20, 0x8f, 0xf2, 0xa4, 0x8d, 0x70, 0x22, 0x79, 0x48, 0xd1,
	0xf1, 0xe2, 0x01, 0x91, 0x78, 0x11, 0x89, 0x88, 0xb4, 0x03, 0x4f, 0x06, 0x9c, 0xb5, 0x62, 0x12,
	0x1d, 0x07, 0x1e, 0x71, 0x44, 0xc4, 0x25, 0x87, 0x33, 0x1a, 0xe5, 0x28, 0x94, 0xa3, 0x51, 0x8e,
	0x41, 0xd9, 0xd7, 0x0d, 0x25, 0x16, 0x01, 0xc2, 0x8c, 0x71, 0x89, 0x53, 0x86, 0x58, 0x43, 0xed,
	0xdc, 0x81, 0xbd, 0xf6, 0x96, 0xc0, 0xa7, 0x94, 0xe3, 0xb6, 0x41, 0x2d, 0xe4, 0xa1, 0xda, 0x58,
	0xe2, 0x56, 0x20, 0x49, 0xd8, 0x99, 0x71, 0x27, 0xaf, 0x3b, 0xe0, 0x83, 0x70, 0x72, 0x41, 0xa2,
	0x3e, 0xdd, 0xb7, 0x4d, 0x37, 0xe5, 0xcc, 0x8f, 0x12, 0xc6, 0x02, 0xe6, 0x9f, 0x6f, 0xba, 0x96,
	0xb1, 0xee, 0xd1, 0x80, 0x30, 0xa9, 0x0b, 0x95, 0x5f, 0x16, 0xf8, 0xbf, 0xa1, 0x6f, 0xb3, 0x49,
	0x5e, 0x27, 0x24, 0x96, 0x10, 0x82, 0x11, 0x86, 0x43, 0x52, 0xb6, 0x66, 0xad, 0xb9, 0x89, 0xa6,
	0xfa, 0x0d, 0x37, 0xc0, 0xb8, 0xf1, 0x5d, 0x2e, 0xcc, 0x5a, 0x73, 0xa5, 0xa5, 0x79, 0x27, 0xe7,
	0xa6, 0x9d, 0x8d, 0x37, 0x38, 0x14, 0x94, 0x34, 0x34, 0xa4, 0xd9, 0xc1, 0xc2, 0x1d, 0x30, 0x26,
	0x70, 0x84, 0xc3, 0xb8, 0x5c, 0x9c, 0x2d, 0xce, 0x95, 0x96, 0x56, 0x72, 0x59, 0xfa, 0x75, 0x39,
	0x0d, 0x85, 0xdc, 0x60, 0x32, 0x3a, 0x6d, 0x1a, 0x1a, 0xfb, 0x11, 0x28, 0x65, 0x3e, 0xc3, 0x8b,
	0xa0, 0xf8, 0x8a, 0x9c, 0x1a, 0xe5, 0xe9, 0x4f, 0x38, 0x0d, 0x46, 0x8f, 0x31, 0x4d, 0x88, 0x92,
	0x3d, 0xd1, 0xd4, 0x07, 0xb7, 0xf0, 0xd0, 0xaa, 0xfc, 0xb4, 0xc0, 0x85, 0xee, 0x84, 0x58, 0x70,
	0x16, 0x13, 0xf8, 0xa4, 0x67, 0xd3, 0x52, 0x02, 0x9d, 0x5c, 0x81, 0xb5, 0xee, 0x56, 0x9c, 0x73,
	0xba, 0x0b, 0xfe, 0x0b, 0x89, 0xc4, 0xe9, 0x06, 0x94, 0x0b, 0x8a, 0xca, 0x1d, 0xcc, 0xab, 0x56,
	0xe2, 0x6c, 0x1b, 0xb0, 0xb6, 0xdb, 0xe5, 0xb2, 0x57, 0xc1, 0x54, 0x5f, 0x69, 0x28, 0xcb, 0x3f,
	0x0a, 0xe0, 0x72, 0x1d, 0x4b, 0xef, 0x68, 0x80, 0x17, 0x7f, 0x01, 0x26, 0x03, 0x26, 0x12, 0xd9,
	0xf2, 0x38, 0x3b, 0x0c, 0xfc, 0x72, 0x51, 0x3d, 0xfb, 0x72, 0xae, 0x89, 0x2c, 0xf7, 0x56, 0x0a,
	0x5e, 0x57, 0xd8, 0x66, 0x29, 0xe8, 0x1d, 0xe0, 0x1e, 0x98, 0xe2, 0x89, 0xcc, 0x30, 0x8f, 0x28,
	0xe6, 0x07, 0x03, 0x33, 0xef, 0x28, 0xb4, 0xa1, 0x9e, 0xe4, 0x99, 0x13, 0x7c, 0xde, 0xdd, 0xaf,
	0x51, 0x75, 0xe7, 0x6b, 0x03, 0x93, 0xfe, 0xa3, 0x25, 0x9b, 0x06, 0xb0, 0x7f, 0x4a, 0x9c, 0x50,
	0xb9, 0xf4, 0xa9, 0x08, 0x2e, 0x35, 0xba, 0x11, 0xf6, 0x4c, 0x27, 0x18, 0xfc, 0x68, 0x81, 0x71,
	0xf3, 0x15, 0xce, 0x0f, 0xf1, 0xc7, 0xb0, 0x17, 0x86, 0xd9, 0xac, 0x4a, 0xfd, 0xfd, 0xd7, 0x6f,
	0x1f, 0x0a, 0x6b, 0x95, 0x95, 0x6e, 0xa4, 0xbc, 0x4d, 0x1f, 0xfc, 0xb1, 0x88, 0xf8, 0x4b, 0xe2,
	0xc9, 0x18, 0x55, 0x11, 0xe5, 0x9e, 0x4e, 0x0f, 0x54, 0x45, 0x21, 0x6f, 0x13, 0x1a, 0xa3, 0xea,
	0x3b, 0xd7, 0x84, 0xae, 0x6b, 0x55, 0x53, 0xa9, 0x93, 0x59, 0x5f, 0xf0, 0xde, 0xb0, 0x17, 0x6d,
	0xdf, 0xe8, 0x20, 0x32, 0xb9, 0xe5, 0xec, 0x74, 0x72, 0xab, 0xb2, 0xa9, 0x54, 0xd6, 0x2a, 0x6b,
	0xc3, 0xaa, 0x3c, 0xc8, 0xcc, 0x72, 0xad, 0xaa, 0xbd, 0x75, 0x56, 0xbb, 0x62, 0xe4, 0xe8, 0x81,
	0x58, 0x04, 0xb1, 0xe3, 0xf1, 0xf0, 0x4b, 0xcd, 0x39, 0x92, 0x52, 0xc4, 0x2e, 0x42, 0x27, 0x27,
	0x27, 0x7f, 0x14, 0xd3, 0xd4, 0x3d, 0xd2, 0x01, 0x7c, 0x57, 0x50, 0x2c, 0x0f, 0x79, 0x14, 0xd6,
	0x3f, 0x5b, 0xe0, 0x96, 0xc7, 0xc3, 0x3c, 0xab, 0xf5, 0xab, 0xe7, 0xde, 0xb5, 0x91, 0xe6, 0x6c,
	0xc3, 0xda, 0xab, 0x19, 0x98, 0xcf, 0x29, 0x66, 0xbe, 0xc3, 0x23, 0x1f, 0xf9, 0x84, 0xa9, 0x14,
	0x46, 0xbd, 0xe1, 0x7f, 0x0d, 0xfd, 0x55, 0x7d, 0x3c, 0x2b, 0xcc, 0x6c, 0xaa, 0xc6, 0xfd, 0xf5,
	0xb4, 0x69, 0xbf, 0x96, 0x48, 0xbe, 0x4d, 0xf7, 0x77, 0x75, 0xd3, 0xf7, 0xc2, 0x4d, 0x5d, 0x75,
	0x5d, 0x55, 0x76, 0x5d, 0x55, 0x7f, 0xea, 0xba, 0xa6, 0xe1, 0x60, 0x4c, 0x0d, 0xbb, 0xff, 0x3b,
	0x00, 0x00, 0xff, 0xff, 0x37, 0x66, 0x03, 0xac, 0x5b, 0x07, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// PredictionServiceClient is the client API for PredictionService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type PredictionServiceClient interface {
	// Perform an online prediction. The prediction result will be directly
	// returned in the response.
	// Available for following ML problems, and their expected request payloads:
	// * Image Classification - Image in .JPEG, .GIF or .PNG format, image_bytes
	//                          up to 30MB.
	// * Image Object Detection - Image in .JPEG, .GIF or .PNG format, image_bytes
	//                            up to 30MB.
	// * Text Classification - TextSnippet, content up to 60,000 characters,
	//                         UTF-8 encoded.
	// * Text Extraction - TextSnippet, content up to 30,000 characters,
	//                     UTF-8 NFC encoded.
	// * Translation - TextSnippet, content up to 25,000 characters, UTF-8
	//                 encoded.
	// * Tables - Row, with column values matching the columns of the model,
	//            up to 5MB. Not available for FORECASTING
	//
	// [prediction_type][google.cloud.automl.v1beta1.TablesModelMetadata.prediction_type].
	Predict(ctx context.Context, in *PredictRequest, opts ...grpc.CallOption) (*PredictResponse, error)
	// Perform a batch prediction. Unlike the online [Predict][google.cloud.automl.v1beta1.PredictionService.Predict], batch
	// prediction result won't be immediately available in the response. Instead,
	// a long running operation object is returned. User can poll the operation
	// result via [GetOperation][google.longrunning.Operations.GetOperation]
	// method. Once the operation is done, [BatchPredictResult][google.cloud.automl.v1beta1.BatchPredictResult] is returned in
	// the [response][google.longrunning.Operation.response] field.
	// Available for following ML problems:
	// * Video Classification
	// * Video Object Tracking
	// * Text Extraction
	// * Tables
	BatchPredict(ctx context.Context, in *BatchPredictRequest, opts ...grpc.CallOption) (*longrunning.Operation, error)
}

type predictionServiceClient struct {
	cc *grpc.ClientConn
}

func NewPredictionServiceClient(cc *grpc.ClientConn) PredictionServiceClient {
	return &predictionServiceClient{cc}
}

func (c *predictionServiceClient) Predict(ctx context.Context, in *PredictRequest, opts ...grpc.CallOption) (*PredictResponse, error) {
	out := new(PredictResponse)
	err := c.cc.Invoke(ctx, "/google.cloud.automl.v1beta1.PredictionService/Predict", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *predictionServiceClient) BatchPredict(ctx context.Context, in *BatchPredictRequest, opts ...grpc.CallOption) (*longrunning.Operation, error) {
	out := new(longrunning.Operation)
	err := c.cc.Invoke(ctx, "/google.cloud.automl.v1beta1.PredictionService/BatchPredict", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PredictionServiceServer is the server API for PredictionService service.
type PredictionServiceServer interface {
	// Perform an online prediction. The prediction result will be directly
	// returned in the response.
	// Available for following ML problems, and their expected request payloads:
	// * Image Classification - Image in .JPEG, .GIF or .PNG format, image_bytes
	//                          up to 30MB.
	// * Image Object Detection - Image in .JPEG, .GIF or .PNG format, image_bytes
	//                            up to 30MB.
	// * Text Classification - TextSnippet, content up to 60,000 characters,
	//                         UTF-8 encoded.
	// * Text Extraction - TextSnippet, content up to 30,000 characters,
	//                     UTF-8 NFC encoded.
	// * Translation - TextSnippet, content up to 25,000 characters, UTF-8
	//                 encoded.
	// * Tables - Row, with column values matching the columns of the model,
	//            up to 5MB. Not available for FORECASTING
	//
	// [prediction_type][google.cloud.automl.v1beta1.TablesModelMetadata.prediction_type].
	Predict(context.Context, *PredictRequest) (*PredictResponse, error)
	// Perform a batch prediction. Unlike the online [Predict][google.cloud.automl.v1beta1.PredictionService.Predict], batch
	// prediction result won't be immediately available in the response. Instead,
	// a long running operation object is returned. User can poll the operation
	// result via [GetOperation][google.longrunning.Operations.GetOperation]
	// method. Once the operation is done, [BatchPredictResult][google.cloud.automl.v1beta1.BatchPredictResult] is returned in
	// the [response][google.longrunning.Operation.response] field.
	// Available for following ML problems:
	// * Video Classification
	// * Video Object Tracking
	// * Text Extraction
	// * Tables
	BatchPredict(context.Context, *BatchPredictRequest) (*longrunning.Operation, error)
}

func RegisterPredictionServiceServer(s *grpc.Server, srv PredictionServiceServer) {
	s.RegisterService(&_PredictionService_serviceDesc, srv)
}

func _PredictionService_Predict_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PredictRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PredictionServiceServer).Predict(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.cloud.automl.v1beta1.PredictionService/Predict",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PredictionServiceServer).Predict(ctx, req.(*PredictRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PredictionService_BatchPredict_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BatchPredictRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PredictionServiceServer).BatchPredict(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.cloud.automl.v1beta1.PredictionService/BatchPredict",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PredictionServiceServer).BatchPredict(ctx, req.(*BatchPredictRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _PredictionService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "google.cloud.automl.v1beta1.PredictionService",
	HandlerType: (*PredictionServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Predict",
			Handler:    _PredictionService_Predict_Handler,
		},
		{
			MethodName: "BatchPredict",
			Handler:    _PredictionService_BatchPredict_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "google/cloud/automl/v1beta1/prediction_service.proto",
}
