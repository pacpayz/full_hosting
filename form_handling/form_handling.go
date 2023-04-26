package form_handling

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type basicData struct {
	Email          string `json:"email"`
	FirstName      string `json:"firstName"`
	LastName       string `json:"lastName"`
	PhoneNumber    string `json:"phoneNumber"`
	BillingAddress string `json:"billingAddress"`
}

type selectPaymentData struct {
	Crypto string `json:"crypto"`
	Card   string `json:"card"`
	Other  string `json:"other"`
}

type cardData struct {
	CardNumber string `json:"cardNumber"`
	CardHolder string `json:"cardHolder"`
	ExpiryDate string `json:"expiryDate"`
	CVV        string `json:"cvv"`
}

type cryptoData struct {
	TXAddress string `json:"cryptoAddress"`
}

type otherData struct {
	PayPal    bool `json:"paypal"`
	ApplePay  bool `json:"applepay"`
	GooglePay bool `json:"googlepay"`
	Venmo     bool `json:"venmo"`
}

func HandleBasicForm(w http.ResponseWriter, r *http.Request) {
	// parse form data
	if err := r.ParseForm(); err != nil {
		http.Error(w, "Failed to parse form data", http.StatusBadRequest)
		return
	}

	// extract form values
	data := basicData{
		Email:          r.FormValue("email"),
		FirstName:      r.FormValue("firstName"),
		LastName:       r.FormValue("lastName"),
		PhoneNumber:    r.FormValue("phoneNumber"),
		BillingAddress: r.FormValue("billingAddress"),
	}

	// write form data to console
	fmt.Printf("Basic Form Data: %+v\n", data)

	// send success response
	fmt.Fprintf(w, "Form data received")
}

func HandleSelectPaymentForm(w http.ResponseWriter, r *http.Request) {
	// parse form data
	if err := r.ParseForm(); err != nil {
		http.Error(w, "Failed to parse form data", http.StatusBadRequest)
		return
	}

	// extract form values
	data := selectPaymentData{
		Crypto: r.FormValue("crypto"),
		Card:   r.FormValue("card"),
		Other:  r.FormValue("other"),
	}

	// write form data to console
	fmt.Printf("Select Payment Form Data: %+v\n", data)

	// send success response
	fmt.Fprintf(w, "Form data received")
}

func HandleCardForm(w http.ResponseWriter, r *http.Request) {
	// read request body
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Failed to read request body", http.StatusBadRequest)
		return
	}

	// parse json data
	var data cardData
	if err := json.Unmarshal(body, &data); err != nil {
		http.Error(w, "Failed to parse json data", http.StatusBadRequest)
		return
	}

	// write form data to console
	fmt.Printf("Card Form Data: %+v\n", data)

	// send success response
	fmt.Fprintf(w, "Form data received")
}

func HandleCryptoForm(w http.ResponseWriter, r *http.Request) {
	// Parse form data
	err := r.ParseForm()
	if err != nil {
		http.Error(w, "Could not parse form data", http.StatusBadRequest)
		return
	}

	// Extract crypto data from form
	cryptoData := cryptoData{
		TXAddress: r.Form.Get("tx-address"),
	}

	// Do something with crypto data
	fmt.Printf("Crypto data received: %+v\n", cryptoData)

	// Respond with success message
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Crypto data received")
}

func HandleOtherForm(w http.ResponseWriter, r *http.Request) {
	// Parse the form data
	err := r.ParseForm()
	if err != nil {
		http.Error(w, "Error parsing form data", http.StatusBadRequest)
		return
	}

	// Get the form values
	paypal := r.Form.Get("paypal") == "on"
	applepay := r.Form.Get("applepay") == "on"
	googlepay := r.Form.Get("googlepay") == "on"
	venmo := r.Form.Get("venmo") == "on"

	// Process the form data
	if paypal {
		log.Println("Using Paypal...")
	} else if applepay {
		log.Println("Using Apple Pay...")
	} else if googlepay {
		log.Println("Using Google Pay...")
	} else if venmo {
		log.Println("Using Venmo...")
	} else {
		log.Println("No payment method selected")
	}

	// Send the response
	w.Write([]byte("Other form submitted successfully"))
}

func HandleProcessingForm(w http.ResponseWriter, r *http.Request) {
	// Parse form data
	err := r.ParseForm()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Get form values
	orderID := r.FormValue("orderID")
	amount := r.FormValue("amount")
	paymentMethod := r.FormValue("paymentMethod")

	// Perform processing logic
	// ...

	// Respond to client
	fmt.Fprintf(w, "Processing form submitted with order ID: %s, amount: %s, and payment method: %s", orderID, amount, paymentMethod)
}
