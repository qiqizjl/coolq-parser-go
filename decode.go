package coolqparser

import "strings"

//DeCode 对收到的数据进行解码
func DeCode(data string) []map[string]string {
	//切割
	//先切割]并保留
	output := make([]map[string]string, 0)
	tmp1 := make([]string, 0)
	str := strings.SplitAfter(data, "]")
	for _, val := range str {
		//根据[切片
		tmpStr := strings.Split(val, "[")
		tmp1 = append(tmp1, tmpStr[0])
		if len(tmpStr) > 1 {
			tmp1 = append(tmp1, tmpStr[1])
		}
	}
	//处理切片之后的数据
	for _, val := range tmp1 {
		//判断是否是CQ码
		tmpCQ := make(map[string]string)
		if strings.Contains(val, "]") {
			val = strings.Replace(val, "]", "", -1)
			//先用,做切割
			cqTmp := strings.Split(val, ",")
			//继续切割
			for _, cq := range cqTmp {
				//先尝试用=做切割
				tmpArr := strings.Split(cq, "=")
				if len(tmpArr) == 1 {
					tmpArr = strings.Split(cq, ":")
				}
				tmpCQ[tmpArr[0]] = decodeValue(tmpArr[1])
			}
		} else {
			tmpCQ["CQ"] = TEXT
			tmpCQ["value"] = decodeText(val)
		}
		output = append(output, tmpCQ)
	}
	return output
}

//decodeText 解码文字
func decodeText(text string) string {
	text = strings.Replace(text, "&amp;", "&", -1)
	text = strings.Replace(text, "&#91;", "[", -1)
	text = strings.Replace(text, "&#93;", "]", -1)
	return text

}

//decodeValue 解码CQ里的值
func decodeValue(text string) string {
	text = decodeText(text)
	text = strings.Replace(text, "&#44;", ",", -1)
	return text
}
