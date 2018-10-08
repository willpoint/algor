package greedy

// The problem of scheduling several competitive activities
// that require exclusive use of a common resource is the first
// problem to proffer an algorithm for under this package
// with the goal of selecting a maximum-size of mutually compatible
// activities.
// Suppose that S = <a1, a2, ...aN> of n proposed activities that
// wish to use a resource, such as a lecture hall, which can serve only
// one activity at a time.
// each activity a(i) has a *start time* s and a *finish time* f
// where 0 <= si < f1 < inf. If selected activity a(i) takes place
// during the half-open time interval [si, fi). Activities ai and aj
// are compatible if the intervals [si, fi) and [sj, fj) do not
// overlap. That is ai and aj are compatible
// if si >= fj or sj >= fi in the activity-selection problem,
// we wish to select a maximum-size subset of mutually compatible
// activities. We assume that activities are sorted monotonically
// in increasing order of finish time:
// f1 <= f2 <= f3 <= ... <= f(n-1) <= fn
