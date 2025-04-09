package dto

type ToDoRequest struct {
    ID   int    `json:"id"`
    Nama string `json:"nama"`
    Hari string `json:"hari"`
    ToDo string `json:"todo"`
}

type ToDoData struct {
    ID   int    `json:"id"`
    Nama string `json:"nama"`
    Hari string `json:"hari"`
    ToDo string `json:"todo"`
}

type ToDoResponse struct {
    StatusCode int      `json:"status_code"`
    Message    string   `json:"message"`
    Data       ToDoData `json:"data"`
}

func (t *ToDoResponse) Error() string {
    return t.Message
}