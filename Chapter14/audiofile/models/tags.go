package models

type Tags struct {
	Title       string `json:"title",header:"Title"`
	Album       string `json:"album",header:"Album"`
	Artist      string `json:"artist",header:"Artist"`
	AlbumArtist string `json:"album_artist",header:"Album Artist"`
	Composer    string `json:"composer",header:"Composer"`
	Genre       string `json:"genre",header:"Genre"`
	Year        int    `json:"year",header:"Year"`
	Lyrics      string `json:"lyrics",header:"Lyrics"`
	Comment     string `json:"comment",header:"Comment"`
}
