package main

func isUppercase(i byte) bool{
	return i >= 'A' && i<='Z'
}

func isDigit(i byte) bool{
	return i >='0' && i<='9'
}

func isValidPassword(password string) bool {
	hasUpper, hasDigit := false, false
	if len(password)<=12 && len(password)>=5 {
		for _, i := range password {
			if hasDigit && hasUpper {
				break // if we found the upper and lower we can stop.
			}
			if isUppercase(byte(i)){
				hasUpper  = true
			}
			if isDigit(byte(i)){
				hasDigit = true
			}
		}
	}
	return hasUpper && hasDigit
}
