package model

type GetUploadURLRequest struct {
	File        string `json:"file,omitempty"`
	ContentType string `json:"contentType,omitempty"`
	Width       int64  `json:"width,omitempty"`
	Height      int64  `json:"height,omitempty"`
}

type GetUploadURLResponse struct {
	URL     string            `json:"url,omitempty"`
	Headers map[string]string `json:"headers,omitempty"`
}

type GetDownloadURLRequest struct {
	File string `json:"file,omitempty"`
}

type GetDownloadURLResponse struct {
	URL string `json:"url,omitempty"`
}
