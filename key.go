/**
 * @author Jongmin Kim (@golanger)
 */

package main

import (
    "rand"
    "time"
)

var keyChar = []byte("0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func genKey() string {
    rand.Seed(time.Nanoseconds())

    var k int
    key := make([]byte, 5)

    for i := 0; i < 5; i++ {
        k = rand.Intn(len(keyChar))
        key[i] = keyChar[k]
    }

    return string(key[:])
}

/* EOF */
