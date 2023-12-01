package user

import "github.com/thtrangphu/bookStoreSvc/entity"

// innem in memory repo
type inmem struct {
	m map[entity.ID]*entity.User
}

func newInmem() *inmem {
	var m = map[entity.ID]*entity.User{} // init

	return &inmem{
		m: m,
	}
}

func (i *inmem) Create(e *entity.User) (entity.ID, error) {
	i.m[e.ID] = e
	return e.ID, nil
}

func (i *inmem) Get(id entity.ID) (*entity.User, error) {
	if i.m[id] == nil {
		return nil, entity.ErrNotFound
	}
	return i.m[id], nil
}

func (i *inmem) Update(e *entity.User) error {
	_, err := i.Get(e.ID)
	if err != nil {
		return err
	}
	i.m[e.ID] = e
	return nil
}

func (i *inmem) List() ([]*entity.User, error) {
	var list []*entity.User
	for _, j := range i.m {
		list = append(list, j)
	}
	return list, nil
}

func (i *inmem) Delete(id entity.ID) error {
	if i.m[id] == nil {
		return entity.ErrNotFound
	}
	i.m[id] = nil
	return nil
}
