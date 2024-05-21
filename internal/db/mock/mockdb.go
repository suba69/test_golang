package mock

type MockDatabase struct {
	InsertExchangeRateFn func(rate float64) error
}

func (m *MockDatabase) InsertExchangeRate(rate float64) error {
	if m.InsertExchangeRateFn != nil {
		return m.InsertExchangeRateFn(rate)
	}
	return nil
}
