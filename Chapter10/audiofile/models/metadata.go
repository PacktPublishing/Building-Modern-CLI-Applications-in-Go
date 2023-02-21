package models

type Metadata struct {
	Tags       Tags   `json:"tags",header:"Tags"`
	Transcript string `json:"transcript",header:"Transcript"`
}
