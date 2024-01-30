package main

import (
	"fmt"
	"net/url"
	"reflect"

	"github.com/google/go-querystring/query"
)

type RepositoryListByOrgOptions struct {
	// Type of repositories to list. Possible values are: all, public, private,
	// forks, sources, member. Default is "all".
	Type string `url:"type,omitempty"`

	// How to sort the repository list. Can be one of created, updated, pushed,
	// full_name. Default is "created".
	Sort string `url:"sort,omitempty"`

	// Direction in which to sort repositories. Can be one of asc or desc.
	// Default when using full_name: asc; otherwise desc.
	Direction string `url:"direction,omitempty"`

	ListOptions
}

// ListOptions specifies the optional parameters to various List methods that
// support offset pagination.
type ListOptions struct {
	// For paginated result sets, page of results to retrieve.
	Page int `url:"page,omitempty"`

	// For paginated result sets, the number of results to include per page.
	PerPage int `url:"per_page,omitempty"`
}


type DeleteOptions struct {
	FilterBy  string `url:"filter_by,omitempty"`
	BatchSize int    `url:"batch_size,omitempty"`
}

type ExportDocumentsParams struct {
	ExportDocumentsParameters struct {
		// ExcludeFields List of fields from the document to exclude in the search result
		ExcludeFields string `url:"exclude_fields"`

		// FilterBy Filter conditions for refining your search results. Separate multiple conditions with &&.
		FilterBy string `url:"filter_by,omitempty"`

		// IncludeFields List of fields from the document to include in the search result
		IncludeFields string `url:"include_fields"`
	} 
}

func addOptions(s string, opts interface{}) (string, error) {
	v := reflect.ValueOf(opts)
	if v.Kind() == reflect.Ptr && v.IsNil() {
		return s, nil
	}

	u, err := url.Parse(s)
	if err != nil {
		return s, err
	}

	qs, err := query.Values(opts)
	if err != nil {
		return s, err
	}

	u.RawQuery = qs.Encode()
	return u.String(), nil
}

func main() {
	var (
		org    = "awesome-org"
		
		opts1  = &RepositoryListByOrgOptions{
			Type: "public",
			Sort: "created",
			Direction: "desc",
			ListOptions: ListOptions{
				Page: 5,
				PerPage: 10,
			},
		}
		
		opts2  = &DeleteOptions{
			FilterBy: "num_students:>100",
			BatchSize: 100,
		}

		opts3  = &ExportDocumentsParams{
			ExportDocumentsParameters: struct{
					ExcludeFields 	string 	`url:"exclude_fields"`
					FilterBy 		string `url:"filter_by,omitempty"`
					IncludeFields 	string 	`url:"include_fields"`
				}{
					ExcludeFields: "name",
					IncludeFields: "age",
				},
		}
	)

	u1 := fmt.Sprintf("orgs/%v/repos", org)
	u1, err := addOptions(u1, opts1)
	if err != nil {
		fmt.Printf("err: %v\n", err)
	}
	fmt.Printf("url1: %v\n", u1)

	u2 := "/collections"
	u2, err = addOptions(u2, opts2)
	if err != nil {
		fmt.Printf("err: %v\n", err)
	}
	fmt.Printf("url2: %v\n", u2)

	u3 := "/collections"
	u3, err = addOptions(u3, opts3)
	if err != nil {
		fmt.Printf("err: %v\n", err)
	}
	fmt.Printf("url3: %v\n", u3)
}