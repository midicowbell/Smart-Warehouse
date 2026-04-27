package warehouse

import "sync"

type Repository struct {
	data map[string]Item
	mtx  sync.RWMutex
}

func NewRepository() *Repository {
	return &Repository{
		data: make(map[string]Item),
	}
}

func (r *Repository) AddItem(item Item) error {
	r.mtx.Lock()
	defer r.mtx.Unlock()
	if _, ok := r.data[item.Name]; ok {
		return ErrItemAlreadyExist
	}
	r.data[item.Name] = item
	return nil
}

func (r *Repository) DeleteItem(name string) error {
	r.mtx.Lock()
	defer r.mtx.Unlock()
	if _, ok := r.data[name]; !ok {
		return ErrItemNotFound
	}
	delete(r.data, name)
	return nil
}

func (r *Repository) UpdateQuantity(name string, quantity int) error {
	r.mtx.Lock()
	defer r.mtx.Unlock()
	if _, ok := r.data[name]; !ok {
		return ErrItemNotFound
	}
	item := r.data[name]
	item.Quantity = quantity
	r.data[name] = item
	return nil
}

func (r *Repository) GetItem(name string) (Item, error) {
	r.mtx.RLock()
	defer r.mtx.RUnlock()
	if _, ok := r.data[name]; !ok {
		return Item{}, ErrItemNotFound
	}
	return r.data[name], nil
}

func (r *Repository) ListItems() map[string]Item {
	r.mtx.RLock()
	defer r.mtx.RUnlock()
	copy := make(map[string]Item, len(r.data))
	for key, value := range r.data {
		copy[key] = value
	}
	return copy
}

func (r *Repository) ListLowStockItems() map[string]Item {
	r.mtx.RLock()
	defer r.mtx.RUnlock()
	lowItemsList := make(map[string]Item, len(r.data))
	for key, value := range r.data {
		if value.Quantity < value.MinThreshold {
			lowItemsList[key] = value
		}
	}
	return lowItemsList
}
