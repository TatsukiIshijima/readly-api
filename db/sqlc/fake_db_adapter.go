package db

type FakeAdapter struct{}

func (f *FakeAdapter) Connect(string, string) (DBTX, Querier) {
	db := FakeDB{}
	q := &FakeQuerier{}
	return db, q
}
