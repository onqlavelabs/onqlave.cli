package common

type BaseErrorResponse struct {
	Error Error `json:"error"`
}

type Error struct {
	Code          int           `json:"code" yaml:"code"`                     // Code represent codes in response
	Status        string        `json:"status" yaml:"status"`                 // Status represent string value of code
	Message       string        `json:"message" yaml:"message"`               // Message represent detail message
	CorrelationID string        `json:"correlation_id" yaml:"correlation_id"` // The RequestId that's also set in the header
	Details       []interface{} `json:"details" yaml:"details"`               // Details is a list of details in any types in string
}
