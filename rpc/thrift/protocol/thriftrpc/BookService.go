// Autogenerated by Thrift Compiler (0.10.0)
// DO NOT EDIT UNLESS YOU ARE SURE THAT YOU KNOW WHAT YOU ARE DOING

package thriftrpc

import (
	"bytes"
	"fmt"
	"git.apache.org/thrift.git/lib/go/thrift"
)

// (needed to ensure safety because of naive import list construction.)
var _ = thrift.ZERO
var _ = fmt.Printf
var _ = bytes.Equal

// Attributes:
//  - Num1
//  - Num2
//  - Comment
type Work struct {
	Num1    int32   `thrift:"num1,1" db:"num1" json:"num1"`
	Num2    int32   `thrift:"num2,2" db:"num2" json:"num2"`
	Comment *string `thrift:"comment,3" db:"comment" json:"comment,omitempty"`
}

func NewWork() *Work {
	return &Work{}
}

func (p *Work) GetNum1() int32 {
	return p.Num1
}

func (p *Work) GetNum2() int32 {
	return p.Num2
}

var Work_Comment_DEFAULT string

func (p *Work) GetComment() string {
	if !p.IsSetComment() {
		return Work_Comment_DEFAULT
	}
	return *p.Comment
}
func (p *Work) IsSetComment() bool {
	return p.Comment != nil
}

func (p *Work) Read(iprot thrift.TProtocol) error {
	if _, err := iprot.ReadStructBegin(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read error: ", p), err)
	}

	for {
		_, fieldTypeId, fieldId, err := iprot.ReadFieldBegin()
		if err != nil {
			return thrift.PrependError(fmt.Sprintf("%T field %d read error: ", p, fieldId), err)
		}
		if fieldTypeId == thrift.STOP {
			break
		}
		switch fieldId {
		case 1:
			if err := p.ReadField1(iprot); err != nil {
				return err
			}
		case 2:
			if err := p.ReadField2(iprot); err != nil {
				return err
			}
		case 3:
			if err := p.ReadField3(iprot); err != nil {
				return err
			}
		default:
			if err := iprot.Skip(fieldTypeId); err != nil {
				return err
			}
		}
		if err := iprot.ReadFieldEnd(); err != nil {
			return err
		}
	}
	if err := iprot.ReadStructEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
	}
	return nil
}

func (p *Work) ReadField1(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadI32(); err != nil {
		return thrift.PrependError("error reading field 1: ", err)
	} else {
		p.Num1 = v
	}
	return nil
}

func (p *Work) ReadField2(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadI32(); err != nil {
		return thrift.PrependError("error reading field 2: ", err)
	} else {
		p.Num2 = v
	}
	return nil
}

func (p *Work) ReadField3(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadString(); err != nil {
		return thrift.PrependError("error reading field 3: ", err)
	} else {
		p.Comment = &v
	}
	return nil
}

func (p *Work) Write(oprot thrift.TProtocol) error {
	if err := oprot.WriteStructBegin("Work"); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err)
	}
	if p != nil {
		if err := p.writeField1(oprot); err != nil {
			return err
		}
		if err := p.writeField2(oprot); err != nil {
			return err
		}
		if err := p.writeField3(oprot); err != nil {
			return err
		}
	}
	if err := oprot.WriteFieldStop(); err != nil {
		return thrift.PrependError("write field stop error: ", err)
	}
	if err := oprot.WriteStructEnd(); err != nil {
		return thrift.PrependError("write struct stop error: ", err)
	}
	return nil
}

func (p *Work) writeField1(oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin("num1", thrift.I32, 1); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 1:num1: ", p), err)
	}
	if err := oprot.WriteI32(int32(p.Num1)); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T.num1 (1) field write error: ", p), err)
	}
	if err := oprot.WriteFieldEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 1:num1: ", p), err)
	}
	return err
}

func (p *Work) writeField2(oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin("num2", thrift.I32, 2); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 2:num2: ", p), err)
	}
	if err := oprot.WriteI32(int32(p.Num2)); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T.num2 (2) field write error: ", p), err)
	}
	if err := oprot.WriteFieldEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 2:num2: ", p), err)
	}
	return err
}

func (p *Work) writeField3(oprot thrift.TProtocol) (err error) {
	if p.IsSetComment() {
		if err := oprot.WriteFieldBegin("comment", thrift.STRING, 3); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field begin error 3:comment: ", p), err)
		}
		if err := oprot.WriteString(string(*p.Comment)); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T.comment (3) field write error: ", p), err)
		}
		if err := oprot.WriteFieldEnd(); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field end error 3:comment: ", p), err)
		}
	}
	return err
}

func (p *Work) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("Work(%+v)", *p)
}

type BookService interface {
	// Parameters:
	//  - Name
	//  - Work
	ReadBook(name string, work *Work) (r string, err error)
}

type BookServiceClient struct {
	Transport       thrift.TTransport
	ProtocolFactory thrift.TProtocolFactory
	InputProtocol   thrift.TProtocol
	OutputProtocol  thrift.TProtocol
	SeqId           int32
}

func NewBookServiceClientFactory(t thrift.TTransport, f thrift.TProtocolFactory) *BookServiceClient {
	return &BookServiceClient{Transport: t,
		ProtocolFactory: f,
		InputProtocol:   f.GetProtocol(t),
		OutputProtocol:  f.GetProtocol(t),
		SeqId:           0,
	}
}

func NewBookServiceClientProtocol(t thrift.TTransport, iprot thrift.TProtocol, oprot thrift.TProtocol) *BookServiceClient {
	return &BookServiceClient{Transport: t,
		ProtocolFactory: nil,
		InputProtocol:   iprot,
		OutputProtocol:  oprot,
		SeqId:           0,
	}
}

// Parameters:
//  - Name
//  - Work
func (p *BookServiceClient) ReadBook(name string, work *Work) (r string, err error) {
	if err = p.sendReadBook(name, work); err != nil {
		return
	}
	return p.recvReadBook()
}

func (p *BookServiceClient) sendReadBook(name string, work *Work) (err error) {
	oprot := p.OutputProtocol
	if oprot == nil {
		oprot = p.ProtocolFactory.GetProtocol(p.Transport)
		p.OutputProtocol = oprot
	}
	p.SeqId++
	if err = oprot.WriteMessageBegin("readBook", thrift.CALL, p.SeqId); err != nil {
		return
	}
	args := BookServiceReadBookArgs{
		Name: name,
		Work: work,
	}
	if err = args.Write(oprot); err != nil {
		return
	}
	if err = oprot.WriteMessageEnd(); err != nil {
		return
	}
	return oprot.Flush()
}

func (p *BookServiceClient) recvReadBook() (value string, err error) {
	iprot := p.InputProtocol
	if iprot == nil {
		iprot = p.ProtocolFactory.GetProtocol(p.Transport)
		p.InputProtocol = iprot
	}
	method, mTypeId, seqId, err := iprot.ReadMessageBegin()
	if err != nil {
		return
	}
	if method != "readBook" {
		err = thrift.NewTApplicationException(thrift.WRONG_METHOD_NAME, "readBook failed: wrong method name")
		return
	}
	if p.SeqId != seqId {
		err = thrift.NewTApplicationException(thrift.BAD_SEQUENCE_ID, "readBook failed: out of sequence response")
		return
	}
	if mTypeId == thrift.EXCEPTION {
		error0 := thrift.NewTApplicationException(thrift.UNKNOWN_APPLICATION_EXCEPTION, "Unknown Exception")
		var error1 error
		error1, err = error0.Read(iprot)
		if err != nil {
			return
		}
		if err = iprot.ReadMessageEnd(); err != nil {
			return
		}
		err = error1
		return
	}
	if mTypeId != thrift.REPLY {
		err = thrift.NewTApplicationException(thrift.INVALID_MESSAGE_TYPE_EXCEPTION, "readBook failed: invalid message type")
		return
	}
	result := BookServiceReadBookResult{}
	if err = result.Read(iprot); err != nil {
		return
	}
	if err = iprot.ReadMessageEnd(); err != nil {
		return
	}
	value = result.GetSuccess()
	return
}

type BookServiceProcessor struct {
	processorMap map[string]thrift.TProcessorFunction
	handler      BookService
}

func (p *BookServiceProcessor) AddToProcessorMap(key string, processor thrift.TProcessorFunction) {
	p.processorMap[key] = processor
}

func (p *BookServiceProcessor) GetProcessorFunction(key string) (processor thrift.TProcessorFunction, ok bool) {
	processor, ok = p.processorMap[key]
	return processor, ok
}

func (p *BookServiceProcessor) ProcessorMap() map[string]thrift.TProcessorFunction {
	return p.processorMap
}

func NewBookServiceProcessor(handler BookService) *BookServiceProcessor {

	self2 := &BookServiceProcessor{handler: handler, processorMap: make(map[string]thrift.TProcessorFunction)}
	self2.processorMap["readBook"] = &bookServiceProcessorReadBook{handler: handler}
	return self2
}

func (p *BookServiceProcessor) Process(iprot, oprot thrift.TProtocol) (success bool, err thrift.TException) {
	name, _, seqId, err := iprot.ReadMessageBegin()
	if err != nil {
		return false, err
	}
	if processor, ok := p.GetProcessorFunction(name); ok {
		return processor.Process(seqId, iprot, oprot)
	}
	iprot.Skip(thrift.STRUCT)
	iprot.ReadMessageEnd()
	x3 := thrift.NewTApplicationException(thrift.UNKNOWN_METHOD, "Unknown function "+name)
	oprot.WriteMessageBegin(name, thrift.EXCEPTION, seqId)
	x3.Write(oprot)
	oprot.WriteMessageEnd()
	oprot.Flush()
	return false, x3

}

type bookServiceProcessorReadBook struct {
	handler BookService
}

func (p *bookServiceProcessorReadBook) Process(seqId int32, iprot, oprot thrift.TProtocol) (success bool, err thrift.TException) {
	args := BookServiceReadBookArgs{}
	if err = args.Read(iprot); err != nil {
		iprot.ReadMessageEnd()
		x := thrift.NewTApplicationException(thrift.PROTOCOL_ERROR, err.Error())
		oprot.WriteMessageBegin("readBook", thrift.EXCEPTION, seqId)
		x.Write(oprot)
		oprot.WriteMessageEnd()
		oprot.Flush()
		return false, err
	}

	iprot.ReadMessageEnd()
	result := BookServiceReadBookResult{}
	var retval string
	var err2 error
	if retval, err2 = p.handler.ReadBook(args.Name, args.Work); err2 != nil {
		x := thrift.NewTApplicationException(thrift.INTERNAL_ERROR, "Internal error processing readBook: "+err2.Error())
		oprot.WriteMessageBegin("readBook", thrift.EXCEPTION, seqId)
		x.Write(oprot)
		oprot.WriteMessageEnd()
		oprot.Flush()
		return true, err2
	} else {
		result.Success = &retval
	}
	if err2 = oprot.WriteMessageBegin("readBook", thrift.REPLY, seqId); err2 != nil {
		err = err2
	}
	if err2 = result.Write(oprot); err == nil && err2 != nil {
		err = err2
	}
	if err2 = oprot.WriteMessageEnd(); err == nil && err2 != nil {
		err = err2
	}
	if err2 = oprot.Flush(); err == nil && err2 != nil {
		err = err2
	}
	if err != nil {
		return
	}
	return true, err
}

// HELPER FUNCTIONS AND STRUCTURES

// Attributes:
//  - Name
//  - Work
type BookServiceReadBookArgs struct {
	Name string `thrift:"name,1" db:"name" json:"name"`
	Work *Work  `thrift:"work,2" db:"work" json:"work"`
}

func NewBookServiceReadBookArgs() *BookServiceReadBookArgs {
	return &BookServiceReadBookArgs{}
}

func (p *BookServiceReadBookArgs) GetName() string {
	return p.Name
}

var BookServiceReadBookArgs_Work_DEFAULT *Work

func (p *BookServiceReadBookArgs) GetWork() *Work {
	if !p.IsSetWork() {
		return BookServiceReadBookArgs_Work_DEFAULT
	}
	return p.Work
}
func (p *BookServiceReadBookArgs) IsSetWork() bool {
	return p.Work != nil
}

func (p *BookServiceReadBookArgs) Read(iprot thrift.TProtocol) error {
	if _, err := iprot.ReadStructBegin(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read error: ", p), err)
	}

	for {
		_, fieldTypeId, fieldId, err := iprot.ReadFieldBegin()
		if err != nil {
			return thrift.PrependError(fmt.Sprintf("%T field %d read error: ", p, fieldId), err)
		}
		if fieldTypeId == thrift.STOP {
			break
		}
		switch fieldId {
		case 1:
			if err := p.ReadField1(iprot); err != nil {
				return err
			}
		case 2:
			if err := p.ReadField2(iprot); err != nil {
				return err
			}
		default:
			if err := iprot.Skip(fieldTypeId); err != nil {
				return err
			}
		}
		if err := iprot.ReadFieldEnd(); err != nil {
			return err
		}
	}
	if err := iprot.ReadStructEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
	}
	return nil
}

func (p *BookServiceReadBookArgs) ReadField1(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadString(); err != nil {
		return thrift.PrependError("error reading field 1: ", err)
	} else {
		p.Name = v
	}
	return nil
}

func (p *BookServiceReadBookArgs) ReadField2(iprot thrift.TProtocol) error {
	p.Work = &Work{}
	if err := p.Work.Read(iprot); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T error reading struct: ", p.Work), err)
	}
	return nil
}

func (p *BookServiceReadBookArgs) Write(oprot thrift.TProtocol) error {
	if err := oprot.WriteStructBegin("readBook_args"); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err)
	}
	if p != nil {
		if err := p.writeField1(oprot); err != nil {
			return err
		}
		if err := p.writeField2(oprot); err != nil {
			return err
		}
	}
	if err := oprot.WriteFieldStop(); err != nil {
		return thrift.PrependError("write field stop error: ", err)
	}
	if err := oprot.WriteStructEnd(); err != nil {
		return thrift.PrependError("write struct stop error: ", err)
	}
	return nil
}

func (p *BookServiceReadBookArgs) writeField1(oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin("name", thrift.STRING, 1); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 1:name: ", p), err)
	}
	if err := oprot.WriteString(string(p.Name)); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T.name (1) field write error: ", p), err)
	}
	if err := oprot.WriteFieldEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 1:name: ", p), err)
	}
	return err
}

func (p *BookServiceReadBookArgs) writeField2(oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin("work", thrift.STRUCT, 2); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 2:work: ", p), err)
	}
	if err := p.Work.Write(oprot); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T error writing struct: ", p.Work), err)
	}
	if err := oprot.WriteFieldEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 2:work: ", p), err)
	}
	return err
}

func (p *BookServiceReadBookArgs) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("BookServiceReadBookArgs(%+v)", *p)
}

// Attributes:
//  - Success
type BookServiceReadBookResult struct {
	Success *string `thrift:"success,0" db:"success" json:"success,omitempty"`
}

func NewBookServiceReadBookResult() *BookServiceReadBookResult {
	return &BookServiceReadBookResult{}
}

var BookServiceReadBookResult_Success_DEFAULT string

func (p *BookServiceReadBookResult) GetSuccess() string {
	if !p.IsSetSuccess() {
		return BookServiceReadBookResult_Success_DEFAULT
	}
	return *p.Success
}
func (p *BookServiceReadBookResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *BookServiceReadBookResult) Read(iprot thrift.TProtocol) error {
	if _, err := iprot.ReadStructBegin(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read error: ", p), err)
	}

	for {
		_, fieldTypeId, fieldId, err := iprot.ReadFieldBegin()
		if err != nil {
			return thrift.PrependError(fmt.Sprintf("%T field %d read error: ", p, fieldId), err)
		}
		if fieldTypeId == thrift.STOP {
			break
		}
		switch fieldId {
		case 0:
			if err := p.ReadField0(iprot); err != nil {
				return err
			}
		default:
			if err := iprot.Skip(fieldTypeId); err != nil {
				return err
			}
		}
		if err := iprot.ReadFieldEnd(); err != nil {
			return err
		}
	}
	if err := iprot.ReadStructEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
	}
	return nil
}

func (p *BookServiceReadBookResult) ReadField0(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadString(); err != nil {
		return thrift.PrependError("error reading field 0: ", err)
	} else {
		p.Success = &v
	}
	return nil
}

func (p *BookServiceReadBookResult) Write(oprot thrift.TProtocol) error {
	if err := oprot.WriteStructBegin("readBook_result"); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err)
	}
	if p != nil {
		if err := p.writeField0(oprot); err != nil {
			return err
		}
	}
	if err := oprot.WriteFieldStop(); err != nil {
		return thrift.PrependError("write field stop error: ", err)
	}
	if err := oprot.WriteStructEnd(); err != nil {
		return thrift.PrependError("write struct stop error: ", err)
	}
	return nil
}

func (p *BookServiceReadBookResult) writeField0(oprot thrift.TProtocol) (err error) {
	if p.IsSetSuccess() {
		if err := oprot.WriteFieldBegin("success", thrift.STRING, 0); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field begin error 0:success: ", p), err)
		}
		if err := oprot.WriteString(string(*p.Success)); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T.success (0) field write error: ", p), err)
		}
		if err := oprot.WriteFieldEnd(); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field end error 0:success: ", p), err)
		}
	}
	return err
}

func (p *BookServiceReadBookResult) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("BookServiceReadBookResult(%+v)", *p)
}