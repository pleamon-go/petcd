package petcd_test

import (
	"context"
	"log"
	"testing"

	"git.pleamon.com/p/petcd"
)

func TestMutext(t *testing.T) {
	etcd, err := petcd.New("http://localhost:2379")
	if err != nil {
		t.Fatal(err)
	}
	mutext, err := etcd.NewMutex("test")
	if err != nil {
		t.Fatal(err)
	}
	err = mutext.Lock(context.TODO())
	if err != nil {
		t.Fatal(err)
	}
	t.Log("success")
}

func TestSet(t *testing.T) {
	etcd, err := petcd.New("http://127.0.0.1:2379")
	if err != nil {
		t.Fatal(err)
	}
	resp, err := etcd.Set("a", "b")
	if err != nil {
		t.Fatal(err)
	}
	log.Println(resp)
	// resp, err := etcd.Get("a")
	// log.Println(resp, err)
}
