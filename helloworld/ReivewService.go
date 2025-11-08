package main

import (
	"encoding/csv"
	"fmt"
	"github.com/fatih/color"
	"os"
	"reflect"
	"time"
)

// fucntions to ask and udpate the review
// functions to get the review from the CSV
// functions to print the review with table format

func reviewSystem(r rating) rating {

	fmt.Println("Enter the product id")
	fmt.Scanln(&r.Id)
	fmt.Println("Enter the rating ")
	fmt.Scanln(&r.Stars)
	fmt.Println("Enter the comment ")
	fmt.Scanln(&r.Comment)
	fmt.Println("Enter the username ")
	fmt.Scanln(&r.Username)
	if r.Stars > 3 {
		color.Green("Thanks for the review ")
	} else {
		color.Red("will look into the feedback  ")
	}
	return r
}

func printStarsAfterFeedBack(r rating) {
	switch r.Stars {
	case 1:
		fmt.Println("😭")
	case 2:
		fmt.Println("😥")
	case 3:
		fmt.Println("😕")
	case 4:
		fmt.Println("😊")
	case 5:
		fmt.Println("😃")
	}
}

func storeReviewInCSV(r rating) error {
	const path = "users.csv"

	// 1) Open (or create) the file for appending
	file, err := os.OpenFile(path, os.O_CREATE|os.O_RDWR|os.O_APPEND, 0644)
	if err != nil {
		return fmt.Errorf("open file: %w", err)
	}
	defer file.Close()

	// 2) Detect if the file is empty to decide whether to write headers
	info, err := file.Stat()
	if err != nil {
		return fmt.Errorf("stat file: %w", err)
	}
	writeHeader := info.Size() == 0

	// 3) Prepare CSV writer
	w := csv.NewWriter(file)
	defer w.Flush()

	typ := reflect.TypeOf(r)
	val := reflect.ValueOf(r)
	if typ.Kind() == reflect.Pointer {
		typ = typ.Elem()
		val = val.Elem()
	}
	if typ.Kind() != reflect.Struct {
		return fmt.Errorf("storeReviewInCSV expects a struct or *struct, got %s", typ.Kind())
	}

	// 4) Build header (once)
	if writeHeader {
		header := make([]string, typ.NumField())
		for i := 0; i < typ.NumField(); i++ {
			f := typ.Field(i)
			if tag := f.Tag.Get("csv"); tag != "" && tag != "-" {
				header[i] = tag
			} else if f.Tag.Get("csv") == "-" {
				// skip field entirely
				header[i] = "" // placeholder; will be dropped when writing record
			} else {
				header[i] = f.Name
			}
		}
		// Drop any fields tagged with csv:"-"
		header = dropBlank(header)
		if err := w.Write(header); err != nil {
			return fmt.Errorf("write header: %w", err)
		}
	}

	// 5) Build record (all fields)
	record := make([]string, 0, typ.NumField())
	for i := 0; i < typ.NumField(); i++ {
		f := typ.Field(i)
		if f.Tag.Get("csv") == "-" {
			continue // skip
		}
		v := val.Field(i)
		record = append(record, stringify(v))
	}

	if err := w.Write(record); err != nil {
		return fmt.Errorf("write record: %w", err)
	}

	// Ensure any write error is caught
	if err := w.Error(); err != nil {
		return fmt.Errorf("flush error: %w", err)
	}
	return nil
}

// stringify converts common kinds to a CSV-friendly string.
func stringify(v reflect.Value) string {
	if !v.IsValid() {
		return ""
	}
	// Handle pointers
	if v.Kind() == reflect.Pointer {
		if v.IsNil() {
			return ""
		}
		v = v.Elem()
	}
	switch val := v.Interface().(type) {
	case time.Time:
		// ISO 8601 is a good default; change to your preferred layout if needed
		return val.Format(time.RFC3339)
	case time.Duration:
		return val.String()
	default:
		return fmt.Sprint(val)
	}
}

// dropBlank removes blank columns (used for csv:"-")
func dropBlank(ss []string) []string {
	out := make([]string, 0, len(ss))
	for _, s := range ss {
		if s != "" {
			out = append(out, s)
		}
	}
	return out
}
