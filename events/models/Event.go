package models

type Event struct {
	UUID        string `json:"uuid"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Date        string `json:"date"` //RFC3339 2023-06-19T10:13:58+01:00
}

var EventsTests = []Event{
	{
		UUID:        "1",
		Title:       "Conference 2025",
		Description: "Annual tech conference",
		Date:        "31.05.2025 10:00",
	},
	{
		UUID:        "2",
		Title:       "Music Festival",
		Description: "Outdoor music festival",
		Date:        "15.06.2025 18:00",
	},
	{
		UUID:        "3",
		Title:       "Art Exhibition",
		Description: "Modern art gallery opening",
		Date:        "20.06.2025 14:00",
	},
	{
		UUID:        "4",
		Title:       "Startup Pitch",
		Description: "Pitch event for startups",
		Date:        "25.06.2025 09:00",
	},
	{
		UUID:        "5",
		Title:       "Coding Workshop",
		Description: "Hands-on coding session",
		Date:        "30.06.2025 16:00",
	},
	{
		UUID:        "6",
		Title:       "Book Launch",
		Description: "Launch event for a new book",
		Date:        "05.07.2025 11:00",
	},
	{
		UUID:        "7",
		Title:       "Charity Run",
		Description: "5K run for charity",
		Date:        "10.07.2025 07:00",
	},
	{
		UUID:        "8",
		Title:       "Tech Meetup",
		Description: "Networking event for tech enthusiasts",
		Date:        "15.07.2025 19:00",
	},
	{
		UUID:        "9",
		Title:       "Film Screening",
		Description: "Premiere of an indie film",
		Date:        "20.07.2025 20:00",
	},
	{
		UUID:        "10",
		Title:       "Cooking Class",
		Description: "Learn to cook Italian cuisine",
		Date:        "25.07.2025 17:00",
	},
}
