package trace

type Tag struct {
	Key string
	Value interface{}
}

// TagString new string tag.
func TagString(key string, val string) Tag {
	return Tag{Key: key, Value: val}
}

type LogFiled struct {
	Key string
	Value string
}

func Log(key string, val string) LogFiled {
	return LogFiled{Key: key, Value: val}
}
