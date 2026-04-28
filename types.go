package main

type GeminiRequest struct {
	Contents []Content `json:"contents"`
}

type GeminiResponse struct {
	Candidates []Candidate `json:"candidates"`
}

type Candidate struct {
	Content Content `json:"content"`
}

type Content struct {
	Parts []Part `json:"parts"`
}

type Part struct {
	Text string `json:"text"`
}

func BuildRequest(query string) GeminiRequest {
	return GeminiRequest{
		Contents: []Content{
			{
				Parts: []Part{
					{Text: query},
				},
			},
		},
	}
}
