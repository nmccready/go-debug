package debug

import (
	"fmt"
	"runtime"
	"strconv"

	"github.com/fatih/color"
	"github.com/nmccready/colorjson"
)

// JSONFormatter formats logs into parsable json
type JSONFormatter struct {
	// TimestampFormat sets the format used for marshaling timestamps.
	TimestampFormat string

	// DisableTimestamp allows disabling automatic timestamps in output
	// DisableTimestamp bool

	// HTMLEscape allows disabling html escaping in output
	HTMLEscape bool

	Indent int
	// FieldMap allows users to customize the names of keys for default fields.
	// As an example:
	// formatter := &JSONFormatter{
	//   	FieldMap: FieldMap{
	// 		 FieldKeyTime:  "@timestamp",
	// 		 FieldKeyLevel: "@level",
	// 		 FieldKeyMsg:   "@message",
	// 		 FieldKeyFunc:  "@caller",
	//    },
	// }
	// FieldMap FieldMap

	// CallerPrettyfier can be set by the user to modify the content
	// of the function and file keys in the json data when ReportCaller is
	// activated. If any of the returned value is the empty string the
	// corresponding key will be removed from json fields.
	CallerPrettyfier func(*runtime.Frame) (function string, file string)

	// PrettyPrint will indent all json logs
	PrettyPrint bool

	FlattenMsgFields bool
}

func (f *JSONFormatter) Format(dbg Debugger, _msg interface{}) string {
	var msg string
	var msgFields *Fields

	switch v := _msg.(type) {
	case Fields:
		msgFields = &v
	case colorjson.Object:
		o := Fields(v)
		msgFields = &o
	case map[string]interface{}:
		o := Fields(v)
		msgFields = &o
	case string:
		msg = v
	default:
		msg = "JSONFormatter Invalid Msg Type must be a string or map[string]interface{}"
	}
	/*
		 Colors disabled due to:
		 {
				"delta": "252ms",
				"msg": "oops",
				"namespace": "\u001b[31merror:example:multiple:b\u001b[0m",
				"time": "20:53:24.798"
			}
	*/
	finalized := finalizeFields(dbg, msg, false, f.FlattenMsgFields, nil)

	if f.PrettyPrint && f.Indent == 0 {
		f.Indent = 2
	}

	if msg != "" && msgFields == nil {
		msgFields = &Fields{"msg": msg}
	}

	if msgFields != nil && f.FlattenMsgFields {
		fields := *msgFields

		for k, v := range fields {
			finalized.Fields[k] = v
		}
	}
	_f := colorjson.NewFormatter()
	_f.DisabledColor = !HAS_COLORS
	_f.HTMLEscape = f.HTMLEscape
	_f.Indent = f.Indent

	intColor, _ := strconv.Atoi(dbg.color)
	_f.KeyMapColors["namespace"] = color.New(color.Attribute(intColor))
	b, err := _f.Marshal(map[string]interface{}(finalized.Fields))
	if err != nil {
		return fmt.Sprintf("failed to marshal fields to JSON, %v", err)
	}
	return string(b) + "\n"
	// return b.String()
}

func (f *JSONFormatter) GetHasFieldsOnly() bool {
	return true
}
