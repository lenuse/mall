package utils

// CheckErr 验证err
func CheckErr(err error) {
	if err != nil {
		panic(err)
	}
}
