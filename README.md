# go-mylar-client

Go Client SDK for Mylar Comicbook Management Application

## Usage

```go
client, err := mylar.New(mylarURL, apiKey)

// List comics in the index
comics, err := client.GetIndex()
/* Returns
[]mylar.Comic{
	{
		Status:      "Active",
		Publisher:   "Les Humanoïdes Associés",
		Name:        "After the Incal",
		Total:       1,
		ImageURL:    "https://comicvine1.cbsistatic.com/uploads/scale_large/6/67663/4361217-01.jpg",
		DetailsURL:  "https://comicvine.gamespot.com/after-the-incal/4050-79767/",
		LatestIssue: "1",
		Year:        "2015",
		ID:          "79767",
	},
    ...
}
*/


// Get a single comic's details
comicDetails, err := client.Getcomic("123")
/* Returns
mylar.ComicDetails{
    Comic:   []Comic{...}
    Annuals: []Issue{...}
    Issues:  []Issue{...}
}
*/

// Get wanted issues
wantedIssues, err := client.GetWanted()

// Get history
history, err := client.GetHistory()
/*
```