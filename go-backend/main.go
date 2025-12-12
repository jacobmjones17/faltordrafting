package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
	"gopkg.in/gomail.v2"
)

type ContactForm struct {
	Name        string `json:"name"`
	Email       string `json:"email"`
	Phone       string `json:"phone"`
	Location    string `json:"location"`
	ProjectType string `json:"projectType"`
	Timeline    string `json:"timeline"`
	Budget      string `json:"budget"`
	Message     string `json:"message"`
}

type EmailConfig struct {
	SMTPHost     string
	SMTPPort     int
	SMTPUsername string
	SMTPPassword string
	ToEmail      string
}

var emailConfig = EmailConfig{
	SMTPHost:     "smtp.gmail.com",
	SMTPPort:     587,
	SMTPUsername: os.Getenv("SMTP_USERNAME"), // From environment variable
	SMTPPassword: os.Getenv("SMTP_PASSWORD"), // From environment variable  
	ToEmail:      "faltorco.ben@gmail.com",   // Form submissions go to your BROTHER'S email
}

func main() {
	r := mux.NewRouter()

	// API routes
	api := r.PathPrefix("/api").Subrouter()
	api.HandleFunc("/contact", handleContactForm).Methods("POST")
	api.HandleFunc("/health", handleHealth).Methods("GET")

	// CORS configuration
	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:3000", "http://localhost:5173", "https://*.netlify.app", "https://faltorconsulting.netlify.app"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"*"},
		AllowCredentials: true,
	})

	handler := c.Handler(r)

	// Get port from environment variable (Render requirement) or default to 8080
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	fmt.Printf("🚀 Server starting on port %s\n", port)
	log.Fatal(http.ListenAndServe(":"+port, handler))
}

func handleHealth(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{
		"status": "healthy",
		"time":   time.Now().Format(time.RFC3339),
	})
}

func handleContactForm(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// Parse multipart form
	err := r.ParseMultipartForm(10 << 20) // 10 MB limit
	if err != nil {
		http.Error(w, "Error parsing form", http.StatusBadRequest)
		return
	}

	// Extract form data
	form := ContactForm{
		Name:        r.FormValue("name"),
		Email:       r.FormValue("email"),
		Phone:       r.FormValue("phone"),
		Location:    r.FormValue("location"),
		ProjectType: r.FormValue("projectType"),
		Timeline:    r.FormValue("timeline"),
		Budget:      r.FormValue("budget"),
		Message:     r.FormValue("message"),
	}

	// Validate required fields
	if form.Name == "" || form.Email == "" || form.ProjectType == "" || form.Message == "" {
		http.Error(w, "Missing required fields", http.StatusBadRequest)
		return
	}

	// Handle file uploads
	var attachments []string
	if files, ok := r.MultipartForm.File["files"]; ok {
		for _, fileHeader := range files {
			file, err := fileHeader.Open()
			if err != nil {
				continue
			}
			defer file.Close()

			// Create uploads directory if it doesn't exist
			uploadsDir := "./uploads"
			if _, err := os.Stat(uploadsDir); os.IsNotExist(err) {
				os.Mkdir(uploadsDir, 0755)
			}

			// Save file
			filename := fmt.Sprintf("%d_%s", time.Now().Unix(), fileHeader.Filename)
			filepath := filepath.Join(uploadsDir, filename)
			
			outFile, err := os.Create(filepath)
			if err != nil {
				continue
			}
			defer outFile.Close()

			_, err = io.Copy(outFile, file)
			if err != nil {
				continue
			}

			attachments = append(attachments, filepath)
		}
	}

	// Send email
	err = sendEmail(form, attachments)
	if err != nil {
		log.Printf("Error sending email: %v", err)
		http.Error(w, "Error sending email", http.StatusInternalServerError)
		return
	}

	// Clean up uploaded files after sending email
	for _, filepath := range attachments {
		os.Remove(filepath)
	}

	// Return success response
	response := map[string]interface{}{
		"success": true,
		"message": "Form submitted successfully",
	}
	
	json.NewEncoder(w).Encode(response)
}

func sendEmail(form ContactForm, attachments []string) error {
	m := gomail.NewMessage()
	
	// Set email headers
	m.SetHeader("From", emailConfig.SMTPUsername)
	m.SetHeader("To", emailConfig.ToEmail)
	m.SetHeader("Subject", fmt.Sprintf("New Project Inquiry from %s", form.Name))
	
	// Create email body
	body := fmt.Sprintf(`
New project inquiry received from the FALTOR Consulting website.

Contact Information:
- Name: %s
- Email: %s
- Phone: %s
- Location: %s

Project Details:
- Project Type: %s
- Timeline: %s
- Budget: %s

Message:
%s

---
This email was automatically generated from the contact form at faltordrafting.com
`, form.Name, form.Email, form.Phone, form.Location, form.ProjectType, form.Timeline, form.Budget, form.Message)

	m.SetBody("text/plain", body)
	
	// Add attachments
	for _, filepath := range attachments {
		m.Attach(filepath)
	}

	// Configure SMTP dialer
	d := gomail.NewDialer(emailConfig.SMTPHost, emailConfig.SMTPPort, emailConfig.SMTPUsername, emailConfig.SMTPPassword)
	
	// Send email
	return d.DialAndSend(m)
}