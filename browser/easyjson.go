// Code generated by easyjson for marshaling/unmarshaling. DO NOT EDIT.

package browser

import (
	json "encoding/json"
	target "github.com/chromedp/cdproto/target"
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

func easyjsonC5a4559bDecodeGithubComChromedpCdprotoBrowser(in *jlexer.Lexer, out *SetWindowBoundsParams) {
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
		case "windowId":
			out.WindowID = WindowID(in.Int64())
		case "bounds":
			if in.IsNull() {
				in.Skip()
				out.Bounds = nil
			} else {
				if out.Bounds == nil {
					out.Bounds = new(Bounds)
				}
				(*out.Bounds).UnmarshalEasyJSON(in)
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
func easyjsonC5a4559bEncodeGithubComChromedpCdprotoBrowser(out *jwriter.Writer, in SetWindowBoundsParams) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"windowId\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int64(int64(in.WindowID))
	}
	{
		const prefix string = ",\"bounds\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		if in.Bounds == nil {
			out.RawString("null")
		} else {
			(*in.Bounds).MarshalEasyJSON(out)
		}
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v SetWindowBoundsParams) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonC5a4559bEncodeGithubComChromedpCdprotoBrowser(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v SetWindowBoundsParams) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonC5a4559bEncodeGithubComChromedpCdprotoBrowser(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *SetWindowBoundsParams) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonC5a4559bDecodeGithubComChromedpCdprotoBrowser(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *SetWindowBoundsParams) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonC5a4559bDecodeGithubComChromedpCdprotoBrowser(l, v)
}
func easyjsonC5a4559bDecodeGithubComChromedpCdprotoBrowser1(in *jlexer.Lexer, out *GetWindowForTargetReturns) {
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
		case "windowId":
			out.WindowID = WindowID(in.Int64())
		case "bounds":
			if in.IsNull() {
				in.Skip()
				out.Bounds = nil
			} else {
				if out.Bounds == nil {
					out.Bounds = new(Bounds)
				}
				(*out.Bounds).UnmarshalEasyJSON(in)
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
func easyjsonC5a4559bEncodeGithubComChromedpCdprotoBrowser1(out *jwriter.Writer, in GetWindowForTargetReturns) {
	out.RawByte('{')
	first := true
	_ = first
	if in.WindowID != 0 {
		const prefix string = ",\"windowId\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int64(int64(in.WindowID))
	}
	if in.Bounds != nil {
		const prefix string = ",\"bounds\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		(*in.Bounds).MarshalEasyJSON(out)
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v GetWindowForTargetReturns) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonC5a4559bEncodeGithubComChromedpCdprotoBrowser1(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v GetWindowForTargetReturns) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonC5a4559bEncodeGithubComChromedpCdprotoBrowser1(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *GetWindowForTargetReturns) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonC5a4559bDecodeGithubComChromedpCdprotoBrowser1(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *GetWindowForTargetReturns) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonC5a4559bDecodeGithubComChromedpCdprotoBrowser1(l, v)
}
func easyjsonC5a4559bDecodeGithubComChromedpCdprotoBrowser2(in *jlexer.Lexer, out *GetWindowForTargetParams) {
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
		case "targetId":
			out.TargetID = target.ID(in.String())
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
func easyjsonC5a4559bEncodeGithubComChromedpCdprotoBrowser2(out *jwriter.Writer, in GetWindowForTargetParams) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"targetId\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.TargetID))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v GetWindowForTargetParams) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonC5a4559bEncodeGithubComChromedpCdprotoBrowser2(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v GetWindowForTargetParams) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonC5a4559bEncodeGithubComChromedpCdprotoBrowser2(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *GetWindowForTargetParams) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonC5a4559bDecodeGithubComChromedpCdprotoBrowser2(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *GetWindowForTargetParams) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonC5a4559bDecodeGithubComChromedpCdprotoBrowser2(l, v)
}
func easyjsonC5a4559bDecodeGithubComChromedpCdprotoBrowser3(in *jlexer.Lexer, out *GetWindowBoundsReturns) {
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
		case "bounds":
			if in.IsNull() {
				in.Skip()
				out.Bounds = nil
			} else {
				if out.Bounds == nil {
					out.Bounds = new(Bounds)
				}
				(*out.Bounds).UnmarshalEasyJSON(in)
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
func easyjsonC5a4559bEncodeGithubComChromedpCdprotoBrowser3(out *jwriter.Writer, in GetWindowBoundsReturns) {
	out.RawByte('{')
	first := true
	_ = first
	if in.Bounds != nil {
		const prefix string = ",\"bounds\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		(*in.Bounds).MarshalEasyJSON(out)
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v GetWindowBoundsReturns) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonC5a4559bEncodeGithubComChromedpCdprotoBrowser3(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v GetWindowBoundsReturns) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonC5a4559bEncodeGithubComChromedpCdprotoBrowser3(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *GetWindowBoundsReturns) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonC5a4559bDecodeGithubComChromedpCdprotoBrowser3(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *GetWindowBoundsReturns) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonC5a4559bDecodeGithubComChromedpCdprotoBrowser3(l, v)
}
func easyjsonC5a4559bDecodeGithubComChromedpCdprotoBrowser4(in *jlexer.Lexer, out *GetWindowBoundsParams) {
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
		case "windowId":
			out.WindowID = WindowID(in.Int64())
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
func easyjsonC5a4559bEncodeGithubComChromedpCdprotoBrowser4(out *jwriter.Writer, in GetWindowBoundsParams) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"windowId\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int64(int64(in.WindowID))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v GetWindowBoundsParams) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonC5a4559bEncodeGithubComChromedpCdprotoBrowser4(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v GetWindowBoundsParams) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonC5a4559bEncodeGithubComChromedpCdprotoBrowser4(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *GetWindowBoundsParams) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonC5a4559bDecodeGithubComChromedpCdprotoBrowser4(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *GetWindowBoundsParams) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonC5a4559bDecodeGithubComChromedpCdprotoBrowser4(l, v)
}
func easyjsonC5a4559bDecodeGithubComChromedpCdprotoBrowser5(in *jlexer.Lexer, out *GetVersionReturns) {
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
		case "protocolVersion":
			out.ProtocolVersion = string(in.String())
		case "product":
			out.Product = string(in.String())
		case "revision":
			out.Revision = string(in.String())
		case "userAgent":
			out.UserAgent = string(in.String())
		case "jsVersion":
			out.JsVersion = string(in.String())
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
func easyjsonC5a4559bEncodeGithubComChromedpCdprotoBrowser5(out *jwriter.Writer, in GetVersionReturns) {
	out.RawByte('{')
	first := true
	_ = first
	if in.ProtocolVersion != "" {
		const prefix string = ",\"protocolVersion\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.ProtocolVersion))
	}
	if in.Product != "" {
		const prefix string = ",\"product\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.Product))
	}
	if in.Revision != "" {
		const prefix string = ",\"revision\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.Revision))
	}
	if in.UserAgent != "" {
		const prefix string = ",\"userAgent\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.UserAgent))
	}
	if in.JsVersion != "" {
		const prefix string = ",\"jsVersion\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.JsVersion))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v GetVersionReturns) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonC5a4559bEncodeGithubComChromedpCdprotoBrowser5(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v GetVersionReturns) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonC5a4559bEncodeGithubComChromedpCdprotoBrowser5(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *GetVersionReturns) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonC5a4559bDecodeGithubComChromedpCdprotoBrowser5(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *GetVersionReturns) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonC5a4559bDecodeGithubComChromedpCdprotoBrowser5(l, v)
}
func easyjsonC5a4559bDecodeGithubComChromedpCdprotoBrowser6(in *jlexer.Lexer, out *GetVersionParams) {
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
func easyjsonC5a4559bEncodeGithubComChromedpCdprotoBrowser6(out *jwriter.Writer, in GetVersionParams) {
	out.RawByte('{')
	first := true
	_ = first
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v GetVersionParams) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonC5a4559bEncodeGithubComChromedpCdprotoBrowser6(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v GetVersionParams) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonC5a4559bEncodeGithubComChromedpCdprotoBrowser6(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *GetVersionParams) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonC5a4559bDecodeGithubComChromedpCdprotoBrowser6(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *GetVersionParams) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonC5a4559bDecodeGithubComChromedpCdprotoBrowser6(l, v)
}
func easyjsonC5a4559bDecodeGithubComChromedpCdprotoBrowser7(in *jlexer.Lexer, out *CloseParams) {
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
func easyjsonC5a4559bEncodeGithubComChromedpCdprotoBrowser7(out *jwriter.Writer, in CloseParams) {
	out.RawByte('{')
	first := true
	_ = first
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v CloseParams) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonC5a4559bEncodeGithubComChromedpCdprotoBrowser7(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v CloseParams) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonC5a4559bEncodeGithubComChromedpCdprotoBrowser7(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *CloseParams) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonC5a4559bDecodeGithubComChromedpCdprotoBrowser7(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *CloseParams) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonC5a4559bDecodeGithubComChromedpCdprotoBrowser7(l, v)
}
func easyjsonC5a4559bDecodeGithubComChromedpCdprotoBrowser8(in *jlexer.Lexer, out *Bounds) {
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
		case "left":
			out.Left = int64(in.Int64())
		case "top":
			out.Top = int64(in.Int64())
		case "width":
			out.Width = int64(in.Int64())
		case "height":
			out.Height = int64(in.Int64())
		case "windowState":
			(out.WindowState).UnmarshalEasyJSON(in)
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
func easyjsonC5a4559bEncodeGithubComChromedpCdprotoBrowser8(out *jwriter.Writer, in Bounds) {
	out.RawByte('{')
	first := true
	_ = first
	if in.Left != 0 {
		const prefix string = ",\"left\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int64(int64(in.Left))
	}
	if in.Top != 0 {
		const prefix string = ",\"top\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int64(int64(in.Top))
	}
	if in.Width != 0 {
		const prefix string = ",\"width\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int64(int64(in.Width))
	}
	if in.Height != 0 {
		const prefix string = ",\"height\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int64(int64(in.Height))
	}
	if in.WindowState != "" {
		const prefix string = ",\"windowState\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		(in.WindowState).MarshalEasyJSON(out)
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v Bounds) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonC5a4559bEncodeGithubComChromedpCdprotoBrowser8(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v Bounds) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonC5a4559bEncodeGithubComChromedpCdprotoBrowser8(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *Bounds) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonC5a4559bDecodeGithubComChromedpCdprotoBrowser8(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *Bounds) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonC5a4559bDecodeGithubComChromedpCdprotoBrowser8(l, v)
}
