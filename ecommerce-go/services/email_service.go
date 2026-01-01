package services

import "log"

func SendOrderConfirmation(email string, orderID uint) {
	log.Println("📧 Sending email to:", email)
	log.Println("✅ Order confirmed! Order ID:", orderID)
}
