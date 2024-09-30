package design

import (
	"container/heap"
)

// https://leetcode.com/problems/design-twitter/

type tweet struct {
	id       int
	uniqueId int
}

type Twitter struct {
	tweets    map[int][]*tweet
	followees map[int]map[int]bool
	nextId    int
}

func NewTwitter() Twitter {
	return Twitter{
		tweets:    make(map[int][]*tweet),
		followees: make(map[int]map[int]bool),
		nextId:    1,
	}
}

func (t *Twitter) PostTweet(userId int, tweetId int) { // O(1)
	var newTweet = &tweet{id: tweetId, uniqueId: t.nextId}
	t.nextId++
	if t.tweets[userId] == nil {
		t.tweets[userId] = []*tweet{newTweet}
	} else {
		t.tweets[userId] = append(t.tweets[userId], newTweet)
	}
}

func (t *Twitter) GetNewsFeed(userId int) []int { // O(N Log N)
	var posts = make([]*tweet, len(t.tweets[userId]))
	for followeeId := range t.followees[userId] {
		posts = append(posts, t.tweets[followeeId]...)
	}
	copy(posts, t.tweets[userId])
	var tweets = NewHeap(posts, func(x, y *tweet) bool { return x.uniqueId > y.uniqueId })
	heap.Init(tweets)
	var result = make([]int, 0)
	var count = 10
	for tweets.Len() > 0 && count > 0 {
		count--
		tweet := heap.Pop(tweets).(*tweet)
		result = append(result, tweet.id)
	}
	return result
}

func (t *Twitter) Follow(followerId int, followeeId int) { // O(1)
	if t.followees[followerId] == nil {
		t.followees[followerId] = make(map[int]bool)
	}
	t.followees[followerId][followeeId] = true
}

func (t *Twitter) Unfollow(followerId int, followeeId int) { // O(1)
	if t.followees[followerId] != nil && t.followees[followerId][followeeId] {
		delete(t.followees[followerId], followeeId)
	}
}

/**
 * Your Twitter object will be instantiated and called as such:
 * obj := Constructor();
 * obj.PostTweet(userId,tweetId);
 * param_2 := obj.GetNewsFeed(userId);
 * obj.Follow(followerId,followeeId);
 * obj.Unfollow(followerId,followeeId);
 */
