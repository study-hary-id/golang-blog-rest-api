package main

import "time"

var articles = payload{
	Data: []article{
		{
			Type: "article",
			ID:   "1",
			Attributes: attributes{
				Title:    "Finding simplicity in life",
				Body:     "Life can get complicated really quickly, but it doesn't have to be! There are many ways to simplify your life, a few of which we've explored in the past.This week we're taking a bit of an approach though, in how you can find simplicity in the life you already living.",
				Date:     time.Date(2019, time.July, 1, 0, 0, 0, 0, time.UTC),
				Featured: true,
			},
			Links: links{
				Self:  "http://localhost:8000/articles/1",
				Cover: "http://localhost:8000/static/img/1.jpg",
			},
		},
		{
			Type: "article",
			ID:   "2",
			Attributes: attributes{
				Title:    "Simple decorations",
				Body:     "A home isn't a home until you've decorated a little. People either don't decorate, or they go overboard and it doesn't have the impact they were hoping for. Staying simple will help draw the eye where you want it to and make things pop like never before.",
				Date:     time.Date(2019, time.July, 3, 0, 0, 0, 0, time.UTC),
				Featured: false,
			},
			Links: links{
				Self:  "http://localhost:8000/articles/2",
				Cover: "http://localhost:8000/static/img/2.jpg",
			},
		},
		{
			Type: "article",
			ID:   "3",
			Attributes: attributes{
				Title:    "Simplicity and work",
				Body:     "Work is often a major source of stress. People get frustrated, it ruins their relationship with others and it leads to burnout. By keeping your work life as simple as possible, it will help balance everything out.",
				Date:     time.Date(2019, time.July, 12, 0, 0, 0, 0, time.UTC),
				Featured: false,
			},
			Links: links{
				Self:  "http://localhost:8000/articles/3",
				Cover: "http://localhost:8000/static/img/3.jpg",
			},
		},
		{
			Type: "article",
			ID:   "4",
			Attributes: attributes{
				Title:    "Keeping cooking simple",
				Body:     "Food is a very important part of everyone's life. If you want to be healthy, you have to eat healthy. One of the easiest ways to do that is to keep your cooking nice and simple.",
				Date:     time.Date(2019, time.July, 19, 0, 0, 0, 0, time.UTC),
				Featured: false,
			},
			Links: links{
				Self:  "http://localhost:8000/articles/4",
				Cover: "http://localhost:8000/static/img/4.jpg",
			},
		},
	},
}
