package mongodb

import (
	"fmt"
	"testing"
	"time"

	"github.com/globalsign/mgo/bson"

	"github.com/globalsign/mgo"
)

func TestPoolLimitMany(t *testing.T) {

	url := "mongodb://127.0.0.1:27017/?minPoolSize=1&maxPoolSize=10&maxIdleTimeMS=2000"

	session, _ := mgo.Dial(url)
	defer session.Close()

	//session.SetPoolTimeout(1 * time.Second)

	const poolLimit = 10
	//session.SetPoolLimit(poolLimit)

	// Consume the whole limit for the master.
	var master []*mgo.Session
	for i := 0; i < poolLimit; i++ {
		func() {
			s := session.Copy()
			//defer s.Close()
			s.Ping()
			master = append(master, s)
			//s.Close()
			t.Log(i)
		}()

	}

	//for _, m := range master {
	//	m.Refresh()
	//}

	before := time.Now()
	go func() {
		time.Sleep(time.Second * 3)
		master[0].Refresh()
	}()

	// Then, a single ping must block, since it would need another
	// connection to the master, over the limit. Once the goroutine
	// above releases its socket, it should move on.
	session.Ping()
	delay := time.Since(before)
	t.Log(delay)

	time.Sleep(time.Minute * 5)
}

func TestPoolLimitTimeout(t *testing.T) {

	session, _ := mgo.Dial("localhost")
	defer session.Close()
	session.SetPoolTimeout(1 * time.Second)
	session.SetPoolLimit(1)

	// Put one socket in use.
	session.Ping()

	// Now block trying to get another one due to the pool limit.
	copy := session.Copy()
	defer copy.Close()
	started := time.Now()
	_ = copy.Ping()
	delay := time.Since(started)
	t.Log(delay)

	//time.Sleep(time.Minute * 5)
}

func TestMgo(t *testing.T) {
	url := "mongodb://127.0.0.1:27017/?minPoolSize=2&maxPoolSize=5&maxIdleTimeMS=2000"
	//url := "mongodb://127.0.0.1:27017/?maxPoolSize=10&maxIdleTimeMS=2000"

	c, err := Dial(url, 30)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer c.Close()

	//session.SetPoolTimeout(1 * time.Second)

	const poolLimit = 10000
	//session.SetPoolLimit(poolLimit)

	//skeleton := base.NewSkeleton()

	// Consume the whole limit for the master.
	//var master []*Session
	for i := 0; i < poolLimit; i++ {
		//s := c.RefNo()
		//defer c.UnRefNo(s)
		//s.Ping()

		go func(x int) {
			//t.Log(x)

			s := c.Ref()
			defer c.UnRef(s)
			s.Ping()

			dbPing(c)

		}(i)

		//go dbPing(c)

		//go dbPing(c)

		//skeleton.Go(func() {
		//	dbPing(c)
		//}, func() {
		//	//dbPing(c)
		//})

	}

	//for _, m := range master {
	//	m.Refresh()
	//}

	before := time.Now()
	go func() {
		time.Sleep(time.Second * 3)
		//master[0].Refresh()
		//for _, m := range master {
		//	m.Refresh()
		//}
		t.Log("sssss")
	}()

	t.Log("aaaaa")

	// Then, a single ping must block, since it would need another
	// connection to the master, over the limit. Once the goroutine
	// above releases its socket, it should move on.
	//session.Ping()
	delay := time.Since(before)
	t.Log(delay)

	time.Sleep(time.Minute * 20)
}

func dbPing(c *DialContext) {
	s := c.Ref()
	defer c.UnRef(s)
	//s.Ping()

	tbl := s.DB("dbn").C("test")
	tbl.Insert(bson.M{"_id": bson.NewObjectId()})

	//time.Sleep(time.Millisecond * 100)
}
