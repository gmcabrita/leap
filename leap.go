// Package leap implements implements the jump consistent hashing algorithm proposed by John Lamping and Eric Veach.
package leap

// Hash takes a 64 bit key and the number of existing buckets and returns the bucket placement for the given key.
func Hash(key uint64, numBuckets int) int32 {
	var previous int64 = -1
	var current int64

	for current < int64(numBuckets) {
		previous = current
		key = key*2862933555777941757 + 1
		current = int64(float64(previous+1) * (float64(1<<31) / float64((key>>33)+1)))
	}

	return int32(previous)
}
