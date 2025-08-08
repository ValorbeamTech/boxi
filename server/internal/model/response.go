package model

type APIResponse struct {
    Success bool        `json:"success"`
    Message string      `json:"message,omitempty"`
    Data    interface{} `json:"data,omitempty"`
    Error   string      `json:"error,omitempty"`
}

type PaginatedResponse struct {
    Data       interface{} `json:"data"`
    Page       int         `json:"page"`
    Limit      int         `json:"limit"`
    Total      int64       `json:"total"`
    TotalPages int         `json:"total_pages"`
}

func SuccessResponse(data interface{}, message string) APIResponse {
    return APIResponse{
        Success: true,
        Message: message,
        Data:    data,
    }
}

func ErrorResponse(err string) APIResponse {
    return APIResponse{
        Success: false,
        Error:   err,
    }
}