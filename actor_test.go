package govalhalla

import (
	"encoding/json"
	"os"
	"strings"
	"testing"
)

// TestConfig provides test configuration for the actor
var TestConfig = `{
	"mjolnir": {
		"tile_dir": "./test/traffic_matcher_tiles"
	},
	"loki": {
		"actions": [
			"route",
			"locate",
			"sources_to_targets",
			"optimized_route",
			"isochrone",
			"trace_route",
			"trace_attributes",
			"transit_available",
			"expansion",
			"centroid"
		],
		"logging": {
			"long_request": 100.0
		},
		"service": {
			"proxy": "ipc:///tmp/loki"
		}
	},
	"thor": {
		"logging": {
			"long_request": 110.0
		}
	}
}`

func TestNewActor(t *testing.T) {
	actor := NewActor(TestConfig, true)
	if actor == nil {
		t.Fatal("Failed to create actor")
	}
	defer actor.Cleanup()
}

func TestBasicRouting(t *testing.T) {
	actor := NewActor(TestConfig, true)
	if actor == nil {
		t.Fatal("Failed to create actor")
	}
	defer actor.Cleanup()

	request := `{
		"locations": [
			{"lat": 40.546115, "lon": -76.385076, "type": "break"},
			{"lat": 40.544232, "lon": -76.385752, "type": "break"}
		],
		"costing": "auto"
	}`

	// Test route functionality
	result := actor.Route(request)
	if result.GetSecond() != "" {
		t.Errorf("Route request failed: %s", result.GetSecond())
	}
	if !strings.Contains(result.GetFirst(), "Tulpehocken") {
		t.Error("Expected 'Tulpehocken' in route response")
	}
}

func TestTraceAttributes(t *testing.T) {
	actor := NewActor(TestConfig, true)
	if actor == nil {
		t.Fatal("Failed to create actor")
	}
	defer actor.Cleanup()

	request := `{
		"shape": [
			{"lat": 40.546115, "lon": -76.385076},
			{"lat": 40.544232, "lon": -76.385752}
		],
		"costing": "auto",
		"shape_match": "map_snap"
	}`

	result := actor.TraceAttributes(request)
	if result.GetSecond() != "" {
		t.Errorf("Trace attributes request failed: %s", result.GetSecond())
	}
	if !strings.Contains(result.GetFirst(), "Tulpehocken") {
		t.Error("Expected 'Tulpehocken' in trace response")
	}
}

func TestTransitAvailable(t *testing.T) {
	actor := NewActor(TestConfig, true)
	if actor == nil {
		t.Fatal("Failed to create actor")
	}
	defer actor.Cleanup()

	request := `{
		"locations": [
			{"lat": 35.647452, "lon": -79.597477, "radius": 20},
			{"lat": 34.766908, "lon": -80.325936, "radius": 10}
		]
	}`

	result := actor.TransitAvailable(request)
	if result.GetSecond() != "" {
		t.Errorf("Transit available request failed: %s", result.GetSecond())
	}

	// Parse response to check for true value
	var response struct {
		Available bool `json:"available"`
	}
	if err := json.Unmarshal([]byte(result.GetFirst()), &response); err != nil {
		t.Errorf("Failed to parse transit response: %v", err)
	}
	if response.Available {
		t.Error("Expected transit to be unavailable in test area")
	}
}

func TestStatus(t *testing.T) {
	actor := NewActor(TestConfig, true)
	if actor == nil {
		t.Fatal("Failed to create actor")
	}
	defer actor.Cleanup()

	// Test basic status
	result := actor.Status("")
	if result.GetSecond() != "" {
		t.Errorf("Status request failed: %s", result.GetSecond())
	}
	if !strings.Contains(result.GetFirst(), "tileset_last_modified") {
		t.Error("Expected 'tileset_last_modified' in status response")
	}

	// Test verbose status
	verboseResult := actor.Status(`{"verbose": true}`)
	if verboseResult.GetSecond() != "" {
		t.Errorf("Verbose status request failed: %s", verboseResult.GetSecond())
	}
	if !strings.Contains(verboseResult.GetFirst(), "Polygon") {
		t.Error("Expected 'Polygon' in verbose status response")
	}
}

// Helper function for integration tests
func skipIfNoData(t *testing.T) {
	if _, err := os.Stat("./test/traffic_matcher_tiles"); os.IsNotExist(err) {
		t.Skip("Test data directory not found, skipping integration test")
	}
}

// TestMain handles test setup and teardown
func TestMain(m *testing.M) {
	// Setup code if needed

	// Run tests
	code := m.Run()

	// Cleanup code if needed

	os.Exit(code)
}
