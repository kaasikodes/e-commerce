package utils

import "net/smtp"

func SendMail(to []string, subject string, body string) error {
	err := smtp.SendMail("smtp.mailtrap.io:2525", smtp.PlainAuth("", "6c53d765680ca4", "83175273732073", "smtp.mailtrap.io"), "hello@caasimedia.com", to, []byte("Subject: "+subject+"\r\n\r\n"+body))
	if err != nil {
		return err
		
	}
	
	return nil
}