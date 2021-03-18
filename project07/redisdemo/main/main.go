package main

import (
	"fmt"
	"github.com/garyburd/redigo/redis"
)

func main()  {
	conn , err := redis.Dial("tcp","127.0.0.1:6379")
	if err != nil {
		fmt.Println("redis.Dial err=", err)
		return 
	}
	defer conn.Close()


	_, err = conn.Do("Set","name","jerry")
	if err != nil {
		fmt.Println("set  err=", err)
		return 
	}
	r, err := redis.String(conn.Do("Get", "name"))
	if err != nil {
		fmt.Println("set  err=", err)
		return 
	}
	fmt.Println("操作ok ", r)



	_, err = conn.Do("HSet", "user01", "name", "john")
	if err != nil {
		fmt.Println("hset  err=", err)
		return 
	}
	_, err = conn.Do("HSet", "user01", "age", 18)
	if err != nil {
		fmt.Println("hset  err=", err)
		return 
	}
	r1, err := redis.String(conn.Do("HGet","user01", "name"))
	if err != nil {
		fmt.Println("hget  err=", err)
		return 
	}
	r2, err := redis.Int(conn.Do("HGet","user01", "age"))
	if err != nil {
		fmt.Println("hget  err=", err)
		return 
	}
	fmt.Printf("操作ok r1=%v r2=%v \n", r1, r2)


	
	_, err = conn.Do("HMSet", "user02", "name", "so", "age", 190)
	if err != nil {
		fmt.Println("HMSet  err=", err)
		return 
	}
	x, err := redis.Strings(conn.Do("HMGet","user02", "name", "age"))
	if err != nil {
		fmt.Println("hget  err=", err)
		return 
	}
	for i, v := range x {
		fmt.Printf("r[%d]=%s\n", i, v)
	}
}
	