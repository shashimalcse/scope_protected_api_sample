package tweet

import (
	"context"
	"errors"
	"sample_api/internal/entity"
	"strconv"
)

var (
	tweets = map[string]interface{}{}
)

type Service interface {
	Get(ctx context.Context, id string) (Tweet, error)
	Query(ctx context.Context) ([]Tweet, error)
	Create(ctx context.Context, req CreateTweetRequest) (Tweet, error)
	Delete(ctx context.Context, id string) error
}

type Tweet struct {
	entity.Tweet
}

type CreateTweetRequest struct {
	Text string `json:"text"`
	User string `json:"user"`
}

type service struct {
}

func NewService() Service {
	return service{}
}

func (s service) Create(ctx context.Context, req CreateTweetRequest) (Tweet, error) {

	if req.Text == "" {
		return Tweet{}, errors.New("Invalid text")
	}
	if req.User == "" {
		return Tweet{}, errors.New("Invalid user")
	}
	id := strconv.Itoa(len(tweets) + 1)
	tweet := entity.Tweet{ID: id, Text: req.Text, User: req.User}
	tweets[id] = tweet
	return s.Get(ctx, id)
}

func (s service) Get(ctx context.Context, id string) (Tweet, error) {

	value, ok := tweets[id]
	if !ok {
		return Tweet{}, errors.New("Tweet not found")
	}
	tweet, ok := value.(entity.Tweet)
	if !ok {
		return Tweet{}, errors.New("Invalid tweet id")
	}
	return Tweet{tweet}, nil
}

func (s service) Query(ctx context.Context) ([]Tweet, error) {

	var result []Tweet
	for _, item := range tweets {
		if tweet, ok := item.(entity.Tweet); ok {
			result = append(result, Tweet{tweet})
		}
	}
	return result, nil
}

func (s service) Delete(ctx context.Context, id string) error {

	value, ok := tweets[id]
	if !ok {
		return errors.New("Tweet not found")
	}
	_, ok = value.(entity.Tweet)
	if !ok {
		return errors.New("Invalid tweet id")
	}

	delete(tweets, id)
	return nil
}
