package dtos

type StatusMsg struct {
	Status  int    `json:"status"`
	Message string `json:"message,omitempty"`
}

var (
	Created StatusMsg = StatusMsg{
		Status:  201,
		Message: "Created",
	}
	Success StatusMsg = StatusMsg{
		Status:  200,
		Message: "Success",
	}
	Accepted StatusMsg = StatusMsg{
		Status:  202,
		Message: "Accepted",
	}
	NotFound StatusMsg = StatusMsg{
		Status:  404,
		Message: "Not Found",
	}
	BadRequest StatusMsg = StatusMsg{
		Status:  400,
		Message: "Bad Request",
	}
	Forbidden StatusMsg = StatusMsg{
		Status:  403,
		Message: "Forbidden",
	}
	ServerError StatusMsg = StatusMsg{
		Status:  500,
		Message: "Internal Server Error",
	}
)
