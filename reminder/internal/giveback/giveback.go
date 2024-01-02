package giveback

// RobotResponse
// @Param meta.code: business level code(often self defined)
// @Param meta.message: business level message
// @Param data: api data
type RobotResponse struct {
	Meta struct {
		Code    int    `json:"code"`
		Message string `json:"message"`
	} `json:"meta"`
	Total *int64 `json:"total"`
	Data  any    `json:"data"`
}
