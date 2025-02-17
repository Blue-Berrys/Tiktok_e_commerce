// Code generated by Fastpb v0.0.2. DO NOT EDIT.

package ai

import (
	fmt "fmt"
	fastpb "github.com/cloudwego/fastpb"
)

var (
	_ = fmt.Errorf
	_ = fastpb.Skip
)

func (x *OrderQueryRequest) FastRead(buf []byte, _type int8, number int32) (offset int, err error) {
	switch number {
	case 1:
		offset, err = x.fastReadField1(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	case 2:
		offset, err = x.fastReadField2(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	default:
		offset, err = fastpb.Skip(buf, _type, number)
		if err != nil {
			goto SkipFieldError
		}
	}
	return offset, nil
SkipFieldError:
	return offset, fmt.Errorf("%T cannot parse invalid wire-format data, error: %s", x, err)
ReadFieldError:
	return offset, fmt.Errorf("%T read field %d '%s' error: %s", x, number, fieldIDToName_OrderQueryRequest[number], err)
}

func (x *OrderQueryRequest) fastReadField1(buf []byte, _type int8) (offset int, err error) {
	x.Message, offset, err = fastpb.ReadString(buf, _type)
	return offset, err
}

func (x *OrderQueryRequest) fastReadField2(buf []byte, _type int8) (offset int, err error) {
	x.UserId, offset, err = fastpb.ReadInt32(buf, _type)
	return offset, err
}

func (x *OrderQueryResponse) FastRead(buf []byte, _type int8, number int32) (offset int, err error) {
	switch number {
	case 1:
		offset, err = x.fastReadField1(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	default:
		offset, err = fastpb.Skip(buf, _type, number)
		if err != nil {
			goto SkipFieldError
		}
	}
	return offset, nil
SkipFieldError:
	return offset, fmt.Errorf("%T cannot parse invalid wire-format data, error: %s", x, err)
ReadFieldError:
	return offset, fmt.Errorf("%T read field %d '%s' error: %s", x, number, fieldIDToName_OrderQueryResponse[number], err)
}

func (x *OrderQueryResponse) fastReadField1(buf []byte, _type int8) (offset int, err error) {
	x.Data, offset, err = fastpb.ReadString(buf, _type)
	return offset, err
}

func (x *AutoOrderRequest) FastRead(buf []byte, _type int8, number int32) (offset int, err error) {
	switch number {
	case 1:
		offset, err = x.fastReadField1(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	case 2:
		offset, err = x.fastReadField2(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	default:
		offset, err = fastpb.Skip(buf, _type, number)
		if err != nil {
			goto SkipFieldError
		}
	}
	return offset, nil
SkipFieldError:
	return offset, fmt.Errorf("%T cannot parse invalid wire-format data, error: %s", x, err)
ReadFieldError:
	return offset, fmt.Errorf("%T read field %d '%s' error: %s", x, number, fieldIDToName_AutoOrderRequest[number], err)
}

func (x *AutoOrderRequest) fastReadField1(buf []byte, _type int8) (offset int, err error) {
	x.Message, offset, err = fastpb.ReadString(buf, _type)
	return offset, err
}

func (x *AutoOrderRequest) fastReadField2(buf []byte, _type int8) (offset int, err error) {
	x.UserId, offset, err = fastpb.ReadInt32(buf, _type)
	return offset, err
}

func (x *AutoOrderResponse) FastRead(buf []byte, _type int8, number int32) (offset int, err error) {
	switch number {
	case 1:
		offset, err = x.fastReadField1(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	default:
		offset, err = fastpb.Skip(buf, _type, number)
		if err != nil {
			goto SkipFieldError
		}
	}
	return offset, nil
SkipFieldError:
	return offset, fmt.Errorf("%T cannot parse invalid wire-format data, error: %s", x, err)
ReadFieldError:
	return offset, fmt.Errorf("%T read field %d '%s' error: %s", x, number, fieldIDToName_AutoOrderResponse[number], err)
}

func (x *AutoOrderResponse) fastReadField1(buf []byte, _type int8) (offset int, err error) {
	x.Data, offset, err = fastpb.ReadString(buf, _type)
	return offset, err
}

func (x *OrderQueryRequest) FastWrite(buf []byte) (offset int) {
	if x == nil {
		return offset
	}
	offset += x.fastWriteField1(buf[offset:])
	offset += x.fastWriteField2(buf[offset:])
	return offset
}

func (x *OrderQueryRequest) fastWriteField1(buf []byte) (offset int) {
	if x.Message == "" {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 1, x.GetMessage())
	return offset
}

func (x *OrderQueryRequest) fastWriteField2(buf []byte) (offset int) {
	if x.UserId == 0 {
		return offset
	}
	offset += fastpb.WriteInt32(buf[offset:], 2, x.GetUserId())
	return offset
}

func (x *OrderQueryResponse) FastWrite(buf []byte) (offset int) {
	if x == nil {
		return offset
	}
	offset += x.fastWriteField1(buf[offset:])
	return offset
}

func (x *OrderQueryResponse) fastWriteField1(buf []byte) (offset int) {
	if x.Data == "" {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 1, x.GetData())
	return offset
}

func (x *AutoOrderRequest) FastWrite(buf []byte) (offset int) {
	if x == nil {
		return offset
	}
	offset += x.fastWriteField1(buf[offset:])
	offset += x.fastWriteField2(buf[offset:])
	return offset
}

func (x *AutoOrderRequest) fastWriteField1(buf []byte) (offset int) {
	if x.Message == "" {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 1, x.GetMessage())
	return offset
}

func (x *AutoOrderRequest) fastWriteField2(buf []byte) (offset int) {
	if x.UserId == 0 {
		return offset
	}
	offset += fastpb.WriteInt32(buf[offset:], 2, x.GetUserId())
	return offset
}

func (x *AutoOrderResponse) FastWrite(buf []byte) (offset int) {
	if x == nil {
		return offset
	}
	offset += x.fastWriteField1(buf[offset:])
	return offset
}

func (x *AutoOrderResponse) fastWriteField1(buf []byte) (offset int) {
	if x.Data == "" {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 1, x.GetData())
	return offset
}

func (x *OrderQueryRequest) Size() (n int) {
	if x == nil {
		return n
	}
	n += x.sizeField1()
	n += x.sizeField2()
	return n
}

func (x *OrderQueryRequest) sizeField1() (n int) {
	if x.Message == "" {
		return n
	}
	n += fastpb.SizeString(1, x.GetMessage())
	return n
}

func (x *OrderQueryRequest) sizeField2() (n int) {
	if x.UserId == 0 {
		return n
	}
	n += fastpb.SizeInt32(2, x.GetUserId())
	return n
}

func (x *OrderQueryResponse) Size() (n int) {
	if x == nil {
		return n
	}
	n += x.sizeField1()
	return n
}

func (x *OrderQueryResponse) sizeField1() (n int) {
	if x.Data == "" {
		return n
	}
	n += fastpb.SizeString(1, x.GetData())
	return n
}

func (x *AutoOrderRequest) Size() (n int) {
	if x == nil {
		return n
	}
	n += x.sizeField1()
	n += x.sizeField2()
	return n
}

func (x *AutoOrderRequest) sizeField1() (n int) {
	if x.Message == "" {
		return n
	}
	n += fastpb.SizeString(1, x.GetMessage())
	return n
}

func (x *AutoOrderRequest) sizeField2() (n int) {
	if x.UserId == 0 {
		return n
	}
	n += fastpb.SizeInt32(2, x.GetUserId())
	return n
}

func (x *AutoOrderResponse) Size() (n int) {
	if x == nil {
		return n
	}
	n += x.sizeField1()
	return n
}

func (x *AutoOrderResponse) sizeField1() (n int) {
	if x.Data == "" {
		return n
	}
	n += fastpb.SizeString(1, x.GetData())
	return n
}

var fieldIDToName_OrderQueryRequest = map[int32]string{
	1: "Message",
	2: "UserId",
}

var fieldIDToName_OrderQueryResponse = map[int32]string{
	1: "Data",
}

var fieldIDToName_AutoOrderRequest = map[int32]string{
	1: "Message",
	2: "UserId",
}

var fieldIDToName_AutoOrderResponse = map[int32]string{
	1: "Data",
}
