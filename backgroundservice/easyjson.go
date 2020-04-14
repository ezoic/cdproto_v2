// Code generated by easyjson for marshaling/unmarshaling. DO NOT EDIT.

package backgroundservice

import (
	json "encoding/json"
	cdp "github.com/ezoic/cdproto_v2/cdp"
	serviceworker "github.com/ezoic/cdproto_v2/serviceworker"
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

func easyjsonC5a4559bDecodeGithubComChromedpCdprotoBackgroundservice(in *jlexer.Lexer, out *StopObservingParams) {
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
		case "service":
			(out.Service).UnmarshalEasyJSON(in)
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
func easyjsonC5a4559bEncodeGithubComChromedpCdprotoBackgroundservice(out *jwriter.Writer, in StopObservingParams) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"service\":"
		out.RawString(prefix[1:])
		(in.Service).MarshalEasyJSON(out)
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v StopObservingParams) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonC5a4559bEncodeGithubComChromedpCdprotoBackgroundservice(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v StopObservingParams) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonC5a4559bEncodeGithubComChromedpCdprotoBackgroundservice(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *StopObservingParams) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonC5a4559bDecodeGithubComChromedpCdprotoBackgroundservice(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *StopObservingParams) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonC5a4559bDecodeGithubComChromedpCdprotoBackgroundservice(l, v)
}
func easyjsonC5a4559bDecodeGithubComChromedpCdprotoBackgroundservice1(in *jlexer.Lexer, out *StartObservingParams) {
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
		case "service":
			(out.Service).UnmarshalEasyJSON(in)
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
func easyjsonC5a4559bEncodeGithubComChromedpCdprotoBackgroundservice1(out *jwriter.Writer, in StartObservingParams) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"service\":"
		out.RawString(prefix[1:])
		(in.Service).MarshalEasyJSON(out)
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v StartObservingParams) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonC5a4559bEncodeGithubComChromedpCdprotoBackgroundservice1(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v StartObservingParams) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonC5a4559bEncodeGithubComChromedpCdprotoBackgroundservice1(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *StartObservingParams) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonC5a4559bDecodeGithubComChromedpCdprotoBackgroundservice1(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *StartObservingParams) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonC5a4559bDecodeGithubComChromedpCdprotoBackgroundservice1(l, v)
}
func easyjsonC5a4559bDecodeGithubComChromedpCdprotoBackgroundservice2(in *jlexer.Lexer, out *SetRecordingParams) {
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
		case "shouldRecord":
			out.ShouldRecord = bool(in.Bool())
		case "service":
			(out.Service).UnmarshalEasyJSON(in)
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
func easyjsonC5a4559bEncodeGithubComChromedpCdprotoBackgroundservice2(out *jwriter.Writer, in SetRecordingParams) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"shouldRecord\":"
		out.RawString(prefix[1:])
		out.Bool(bool(in.ShouldRecord))
	}
	{
		const prefix string = ",\"service\":"
		out.RawString(prefix)
		(in.Service).MarshalEasyJSON(out)
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v SetRecordingParams) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonC5a4559bEncodeGithubComChromedpCdprotoBackgroundservice2(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v SetRecordingParams) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonC5a4559bEncodeGithubComChromedpCdprotoBackgroundservice2(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *SetRecordingParams) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonC5a4559bDecodeGithubComChromedpCdprotoBackgroundservice2(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *SetRecordingParams) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonC5a4559bDecodeGithubComChromedpCdprotoBackgroundservice2(l, v)
}
func easyjsonC5a4559bDecodeGithubComChromedpCdprotoBackgroundservice3(in *jlexer.Lexer, out *EventRecordingStateChanged) {
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
		case "isRecording":
			out.IsRecording = bool(in.Bool())
		case "service":
			(out.Service).UnmarshalEasyJSON(in)
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
func easyjsonC5a4559bEncodeGithubComChromedpCdprotoBackgroundservice3(out *jwriter.Writer, in EventRecordingStateChanged) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"isRecording\":"
		out.RawString(prefix[1:])
		out.Bool(bool(in.IsRecording))
	}
	{
		const prefix string = ",\"service\":"
		out.RawString(prefix)
		(in.Service).MarshalEasyJSON(out)
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v EventRecordingStateChanged) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonC5a4559bEncodeGithubComChromedpCdprotoBackgroundservice3(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v EventRecordingStateChanged) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonC5a4559bEncodeGithubComChromedpCdprotoBackgroundservice3(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *EventRecordingStateChanged) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonC5a4559bDecodeGithubComChromedpCdprotoBackgroundservice3(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *EventRecordingStateChanged) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonC5a4559bDecodeGithubComChromedpCdprotoBackgroundservice3(l, v)
}
func easyjsonC5a4559bDecodeGithubComChromedpCdprotoBackgroundservice4(in *jlexer.Lexer, out *EventMetadata) {
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
		case "key":
			out.Key = string(in.String())
		case "value":
			out.Value = string(in.String())
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
func easyjsonC5a4559bEncodeGithubComChromedpCdprotoBackgroundservice4(out *jwriter.Writer, in EventMetadata) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"key\":"
		out.RawString(prefix[1:])
		out.String(string(in.Key))
	}
	{
		const prefix string = ",\"value\":"
		out.RawString(prefix)
		out.String(string(in.Value))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v EventMetadata) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonC5a4559bEncodeGithubComChromedpCdprotoBackgroundservice4(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v EventMetadata) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonC5a4559bEncodeGithubComChromedpCdprotoBackgroundservice4(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *EventMetadata) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonC5a4559bDecodeGithubComChromedpCdprotoBackgroundservice4(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *EventMetadata) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonC5a4559bDecodeGithubComChromedpCdprotoBackgroundservice4(l, v)
}
func easyjsonC5a4559bDecodeGithubComChromedpCdprotoBackgroundservice5(in *jlexer.Lexer, out *EventBackgroundServiceEventReceived) {
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
		case "backgroundServiceEvent":
			if in.IsNull() {
				in.Skip()
				out.BackgroundServiceEvent = nil
			} else {
				if out.BackgroundServiceEvent == nil {
					out.BackgroundServiceEvent = new(Event)
				}
				(*out.BackgroundServiceEvent).UnmarshalEasyJSON(in)
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
func easyjsonC5a4559bEncodeGithubComChromedpCdprotoBackgroundservice5(out *jwriter.Writer, in EventBackgroundServiceEventReceived) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"backgroundServiceEvent\":"
		out.RawString(prefix[1:])
		if in.BackgroundServiceEvent == nil {
			out.RawString("null")
		} else {
			(*in.BackgroundServiceEvent).MarshalEasyJSON(out)
		}
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v EventBackgroundServiceEventReceived) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonC5a4559bEncodeGithubComChromedpCdprotoBackgroundservice5(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v EventBackgroundServiceEventReceived) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonC5a4559bEncodeGithubComChromedpCdprotoBackgroundservice5(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *EventBackgroundServiceEventReceived) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonC5a4559bDecodeGithubComChromedpCdprotoBackgroundservice5(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *EventBackgroundServiceEventReceived) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonC5a4559bDecodeGithubComChromedpCdprotoBackgroundservice5(l, v)
}
func easyjsonC5a4559bDecodeGithubComChromedpCdprotoBackgroundservice6(in *jlexer.Lexer, out *Event) {
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
		case "timestamp":
			if in.IsNull() {
				in.Skip()
				out.Timestamp = nil
			} else {
				if out.Timestamp == nil {
					out.Timestamp = new(cdp.TimeSinceEpoch)
				}
				(*out.Timestamp).UnmarshalEasyJSON(in)
			}
		case "origin":
			out.Origin = string(in.String())
		case "serviceWorkerRegistrationId":
			out.ServiceWorkerRegistrationID = serviceworker.RegistrationID(in.String())
		case "service":
			(out.Service).UnmarshalEasyJSON(in)
		case "eventName":
			out.EventName = string(in.String())
		case "instanceId":
			out.InstanceID = string(in.String())
		case "eventMetadata":
			if in.IsNull() {
				in.Skip()
				out.EventMetadata = nil
			} else {
				in.Delim('[')
				if out.EventMetadata == nil {
					if !in.IsDelim(']') {
						out.EventMetadata = make([]*EventMetadata, 0, 8)
					} else {
						out.EventMetadata = []*EventMetadata{}
					}
				} else {
					out.EventMetadata = (out.EventMetadata)[:0]
				}
				for !in.IsDelim(']') {
					var v1 *EventMetadata
					if in.IsNull() {
						in.Skip()
						v1 = nil
					} else {
						if v1 == nil {
							v1 = new(EventMetadata)
						}
						(*v1).UnmarshalEasyJSON(in)
					}
					out.EventMetadata = append(out.EventMetadata, v1)
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
func easyjsonC5a4559bEncodeGithubComChromedpCdprotoBackgroundservice6(out *jwriter.Writer, in Event) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"timestamp\":"
		out.RawString(prefix[1:])
		if in.Timestamp == nil {
			out.RawString("null")
		} else {
			(*in.Timestamp).MarshalEasyJSON(out)
		}
	}
	{
		const prefix string = ",\"origin\":"
		out.RawString(prefix)
		out.String(string(in.Origin))
	}
	{
		const prefix string = ",\"serviceWorkerRegistrationId\":"
		out.RawString(prefix)
		out.String(string(in.ServiceWorkerRegistrationID))
	}
	{
		const prefix string = ",\"service\":"
		out.RawString(prefix)
		(in.Service).MarshalEasyJSON(out)
	}
	{
		const prefix string = ",\"eventName\":"
		out.RawString(prefix)
		out.String(string(in.EventName))
	}
	{
		const prefix string = ",\"instanceId\":"
		out.RawString(prefix)
		out.String(string(in.InstanceID))
	}
	{
		const prefix string = ",\"eventMetadata\":"
		out.RawString(prefix)
		if in.EventMetadata == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
			out.RawString("null")
		} else {
			out.RawByte('[')
			for v2, v3 := range in.EventMetadata {
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
func (v Event) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonC5a4559bEncodeGithubComChromedpCdprotoBackgroundservice6(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v Event) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonC5a4559bEncodeGithubComChromedpCdprotoBackgroundservice6(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *Event) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonC5a4559bDecodeGithubComChromedpCdprotoBackgroundservice6(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *Event) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonC5a4559bDecodeGithubComChromedpCdprotoBackgroundservice6(l, v)
}
func easyjsonC5a4559bDecodeGithubComChromedpCdprotoBackgroundservice7(in *jlexer.Lexer, out *ClearEventsParams) {
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
		case "service":
			(out.Service).UnmarshalEasyJSON(in)
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
func easyjsonC5a4559bEncodeGithubComChromedpCdprotoBackgroundservice7(out *jwriter.Writer, in ClearEventsParams) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"service\":"
		out.RawString(prefix[1:])
		(in.Service).MarshalEasyJSON(out)
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v ClearEventsParams) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonC5a4559bEncodeGithubComChromedpCdprotoBackgroundservice7(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v ClearEventsParams) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonC5a4559bEncodeGithubComChromedpCdprotoBackgroundservice7(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *ClearEventsParams) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonC5a4559bDecodeGithubComChromedpCdprotoBackgroundservice7(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *ClearEventsParams) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonC5a4559bDecodeGithubComChromedpCdprotoBackgroundservice7(l, v)
}
