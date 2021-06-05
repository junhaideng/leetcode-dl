package code 

func GetTplSuffix(lang string) string{
	switch lang{
	case "Python3", "Python":
		return ".py"
	case "Go":
		return ".go"
	case "C++":
		return ".cpp"
	case "C":
		return ".c"
	case "Java":
		return ".java"
	case "JavaScript":
		return ".js"
	case "Rust":
		return ".rs"
	default:
		return ""
	}
}