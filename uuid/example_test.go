package uuid_test

import (
	"fmt"
	"log"
	"time"

	"github.com/go4x/goal/uuid"
)

// ExampleUUID demonstrates basic UUID generation
func ExampleUUID() {
	// Generate a standard UUID
	id, err := uuid.UUID()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Standard UUID: %s\n", id)

}

// ExampleUUID32 demonstrates UUID generation without hyphens
func ExampleUUID32() {
	// Generate a UUID without hyphens
	id := uuid.UUID32()
	fmt.Printf("UUID32: %s\n", id)

}

// ExampleNewSid demonstrates creating a Sonyflake ID generator
func ExampleNewSid() {
	// Create a new Sonyflake ID generator
	_ = uuid.NewSid()
	fmt.Printf("Sonyflake generator created\n")

}

// ExampleSid_GenString demonstrates generating string IDs
func ExampleSid_GenString() {
	// Create Sonyflake generator
	sid := uuid.NewSid()

	// Generate string ID
	id, err := sid.GenString()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("String ID: %s\n", id)

}

// ExampleSid_GenUint64 demonstrates generating raw uint64 IDs
func ExampleSid_GenUint64() {
	// Create Sonyflake generator
	sid := uuid.NewSid()

	// Generate raw uint64 ID
	id, err := sid.GenUint64()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Raw ID: %d\n", id)

}

// ExampleSid_GenString_multiple demonstrates generating multiple string IDs
func ExampleSid_GenString_multiple() {
	// Create Sonyflake generator
	sid := uuid.NewSid()

	// Generate multiple string IDs
	for i := 0; i < 3; i++ {
		id, err := sid.GenString()
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("ID %d: %s\n", i+1, id)
	}

}

// ExampleSid_GenUint64_multiple demonstrates generating multiple raw IDs
func ExampleSid_GenUint64_multiple() {
	// Create Sonyflake generator
	sid := uuid.NewSid()

	// Generate multiple raw uint64 IDs
	for i := 0; i < 3; i++ {
		id, err := sid.GenUint64()
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("Raw ID %d: %d\n", i+1, id)
	}

}

// ExampleSid_GenString_batch demonstrates batch ID generation
func ExampleSid_GenString_batch() {
	// Create Sonyflake generator
	sid := uuid.NewSid()

	// Generate batch of string IDs
	ids := make([]string, 5)
	for i := 0; i < 5; i++ {
		id, err := sid.GenString()
		if err != nil {
			log.Fatal(err)
		}
		ids[i] = id
	}

	fmt.Printf("Generated %d IDs:\n", len(ids))
	for i, id := range ids {
		fmt.Printf("  %d: %s\n", i+1, id)
	}

}

// ExampleSid_GenUint64_batch demonstrates batch raw ID generation
func ExampleSid_GenUint64_batch() {
	// Create Sonyflake generator
	sid := uuid.NewSid()

	// Generate batch of raw uint64 IDs
	ids := make([]uint64, 5)
	for i := 0; i < 5; i++ {
		id, err := sid.GenUint64()
		if err != nil {
			log.Fatal(err)
		}
		ids[i] = id
	}

	fmt.Printf("Generated %d raw IDs:\n", len(ids))
	for i, id := range ids {
		fmt.Printf("  %d: %d\n", i+1, id)
	}

}

// ExampleSid_GenString_concurrent demonstrates concurrent ID generation
func ExampleSid_GenString_concurrent() {
	// Create Sonyflake generator
	sid := uuid.NewSid()

	// Channel for collecting results
	results := make(chan string, 3)

	// Generate IDs concurrently
	for i := 0; i < 3; i++ {
		go func() {
			id, err := sid.GenString()
			if err != nil {
				log.Fatal(err)
			}
			results <- id
		}()
	}

	// Collect results
	fmt.Println("Concurrent ID generation:")
	for i := 0; i < 3; i++ {
		id := <-results
		fmt.Printf("  ID: %s\n", id)
	}

}

// ExampleSid_GenUint64_concurrent demonstrates concurrent raw ID generation
func ExampleSid_GenUint64_concurrent() {
	// Create Sonyflake generator
	sid := uuid.NewSid()

	// Channel for collecting results
	results := make(chan uint64, 3)

	// Generate IDs concurrently
	for i := 0; i < 3; i++ {
		go func() {
			id, err := sid.GenUint64()
			if err != nil {
				log.Fatal(err)
			}
			results <- id
		}()
	}

	// Collect results
	fmt.Println("Concurrent raw ID generation:")
	for i := 0; i < 3; i++ {
		id := <-results
		fmt.Printf("  Raw ID: %d\n", id)
	}

}

// ExampleSid_GenString_performance demonstrates performance characteristics
func ExampleSid_GenString_performance() {
	// Create Sonyflake generator
	sid := uuid.NewSid()

	// Measure generation time
	start := time.Now()
	count := 1000

	for i := 0; i < count; i++ {
		_, err := sid.GenString()
		if err != nil {
			log.Fatal(err)
		}
	}

	duration := time.Since(start)
	fmt.Printf("Generated %d string IDs in %v\n", count, duration)
	fmt.Printf("Average time per ID: %v\n", duration/time.Duration(count))

}

// ExampleSid_GenUint64_performance demonstrates raw ID performance
func ExampleSid_GenUint64_performance() {
	// Create Sonyflake generator
	sid := uuid.NewSid()

	// Measure generation time
	start := time.Now()
	count := 1000

	for i := 0; i < count; i++ {
		_, err := sid.GenUint64()
		if err != nil {
			log.Fatal(err)
		}
	}

	duration := time.Since(start)
	fmt.Printf("Generated %d raw IDs in %v\n", count, duration)
	fmt.Printf("Average time per ID: %v\n", duration/time.Duration(count))

}

// ExampleSid_GenString_error_handling demonstrates error handling
func ExampleSid_GenString_error_handling() {
	// Create Sonyflake generator
	sid := uuid.NewSid()

	// Generate ID with error handling
	id, err := sid.GenString()
	if err != nil {
		log.Printf("Failed to generate ID: %v", err)
		return
	}

	fmt.Printf("Generated ID: %s\n", id)

}

// ExampleSid_GenUint64_error_handling demonstrates raw ID error handling
func ExampleSid_GenUint64_error_handling() {
	// Create Sonyflake generator
	sid := uuid.NewSid()

	// Generate raw ID with error handling
	id, err := sid.GenUint64()
	if err != nil {
		log.Printf("Failed to generate raw ID: %v", err)
		return
	}

	fmt.Printf("Generated raw ID: %d\n", id)

}

// ExampleSid_GenString_database demonstrates database usage
func ExampleSid_GenString_database() {
	// Create Sonyflake generator
	sid := uuid.NewSid()

	// Generate ID for database record
	id, err := sid.GenString()
	if err != nil {
		log.Fatal(err)
	}

	// Simulate database record
	user := struct {
		ID   string
		Name string
	}{
		ID:   id,
		Name: "John Doe",
	}

	fmt.Printf("User record: ID=%s, Name=%s\n", user.ID, user.Name)

}

// ExampleSid_GenUint64_database demonstrates raw ID database usage
func ExampleSid_GenUint64_database() {
	// Create Sonyflake generator
	sid := uuid.NewSid()

	// Generate raw ID for database record
	id, err := sid.GenUint64()
	if err != nil {
		log.Fatal(err)
	}

	// Simulate database record
	user := struct {
		ID   uint64
		Name string
	}{
		ID:   id,
		Name: "John Doe",
	}

	fmt.Printf("User record: ID=%d, Name=%s\n", user.ID, user.Name)

}

// ExampleSid_GenString_microservice demonstrates microservice usage
func ExampleSid_GenString_microservice() {
	// Create Sonyflake generator
	sid := uuid.NewSid()

	// Generate request ID for microservice
	requestID, err := sid.GenString()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Microservice request ID: %s\n", requestID)
	fmt.Printf("Processing request %s\n", requestID)

}

// ExampleSid_GenUint64_microservice demonstrates raw ID microservice usage
func ExampleSid_GenUint64_microservice() {
	// Create Sonyflake generator
	sid := uuid.NewSid()

	// Generate raw request ID for microservice
	requestID, err := sid.GenUint64()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Microservice raw request ID: %d\n", requestID)
	fmt.Printf("Processing request %d\n", requestID)

}
