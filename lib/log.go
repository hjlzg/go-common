package lib

import (
	"fmt"
	log "github.com/hjlzg/go-common/log"
	"strings"
)

// 通用DLTag常量定义
const(
	DLTagHTTPSuccess	="_com_http_success"
	DLTagHTTPFailed		="_com_http_failed"
	DLTagMySqlFailed	="_com_mysql_failure"
	DLTagMySqlSuccess	="_com_mysql_success"
	DLTagRedisFailed	="_com_redis_failure"
	DLTagRedisSuccess	="_com_redis_success"
	DLTagRequestIn		="_com_request_in"
	DLTagRequestOut		="_com_request_out"
	DLTagThriftFailed	="_com_thrift_failed"
	DLTagThriftSuccess	="_com_thrift_success"
	DLTagTCPFailed		="_com_tcp_failed"
	DLTagUndefind		="_undef"
)

const(
	_dlTag			="dltag"
	_traceId		="traceid"
	_spanId			="spanid"
	_childSpanId	="cspanid"
	_dlTagBizPrefix	="_com_"
	_dlTagBizUndef	="_com_undef"
)

type Logger struct {
}

var Log	*Logger

type Trace struct{
	Caller		string
	HintCode	string
	HintContent	string
	SpanId		string
	SrcMethod	string
	TraceId		string
}

type TraceContext struct{
	Trace
	CSpanId string
}

func (l *Logger) TagInfo(trace *TraceContext,dltag string,m map[string]interface{}){
	m[_dlTag]=checkDLTag(dltag)
	m[_traceId]=trace.TraceId
	m[_childSpanId]=trace.CSpanId
	m[_spanId]=trace.SpanId
	log.Warn(parseParams(m))
}

func (l *Logger) TagWarn(trace *TraceContext,dltag string,m map[string]interface{}){
	m[_dlTag] = checkDLTag(dltag)
	m[_traceId] = trace.TraceId
	m[_childSpanId] = trace.CSpanId
	m[_spanId] = trace.SpanId
	log.Warn(parseParams(m))
}

func (l *Logger) TagError(trace *TraceContext, dltag string, m map[string]interface{}) {
	m[_dlTag] = checkDLTag(dltag)
	m[_traceId] = trace.TraceId
	m[_childSpanId] = trace.CSpanId
	m[_spanId] = trace.SpanId
	dlog.Error(parseParams(m))
}

func (l *Logger) TagTrace(trace *TraceContext, dltag string, m map[string]interface{}) {
	m[_dlTag] = checkDLTag(dltag)
	m[_traceId] = trace.TraceId
	m[_childSpanId] = trace.CSpanId
	m[_spanId] = trace.SpanId
	log.Trace(parseParams(m))
}

func (l *Logger) TagDebug(trace *TraceContext, dltag string, m map[string]interface{}) {
	m[_dlTag] = checkDLTag(dltag)
	m[_traceId] = trace.TraceId
	m[_childSpanId] = trace.CSpanId
	m[_spanId] = trace.SpanId
	log.Debug(parseParams(m))
}

func (l *Logger) Close(){
	log.Close()
}

func CreateBizDLTag(tagName string) string{
	if tagName==""{
		return _dlTagBizUndef
	}
	return _dlTagBizPrefix+tagName
}

//校验dltag合法性
func checkDLTag(dltag string) string{
	if strings.HasPrefix(dltag,_dlTagBizPrefix){
		return dltag
	}

	if strings.HasPrefix(dltag,"_com_"){
		return dltag
	}

	if dltag==DLTagUndefind{
		return dltag
	}
	return dltag
}

//map 格式化为string
func parseParams(m map[string]interface{}) string{
	var dltag string="_undef"
	if _dlTag,_have:=m["dltag"];_have{
		if _val,_ok:=_dlTag.(string);_ok{
			dltag=_val
		}
	}
	for _key, _val := range m {
		if _key == "dltag" {
			continue
		}
		dltag = dltag + "||" + fmt.Sprintf("%v=%+v", _key, _val)
	}
	dltag = strings.Trim(fmt.Sprintf("%q", dltag), "\"")
	return dltag
}
