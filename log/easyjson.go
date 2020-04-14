// Code generated by easyjson for marshaling/unmarshaling. DO NOT EDIT.

package log

import (
	json "encoding/json"
	network "github.com/ezoic/cdproto_v2/network"
	runtime "github.com/ezoic/cdproto_v2/runtime"
	easyjson "github.com/mailru/easyjson"
	jlexer "github.com/mailru/easyjson/jlexer"
	jwriter "github.com/mailru/easyjson/jwriter"
)

// suppress unused package warning
var (
	_ *json.RawMessage
	_ *jlexer.Lexer
	_ *jwriter.Writer
	_ easyjson.Marshaler
)

func easyjsonC5a4559bDecodeGithubComChromedpCdprotoLog(in *jlexer.Lexer, out *ViolationSetting) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeString()
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "name":
			(out.Name).UnmarshalEasyJSON(in)
		case "threshold":
			out.Threshold = float64(in.Float64())
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjsonC5a4559bEncodeGithubComChromedpCdprotoLog(out *jwriter.Writer, in ViolationSetting) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"name\":"
		out.RawString(prefix[1:])
		(in.Name).MarshalEasyJSON(out)
	}
	{
		const prefix string = ",\"threshold\":"
		out.RawString(prefix)
		out.Float64(float64(in.Threshold))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v ViolationSetting) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonC5a4559bEncodeGithubComChromedpCdprotoLog(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v ViolationSetting) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonC5a4559bEncodeGithubComChromedpCdprotoLog(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *ViolationSetting) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonC5a4559bDecodeGithubComChromedpCdprotoLog(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *ViolationSetting) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonC5a4559bDecodeGithubComChromedpCdprotoLog(l, v)
}
func easyjsonC5a4559bDecodeGithubComChromedpCdprotoLog1(in *jlexer.Lexer, out *StopViolationsReportParams) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeString()
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjsonC5a4559bEncodeGithubComChromedpCdprotoLog1(out *jwriter.Writer, in StopViolationsReportParams) {
	out.RawByte('{')
	first := true
	_ = first
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v StopViolationsReportParams) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonC5a4559bEncodeGithubComChromedpCdprotoLog1(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v StopViolationsReportParams) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonC5a4559bEncodeGithubComChromedpCdprotoLog1(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *StopViolationsReportParams) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonC5a4559bDecodeGithubComChromedpCdprotoLog1(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *StopViolationsReportParams) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonC5a4559bDecodeGithubComChromedpCdprotoLog1(l, v)
}
func easyjsonC5a4559bDecodeGithubComChromedpCdprotoLog2(in *jlexer.Lexer, out *StartViolationsReportParams) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeString()
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "config":
			if in.IsNull() {
				in.Skip()
				out.Config = nil
			} else {
				in.Delim('[')
				if out.Config == nil {
					if !in.IsDelim(']') {
						out.Config = make([]*ViolationSetting, 0, 8)
					} else {
						out.Config = []*ViolationSetting{}
					}
				} else {
					out.Config = (out.Config)[:0]
				}
				for !in.IsDelim(']') {
					var v1 *ViolationSetting
					if in.IsNull() {
						in.Skip()
						v1 = nil
					} else {
						if v1 == nil {
							v1 = new(ViolationSetting)
						}
						(*v1).UnmarshalEasyJSON(in)
					}
					out.Config = append(out.Config, v1)
					in.WantComma()
				}
				in.Delim(']')
			}
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjsonC5a4559bEncodeGithubComChromedpCdprotoLog2(out *jwriter.Writer, in StartViolationsReportParams) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"config\":"
		out.RawString(prefix[1:])
		if in.Config == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
			out.RawString("null")
		} else {
			out.RawByte('[')
			for v2, v3 := range in.Config {
				if v2 > 0 {
					out.RawByte(',')
				}
				if v3 == nil {
					out.RawString("null")
				} else {
					(*v3).MarshalEasyJSON(out)
				}
			}
			out.RawByte(']')
		}
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v StartViolationsReportParams) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonC5a4559bEncodeGithubComChromedpCdprotoLog2(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v StartViolationsReportParams) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonC5a4559bEncodeGithubComChromedpCdprotoLog2(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *StartViolationsReportParams) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonC5a4559bDecodeGithubComChromedpCdprotoLog2(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *StartViolationsReportParams) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonC5a4559bDecodeGithubComChromedpCdprotoLog2(l, v)
}
func easyjsonC5a4559bDecodeGithubComChromedpCdprotoLog3(in *jlexer.Lexer, out *EventEntryAdded) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeString()
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "entry":
			if in.IsNull() {
				in.Skip()
				out.Entry = nil
			} else {
				if out.Entry == nil {
					out.Entry = new(Entry)
				}
				(*out.Entry).UnmarshalEasyJSON(in)
			}
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjsonC5a4559bEncodeGithubComChromedpCdprotoLog3(out *jwriter.Writer, in EventEntryAdded) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"entry\":"
		out.RawString(prefix[1:])
		if in.Entry == nil {
			out.RawString("null")
		} else {
			(*in.Entry).MarshalEasyJSON(out)
		}
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v EventEntryAdded) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonC5a4559bEncodeGithubComChromedpCdprotoLog3(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v EventEntryAdded) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonC5a4559bEncodeGithubComChromedpCdprotoLog3(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *EventEntryAdded) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonC5a4559bDecodeGithubComChromedpCdprotoLog3(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *EventEntryAdded) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonC5a4559bDecodeGithubComChromedpCdprotoLog3(l, v)
}
func easyjsonC5a4559bDecodeGithubComChromedpCdprotoLog4(in *jlexer.Lexer, out *Entry) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeString()
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "source":
			(out.Source).UnmarshalEasyJSON(in)
		case "level":
			(out.Level).UnmarshalEasyJSON(in)
		case "text":
			out.Text = string(in.String())
		case "timestamp":
			if in.IsNull() {
				in.Skip()
				out.Timestamp = nil
			} else {
				if out.Timestamp == nil {
					out.Timestamp = new(runtime.Timestamp)
				}
				(*out.Timestamp).UnmarshalEasyJSON(in)
			}
		case "url":
			out.URL = string(in.String())
		case "lineNumber":
			out.LineNumber = int64(in.Int64())
		case "stackTrace":
			if in.IsNull() {
				in.Skip()
				out.StackTrace = nil
			} else {
				if out.StackTrace == nil {
					out.StackTrace = new(runtime.StackTrace)
				}
				(*out.StackTrace).UnmarshalEasyJSON(in)
			}
		case "networkRequestId":
			out.NetworkRequestID = network.RequestID(in.String())
		case "workerId":
			out.WorkerID = string(in.String())
		case "args":
			if in.IsNull() {
				in.Skip()
				out.Args = nil
			} else {
				in.Delim('[')
				if out.Args == nil {
					if !in.IsDelim(']') {
						out.Args = make([]*runtime.RemoteObject, 0, 8)
					} else {
						out.Args = []*runtime.RemoteObject{}
					}
				} else {
					out.Args = (out.Args)[:0]
				}
				for !in.IsDelim(']') {
					var v4 *runtime.RemoteObject
					if in.IsNull() {
						in.Skip()
						v4 = nil
					} else {
						if v4 == nil {
							v4 = new(runtime.RemoteObject)
						}
						(*v4).UnmarshalEasyJSON(in)
					}
					out.Args = append(out.Args, v4)
					in.WantComma()
				}
				in.Delim(']')
			}
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjsonC5a4559bEncodeGithubComChromedpCdprotoLog4(out *jwriter.Writer, in Entry) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"source\":"
		out.RawString(prefix[1:])
		(in.Source).MarshalEasyJSON(out)
	}
	{
		const prefix string = ",\"level\":"
		out.RawString(prefix)
		(in.Level).MarshalEasyJSON(out)
	}
	{
		const prefix string = ",\"text\":"
		out.RawString(prefix)
		out.String(string(in.Text))
	}
	{
		const prefix string = ",\"timestamp\":"
		out.RawString(prefix)
		if in.Timestamp == nil {
			out.RawString("null")
		} else {
			(*in.Timestamp).MarshalEasyJSON(out)
		}
	}
	if in.URL != "" {
		const prefix string = ",\"url\":"
		out.RawString(prefix)
		out.String(string(in.URL))
	}
	if in.LineNumber != 0 {
		const prefix string = ",\"lineNumber\":"
		out.RawString(prefix)
		out.Int64(int64(in.LineNumber))
	}
	if in.StackTrace != nil {
		const prefix string = ",\"stackTrace\":"
		out.RawString(prefix)
		(*in.StackTrace).MarshalEasyJSON(out)
	}
	if in.NetworkRequestID != "" {
		const prefix string = ",\"networkRequestId\":"
		out.RawString(prefix)
		out.String(string(in.NetworkRequestID))
	}
	if in.WorkerID != "" {
		const prefix string = ",\"workerId\":"
		out.RawString(prefix)
		out.String(string(in.WorkerID))
	}
	if len(in.Args) != 0 {
		const prefix string = ",\"args\":"
		out.RawString(prefix)
		{
			out.RawByte('[')
			for v5, v6 := range in.Args {
				if v5 > 0 {
					out.RawByte(',')
				}
				if v6 == nil {
					out.RawString("null")
				} else {
					(*v6).MarshalEasyJSON(out)
				}
			}
			out.RawByte(']')
		}
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v Entry) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonC5a4559bEncodeGithubComChromedpCdprotoLog4(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v Entry) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonC5a4559bEncodeGithubComChromedpCdprotoLog4(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *Entry) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonC5a4559bDecodeGithubComChromedpCdprotoLog4(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *Entry) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonC5a4559bDecodeGithubComChromedpCdprotoLog4(l, v)
}
func easyjsonC5a4559bDecodeGithubComChromedpCdprotoLog5(in *jlexer.Lexer, out *EnableParams) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeString()
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjsonC5a4559bEncodeGithubComChromedpCdprotoLog5(out *jwriter.Writer, in EnableParams) {
	out.RawByte('{')
	first := true
	_ = first
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v EnableParams) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonC5a4559bEncodeGithubComChromedpCdprotoLog5(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v EnableParams) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonC5a4559bEncodeGithubComChromedpCdprotoLog5(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *EnableParams) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonC5a4559bDecodeGithubComChromedpCdprotoLog5(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *EnableParams) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonC5a4559bDecodeGithubComChromedpCdprotoLog5(l, v)
}
func easyjsonC5a4559bDecodeGithubComChromedpCdprotoLog6(in *jlexer.Lexer, out *DisableParams) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeString()
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjsonC5a4559bEncodeGithubComChromedpCdprotoLog6(out *jwriter.Writer, in DisableParams) {
	out.RawByte('{')
	first := true
	_ = first
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v DisableParams) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonC5a4559bEncodeGithubComChromedpCdprotoLog6(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v DisableParams) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonC5a4559bEncodeGithubComChromedpCdprotoLog6(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *DisableParams) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonC5a4559bDecodeGithubComChromedpCdprotoLog6(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *DisableParams) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonC5a4559bDecodeGithubComChromedpCdprotoLog6(l, v)
}
func easyjsonC5a4559bDecodeGithubComChromedpCdprotoLog7(in *jlexer.Lexer, out *ClearParams) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeString()
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjsonC5a4559bEncodeGithubComChromedpCdprotoLog7(out *jwriter.Writer, in ClearParams) {
	out.RawByte('{')
	first := true
	_ = first
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v ClearParams) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonC5a4559bEncodeGithubComChromedpCdprotoLog7(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v ClearParams) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonC5a4559bEncodeGithubComChromedpCdprotoLog7(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *ClearParams) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonC5a4559bDecodeGithubComChromedpCdprotoLog7(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *ClearParams) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonC5a4559bDecodeGithubComChromedpCdprotoLog7(l, v)
}
