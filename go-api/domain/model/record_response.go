package model

type Record struct {
	Id        int    `json:"id"`
	Item      string `json:"item"`
	Content   string `json:"content"`
	StudyDate string `json:"study_date"`
}
