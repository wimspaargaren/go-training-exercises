package repository

type Info struct {
	ID, X, Y int
}

type DataStore interface {
	Store(*Info) error
}

type infoHandler struct {
	dataStore DataStore
}

func NewInfoHandler(dataStore DataStore) *infoHandler {
	return &infoHandler{
		dataStore: dataStore,
	}
}

func (i *infoHandler) CreateNewInfo(x, y int) (*Info, error) {
	info := Info{
		X: x,
		Y: y,
	}
	err := i.dataStore.Store(&info)
	if err != nil {
		return nil, err
	}
	return &info, nil
}
