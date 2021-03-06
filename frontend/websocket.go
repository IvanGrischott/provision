package frontend

import (
	"bytes"
	"encoding/json"
	"strings"
	"sync"

	"github.com/digitalrebar/provision/models"
	"github.com/gin-gonic/gin"
	"gopkg.in/olahol/melody.v1"
)

var wsLock = &sync.Mutex{}

func (fe *Frontend) InitWebSocket() {
	fe.melody = melody.New()

	fe.ApiGroup.GET("/ws", func(c *gin.Context) {
		fe.melody.HandleRequest(c.Writer, c.Request)
	})

	fe.melody.HandleMessage(fe.websocketHandler)
}

// Callers register or deregister values.
// type.action.key = with * as wildcard spots.

func filterFunction(s *melody.Session, e *models.Event) bool {
	val, ok := s.Get("EventMap")
	if !ok {
		return false
	}
	emap := val.([]string)
	for _, test := range emap {
		arr := strings.SplitN(test, ".", 3)

		if arr[0] != "*" && arr[0] != e.Type {
			continue
		}
		if arr[1] != "*" && arr[1] != e.Action {
			continue
		}
		if arr[2] != "*" && arr[2] != e.Key {
			continue
		}
		return true
	}
	return false
}

func (f *Frontend) Publish(e *models.Event) error {
	if msg, err := json.Marshal(e); err != nil {
		return err
	} else {
		return f.melody.BroadcastFilter(msg, func(s *melody.Session) bool {
			wsLock.Lock()
			defer wsLock.Unlock()
			return filterFunction(s, e)
		})
	}
}

// This never gets unloaded.
func (f *Frontend) Reserve() error {
	return nil
}
func (f *Frontend) Release() {}
func (f *Frontend) Unload()  {}

func (f *Frontend) websocketHandler(s *melody.Session, buf []byte) {
	splitMsg := bytes.SplitN(bytes.TrimSpace(buf), []byte(" "), 2)
	if len(splitMsg) != 2 {
		f.Logger.Printf("WS: Unknown: Received message: %s\n", string(buf))
		return
	}
	prefix, msg := string(splitMsg[0]), string(splitMsg[1])
	if !(prefix == "register" || prefix == "deregister") {
		f.Logger.Printf("WS: Invalid msg prefix %s", prefix)
		return
	}
	wsLock.Lock()
	defer wsLock.Unlock()
	val, ok := s.Get("EventMap")
	if !ok {
		val = []string{}
	}
	emap := val.([]string)
	switch prefix {
	case "register":
		for _, test := range emap {
			if test == msg {
				return
			}
		}
		emap = append(emap, msg)
	case "deregister":
		res := make([]string, 0, len(emap))
		for _, test := range emap {
			if test == msg {
				continue
			}
		}
		emap = res
	}
	s.Set("EventMap", emap)
}
