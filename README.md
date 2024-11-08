# __In-memory-cache__

Этот простой проект был написан для того, чтобы создавать и управлять кэшем на основе map в Go.

## Описание

* ```c.Set``` - добавляет значение в кэш
* ```c.Get``` - извлекает значение по ключу
* ```c.Delete``` - удаляет значение по ключу
* ```c.Clear``` - очищает весь кэш

## Пример
```
package main

import (
	"awesomeProject/cache"
	"fmt"
)

func main() {
	c := cache.NewCache()

  	// Добавляем 
	c.Set("key", "value")

  	// Извлекаем значение по ключу
	if val, ok := c.Get("key"); ok {
		fmt.Println("Found:", val)
	} else {
		fmt.Println("Not Found")
	}

  	// Удаляем значение по ключу
	c.Delete("key")
	if _, ok := c.Get("key"); ok {
		fmt.Println("Not deleted")
	} else {
		fmt.Println("Deleted")
	}

  	// Очищаем весь кэш
	c.Set("Key1", "Value1")
	c.Clear()
	if _, ok := c.Get("Key1"); !ok {
		fmt.Println("Cache cleared")
	}
}
```
