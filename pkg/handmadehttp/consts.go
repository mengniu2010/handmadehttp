package handmadehttp

type Method string

var (
	ReqMethodGet Method = "get"
)

const (
	MaxChanSize = 1000
)

const (
	DefaultProtocal    = "HTTP/1.1"
	DefaultContentType = "text/html; charset=UTF-8"
	BuffSize           = 1024
	MaxHeaderLine      = 1024

	KeyContentType   = "Content-Type"
	KeyContentLength = "Content-Length"

	LineBreak = "\r\n"
)

var CodeToStatus = map[int]string{
	100: "Continue",
	101: "Switching Protocols",
	200: "OK",
	201: "Created",
	202: "Accepted",
	203: "Non-Authoritative Information",
	204: "No Content",
	205: "Reset Content",
	206: "Partial Content",
	300: "Multiple Choices",
	301: "Moved Permanently",
	302: "Found",
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
	408: "Request Time-Out",
	409: "Conflict",
	410: "Gone",
	411: "Length Required",
	412: "Precondition Failed",
	413: "Request Entity Too Large",
	414: "Request-Uri Too Large",
	415: "Unsupported Media Type",
	416: "Requested Range Not Satisfiable",
	417: "Expectation Failed",
	500: "Internal Server Error",
	501: "Not Implemented",
	502: "Bad Gateway",
	503: "Service Unavailable",
	504: "Gateway Time-Out",
	505: "Http Version Not Supported",
}
