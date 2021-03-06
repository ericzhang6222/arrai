package syntax

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"

	"github.com/arr-ai/arrai/rel"
	"github.com/arr-ai/arrai/translate"
)

func stdEncodingJSON() rel.Attr {
	return rel.NewTupleAttr(
		"json",
		rel.NewNativeFunctionAttr("decode", func(_ context.Context, v rel.Value) (rel.Value, error) {
			var bytes []byte
			switch v := v.(type) {
			case rel.String:
				bytes = []byte(v.String())
			case rel.Bytes:
				bytes = v.Bytes()
			default:
				return nil, fmt.Errorf("unexpected arrai object type: %s", rel.ValueTypeAsString(v))
			}
			return bytesJSONToArrai(bytes)
		}),
		rel.NewNativeFunctionAttr("encode", func(_ context.Context, v rel.Value) (rel.Value, error) {
			data, err := translate.FromArrai(v)
			if err != nil {
				return nil, err
			}
			result := bytes.NewBuffer([]byte{})
			enc := json.NewEncoder(result)
			enc.SetEscapeHTML(false)
			err = enc.Encode(data)
			if err != nil {
				return nil, err
			}
			return rel.NewBytes(result.Bytes()), nil
		}),
		rel.NewNativeFunctionAttr("encode_indent", func(_ context.Context, v rel.Value) (rel.Value, error) {
			data, err := translate.FromArrai(v)
			if err != nil {
				return nil, err
			}
			result := bytes.NewBuffer([]byte{})
			enc := json.NewEncoder(result)
			enc.SetEscapeHTML(false)
			enc.SetIndent("", "  ")
			err = enc.Encode(data)
			if err != nil {
				return nil, err
			}
			return rel.NewBytes(result.Bytes()), nil
		}),
	)
}

func bytesJSONToArrai(bytes []byte) (rel.Value, error) {
	var data interface{}
	var err error
	if err = json.Unmarshal(bytes, &data); err == nil {
		var d rel.Value
		if d, err = translate.ToArrai(data); err == nil {
			return d, nil
		}
	}
	return nil, err
}
