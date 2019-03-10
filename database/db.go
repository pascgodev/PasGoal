package database

import (
	"fmt"
	"github.com/syndtr/goleveldb/leveldb"
	"github.com/syndtr/goleveldb/leveldb/errors"
	"github.com/syndtr/goleveldb/leveldb/filter"
	"github.com/syndtr/goleveldb/leveldb/iterator"
	"github.com/syndtr/goleveldb/leveldb/opt"
	"github.com/syndtr/goleveldb/leveldb/util"
	"sync"
)

type Database struct {
	file string
	db   *leveldb.DB

	quitLock sync.Mutex      // Mutex protecting the quit channel access
	quitChan chan chan error // Quit channel to stop the metrics collection before closing the database
}

// NewLDBDatabase returns a LevelDB wrapped object.
func NewDatabase(file string, cache int, handles int) (*Database, error) {
	// Open the db and recover any potential corruptions
	if cache < 16 {
		cache = 16
	}
	if handles < 16 {
		handles = 16
	}
	db, err := leveldb.OpenFile(file, &opt.Options{
		OpenFilesCacheCapacity: handles,
		BlockCacheCapacity:     cache / 2 * opt.MiB,
		WriteBuffer:            cache / 4 * opt.MiB, // Two of these are used internally
		Filter:                 filter.NewBloomFilter(10),
	})
	if _, corrupted := err.(*errors.ErrCorrupted); corrupted {
		db, err = leveldb.RecoverFile(file, nil)
	}
	// (Re)check for errors and abort if opening of the db failed
	if err != nil {
		return nil, err
	}
	return &Database{
		db: db,
	}, nil
}

// Put puts the given key / value to the queue
func (DB *Database) Put(key []byte, value []byte) error {
	return DB.db.Put(key, value, nil)
}

// Get returns the given key if it's present.
func (DB *Database) Get(key []byte) ([]byte, error) {
	// Retrieve the key and increment the miss counter if not found
	dat, err := DB.db.Get(key, nil)
	if err != nil {
		return nil, err
	}
	return dat, nil
}

func (DB *Database) Has(key []byte) (bool, error) {
	return DB.db.Has(key, nil)
}

// Delete deletes the key from the queue and database
func (DB *Database) Delete(key []byte) error {
	// Execute the actual operation
	return DB.db.Delete(key, nil)
}

func (DB *Database) NewIterator() iterator.Iterator {
	return DB.db.NewIterator(nil, nil)
}

func (DB *Database) NewIteratorRange(slice *util.Range) iterator.Iterator {
	return DB.db.NewIterator(slice, nil)
}

func NewBytesPrefix(prefix []byte) *util.Range {
	return util.BytesPrefix(prefix)
}

func (DB *Database) Close() {
	if err := DB.db.Close(); err != nil {
		fmt.Errorf("DB %s: %s", DB.file, err)
	}
}

func (DB *Database) LDB() *leveldb.DB {
	return DB.db
}
