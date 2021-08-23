package pars

import "golang.org/x/net/html"

func ParseAttributes(z *html.Tokenizer) map[string]string {
	retval := map[string]string{}

	for {
		key, value, hasMore := z.TagAttr()

		if string(key) != "" {
			retval[string(key)] = string(value)
		}

		if !hasMore {
			break
		}
	}

	return retval
}
