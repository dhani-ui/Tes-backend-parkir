package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

// 1. Tipe Data Payload dari React
type WebhookPayload struct {
	Phone   string `json:"phone"`
	Message string `json:"message"`
}

// 2. Kredensial Fonnte
// Nanti ganti dengan Token Fonnte asli Anda
const fonnteToken = "utzi75V5vjCTQcShFN2k"

// 3. Fungsi Handler untuk Endpoint /api/send-wa
func handleSendWA(w http.ResponseWriter, r *http.Request) {
	// Pengaturan CORS agar React (port 5173) bisa mengakses endpoint ini
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	// Handle preflight request dari browser
	if r.Method == http.MethodOptions {
		w.WriteHeader(http.StatusOK)
		return
	}

	// Pastikan hanya menerima method POST
	if r.Method != http.MethodPost {
		http.Error(w, "Metode tidak diizinkan", http.StatusMethodNotAllowed)
		return
	}

	// Parsing JSON dari React
	var payload WebhookPayload
	err := json.NewDecoder(r.Body).Decode(&payload)
	if err != nil {
		http.Error(w, "Format JSON tidak valid", http.StatusBadRequest)
		return
	}

	// Kirim ke Fonnte
	err = sendToFonnte(payload.Phone, payload.Message)
	if err != nil {
		log.Printf("Gagal mengirim WA ke %s: %v", payload.Phone, err)
		http.Error(w, "Gagal mengirim pesan", http.StatusInternalServerError)
		return
	}

	// Respon sukses ke React
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{
		"status": "success",
		"detail": "Pesan WA berhasil masuk antrean",
	})
	
	log.Printf("Pesan berhasil dikirim ke: %s", payload.Phone)
}

// 4. Fungsi Integrasi ke Fonnte API
func sendToFonnte(phone string, message string) error {
	fonnteURL := "https://api.fonnte.com/send"

	// Fonnte menerima data berupa JSON
	requestBody, err := json.Marshal(map[string]string{
		"target": phone,
		"message": message,
	})
	if err != nil {
		return err
	}

	req, err := http.NewRequest("POST", fonnteURL, bytes.NewBuffer(requestBody))
	if err != nil {
		return err
	}

	// Header wajib Fonnte
	req.Header.Set("Authorization", fonnteToken)
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// Cek respon dari Fonnte
	bodyBytes, _ := io.ReadAll(resp.Body)
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("Fonnte error: status %d, body: %s", resp.StatusCode, string(bodyBytes))
	}

	return nil
}

// 5. Fungsi Utama
func main() {
	http.HandleFunc("/api/send-wa", handleSendWA)

	port := ":8080"
	fmt.Printf("Backend CRM berjalan di http://localhost%s\n", port)
	fmt.Println("Tekan Ctrl+C untuk menghentikan.")

	// Jalankan server
	if err := http.ListenAndServe(port, nil); err != nil {
		log.Fatalf("Server gagal berjalan: %v", err)
	}
}
