// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import context "context"
import mock "github.com/stretchr/testify/mock"
import s3manager "github.com/IBM/ibm-cos-sdk-go/service/s3/s3manager"

// Uploader is an autogenerated mock type for the Uploader type
type Uploader struct {
	mock.Mock
}

// Upload provides a mock function with given fields: input, options
func (_m *Uploader) Upload(input *s3manager.UploadInput, options ...func(*s3manager.Uploader)) (*s3manager.UploadOutput, error) {
	_va := make([]interface{}, len(options))
	for _i := range options {
		_va[_i] = options[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, input)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *s3manager.UploadOutput
	if rf, ok := ret.Get(0).(func(*s3manager.UploadInput, ...func(*s3manager.Uploader)) *s3manager.UploadOutput); ok {
		r0 = rf(input, options...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*s3manager.UploadOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*s3manager.UploadInput, ...func(*s3manager.Uploader)) error); ok {
		r1 = rf(input, options...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UploadWithContext provides a mock function with given fields: ctx, input, opts
func (_m *Uploader) UploadWithContext(ctx context.Context, input *s3manager.UploadInput, opts ...func(*s3manager.Uploader)) (*s3manager.UploadOutput, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, input)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *s3manager.UploadOutput
	if rf, ok := ret.Get(0).(func(context.Context, *s3manager.UploadInput, ...func(*s3manager.Uploader)) *s3manager.UploadOutput); ok {
		r0 = rf(ctx, input, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*s3manager.UploadOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *s3manager.UploadInput, ...func(*s3manager.Uploader)) error); ok {
		r1 = rf(ctx, input, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UploadWithIterator provides a mock function with given fields: ctx, iter, opts
func (_m *Uploader) UploadWithIterator(ctx context.Context, iter s3manager.BatchUploadIterator, opts ...func(*s3manager.Uploader)) error {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, iter)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, s3manager.BatchUploadIterator, ...func(*s3manager.Uploader)) error); ok {
		r0 = rf(ctx, iter, opts...)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
