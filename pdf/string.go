package pdf

type String string

func (s String) String() string {
	return "(" + string(s) + ")"
}
func (s String) ToUtf8() string {
	if len([]byte(s)) == 2 {

		uts8, err := decoder.Bytes([]byte(s))
		if err != nil {
			return string(s)
		}

		return string(uts8)
	}
	return string(s)
}
