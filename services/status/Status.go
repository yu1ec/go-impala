// Code generated by Thrift Compiler (0.17.0). DO NOT EDIT.

package status

import (
	"bytes"
	"context"
	"database/sql/driver"
	"errors"
	"fmt"
	"time"
	thrift "github.com/apache/thrift/lib/go/thrift"
)

// (needed to ensure safety because of naive import list construction.)
var _ = thrift.ZERO
var _ = fmt.Printf
var _ = errors.New
var _ = context.Background
var _ = time.Now
var _ = bytes.Equal

type TStatusCode int64
const (
  TStatusCode_OK TStatusCode = 0
  TStatusCode_CANCELLED TStatusCode = 1
  TStatusCode_ANALYSIS_ERROR TStatusCode = 2
  TStatusCode_NOT_IMPLEMENTED_ERROR TStatusCode = 3
  TStatusCode_RUNTIME_ERROR TStatusCode = 4
  TStatusCode_MEM_LIMIT_EXCEEDED TStatusCode = 5
  TStatusCode_INTERNAL_ERROR TStatusCode = 6
)

func (p TStatusCode) String() string {
  switch p {
  case TStatusCode_OK: return "OK"
  case TStatusCode_CANCELLED: return "CANCELLED"
  case TStatusCode_ANALYSIS_ERROR: return "ANALYSIS_ERROR"
  case TStatusCode_NOT_IMPLEMENTED_ERROR: return "NOT_IMPLEMENTED_ERROR"
  case TStatusCode_RUNTIME_ERROR: return "RUNTIME_ERROR"
  case TStatusCode_MEM_LIMIT_EXCEEDED: return "MEM_LIMIT_EXCEEDED"
  case TStatusCode_INTERNAL_ERROR: return "INTERNAL_ERROR"
  }
  return "<UNSET>"
}

func TStatusCodeFromString(s string) (TStatusCode, error) {
  switch s {
  case "OK": return TStatusCode_OK, nil 
  case "CANCELLED": return TStatusCode_CANCELLED, nil 
  case "ANALYSIS_ERROR": return TStatusCode_ANALYSIS_ERROR, nil 
  case "NOT_IMPLEMENTED_ERROR": return TStatusCode_NOT_IMPLEMENTED_ERROR, nil 
  case "RUNTIME_ERROR": return TStatusCode_RUNTIME_ERROR, nil 
  case "MEM_LIMIT_EXCEEDED": return TStatusCode_MEM_LIMIT_EXCEEDED, nil 
  case "INTERNAL_ERROR": return TStatusCode_INTERNAL_ERROR, nil 
  }
  return TStatusCode(0), fmt.Errorf("not a valid TStatusCode string")
}


func TStatusCodePtr(v TStatusCode) *TStatusCode { return &v }

func (p TStatusCode) MarshalText() ([]byte, error) {
return []byte(p.String()), nil
}

func (p *TStatusCode) UnmarshalText(text []byte) error {
q, err := TStatusCodeFromString(string(text))
if (err != nil) {
return err
}
*p = q
return nil
}

func (p *TStatusCode) Scan(value interface{}) error {
v, ok := value.(int64)
if !ok {
return errors.New("Scan value is not int64")
}
*p = TStatusCode(v)
return nil
}

func (p * TStatusCode) Value() (driver.Value, error) {
  if p == nil {
    return nil, nil
  }
return int64(*p), nil
}
// Attributes:
//  - StatusCode
//  - ErrorMsgs
type TStatus struct {
  StatusCode TStatusCode `thrift:"status_code,1,required" db:"status_code" json:"status_code"`
  ErrorMsgs []string `thrift:"error_msgs,2" db:"error_msgs" json:"error_msgs"`
}

func NewTStatus() *TStatus {
  return &TStatus{}
}


func (p *TStatus) GetStatusCode() TStatusCode {
  return p.StatusCode
}

func (p *TStatus) GetErrorMsgs() []string {
  return p.ErrorMsgs
}
func (p *TStatus) Read(ctx context.Context, iprot thrift.TProtocol) error {
  if _, err := iprot.ReadStructBegin(ctx); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T read error: ", p), err)
  }

  var issetStatusCode bool = false;

  for {
    _, fieldTypeId, fieldId, err := iprot.ReadFieldBegin(ctx)
    if err != nil {
      return thrift.PrependError(fmt.Sprintf("%T field %d read error: ", p, fieldId), err)
    }
    if fieldTypeId == thrift.STOP { break; }
    switch fieldId {
    case 1:
      if fieldTypeId == thrift.I32 {
        if err := p.ReadField1(ctx, iprot); err != nil {
          return err
        }
        issetStatusCode = true
      } else {
        if err := iprot.Skip(ctx, fieldTypeId); err != nil {
          return err
        }
      }
    case 2:
      if fieldTypeId == thrift.LIST {
        if err := p.ReadField2(ctx, iprot); err != nil {
          return err
        }
      } else {
        if err := iprot.Skip(ctx, fieldTypeId); err != nil {
          return err
        }
      }
    default:
      if err := iprot.Skip(ctx, fieldTypeId); err != nil {
        return err
      }
    }
    if err := iprot.ReadFieldEnd(ctx); err != nil {
      return err
    }
  }
  if err := iprot.ReadStructEnd(ctx); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
  }
  if !issetStatusCode{
    return thrift.NewTProtocolExceptionWithType(thrift.INVALID_DATA, fmt.Errorf("Required field StatusCode is not set"));
  }
  return nil
}

func (p *TStatus)  ReadField1(ctx context.Context, iprot thrift.TProtocol) error {
  if v, err := iprot.ReadI32(ctx); err != nil {
  return thrift.PrependError("error reading field 1: ", err)
} else {
  temp := TStatusCode(v)
  p.StatusCode = temp
}
  return nil
}

func (p *TStatus)  ReadField2(ctx context.Context, iprot thrift.TProtocol) error {
  _, size, err := iprot.ReadListBegin(ctx)
  if err != nil {
    return thrift.PrependError("error reading list begin: ", err)
  }
  tSlice := make([]string, 0, size)
  p.ErrorMsgs =  tSlice
  for i := 0; i < size; i ++ {
var _elem0 string
    if v, err := iprot.ReadString(ctx); err != nil {
    return thrift.PrependError("error reading field 0: ", err)
} else {
    _elem0 = v
}
    p.ErrorMsgs = append(p.ErrorMsgs, _elem0)
  }
  if err := iprot.ReadListEnd(ctx); err != nil {
    return thrift.PrependError("error reading list end: ", err)
  }
  return nil
}

func (p *TStatus) Write(ctx context.Context, oprot thrift.TProtocol) error {
  if err := oprot.WriteStructBegin(ctx, "TStatus"); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err) }
  if p != nil {
    if err := p.writeField1(ctx, oprot); err != nil { return err }
    if err := p.writeField2(ctx, oprot); err != nil { return err }
  }
  if err := oprot.WriteFieldStop(ctx); err != nil {
    return thrift.PrependError("write field stop error: ", err) }
  if err := oprot.WriteStructEnd(ctx); err != nil {
    return thrift.PrependError("write struct stop error: ", err) }
  return nil
}

func (p *TStatus) writeField1(ctx context.Context, oprot thrift.TProtocol) (err error) {
  if err := oprot.WriteFieldBegin(ctx, "status_code", thrift.I32, 1); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write field begin error 1:status_code: ", p), err) }
  if err := oprot.WriteI32(ctx, int32(p.StatusCode)); err != nil {
  return thrift.PrependError(fmt.Sprintf("%T.status_code (1) field write error: ", p), err) }
  if err := oprot.WriteFieldEnd(ctx); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write field end error 1:status_code: ", p), err) }
  return err
}

func (p *TStatus) writeField2(ctx context.Context, oprot thrift.TProtocol) (err error) {
  if err := oprot.WriteFieldBegin(ctx, "error_msgs", thrift.LIST, 2); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write field begin error 2:error_msgs: ", p), err) }
  if err := oprot.WriteListBegin(ctx, thrift.STRING, len(p.ErrorMsgs)); err != nil {
    return thrift.PrependError("error writing list begin: ", err)
  }
  for _, v := range p.ErrorMsgs {
    if err := oprot.WriteString(ctx, string(v)); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T. (0) field write error: ", p), err) }
  }
  if err := oprot.WriteListEnd(ctx); err != nil {
    return thrift.PrependError("error writing list end: ", err)
  }
  if err := oprot.WriteFieldEnd(ctx); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write field end error 2:error_msgs: ", p), err) }
  return err
}

func (p *TStatus) Equals(other *TStatus) bool {
  if p == other {
    return true
  } else if p == nil || other == nil {
    return false
  }
  if p.StatusCode != other.StatusCode { return false }
  if len(p.ErrorMsgs) != len(other.ErrorMsgs) { return false }
  for i, _tgt := range p.ErrorMsgs {
    _src1 := other.ErrorMsgs[i]
    if _tgt != _src1 { return false }
  }
  return true
}

func (p *TStatus) String() string {
  if p == nil {
    return "<nil>"
  }
  return fmt.Sprintf("TStatus(%+v)", *p)
}

