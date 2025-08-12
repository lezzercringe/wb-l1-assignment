package main

import (
	"bytes"
	"encoding/gob"
	"errors"
	"fmt"
)

// Предположим появился новый супер-быстрый кэш разработанный сообществом,
// и нам бы хотелось начать использовать его в своем приложении для хранения чего-нибудь.

type UltraFastWellOptimizedCache struct{}

func (c *UltraFastWellOptimizedCache) Store(k, v []byte) {
	panic("unimplemented")
}

func (c *UltraFastWellOptimizedCache) Load(k, v []byte) []byte {
	panic("unimplemented")
}

// Будем считать что при разработке приложения мы ориентировались на гексагональную или луковую архитектуру,
// поэтому хранение некой сущности абстрагируется паттерном репозиторий:

type Entity = struct {
	ID int
}

type EntityRepository interface {
	Save(*Entity) error
	GetByID(id int) (*Entity, error)
}

// Очевидно, что прямо "из коробки" очень быстрый кэш не реализует этот интерфейс.
// Для того чтобы подогнать кэш под необходимый интерфейс и прокидывать его в приложение
// можно написать адаптер.

var _ EntityRepository = &EntityCache{}

type EntityCache struct {
	cache *UltraFastWellOptimizedCache
}

func (c *EntityCache) GetByID(id int) (*Entity, error) {
	buf := c.cache.Load(serializeID(id), nil)
	if len(buf) == 0 {
		return nil, errors.New("not found")
	}

	var e Entity
	if err := gob.NewDecoder(bytes.NewReader(buf)).Decode(&e); err != nil {
		return nil, fmt.Errorf("gob-deserializing entity: %w", err)
	}

	return &e, nil
}

func (c *EntityCache) Save(e *Entity) error {
	var buf bytes.Buffer

	if err := gob.NewEncoder(&buf).Encode(e); err != nil {
		return fmt.Errorf("gob-serializing entity: %w", err)
	}

	c.cache.Store(serializeID(e.ID), buf.Bytes())
	return nil
}

func serializeID(id int) []byte {
	return fmt.Append(nil, id)
}
