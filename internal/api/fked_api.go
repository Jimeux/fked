package api

import (
	"fmt"
	"log"
	"net/http"
	"net/http/httputil"
	"regexp"
	"strings"

	"github.com/Jimeux/fked/internal/domain/fked"
	"github.com/Jimeux/fked/internal/domain/reaction"
	"github.com/Jimeux/fked/internal/domain/user"
	"github.com/Jimeux/fked/internal/infra/slack"
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
)

var feelPattern = regexp.MustCompile("(?i)i feel :(\\w+): ?")

type FkedAPI struct {
	slackClient *slack.Client
	fkedService fked.Service
}

func NewSlackAPI(slackClient *slack.Client, svc fked.Service) *FkedAPI {
	return &FkedAPI{slackClient, svc}
}

func (s *FkedAPI) Event(c *gin.Context) {
	var req slack.EventRequest
	if err := c.Bind(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	switch req.Type {
	case slack.URLVerificationRequestType:
		c.JSON(http.StatusOK, slack.ChallengeResponse{Challenge: req.Challenge})
		return
	case slack.EventCallbackRequestType:
		s.handleEvent(c, &req.Event)
		c.Status(http.StatusOK)
		return
	default:
		c.Status(http.StatusNotFound)
	}
}

func (s *FkedAPI) handleEvent(c *gin.Context, event *slack.Event) {
	switch event.Type {
	case slack.AppMentionEvent:
		code, err := extractReactionCode(event.Text)
		if err != nil {
			log.Println(err)
			// TODO 28/11/2018 @Jimeux Send err in message to user
		}
		if err := s.fkedService.UpdateFked(user.ID(event.User), code); err != nil {
			log.Println(err)
			// TODO 28/11/2018 @Jimeux Send err in message to user
		}
		if err := s.slackClient.SendMessage(event.Channel, "Okay, cool. Be lucky!"); err != nil {
			log.Println(err)
			// TODO 28/11/2018 @Jimeux Retry?
		}
		c.Status(http.StatusOK)
	default:
		c.Status(http.StatusNotFound)
	}
}

func extractReactionCode(text string) (reaction.Code, error) {
	t := strings.ToLower(strings.TrimSpace(text))
	match := feelPattern.FindStringSubmatch(t)

	if len(match) != 2 {
		return "", errors.New("usages: `I feel :emoji:`")
	}
	return reaction.Code(match[1]), nil
}

func debug(req *http.Request) {
	bytes, _ := httputil.DumpRequest(req, true)
	fmt.Println(string(bytes))
}
