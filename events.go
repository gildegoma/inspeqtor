package inspeqtor

/*
 There are 3 pairs of events we need to handle:
 1. Process does not exist, either on startup or disappearing during a cycle.
 2. Process exists, it appears after previously not existing.

 3. Rule triggered after metric crossed threshold for N cycles (simple alert)
 4. Metric recovered

 5. Service needs restarting due to rule triggering.
 6. Service started correctly (shared with 2?)
*/

/*
 There are several different types of Events:

 * Process disappeared (or did not exist when we started up)
 * Process appeared
 * Rule triggered based on metric check
 * Metric has recovered
 * Process is restarting due to rule trigger
 * Process has restarted
*/
type EventType uint8

const (
	ProcessDoesNotExist EventType = iota
	ProcessExists
	HealthFailure
	HealthRecovered
	ServiceRestarting
	ServiceRestarted
)

type Event struct {
	Check Checkable
	Rule  *Rule
	Type  EventType
}