/**
 *
 */

package main

import (
    "os"
    "launchpad.net/gobson/bson"
    "launchpad.net/mgo"
)

type Store interface {
    Check(key *string) bool
    Put(key, url *string) os.Error
    Get(key, url *string) os.Error
}

type URLStore struct {
    Key string
    Long_url string
}


func (s *URLStore) Check(key *string) bool {
    session, _ := mgo.Mongo("127.0.0.1")
    defer session.Close()

    c := session.DB("yourl").C("urls")

    result := URLStore{}
    if err := c.Find(bson.M{"key": *key}).One(&result); err != nil {
        return true
    }

    return false
}

func (s *URLStore) Put(key, url *string) os.Error {
    session, _ := mgo.Mongo("127.0.0.1")
    defer session.Close()

    c := session.DB("yourl").C("urls")
    err := c.Insert(&URLStore{*key, *url})
    if err != nil {
        panic(err)
    }

    return nil
}

func (s *URLStore) Get(key, url *string) os.Error {
    session, _ := mgo.Mongo("127.0.0.1")
    defer session.Close()

    c := session.DB("yourl").C("urls")

    result := URLStore{}
    err := c.Find(bson.M{"key": *key}).One(&result)
    if err != nil {
        panic(err)
    }

    *url = result.Long_url

    return nil
}

/* EOF */
