package pool

import (
	"fmt"
	"sync"
	"time"

	"golang.org/x/sync/errgroup"
)

func SyncPoolExample() {
	pool := &sync.Pool{
		// New:func() any{
		// fmt.Println("Pool:creating a new connection")
		// db,_:=db
		// return nil
		// }
	}

	var msec time.Duration = 10

	var g errgroup.Group

	for i := 0; i < 100; i++ {
		time.Sleep(msec * time.Millisecond)

		g.Go(func() error {
			var db *MockDB
			item := pool.Get()

			if item == nil {
				fmt.Println("creating new connection")
				db, _ = db.Open("server 1")
			} else {
				db = item.(*MockDB)
				fmt.Println("reusing conn from the pool")
			}
			defer db.Close()
			_, _ = db.Query("SELECT * FROM notes")
			pool.Put(db)
			return nil
		})
	}

}

type MockDB struct {
	name string
}

func (mdb *MockDB) Close()                       {}
func (mdb *MockDB) Query(q string) (bool, error) { return true, nil }
func (mdb *MockDB) Open(name string) (*MockDB, error) {
	return &MockDB{name}, nil
}
