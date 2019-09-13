// Code generated by easyjson for marshaling/unmarshaling. DO NOT EDIT.

package probes

import (
	json "encoding/json"
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

func easyjsonB583f6f9DecodeGithubComSkydiveProjectSkydiveFlowProbes(in *jlexer.Lexer, out *Captures) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		in.Skip()
		*out = nil
	} else {
		in.Delim('[')
		if *out == nil {
			if !in.IsDelim(']') {
				*out = make(Captures, 0, 8)
			} else {
				*out = Captures{}
			}
		} else {
			*out = (*out)[:0]
		}
		for !in.IsDelim(']') {
			var v1 *CaptureMetadata
			if in.IsNull() {
				in.Skip()
				v1 = nil
			} else {
				if v1 == nil {
					v1 = new(CaptureMetadata)
				}
				(*v1).UnmarshalEasyJSON(in)
			}
			*out = append(*out, v1)
			in.WantComma()
		}
		in.Delim(']')
	}
	if isTopLevel {
		in.Consumed()
	}
}
func easyjsonB583f6f9EncodeGithubComSkydiveProjectSkydiveFlowProbes(out *jwriter.Writer, in Captures) {
	if in == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
		out.RawString("null")
	} else {
		out.RawByte('[')
		for v2, v3 := range in {
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

// MarshalJSON supports json.Marshaler interface
func (v Captures) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonB583f6f9EncodeGithubComSkydiveProjectSkydiveFlowProbes(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v Captures) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonB583f6f9EncodeGithubComSkydiveProjectSkydiveFlowProbes(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *Captures) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonB583f6f9DecodeGithubComSkydiveProjectSkydiveFlowProbes(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *Captures) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonB583f6f9DecodeGithubComSkydiveProjectSkydiveFlowProbes(l, v)
}
func easyjsonB583f6f9DecodeGithubComSkydiveProjectSkydiveFlowProbes1(in *jlexer.Lexer, out *CaptureStats) {
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
		case "PacketsReceived":
			out.PacketsReceived = int64(in.Int64())
		case "PacketsDropped":
			out.PacketsDropped = int64(in.Int64())
		case "PacketsIfDropped":
			out.PacketsIfDropped = int64(in.Int64())
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
func easyjsonB583f6f9EncodeGithubComSkydiveProjectSkydiveFlowProbes1(out *jwriter.Writer, in CaptureStats) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"PacketsReceived\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int64(int64(in.PacketsReceived))
	}
	{
		const prefix string = ",\"PacketsDropped\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int64(int64(in.PacketsDropped))
	}
	{
		const prefix string = ",\"PacketsIfDropped\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int64(int64(in.PacketsIfDropped))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v CaptureStats) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonB583f6f9EncodeGithubComSkydiveProjectSkydiveFlowProbes1(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v CaptureStats) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonB583f6f9EncodeGithubComSkydiveProjectSkydiveFlowProbes1(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *CaptureStats) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonB583f6f9DecodeGithubComSkydiveProjectSkydiveFlowProbes1(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *CaptureStats) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonB583f6f9DecodeGithubComSkydiveProjectSkydiveFlowProbes1(l, v)
}
func easyjsonB583f6f9DecodeGithubComSkydiveProjectSkydiveFlowProbes2(in *jlexer.Lexer, out *CaptureMetadata) {
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
		case "ID":
			out.ID = string(in.String())
		case "State":
			out.State = string(in.String())
		case "Name":
			out.Name = string(in.String())
		case "Description":
			out.Description = string(in.String())
		case "BPFFilter":
			out.BPFFilter = string(in.String())
		case "Type":
			out.Type = string(in.String())
		case "PCAPSocket":
			out.PCAPSocket = string(in.String())
		case "MirrorOf":
			out.MirrorOf = string(in.String())
		case "SFlowSocket":
			out.SFlowSocket = string(in.String())
		case "Error":
			out.Error = string(in.String())
		case "PacketsReceived":
			out.PacketsReceived = int64(in.Int64())
		case "PacketsDropped":
			out.PacketsDropped = int64(in.Int64())
		case "PacketsIfDropped":
			out.PacketsIfDropped = int64(in.Int64())
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
func easyjsonB583f6f9EncodeGithubComSkydiveProjectSkydiveFlowProbes2(out *jwriter.Writer, in CaptureMetadata) {
	out.RawByte('{')
	first := true
	_ = first
	if in.ID != "" {
		const prefix string = ",\"ID\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.ID))
	}
	if in.State != "" {
		const prefix string = ",\"State\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.State))
	}
	if in.Name != "" {
		const prefix string = ",\"Name\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.Name))
	}
	if in.Description != "" {
		const prefix string = ",\"Description\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.Description))
	}
	if in.BPFFilter != "" {
		const prefix string = ",\"BPFFilter\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.BPFFilter))
	}
	if in.Type != "" {
		const prefix string = ",\"Type\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.Type))
	}
	if in.PCAPSocket != "" {
		const prefix string = ",\"PCAPSocket\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.PCAPSocket))
	}
	if in.MirrorOf != "" {
		const prefix string = ",\"MirrorOf\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.MirrorOf))
	}
	if in.SFlowSocket != "" {
		const prefix string = ",\"SFlowSocket\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.SFlowSocket))
	}
	if in.Error != "" {
		const prefix string = ",\"Error\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.Error))
	}
	{
		const prefix string = ",\"PacketsReceived\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int64(int64(in.PacketsReceived))
	}
	{
		const prefix string = ",\"PacketsDropped\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int64(int64(in.PacketsDropped))
	}
	{
		const prefix string = ",\"PacketsIfDropped\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int64(int64(in.PacketsIfDropped))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v CaptureMetadata) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonB583f6f9EncodeGithubComSkydiveProjectSkydiveFlowProbes2(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v CaptureMetadata) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonB583f6f9EncodeGithubComSkydiveProjectSkydiveFlowProbes2(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *CaptureMetadata) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonB583f6f9DecodeGithubComSkydiveProjectSkydiveFlowProbes2(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *CaptureMetadata) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonB583f6f9DecodeGithubComSkydiveProjectSkydiveFlowProbes2(l, v)
}
