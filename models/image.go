package models

type Image struct {
	URL     string `bson:"url"`
	Title   string `bson:"title"`
	AltText string `bson:"alt_text"`
}
