package misc

import (
	"bytes"
	"encoding/json"
	"fmt"
)

func ToJSON(x interface{}) []byte {
	buf, _ := json.Marshal(x)
	return buf
}

func ToJSONString(x interface{}) string {
	buf, _ := json.Marshal(x)
	return string(buf)
}

func ToPrettyJSON(x interface{}) []byte {
	buf, _ := json.MarshalIndent(x, "", "\t")
	return buf
}

func ToPrettyJSONString(x interface{}) string {
	buf, _ := json.MarshalIndent(x, "", "\t")
	return string(buf)
}

func ToPlainJSON(v interface{}) []byte {
	var buf bytes.Buffer
	enc := json.NewEncoder(&buf)
	enc.SetEscapeHTML(false)
	if err := enc.Encode(v); err == nil {
		return bytes.TrimSuffix(buf.Bytes(), []byte("\n"))
	}

	buf.Reset()
	_, _ = fmt.Fprintf(&buf, "%+v", v)
	return buf.Bytes()
}

func ToPlainJSONString(v interface{}) string {
	var buf bytes.Buffer
	enc := json.NewEncoder(&buf)
	enc.SetEscapeHTML(false)
	if err := enc.Encode(v); err == nil {
		return string(bytes.TrimSuffix(buf.Bytes(), []byte("\n")))
	}

	buf.Reset()
	_, _ = fmt.Fprintf(&buf, "%+v", v)
	return buf.String()
}
