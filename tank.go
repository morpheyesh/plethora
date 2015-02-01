package detour


import (

  "fmt"
  "github.com/garyburd/redigo/redis"
  )


type Tank struct {
  cram    *Cramshaft

}

func (t Tank) redis_add(user User, item Item) error {

est, err := redis.Bool(t.cram.redis_conn.Do("SISMEMBER", fmt.Sprintf("%s:%s:%s", t.cram.class, item), user))
if err != nil
  return err
  }
 }
}
