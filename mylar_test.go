package mylar_test

import (
	"testing"

	"github.com/SemanticallyNull/golandreporter"
	"github.com/billtomturner/go-mylar-client"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"gopkg.in/h2non/gock.v1"
)

const (
	mylarURL = "http://localhost:8090"
	apiKey   = "testapikey"
)

func TestMylar(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecsWithCustomReporters(t, "Mylar Suite", []Reporter{
		golandreporter.NewGolandReporter(),
	})
}

var (
	mockIndex = map[string]interface{}{
		"success": true,
		"data": []mylar.Comic{
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
			{
				Status:      "Active",
				Publisher:   "Marvel",
				Name:        "Amazing Mary Jane",
				Total:       6,
				ImageURL:    "https://comicvine1.cbsistatic.com/uploads/scale_large/6/67663/7116848-01.jpg",
				DetailsURL:  "https://comicvine.gamespot.com/amazing-mary-jane/4050-122214/",
				LatestIssue: "6",
				Year:        "2019",
				ID:          "122214",
			},
		},
	}

	mockComicDetail = map[string]interface{}{
		"success": true,
		"data": map[string]interface{}{
			"comic": []mylar.Comic{
				{
					Status:      "Active",
					Publisher:   "Marvel",
					Name:        "Amazing Mary Jane",
					Total:       6,
					ImageURL:    "https://comicvine1.cbsistatic.com/uploads/scale_large/6/67663/7116848-01.jpg",
					DetailsURL:  "https://comicvine.gamespot.com/amazing-mary-jane/4050-122214/",
					LatestIssue: "6",
					Year:        "2019",
					ID:          "122214",
				},
			},
			"annuals": []mylar.Issue{},
			"issues": []mylar.Issue{
				{
					Status:      "Downloaded",
					ComicName:   "Amazing Mary Jane",
					Name:        "None",
					ImageURL:    "https://comicvine1.cbsistatic.com/uploads/scale_small/6/67663/7284327-06.jpg",
					Number:      "6",
					ReleaseDate: "2020-03-18",
					IssueDate:   "2020-05-01",
					ID:          "742362",
				},
				{
					Status:      "Snatched",
					ComicName:   "Amazing Mary Jane",
					Name:        "None",
					ImageURL:    "https://comicvine1.cbsistatic.com/uploads/scale_small/6/67663/7250738-05.jpg",
					Number:      "5",
					ReleaseDate: "2020-02-19",
					IssueDate:   "2020-04-01",
					ID:          "737748",
				},
			},
		},
	}

	mockWanted = `[
  {
    "Status": "Wanted",
    "ComicSize": null,
    "ComicName": "Hellblazer Special: Chas",
    "IssueID": "135206",
    "DigitalDate": "0000-00-00",
    "IssueDate": "2008-10-24",
    "ImageURL": "https://comicvine1.cbsistatic.com/uploads/scale_small/6/67663/2788640-02.jpg",
    "inCacheDIR": null,
    "IssueDate_Edit": null,
    "ImageURL_ALT": "https://comicvine1.cbsistatic.com/uploads/scale_medium/6/67663/2788640-02.jpg",
    "ReleaseDate": "0000-00-00",
    "ArtworkURL": null,
    "Issue_Number": "2",
    "Location": null,
    "Int_IssueNumber": 2000,
    "IssueName": "the knowledge, chapter two",
    "ComicID": "22530",
    "Type": null,
    "AltIssueNumber": null,
    "DateAdded": "2020-05-04"
  },
  {
    "Status": "Wanted",
    "ComicSize": null,
    "ComicName": "Hellblazer Special: Chas",
    "IssueID": "137432",
    "DigitalDate": "0000-00-00",
    "IssueDate": "2008-11-24",
    "ImageURL": "https://comicvine1.cbsistatic.com/uploads/scale_small/6/67663/2788642-03.jpg",
    "inCacheDIR": null,
    "IssueDate_Edit": null,
    "ImageURL_ALT": "https://comicvine1.cbsistatic.com/uploads/scale_medium/6/67663/2788642-03.jpg",
    "ReleaseDate": "2008-09-04",
    "ArtworkURL": null,
    "Issue_Number": "3",
    "Location": null,
    "Int_IssueNumber": 3000,
    "IssueName": "chapter three",
    "ComicID": "22530",
    "Type": null,
    "AltIssueNumber": null,
    "DateAdded": "2020-05-04"
  },
  {
    "Status": "Wanted",
    "ComicSize": null,
    "ComicName": "Hellblazer Special: Chas",
    "IssueID": "139705",
    "DigitalDate": "0000-00-00",
    "IssueDate": "2008-12-24",
    "ImageURL": "https://comicvine1.cbsistatic.com/uploads/scale_small/6/67663/2788643-04.jpg",
    "inCacheDIR": null,
    "IssueDate_Edit": null,
    "ImageURL_ALT": "https://comicvine1.cbsistatic.com/uploads/scale_medium/6/67663/2788643-04.jpg",
    "ReleaseDate": "0000-00-00",
    "ArtworkURL": null,
    "Issue_Number": "4",
    "Location": null,
    "Int_IssueNumber": 4000,
    "IssueName": "chapter four",
    "ComicID": "22530",
    "Type": null,
    "AltIssueNumber": null,
    "DateAdded": "2020-05-04"
  }
]`

	mockHistory = map[string]interface{}{
		"success": true,
		"data": []mylar.History{
			{
				Status:      "Post-Processed",
				ComicName:   "Zombie Tramp",
				IssueID:     "454352",
				CheckSum:    "b8f7ebd8c7f700ff253503a467add14b",
				IssueNumber: "0",
				ComicID:     "75908",
				Provider:    "",
				DateAdded:   "2020-04-30 06:38:34",
			},
			{
				Status:      "Snatched",
				ComicName:   "Zombie Tramp",
				IssueID:     "454352",
				CheckSum:    "",
				IssueNumber: "0",
				ComicID:     "75908",
				Provider:    "nzbhydra2 (newznab)",
				DateAdded:   "2020-04-30 06:38:03",
			},
		},
	}
)

var _ = Describe("Mylar", func() {

	Context("#GetIndex", func() {

		Context("successful", func() {
			AfterEach(func() {
				gock.Off()
			})

			BeforeEach(func() {
				gock.New(mylarURL).
					Get("/api").
					MatchParams(map[string]string{
						"cmd":    mylar.GetIndexCommand.String(),
						"apikey": apiKey,
					}).
					Reply(200).
					JSON(mockIndex)
			})

			It("should return two items", func() {
				client, err := mylar.New(mylarURL, apiKey)
				Expect(err).NotTo(HaveOccurred())
				comics, err := client.GetIndex()
				Expect(err).NotTo(HaveOccurred())
				Expect(len(comics)).To(Equal(2))
			})

		})

		Context("failure", func() {
			AfterEach(func() {
				gock.Off()
			})

			BeforeEach(func() {
				gock.New(mylarURL).
					Get("/api").
					MatchParams(map[string]string{
						"cmd":    mylar.GetIndexCommand.String(),
						"apikey": apiKey,
					}).
					Reply(200).
					JSON(map[string]interface{}{
						"success": false,
						"error": mylar.Error{
							Code:    100,
							Message: "this is an error",
						},
					})
			})

			It("should return two items", func() {
				client, err := mylar.New(mylarURL, apiKey)
				Expect(err).NotTo(HaveOccurred())
				_, err = client.GetIndex()
				Expect(err).To(HaveOccurred())
				Expect(err.Error()).To(Equal("error 100: this is an error"))
			})

		})

	})

	Context("#GetComic", func() {
		Context("successful", func() {
			AfterEach(func() {
				gock.Off()
			})

			var comicID = "122214"

			BeforeEach(func() {
				gock.New(mylarURL).
					Get("/api").
					MatchParams(map[string]string{
						"cmd":    mylar.GetComicCommand.String(),
						"apikey": apiKey,
						"id":     comicID,
					}).
					Reply(200).
					JSON(mockComicDetail)
			})

			It("should return two items", func() {
				client, err := mylar.New(mylarURL, apiKey)
				Expect(err).NotTo(HaveOccurred())
				comicDetail, err := client.GetComic(comicID)
				Expect(err).NotTo(HaveOccurred())
				Expect(len(comicDetail.Comic)).To(Equal(1))
				Expect(len(comicDetail.Annuals)).To(Equal(0))
				Expect(len(comicDetail.Issues)).To(Equal(2))
			})

		})
	})

	Context("#GetWanted", func() {
		Context("successful", func() {
			AfterEach(func() {
				gock.Off()
			})

			BeforeEach(func() {
				gock.New(mylarURL).
					Get("/api").
					MatchParams(map[string]string{
						"cmd":    mylar.GetWantedCommand.String(),
						"apikey": apiKey,
					}).
					Reply(200).
					BodyString(mockWanted)
			})

			It("should return three items", func() {
				client, err := mylar.New(mylarURL, apiKey)
				Expect(err).NotTo(HaveOccurred())
				wanted, err := client.GetWanted()
				Expect(err).NotTo(HaveOccurred())
				Expect(len(wanted)).To(Equal(3))
			})

		})
	})

	Context("#GetHistory", func() {
		Context("successful", func() {
			AfterEach(func() {
				gock.Off()
			})

			BeforeEach(func() {
				gock.New(mylarURL).
					Get("/api").
					MatchParams(map[string]string{
						"cmd":    mylar.GetHistoryCommand.String(),
						"apikey": apiKey,
					}).
					Reply(200).
					JSON(mockHistory)
			})

			It("should return two items", func() {
				client, err := mylar.New(mylarURL, apiKey)
				Expect(err).NotTo(HaveOccurred())
				history, err := client.GetHistory()
				Expect(err).NotTo(HaveOccurred())
				Expect(len(history)).To(Equal(2))
			})

		})
	})
})
