package main

// API interface defines the methods required for an API wrapper
type NSEIndia struct {
	BaseURL string
}

//need to separate them into different interface file

// NewNSEIndia creates a new instance of NSEIndia
func NewNSEIndia() *NSEIndia {
	return &NSEIndia{
		BaseURL: "https://www.nseindia.com",
	}
}
