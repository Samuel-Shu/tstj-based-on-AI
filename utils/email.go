package utils

import "gopkg.in/gomail.v2"

func Email(To string, captcha string) {
	m := gomail.NewMessage()
	m.SetHeader("From", "2440039602@qq.com")
	m.SetHeader("To", To)
	m.SetHeader("Subject", "登录验证码：")
	m.SetBody("text/html", "验证码是 <b>"+captcha+"</b>")
	//m.Attach("/home/Alex/lolcat.jpg")

	d := gomail.NewDialer("smtp.qq.com", 587, "2440039602@qq.com", "pcxctbkjmeevdjjd")

	// Send the emailFun to Bob, Cora and Dan.
	if err := d.DialAndSend(m); err != nil {
		panic(err)
	}
}
