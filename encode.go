package coolqparser

import "strings"

//EnCode 把map转成CQ码
func EnCode(data []map[string]string) string {
	output := ""
	//
	for _, value := range data {
		//对特殊进行处理
		switch value["CQ"] {
		case TEXT:
			output += encodeText(value["value"])
		default:
			output += enCodeCQ(value)
		}
	}
	return output
}

func enCodeCQ(Data map[string]string) string {
	output := "[CQ:" + Data["CQ"]
	delete(Data, output)
	for key, value := range Data {
		if value == "" {
			continue
		}
		output += "," + key + "=" + encodeValue(value)
	}
	return output + "]"
}

func encodeText(text string) string {
	text = strings.Replace(text, "&", "&amp;", -1)
	text = strings.Replace(text, "[", "&#91;", -1)
	text = strings.Replace(text, "]", "&#93;", -1)
	return text

}

func encodeValue(text string) string {
	text = encodeText(text)
	text = strings.Replace(text, ",", "&#44;", -1)
	return text
}
