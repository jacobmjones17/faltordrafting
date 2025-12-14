package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/smtp"
	"os"
)

type ContactFormData struct {
	Name        string `json:"name"`
	Email       string `json:"email"`
	Phone       string `json:"phone"`
	Location    string `json:"location"`
	ProjectType string `json:"projectType"`
	Timeline    string `json:"timeline"`
	Budget      string `json:"budget"`
	Message     string `json:"message"`
}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	http.HandleFunc("/api/contact", handleContact)
	http.HandleFunc("/api/health", handleHealth)

	log.Printf("Server starting on port %s", port)
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		log.Fatal(err)
	}
}

func handleHealth(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"status": "ok"})
}

func handleContact(w http.ResponseWriter, r *http.Request) {
	// Set CORS headers
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	if r.Method == "OPTIONS" {
		w.WriteHeader(http.StatusOK)
		return
	}

	if r.Method != "POST" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var formData ContactFormData
	if err := json.NewDecoder(r.Body).Decode(&formData); err != nil {
		log.Printf("Error decoding form data: %v", err)
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	// Validate required fields
	if formData.Name == "" || formData.Email == "" || formData.Message == "" {
		http.Error(w, "Name, email, and message are required", http.StatusBadRequest)
		return
	}

	// Send email
	if err := sendEmail(formData); err != nil {
		log.Printf("Error sending email: %v", err)
		http.Error(w, "Failed to send email", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{
		"status":  "success",
		"message": "Your message has been sent successfully!",
	})
}

func sendEmail(data ContactFormData) error {
	// Email configuration from environment variables
	smtpHost := os.Getenv("SMTP_HOST")     // e.g., smtp.gmail.com
	smtpPort := os.Getenv("SMTP_PORT")     // e.g., 587
	smtpUser := os.Getenv("SMTP_USER")     // Your email
	smtpPass := os.Getenv("SMTP_PASSWORD") // Your email password or app password
	toEmail := os.Getenv("TO_EMAIL")       // Email to receive contact form submissions

	if smtpHost == "" || smtpPort == "" || smtpUser == "" || smtpPass == "" || toEmail == "" {
		return fmt.Errorf("email configuration missing")
	}

	// Compose email
	subject := fmt.Sprintf("New Contact Form Submission from %s", data.Name)
	body := fmt.Sprintf(`
New contact form submission from FALTOR Consulting website:

Name: %s
Email: %s
Phone: %s
Location: %s
Project Type: %s
Timeline: %s
Budget: %s

Message:
%s

---
This email was sent from the FALTOR Consulting contact form.
`, data.Name, data.Email, data.Phone, data.Location, data.ProjectType, data.Timeline, data.Budget, data.Message)

	// Build email message
	message := []byte(fmt.Sprintf(
		"From: %s\r\n"+
			"To: %s\r\n"+
			"Subject: %s\r\n"+
			"\r\n"+
			"%s\r\n",
		smtpUser, toEmail, subject, body,
	))

	// SMTP authentication
	auth := smtp.PlainAuth("", smtpUser, smtpPass, smtpHost)

	// Send email
	err := smtp.SendMail(
		smtpHost+":"+smtpPort,
		auth,
		smtpUser,
		[]string{toEmail},
		message,
	)

	if err != nil {
		return fmt.Errorf("failed to send email: %v", err)
	}

	log.Printf("Email sent successfully to %s", toEmail)
	return nil
}
