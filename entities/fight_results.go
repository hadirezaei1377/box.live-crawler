package entities

type Fight_Results struct {
	Date             interface{}
	Participant1Name string
	Participant2Name string
	FightInfo        interface{}
	EventPlace       Event_Place
	UnderCard_Events interface{}
}

type Event_Place struct {
	Stadium_Name string
	Country_Name string
}
