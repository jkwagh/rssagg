package main

import (
	"time"

	"github.com/google/uuid"
	"github.com/jkwagh/rssagg/internal/database"
)

//want to be able to configure the struct without touching the sqlc generated models.go file

type User struct {
	ID        uuid.UUID `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Name      string    `json:"name"`
	APIKey    string    `json:"api_key"`
}

func databaseUserToUser(dbUser database.User) User {
	return User{
		ID:        dbUser.ID,
		CreatedAt: dbUser.CreatedAt,
		UpdatedAt: dbUser.UpdatedAt,
		Name:      dbUser.Name,
		APIKey:    dbUser.ApiKey,
	}
}

type Feed struct {
	ID 			uuid.UUID	`json:"id"`
	CreatedAt 	time.Time	`json:"created_at"`
	UpdatedAt 	time.Time	`json:"updated_at"`
	Name 		string		`json:"name"`
	URL 		string		`json:"url"`
	UserID 		uuid.UUID	`json:"user_id"`
}

func databaseFeedToFeed(dbFeed database.Feed) Feed {
	return Feed{
		ID: dbFeed.ID,
		CreatedAt: dbFeed.CreatedAt,
		UpdatedAt: dbFeed.UpdatedAt,
		Name: dbFeed.Name,
		URL: dbFeed.Url,
		UserID: dbFeed.UserID,
	}
}

func databaseFeedsToFeeds(dbFeeds []database.Feed) []Feed {
	feeds:= []Feed{}
	for _, dbFeed := range dbFeeds {
		feeds = append(feeds, databaseFeedToFeed(dbFeed))
	}     
	return feeds
}

type FeedFollow struct {
	ID 			uuid.UUID	`json:"id"`
	CreatedAt 	time.Time	`json:"created_at"`
	UpdatedAt 	time.Time	`json:"updated_at"`
	UserID 		uuid.UUID	`json:"user_id"`
	FeedID 		uuid.UUID	`json:"feed_id"`
}

func databaseFeedFollowToFeedFollow(dbFeedFollow database.FeedFollow) FeedFollow {
	return FeedFollow{
		ID: 		dbFeedFollow.ID,
		CreatedAt: 	dbFeedFollow.CreatedAt,
		UpdatedAt: 	dbFeedFollow.UpdatedAt,
		UserID: 	dbFeedFollow.UserID,
		FeedID: 	dbFeedFollow.FeedID,
	}
}

func databaseFeedFollowsToFeedFollows(dbFeedFollows []database.FeedFollow) []FeedFollow {
	feedFollows := []FeedFollow{}
	for _, dbFeedFollow := range dbFeedFollows {
		feedFollows = append(feedFollows, databaseFeedFollowToFeedFollow(dbFeedFollow))
	}     
	return feedFollows
}