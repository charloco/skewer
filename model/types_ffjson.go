// Code generated by ffjson <https://github.com/pquerna/ffjson>. DO NOT EDIT.
// source: types.go

package model

import (
	"bytes"
	"fmt"
	fflib "github.com/pquerna/ffjson/fflib/v1"
)

// MarshalJSON marshal bytes to json - template
func (j *AuditMessageGroup) MarshalJSON() ([]byte, error) {
	var buf fflib.Buffer
	if j == nil {
		buf.WriteString("null")
		return buf.Bytes(), nil
	}
	err := j.MarshalJSONBuf(&buf)
	if err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

// MarshalJSONBuf marshal buff to json - template
func (j *AuditMessageGroup) MarshalJSONBuf(buf fflib.EncodingBuffer) error {
	if j == nil {
		buf.WriteString("null")
		return nil
	}
	var err error
	var obj []byte
	_ = obj
	_ = err
	buf.WriteString(`{"sequence":`)
	fflib.FormatBits2(buf, uint64(j.Seq), 10, j.Seq < 0)
	buf.WriteString(`,"timestamp":`)
	fflib.WriteJsonString(buf, string(j.AuditTime))
	buf.WriteString(`,"messages":`)
	if j.Msgs != nil {
		buf.WriteString(`[`)
		for i, v := range j.Msgs {
			if i != 0 {
				buf.WriteString(`,`)
			}

			{

				err = v.MarshalJSONBuf(buf)
				if err != nil {
					return err
				}

			}
		}
		buf.WriteString(`]`)
	} else {
		buf.WriteString(`null`)
	}
	if j.UidMap == nil {
		buf.WriteString(`,"uid_map":null`)
	} else {
		buf.WriteString(`,"uid_map":{ `)
		for key, value := range j.UidMap {
			fflib.WriteJsonString(buf, key)
			buf.WriteString(`:`)
			fflib.WriteJsonString(buf, string(value))
			buf.WriteByte(',')
		}
		buf.Rewind(1)
		buf.WriteByte('}')
	}
	buf.WriteByte('}')
	return nil
}

const (
	ffjtAuditMessageGroupbase = iota
	ffjtAuditMessageGroupnosuchkey

	ffjtAuditMessageGroupSeq

	ffjtAuditMessageGroupAuditTime

	ffjtAuditMessageGroupMsgs

	ffjtAuditMessageGroupUidMap
)

var ffjKeyAuditMessageGroupSeq = []byte("sequence")

var ffjKeyAuditMessageGroupAuditTime = []byte("timestamp")

var ffjKeyAuditMessageGroupMsgs = []byte("messages")

var ffjKeyAuditMessageGroupUidMap = []byte("uid_map")

// UnmarshalJSON umarshall json - template of ffjson
func (j *AuditMessageGroup) UnmarshalJSON(input []byte) error {
	fs := fflib.NewFFLexer(input)
	return j.UnmarshalJSONFFLexer(fs, fflib.FFParse_map_start)
}

// UnmarshalJSONFFLexer fast json unmarshall - template ffjson
func (j *AuditMessageGroup) UnmarshalJSONFFLexer(fs *fflib.FFLexer, state fflib.FFParseState) error {
	var err error
	currentKey := ffjtAuditMessageGroupbase
	_ = currentKey
	tok := fflib.FFTok_init
	wantedTok := fflib.FFTok_init

mainparse:
	for {
		tok = fs.Scan()
		//	println(fmt.Sprintf("debug: tok: %v  state: %v", tok, state))
		if tok == fflib.FFTok_error {
			goto tokerror
		}

		switch state {

		case fflib.FFParse_map_start:
			if tok != fflib.FFTok_left_bracket {
				wantedTok = fflib.FFTok_left_bracket
				goto wrongtokenerror
			}
			state = fflib.FFParse_want_key
			continue

		case fflib.FFParse_after_value:
			if tok == fflib.FFTok_comma {
				state = fflib.FFParse_want_key
			} else if tok == fflib.FFTok_right_bracket {
				goto done
			} else {
				wantedTok = fflib.FFTok_comma
				goto wrongtokenerror
			}

		case fflib.FFParse_want_key:
			// json {} ended. goto exit. woo.
			if tok == fflib.FFTok_right_bracket {
				goto done
			}
			if tok != fflib.FFTok_string {
				wantedTok = fflib.FFTok_string
				goto wrongtokenerror
			}

			kn := fs.Output.Bytes()
			if len(kn) <= 0 {
				// "" case. hrm.
				currentKey = ffjtAuditMessageGroupnosuchkey
				state = fflib.FFParse_want_colon
				goto mainparse
			} else {
				switch kn[0] {

				case 'm':

					if bytes.Equal(ffjKeyAuditMessageGroupMsgs, kn) {
						currentKey = ffjtAuditMessageGroupMsgs
						state = fflib.FFParse_want_colon
						goto mainparse
					}

				case 's':

					if bytes.Equal(ffjKeyAuditMessageGroupSeq, kn) {
						currentKey = ffjtAuditMessageGroupSeq
						state = fflib.FFParse_want_colon
						goto mainparse
					}

				case 't':

					if bytes.Equal(ffjKeyAuditMessageGroupAuditTime, kn) {
						currentKey = ffjtAuditMessageGroupAuditTime
						state = fflib.FFParse_want_colon
						goto mainparse
					}

				case 'u':

					if bytes.Equal(ffjKeyAuditMessageGroupUidMap, kn) {
						currentKey = ffjtAuditMessageGroupUidMap
						state = fflib.FFParse_want_colon
						goto mainparse
					}

				}

				if fflib.AsciiEqualFold(ffjKeyAuditMessageGroupUidMap, kn) {
					currentKey = ffjtAuditMessageGroupUidMap
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				if fflib.EqualFoldRight(ffjKeyAuditMessageGroupMsgs, kn) {
					currentKey = ffjtAuditMessageGroupMsgs
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				if fflib.EqualFoldRight(ffjKeyAuditMessageGroupAuditTime, kn) {
					currentKey = ffjtAuditMessageGroupAuditTime
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				if fflib.EqualFoldRight(ffjKeyAuditMessageGroupSeq, kn) {
					currentKey = ffjtAuditMessageGroupSeq
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				currentKey = ffjtAuditMessageGroupnosuchkey
				state = fflib.FFParse_want_colon
				goto mainparse
			}

		case fflib.FFParse_want_colon:
			if tok != fflib.FFTok_colon {
				wantedTok = fflib.FFTok_colon
				goto wrongtokenerror
			}
			state = fflib.FFParse_want_value
			continue
		case fflib.FFParse_want_value:

			if tok == fflib.FFTok_left_brace || tok == fflib.FFTok_left_bracket || tok == fflib.FFTok_integer || tok == fflib.FFTok_double || tok == fflib.FFTok_string || tok == fflib.FFTok_bool || tok == fflib.FFTok_null {
				switch currentKey {

				case ffjtAuditMessageGroupSeq:
					goto handle_Seq

				case ffjtAuditMessageGroupAuditTime:
					goto handle_AuditTime

				case ffjtAuditMessageGroupMsgs:
					goto handle_Msgs

				case ffjtAuditMessageGroupUidMap:
					goto handle_UidMap

				case ffjtAuditMessageGroupnosuchkey:
					err = fs.SkipField(tok)
					if err != nil {
						return fs.WrapErr(err)
					}
					state = fflib.FFParse_after_value
					goto mainparse
				}
			} else {
				goto wantedvalue
			}
		}
	}

handle_Seq:

	/* handler: j.Seq type=int kind=int quoted=false*/

	{
		if tok != fflib.FFTok_integer && tok != fflib.FFTok_null {
			return fs.WrapErr(fmt.Errorf("cannot unmarshal %s into Go value for int", tok))
		}
	}

	{

		if tok == fflib.FFTok_null {

		} else {

			tval, err := fflib.ParseInt(fs.Output.Bytes(), 10, 64)

			if err != nil {
				return fs.WrapErr(err)
			}

			j.Seq = int(tval)

		}
	}

	state = fflib.FFParse_after_value
	goto mainparse

handle_AuditTime:

	/* handler: j.AuditTime type=string kind=string quoted=false*/

	{

		{
			if tok != fflib.FFTok_string && tok != fflib.FFTok_null {
				return fs.WrapErr(fmt.Errorf("cannot unmarshal %s into Go value for string", tok))
			}
		}

		if tok == fflib.FFTok_null {

		} else {

			outBuf := fs.Output.Bytes()

			j.AuditTime = string(string(outBuf))

		}
	}

	state = fflib.FFParse_after_value
	goto mainparse

handle_Msgs:

	/* handler: j.Msgs type=[]model.AuditSubMessage kind=slice quoted=false*/

	{

		{
			if tok != fflib.FFTok_left_brace && tok != fflib.FFTok_null {
				return fs.WrapErr(fmt.Errorf("cannot unmarshal %s into Go value for ", tok))
			}
		}

		if tok == fflib.FFTok_null {
			j.Msgs = nil
		} else {

			j.Msgs = []AuditSubMessage{}

			wantVal := true

			for {

				var tmpJMsgs AuditSubMessage

				tok = fs.Scan()
				if tok == fflib.FFTok_error {
					goto tokerror
				}
				if tok == fflib.FFTok_right_brace {
					break
				}

				if tok == fflib.FFTok_comma {
					if wantVal == true {
						// TODO(pquerna): this isn't an ideal error message, this handles
						// things like [,,,] as an array value.
						return fs.WrapErr(fmt.Errorf("wanted value token, but got token: %v", tok))
					}
					continue
				} else {
					wantVal = true
				}

				/* handler: tmpJMsgs type=model.AuditSubMessage kind=struct quoted=false*/

				{
					if tok == fflib.FFTok_null {

						state = fflib.FFParse_after_value
						goto mainparse
					}

					err = tmpJMsgs.UnmarshalJSONFFLexer(fs, fflib.FFParse_want_key)
					if err != nil {
						return err
					}
					state = fflib.FFParse_after_value
				}

				j.Msgs = append(j.Msgs, tmpJMsgs)

				wantVal = false
			}
		}
	}

	state = fflib.FFParse_after_value
	goto mainparse

handle_UidMap:

	/* handler: j.UidMap type=map[string]string kind=map quoted=false*/

	{

		{
			if tok != fflib.FFTok_left_bracket && tok != fflib.FFTok_null {
				return fs.WrapErr(fmt.Errorf("cannot unmarshal %s into Go value for ", tok))
			}
		}

		if tok == fflib.FFTok_null {
			j.UidMap = nil
		} else {

			j.UidMap = make(map[string]string, 0)

			wantVal := true

			for {

				var k string

				var tmpJUidMap string

				tok = fs.Scan()
				if tok == fflib.FFTok_error {
					goto tokerror
				}
				if tok == fflib.FFTok_right_bracket {
					break
				}

				if tok == fflib.FFTok_comma {
					if wantVal == true {
						// TODO(pquerna): this isn't an ideal error message, this handles
						// things like [,,,] as an array value.
						return fs.WrapErr(fmt.Errorf("wanted value token, but got token: %v", tok))
					}
					continue
				} else {
					wantVal = true
				}

				/* handler: k type=string kind=string quoted=false*/

				{

					{
						if tok != fflib.FFTok_string && tok != fflib.FFTok_null {
							return fs.WrapErr(fmt.Errorf("cannot unmarshal %s into Go value for string", tok))
						}
					}

					if tok == fflib.FFTok_null {

					} else {

						outBuf := fs.Output.Bytes()

						k = string(string(outBuf))

					}
				}

				// Expect ':' after key
				tok = fs.Scan()
				if tok != fflib.FFTok_colon {
					return fs.WrapErr(fmt.Errorf("wanted colon token, but got token: %v", tok))
				}

				tok = fs.Scan()
				/* handler: tmpJUidMap type=string kind=string quoted=false*/

				{

					{
						if tok != fflib.FFTok_string && tok != fflib.FFTok_null {
							return fs.WrapErr(fmt.Errorf("cannot unmarshal %s into Go value for string", tok))
						}
					}

					if tok == fflib.FFTok_null {

					} else {

						outBuf := fs.Output.Bytes()

						tmpJUidMap = string(string(outBuf))

					}
				}

				j.UidMap[k] = tmpJUidMap

				wantVal = false
			}

		}
	}

	state = fflib.FFParse_after_value
	goto mainparse

wantedvalue:
	return fs.WrapErr(fmt.Errorf("wanted value token, but got token: %v", tok))
wrongtokenerror:
	return fs.WrapErr(fmt.Errorf("ffjson: wanted token: %v, but got token: %v output=%s", wantedTok, tok, fs.Output.String()))
tokerror:
	if fs.BigError != nil {
		return fs.WrapErr(fs.BigError)
	}
	err = fs.Error.ToError()
	if err != nil {
		return fs.WrapErr(err)
	}
	panic("ffjson-generated: unreachable, please report bug.")
done:

	return nil
}

// MarshalJSON marshal bytes to json - template
func (j *AuditSubMessage) MarshalJSON() ([]byte, error) {
	var buf fflib.Buffer
	if j == nil {
		buf.WriteString("null")
		return buf.Bytes(), nil
	}
	err := j.MarshalJSONBuf(&buf)
	if err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

// MarshalJSONBuf marshal buff to json - template
func (j *AuditSubMessage) MarshalJSONBuf(buf fflib.EncodingBuffer) error {
	if j == nil {
		buf.WriteString("null")
		return nil
	}
	var err error
	var obj []byte
	_ = obj
	_ = err
	buf.WriteString(`{"type":`)
	fflib.FormatBits2(buf, uint64(j.Type), 10, false)
	buf.WriteString(`,"data":`)
	fflib.WriteJsonString(buf, string(j.Data))
	buf.WriteByte('}')
	return nil
}

const (
	ffjtAuditSubMessagebase = iota
	ffjtAuditSubMessagenosuchkey

	ffjtAuditSubMessageType

	ffjtAuditSubMessageData
)

var ffjKeyAuditSubMessageType = []byte("type")

var ffjKeyAuditSubMessageData = []byte("data")

// UnmarshalJSON umarshall json - template of ffjson
func (j *AuditSubMessage) UnmarshalJSON(input []byte) error {
	fs := fflib.NewFFLexer(input)
	return j.UnmarshalJSONFFLexer(fs, fflib.FFParse_map_start)
}

// UnmarshalJSONFFLexer fast json unmarshall - template ffjson
func (j *AuditSubMessage) UnmarshalJSONFFLexer(fs *fflib.FFLexer, state fflib.FFParseState) error {
	var err error
	currentKey := ffjtAuditSubMessagebase
	_ = currentKey
	tok := fflib.FFTok_init
	wantedTok := fflib.FFTok_init

mainparse:
	for {
		tok = fs.Scan()
		//	println(fmt.Sprintf("debug: tok: %v  state: %v", tok, state))
		if tok == fflib.FFTok_error {
			goto tokerror
		}

		switch state {

		case fflib.FFParse_map_start:
			if tok != fflib.FFTok_left_bracket {
				wantedTok = fflib.FFTok_left_bracket
				goto wrongtokenerror
			}
			state = fflib.FFParse_want_key
			continue

		case fflib.FFParse_after_value:
			if tok == fflib.FFTok_comma {
				state = fflib.FFParse_want_key
			} else if tok == fflib.FFTok_right_bracket {
				goto done
			} else {
				wantedTok = fflib.FFTok_comma
				goto wrongtokenerror
			}

		case fflib.FFParse_want_key:
			// json {} ended. goto exit. woo.
			if tok == fflib.FFTok_right_bracket {
				goto done
			}
			if tok != fflib.FFTok_string {
				wantedTok = fflib.FFTok_string
				goto wrongtokenerror
			}

			kn := fs.Output.Bytes()
			if len(kn) <= 0 {
				// "" case. hrm.
				currentKey = ffjtAuditSubMessagenosuchkey
				state = fflib.FFParse_want_colon
				goto mainparse
			} else {
				switch kn[0] {

				case 'd':

					if bytes.Equal(ffjKeyAuditSubMessageData, kn) {
						currentKey = ffjtAuditSubMessageData
						state = fflib.FFParse_want_colon
						goto mainparse
					}

				case 't':

					if bytes.Equal(ffjKeyAuditSubMessageType, kn) {
						currentKey = ffjtAuditSubMessageType
						state = fflib.FFParse_want_colon
						goto mainparse
					}

				}

				if fflib.SimpleLetterEqualFold(ffjKeyAuditSubMessageData, kn) {
					currentKey = ffjtAuditSubMessageData
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				if fflib.SimpleLetterEqualFold(ffjKeyAuditSubMessageType, kn) {
					currentKey = ffjtAuditSubMessageType
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				currentKey = ffjtAuditSubMessagenosuchkey
				state = fflib.FFParse_want_colon
				goto mainparse
			}

		case fflib.FFParse_want_colon:
			if tok != fflib.FFTok_colon {
				wantedTok = fflib.FFTok_colon
				goto wrongtokenerror
			}
			state = fflib.FFParse_want_value
			continue
		case fflib.FFParse_want_value:

			if tok == fflib.FFTok_left_brace || tok == fflib.FFTok_left_bracket || tok == fflib.FFTok_integer || tok == fflib.FFTok_double || tok == fflib.FFTok_string || tok == fflib.FFTok_bool || tok == fflib.FFTok_null {
				switch currentKey {

				case ffjtAuditSubMessageType:
					goto handle_Type

				case ffjtAuditSubMessageData:
					goto handle_Data

				case ffjtAuditSubMessagenosuchkey:
					err = fs.SkipField(tok)
					if err != nil {
						return fs.WrapErr(err)
					}
					state = fflib.FFParse_after_value
					goto mainparse
				}
			} else {
				goto wantedvalue
			}
		}
	}

handle_Type:

	/* handler: j.Type type=uint16 kind=uint16 quoted=false*/

	{
		if tok != fflib.FFTok_integer && tok != fflib.FFTok_null {
			return fs.WrapErr(fmt.Errorf("cannot unmarshal %s into Go value for uint16", tok))
		}
	}

	{

		if tok == fflib.FFTok_null {

		} else {

			tval, err := fflib.ParseUint(fs.Output.Bytes(), 10, 16)

			if err != nil {
				return fs.WrapErr(err)
			}

			j.Type = uint16(tval)

		}
	}

	state = fflib.FFParse_after_value
	goto mainparse

handle_Data:

	/* handler: j.Data type=string kind=string quoted=false*/

	{

		{
			if tok != fflib.FFTok_string && tok != fflib.FFTok_null {
				return fs.WrapErr(fmt.Errorf("cannot unmarshal %s into Go value for string", tok))
			}
		}

		if tok == fflib.FFTok_null {

		} else {

			outBuf := fs.Output.Bytes()

			j.Data = string(string(outBuf))

		}
	}

	state = fflib.FFParse_after_value
	goto mainparse

wantedvalue:
	return fs.WrapErr(fmt.Errorf("wanted value token, but got token: %v", tok))
wrongtokenerror:
	return fs.WrapErr(fmt.Errorf("ffjson: wanted token: %v, but got token: %v output=%s", wantedTok, tok, fs.Output.String()))
tokerror:
	if fs.BigError != nil {
		return fs.WrapErr(fs.BigError)
	}
	err = fs.Error.ToError()
	if err != nil {
		return fs.WrapErr(err)
	}
	panic("ffjson-generated: unreachable, please report bug.")
done:

	return nil
}

// MarshalJSON marshal bytes to json - template
func (j *ExportedMessage) MarshalJSON() ([]byte, error) {
	var buf fflib.Buffer
	if j == nil {
		buf.WriteString("null")
		return buf.Bytes(), nil
	}
	err := j.MarshalJSONBuf(&buf)
	if err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

// MarshalJSONBuf marshal buff to json - template
func (j *ExportedMessage) MarshalJSONBuf(buf fflib.EncodingBuffer) error {
	if j == nil {
		buf.WriteString("null")
		return nil
	}
	var err error
	var obj []byte
	_ = obj
	_ = err
	buf.WriteString(`{ "fields":`)

	{

		err = j.Fields.MarshalJSONBuf(buf)
		if err != nil {
			return err
		}

	}
	buf.WriteByte(',')
	if len(j.Client) != 0 {
		buf.WriteString(`"client":`)
		fflib.WriteJsonString(buf, string(j.Client))
		buf.WriteByte(',')
	}
	buf.WriteString(`"local_port":"`)
	fflib.FormatBits2(buf, uint64(j.LocalPort), 10, j.LocalPort < 0)
	buf.WriteString(`",`)
	if len(j.UnixSocketPath) != 0 {
		buf.WriteString(`"unix_socket_path":`)
		fflib.WriteJsonString(buf, string(j.UnixSocketPath))
		buf.WriteByte(',')
	}
	buf.Rewind(1)
	buf.WriteByte('}')
	return nil
}

// MarshalJSON marshal bytes to json - template
func (j *ParsedMessage) MarshalJSON() ([]byte, error) {
	var buf fflib.Buffer
	if j == nil {
		buf.WriteString("null")
		return buf.Bytes(), nil
	}
	err := j.MarshalJSONBuf(&buf)
	if err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

// MarshalJSONBuf marshal buff to json - template
func (j *ParsedMessage) MarshalJSONBuf(buf fflib.EncodingBuffer) error {
	if j == nil {
		buf.WriteString("null")
		return nil
	}
	var err error
	var obj []byte
	_ = obj
	_ = err
	buf.WriteString(`{ "fields":`)

	{

		err = j.Fields.MarshalJSONBuf(buf)
		if err != nil {
			return err
		}

	}
	buf.WriteByte(',')
	if len(j.Client) != 0 {
		buf.WriteString(`"client":`)
		fflib.WriteJsonString(buf, string(j.Client))
		buf.WriteByte(',')
	}
	buf.WriteString(`"local_port":"`)
	fflib.FormatBits2(buf, uint64(j.LocalPort), 10, j.LocalPort < 0)
	buf.WriteString(`",`)
	if len(j.UnixSocketPath) != 0 {
		buf.WriteString(`"unix_socket_path":`)
		fflib.WriteJsonString(buf, string(j.UnixSocketPath))
		buf.WriteByte(',')
	}
	buf.Rewind(1)
	buf.WriteByte('}')
	return nil
}

// MarshalJSON marshal bytes to json - template
func (j *SyslogMessage) MarshalJSON() ([]byte, error) {
	var buf fflib.Buffer
	if j == nil {
		buf.WriteString("null")
		return buf.Bytes(), nil
	}
	err := j.MarshalJSONBuf(&buf)
	if err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

// MarshalJSONBuf marshal buff to json - template
func (j *SyslogMessage) MarshalJSONBuf(buf fflib.EncodingBuffer) error {
	if j == nil {
		buf.WriteString("null")
		return nil
	}
	var err error
	var obj []byte
	_ = obj
	_ = err
	buf.WriteString(`{ "priority":"`)
	fflib.FormatBits2(buf, uint64(j.Priority), 10, j.Priority < 0)
	buf.WriteString(`","facility":"`)
	fflib.FormatBits2(buf, uint64(j.Facility), 10, j.Facility < 0)
	buf.WriteString(`","severity":"`)
	fflib.FormatBits2(buf, uint64(j.Severity), 10, j.Severity < 0)
	buf.WriteString(`","version":"`)
	fflib.FormatBits2(buf, uint64(j.Version), 10, j.Version < 0)
	buf.WriteString(`","timereported":`)
	fflib.WriteJsonString(buf, string(j.TimeReported))
	buf.WriteString(`,"timegenerated":`)
	fflib.WriteJsonString(buf, string(j.TimeGenerated))
	buf.WriteString(`,"hostname":`)
	fflib.WriteJsonString(buf, string(j.Hostname))
	buf.WriteString(`,"appname":`)
	fflib.WriteJsonString(buf, string(j.Appname))
	buf.WriteByte(',')
	if len(j.Procid) != 0 {
		buf.WriteString(`"procid":`)
		fflib.WriteJsonString(buf, string(j.Procid))
		buf.WriteByte(',')
	}
	if len(j.Msgid) != 0 {
		buf.WriteString(`"msgid":`)
		fflib.WriteJsonString(buf, string(j.Msgid))
		buf.WriteByte(',')
	}
	if len(j.Structured) != 0 {
		buf.WriteString(`"structured":`)
		fflib.WriteJsonString(buf, string(j.Structured))
		buf.WriteByte(',')
	}
	buf.WriteString(`"message":`)
	fflib.WriteJsonString(buf, string(j.Message))
	buf.WriteByte(',')
	if len(j.Properties) != 0 {
		buf.WriteString(`"properties":`)
		/* Falling back. type=map[string]map[string]string kind=map */
		err = buf.Encode(j.Properties)
		if err != nil {
			return err
		}
		buf.WriteByte(',')
	}
	buf.Rewind(1)
	buf.WriteByte('}')
	return nil
}
