package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"

	"github.com/adrg/strutil"
	"github.com/adrg/strutil/metrics"
)

var WeightFirstName = 0.2
var WeightLastName = 0.2
var WeightEmailAddress = 0.3
var WeightZipCode = 0.15
var WeightAddress = 0.15
var ThresholdMatch = 0.8
var ThresholdAccuracy = 0.9

// output is list of contactId source, contactId match, accuracy
func main() {
	duplicates := findDuplicates("assets/SampleCode.csv")
	fmt.Println("Found", len(duplicates), "duplicates")
	for i, duplicate := range duplicates {
		fmt.Println(i+1, "Found possible duplicate: Contact", duplicate.ContactIdSource, "with", duplicate.ContactIdMatch, "with accuracy:", duplicate.Accuracy)
	}
}

func findDuplicates(sourceFile string) []Match {
	contacts := loadAllContacts(sourceFile)
	matches := []Match{}
	for index, sourceContact := range contacts {
		for _, targetContact := range contacts[index+1:] {
			// fmt.Printf("Comparing %s with %s\n", sourceContact.FirstName, targetContact.FirstName)
			valid, accuracy := calculateMatch(sourceContact, targetContact)
			if valid {
				var accuracyStr string
				if accuracy {
					accuracyStr = "High"
				} else {
					accuracyStr = "Low"
				}
				match := Match{ContactIdSource: sourceContact.Id, ContactIdMatch: targetContact.Id, Accuracy: accuracyStr}
				matches = append(matches, match)
			}
		}
	}

	return matches
}

func calculateMatch(source Contact, target Contact) (bool, bool) {
	// fmt.Println("Comparing")
	// fmt.Println(source)
	// fmt.Println(target)
	metricInsensitive := metrics.NewJaroWinkler()
	metricInsensitive.CaseSensitive = false
	metricSensitive := metrics.NewJaroWinkler()
	// Use 3rd-party lib to calculate similarity based on JaroWinkler algorithm.
	// Use two different algorithms, given that fields are case insensitive
	// Weight each individual scoring
	totalScore := strutil.Similarity(source.FirstName, target.FirstName, metricInsensitive)*WeightFirstName +
		strutil.Similarity(source.LastName, target.LastName, metricInsensitive)*WeightLastName +
		strutil.Similarity(source.EmailAddress, target.EmailAddress, metricInsensitive)*WeightEmailAddress +
		strutil.Similarity(source.ZipCode, target.ZipCode, metricSensitive)*WeightZipCode +
		strutil.Similarity(source.Address, target.Address, metricInsensitive)*WeightAddress
	// fmt.Println("Total Score", totalScore)
	// Compare the total scoring with threshold value
	valid := totalScore >= ThresholdMatch
	accuracy := false
	// If threshold is met, compare with high or low expectation
	if valid && totalScore >= ThresholdAccuracy {
		accuracy = true
	}
	return valid, accuracy
}

func loadAllContacts(sourceFile string) []Contact {
	// Load everything to memory
	// Could make a more efficient version, but for this exercise will suffice.
	f, _ := os.Open(sourceFile)
	r := csv.NewReader(f)
	contacts := make([]Contact, 0)
	header := true
	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if header {
			header = false
			continue
		}
		if err != nil {
			panic(err)
		}
		contacts = append(contacts, *LoadContact(record))
	}
	return contacts
}
