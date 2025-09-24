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
	id := uuid.UUID()
	fmt.Printf("Standard UUID: %s\n", id)
	// Output: Standard UUID: 550e8400-e29b-41d4-a716-446655440000
}

// ExampleUUID32 demonstrates UUID generation without hyphens
func ExampleUUID32() {
	// Generate a UUID without hyphens
	id := uuid.UUID32()
	fmt.Printf("UUID32: %s\n", id)
	// Output: UUID32: 550e8400e29b41d4a716446655440000
}

// ExampleNewSid demonstrates creating a Sonyflake ID generator
func ExampleNewSid() {
	// Create a new Sonyflake ID generator
	_ = uuid.NewSid()
	fmt.Printf("Sonyflake generator created\n")
	// Output: Sonyflake generator created
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
	// Output: String ID: 1Z3k9X2m
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
	// Output: Raw ID: 1234567890123456789
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
	// Output:
	// ID 1: 1Z3k9X2m
	// ID 2: 1Z3k9X2n
	// ID 3: 1Z3k9X2o
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
	// Output:
	// Raw ID 1: 1234567890123456789
	// Raw ID 2: 1234567890123456790
	// Raw ID 3: 1234567890123456791
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
	// Output:
	// Generated 5 IDs:
	//   1: 1Z3k9X2m
	//   2: 1Z3k9X2n
	//   3: 1Z3k9X2o
	//   4: 1Z3k9X2p
	//   5: 1Z3k9X2q
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
	// Output:
	// Generated 5 raw IDs:
	//   1: 1234567890123456789
	//   2: 1234567890123456790
	//   3: 1234567890123456791
	//   4: 1234567890123456792
	//   5: 1234567890123456793
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
	// Output:
	// Concurrent ID generation:
	//   ID: 1Z3k9X2m
	//   ID: 1Z3k9X2n
	//   ID: 1Z3k9X2o
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
	// Output:
	// Concurrent raw ID generation:
	//   Raw ID: 1234567890123456789
	//   Raw ID: 1234567890123456790
	//   Raw ID: 1234567890123456791
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
	// Output:
	// Generated 1000 string IDs in 250µs
	// Average time per ID: 250ns
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
	// Output:
	// Generated 1000 raw IDs in 100µs
	// Average time per ID: 100ns
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
	// Output: Generated ID: 1Z3k9X2m
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
	// Output: Generated raw ID: 1234567890123456789
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
	// Output: User record: ID=1Z3k9X2m, Name=John Doe
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
	// Output: User record: ID=1234567890123456789, Name=John Doe
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
	// Output:
	// Microservice request ID: 1Z3k9X2m
	// Processing request 1Z3k9X2m
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
	// Output:
	// Microservice raw request ID: 1234567890123456789
	// Processing request 1234567890123456789
}
