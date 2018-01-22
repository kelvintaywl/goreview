package handler

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/kelvintaywl/goreview/domain"
	er "github.com/kelvintaywl/goreview/domain/errors"
	"github.com/kelvintaywl/goreview/service"
)

const (
	PullRequestOpened string = "opened"
)

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("FIX ME"))
}

func HookHandler(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	b, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	var p domain.PullRequestEventPayload
	err = json.Unmarshal(b, &p)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	switch p.Action {
	case PullRequestOpened:
		cfgPtr, err := service.FetchConfig(ctx, p)
		if err != nil {
			if err == er.ErrJSONParseFailed {
				w.WriteHeader(http.StatusInternalServerError)
				w.Write([]byte("ERROR"))
				return
			}
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("ERROR"))
			return
		}
		cfg := *cfgPtr
		go (func() {
			err := service.AssignReviewers(ctx, p, cfg)
			if err != nil {
				// TODO: better err handling; delegate to service?
				fmt.Printf("assign_reviewers err: %s", err.Error())
			}
		})()
		break
	}
	w.WriteHeader(http.StatusAccepted)
	w.Write([]byte("OK"))
}
