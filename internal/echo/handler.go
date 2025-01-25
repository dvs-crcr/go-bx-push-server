package echo

import (
    "log"
    "net/http"
)

type Handler struct {
    echoService *Service
}

func NewHandler(echoService *Service) *Handler {
    handler := &Handler{
        echoService: echoService,
    }

    return handler
}

func (h *Handler) Echo(w http.ResponseWriter, r *http.Request) {
    // TODO: REFAC. src - https://github.com/gorilla/websocket/blob/main/examples/echo/server.go
    c, err := h.echoService.upgrader.Upgrade(w, r, nil)
    if err != nil {
        log.Print("upgrade:", err)
        return
    }
    defer c.Close()

    for {
        mt, message, err := c.ReadMessage()
        if err != nil {
            log.Println("read:", err)
            break
        }

        log.Printf("recv: %s", message)
        if err = c.WriteMessage(mt, message); err != nil {
            log.Println("write:", err)
            break
        }
    }
}
