package api

// Login 根据password和username验证如果成功返回
func Login(username, password string) error {
	if  username == "wang" && password == "123"{
		return nil 
	}
	return error
}