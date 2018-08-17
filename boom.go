package boom

import "encoding/json"

var statusCodes = map[int]string{
	100: "Continue",
	101: "Switching Protocols",
	102: "Processing",
	200: "OK",
	201: "Created",
	202: "Accepted",
	203: "Non-Authoritative Information",
	204: "No Content",
	205: "Reset Content",
	206: "Partial Content",
	207: "Multi-Status",
	300: "Multiple Choices",
	301: "Moved Permanently",
	302: "Moved Temporarily",
	303: "See Other",
	304: "Not Modified",
	305: "Use Proxy",
	307: "Temporary Redirect",
	400: "Bad Request",
	401: "Unauthorized",
	402: "Payment Required",
	403: "Forbidden",
	404: "Not Found",
	405: "Method Not Allowed",
	406: "Not Acceptable",
	407: "Proxy Authentication Required",
	408: "Request Time-out",
	409: "Conflict",
	410: "Gone",
	411: "Length Required",
	412: "Precondition Failed",
	413: "Request Entity Too Large",
	414: "Request-URI Too Large",
	415: "Unsupported Media Type",
	416: "Requested Range Not Satisfiable",
	417: "Expectation Failed",
	418: "I\"m a teapot",
	422: "Unprocessable Entity",
	423: "Locked",
	424: "Failed Dependency",
	425: "Unordered Collection",
	426: "Upgrade Required",
	428: "Precondition Required",
	429: "Too Many Requests",
	431: "Request Header Fields Too Large",
	451: "Unavailable For Legal Reasons",
	500: "Internal Server Error",
	501: "Not Implemented",
	502: "Bad Gateway",
	503: "Service Unavailable",
	504: "Gateway Time-out",
	505: "HTTP Version Not Supported",
	506: "Variant Also Negotiates",
	507: "Insufficient Storage",
	509: "Bandwidth Limit Exceeded",
	510: "Not Extended",
	511: "Network Authentication Required",
}

// Output should be used to present a boom error to a user over
// HTTP.
type Output struct {
	StatusCode int                    `json:"status_code"`
	Error      string                 `json:"error"`
	Message    string                 `json:"message"`
	Data       map[string]interface{} `json:"data"`
}

// Error ..
type Error struct {
	Err        error
	IsServer   bool
	StatusCode int
	Output     Output
}

// MarshalJSON implements Marshaler and returns only e.Output.
func (e Error) MarshalJSON() ([]byte, error) {
	return json.Marshal(e.Output)
}

// Error returns the internal error.
func (e Error) Error() string {
	if e.Err != nil {
		return e.Err.Error()
	}
	if e.Output.Message == "" {
		return e.Output.Error
	}
	return e.Output.Message
}

func create(statusCode int, message string, data map[string]interface{}, err error) error {
	berr := Error{
		Err:        err,
		StatusCode: statusCode,
		IsServer:   statusCode >= 500,
		Output: Output{
			Data:       make(map[string]interface{}),
			StatusCode: statusCode,
			Error:      statusCodes[statusCode],
			Message:    message,
		},
	}
	if data != nil {
		berr.Output.Data = data
	}
	if berr.Output.Error == "" {
		berr.Output.Error = "Unknown"
	}
	return berr
}

// BadImplementation returns a 500 Internal Server Error.
func BadImplementation(err error) error {
	return create(500, "An internal server error occurred", nil, err)
}

// BadRequest returns 400 Bad Request.
func BadRequest(message string, data map[string]interface{}) error {
	return create(400, message, data, nil)
}

// Unauthorized returns 401 Unauthorized.
func Unauthorized(message string) error {
	return create(401, message, nil, nil)
}

// Forbidden returns 403 Forbidden.
func Forbidden(message string) error {
	return create(403, message, nil, nil)
}

// NotFound returns 404 Not Found.
func NotFound(message string) error {
	return create(404, message, nil, nil)
}

// RangeNotSatisfiable returns 416 Range Not Satisfiable.
func RangeNotSatisfiable(message string, data map[string]interface{}) error {
	return create(416, message, data, nil)
}
