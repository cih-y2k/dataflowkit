// AUTOGENERATED FILE: easyjson marshaller/unmarshallers.

package parser

import (
	json "encoding/json"
	jlexer "github.com/mailru/easyjson/jlexer"
	jwriter "github.com/mailru/easyjson/jwriter"
)

// suppress unused package warning
var (
	_ = json.RawMessage{}
	_ = jlexer.Lexer{}
	_ = jwriter.Writer{}
)

func easyjsonBc4ecbcDecodeDataflowkitParser(in *jlexer.Lexer, out *CSVTableCollection) {
	if in.IsNull() {
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
		case "tables":
			if in.IsNull() {
				in.Skip()
				out.Tables = nil
			} else {
				in.Delim('[')
				if !in.IsDelim(']') {
					out.Tables = make([]CSVTable, 0, 2)
				} else {
					out.Tables = []CSVTable{}
				}
				for !in.IsDelim(']') {
					var v1 CSVTable
					(v1).UnmarshalEasyJSON(in)
					out.Tables = append(out.Tables, v1)
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
}
func easyjsonBc4ecbcEncodeDataflowkitParser(out *jwriter.Writer, in CSVTableCollection) {
	out.RawByte('{')
	first := true
	_ = first
	if !first {
		out.RawByte(',')
	}
	first = false
	out.RawString("\"tables\":")
	if in.Tables == nil {
		out.RawString("null")
	} else {
		out.RawByte('[')
		for v2, v3 := range in.Tables {
			if v2 > 0 {
				out.RawByte(',')
			}
			(v3).MarshalEasyJSON(out)
		}
		out.RawByte(']')
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v CSVTableCollection) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonBc4ecbcEncodeDataflowkitParser(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v CSVTableCollection) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonBc4ecbcEncodeDataflowkitParser(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *CSVTableCollection) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonBc4ecbcDecodeDataflowkitParser(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *CSVTableCollection) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonBc4ecbcDecodeDataflowkitParser(l, v)
}
func easyjsonBc4ecbcDecodeDataflowkitParser1(in *jlexer.Lexer, out *CSVTable) {
	if in.IsNull() {
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
		case "table":
			out.Content = string(in.String())
		case "url":
			out.URL = string(in.String())
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
}
func easyjsonBc4ecbcEncodeDataflowkitParser1(out *jwriter.Writer, in CSVTable) {
	out.RawByte('{')
	first := true
	_ = first
	if !first {
		out.RawByte(',')
	}
	first = false
	out.RawString("\"table\":")
	out.String(string(in.Content))
	if !first {
		out.RawByte(',')
	}
	first = false
	out.RawString("\"url\":")
	out.String(string(in.URL))
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v CSVTable) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonBc4ecbcEncodeDataflowkitParser1(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v CSVTable) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonBc4ecbcEncodeDataflowkitParser1(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *CSVTable) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonBc4ecbcDecodeDataflowkitParser1(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *CSVTable) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonBc4ecbcDecodeDataflowkitParser1(l, v)
}
func easyjsonBc4ecbcDecodeDataflowkitParser2(in *jlexer.Lexer, out *Out) {
	if in.IsNull() {
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
		case "collections":
			if in.IsNull() {
				in.Skip()
				out.Element = nil
			} else {
				in.Delim('[')
				if !in.IsDelim(']') {
					out.Element = make([]outItem, 0, 1)
				} else {
					out.Element = []outItem{}
				}
				for !in.IsDelim(']') {
					var v4 outItem
					(v4).UnmarshalEasyJSON(in)
					out.Element = append(out.Element, v4)
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
}
func easyjsonBc4ecbcEncodeDataflowkitParser2(out *jwriter.Writer, in Out) {
	out.RawByte('{')
	first := true
	_ = first
	if !first {
		out.RawByte(',')
	}
	first = false
	out.RawString("\"collections\":")
	if in.Element == nil {
		out.RawString("null")
	} else {
		out.RawByte('[')
		for v5, v6 := range in.Element {
			if v5 > 0 {
				out.RawByte(',')
			}
			(v6).MarshalEasyJSON(out)
		}
		out.RawByte(']')
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v Out) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonBc4ecbcEncodeDataflowkitParser2(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v Out) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonBc4ecbcEncodeDataflowkitParser2(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *Out) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonBc4ecbcDecodeDataflowkitParser2(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *Out) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonBc4ecbcDecodeDataflowkitParser2(l, v)
}
func easyjsonBc4ecbcDecodeDataflowkitParser3(in *jlexer.Lexer, out *outItem) {
	if in.IsNull() {
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
		case "items":
			if in.IsNull() {
				in.Skip()
				out.Items = nil
			} else {
				in.Delim('[')
				if !in.IsDelim(']') {
					out.Items = make([]interface{}, 0, 4)
				} else {
					out.Items = []interface{}{}
				}
				for !in.IsDelim(']') {
					var v7 interface{}
					v7 = in.Interface()
					out.Items = append(out.Items, v7)
					in.WantComma()
				}
				in.Delim(']')
			}
		case "count":
			out.Count = int(in.Int())
		case "time":
			out.CreatedAt = int64(in.Int64())
		case "name":
			out.Name = string(in.String())
		case "url":
			out.URL = string(in.String())
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
}
func easyjsonBc4ecbcEncodeDataflowkitParser3(out *jwriter.Writer, in outItem) {
	out.RawByte('{')
	first := true
	_ = first
	if !first {
		out.RawByte(',')
	}
	first = false
	out.RawString("\"items\":")
	if in.Items == nil {
		out.RawString("null")
	} else {
		out.RawByte('[')
		for v8, v9 := range in.Items {
			if v8 > 0 {
				out.RawByte(',')
			}
			if m, ok := v9.(json.Marshaler); ok {
				out.Raw(m.MarshalJSON())
			} else {
				out.Raw(json.Marshal(v9))
			}
		}
		out.RawByte(']')
	}
	if !first {
		out.RawByte(',')
	}
	first = false
	out.RawString("\"count\":")
	out.Int(int(in.Count))
	if !first {
		out.RawByte(',')
	}
	first = false
	out.RawString("\"time\":")
	out.Int64(int64(in.CreatedAt))
	if !first {
		out.RawByte(',')
	}
	first = false
	out.RawString("\"name\":")
	out.String(string(in.Name))
	if !first {
		out.RawByte(',')
	}
	first = false
	out.RawString("\"url\":")
	out.String(string(in.URL))
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v outItem) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonBc4ecbcEncodeDataflowkitParser3(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v outItem) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonBc4ecbcEncodeDataflowkitParser3(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *outItem) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonBc4ecbcDecodeDataflowkitParser3(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *outItem) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonBc4ecbcDecodeDataflowkitParser3(l, v)
}
func easyjsonBc4ecbcDecodeDataflowkitParser4(in *jlexer.Lexer, out *Payloads) {
	if in.IsNull() {
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
		case "format":
			out.Format = string(in.String())
		case "collections":
			if in.IsNull() {
				in.Skip()
				out.Collections = nil
			} else {
				in.Delim('[')
				if !in.IsDelim(']') {
					out.Collections = make([]payload, 0, 1)
				} else {
					out.Collections = []payload{}
				}
				for !in.IsDelim(']') {
					var v10 payload
					easyjsonBc4ecbcDecodeDataflowkitParser5(in, &v10)
					out.Collections = append(out.Collections, v10)
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
}
func easyjsonBc4ecbcEncodeDataflowkitParser4(out *jwriter.Writer, in Payloads) {
	out.RawByte('{')
	first := true
	_ = first
	if !first {
		out.RawByte(',')
	}
	first = false
	out.RawString("\"format\":")
	out.String(string(in.Format))
	if !first {
		out.RawByte(',')
	}
	first = false
	out.RawString("\"collections\":")
	if in.Collections == nil {
		out.RawString("null")
	} else {
		out.RawByte('[')
		for v11, v12 := range in.Collections {
			if v11 > 0 {
				out.RawByte(',')
			}
			easyjsonBc4ecbcEncodeDataflowkitParser5(out, v12)
		}
		out.RawByte(']')
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v Payloads) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonBc4ecbcEncodeDataflowkitParser4(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v Payloads) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonBc4ecbcEncodeDataflowkitParser4(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *Payloads) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonBc4ecbcDecodeDataflowkitParser4(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *Payloads) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonBc4ecbcDecodeDataflowkitParser4(l, v)
}
func easyjsonBc4ecbcDecodeDataflowkitParser5(in *jlexer.Lexer, out *payload) {
	if in.IsNull() {
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
		case "fields":
			if in.IsNull() {
				in.Skip()
				out.Fields = nil
			} else {
				in.Delim('[')
				if !in.IsDelim(']') {
					out.Fields = make([]field, 0, 2)
				} else {
					out.Fields = []field{}
				}
				for !in.IsDelim(']') {
					var v13 field
					easyjsonBc4ecbcDecodeDataflowkitParser6(in, &v13)
					out.Fields = append(out.Fields, v13)
					in.WantComma()
				}
				in.Delim(']')
			}
		case "name":
			out.Name = string(in.String())
		case "url":
			out.URL = string(in.String())
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
}
func easyjsonBc4ecbcEncodeDataflowkitParser5(out *jwriter.Writer, in payload) {
	out.RawByte('{')
	first := true
	_ = first
	if !first {
		out.RawByte(',')
	}
	first = false
	out.RawString("\"fields\":")
	if in.Fields == nil {
		out.RawString("null")
	} else {
		out.RawByte('[')
		for v14, v15 := range in.Fields {
			if v14 > 0 {
				out.RawByte(',')
			}
			easyjsonBc4ecbcEncodeDataflowkitParser6(out, v15)
		}
		out.RawByte(']')
	}
	if !first {
		out.RawByte(',')
	}
	first = false
	out.RawString("\"name\":")
	out.String(string(in.Name))
	if !first {
		out.RawByte(',')
	}
	first = false
	out.RawString("\"url\":")
	out.String(string(in.URL))
	out.RawByte('}')
}
func easyjsonBc4ecbcDecodeDataflowkitParser6(in *jlexer.Lexer, out *field) {
	if in.IsNull() {
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
		case "field_name":
			out.FieldName = string(in.String())
		case "css_selector":
			out.CSSSelector = string(in.String())
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
}
func easyjsonBc4ecbcEncodeDataflowkitParser6(out *jwriter.Writer, in field) {
	out.RawByte('{')
	first := true
	_ = first
	if !first {
		out.RawByte(',')
	}
	first = false
	out.RawString("\"field_name\":")
	out.String(string(in.FieldName))
	if !first {
		out.RawByte(',')
	}
	first = false
	out.RawString("\"css_selector\":")
	out.String(string(in.CSSSelector))
	out.RawByte('}')
}
