package highorderfunc

import "testing"

func TestBadBank(t *testing.T) {
	transactions := []Transaction{
		{
			From: "Alice",
			To:   "Bob",
			Sum:  10},
		{
			From: "Bob",
			To:   "Charlie",
			Sum:  20},
		{
			From: "Charlie",
			To:   "Alice",
			Sum:  5},
	}
	AssertEqual(t, BalanceFor(transactions, "Alice"), -5)
}
func AssertEqual[T comparable](t testing.TB, got, want T) {
	t.Helper()
	if got != want {
		t.Errorf("got %v, want %v", got, want)
	}
}
