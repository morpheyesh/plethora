

package detour

import (
  "net"

  "github.com/garyburd/redigo/redis"
)

type Cramshaft struct {
  redis_conn     redis.Conn
  class          string

}

func New(address net.Addr, class  string) (*Cramshaft, error) {
  redis_conn, err := redis.Dial(address.Network(), address.String())
  if err != nil {
    return nil, err
  }

  piston := &Cramshaft{
    redis_conn:     redis_conn,
    class:   class
  }
return piston
}
