# Links to Read

## packages
- [Встроенные в Go пакеты](https://golang.org/pkg/.)
- [Организация go проекта](https://golang.org/doc/code.html)
## pointers
- [про память](https://stepik.org/lesson/266497/step/1?auth=login&unit=247452)
- [как не наступать на грабли в Go *](https://habr.com/ru/post/325468/)
## structures

Структура эта типа класс без методов (`map object`), а только с атрибутами? и экземпляр такого класса? :)

А нет, методы есть :)

```go
type Rectangle struct {
    x1, y1, x2, y2 float64
}
func (r *Rectangle) area() float64 {
    l := distance(r.x1, r.y1, r.x1, r.y2)
    w := distance(r.x1, r.y1, r.x2, r.y1)
    return l * w
}

func main(){
    r := Rectangle{0, 0, 10, 10}
    fmt.Println(r.area())
}
```
## strings

## errors

```go
    result, err := divide(a, b)
	if err != nil {
		fmt.Println("aka <except> в python")
	} else {
		fmt.Printf("My cool result %v", result)  // try:except:<else> 
	}
    defer closeDBConn()  // aka <finally>
    if result == "MyError" {
        panic("MyError hit, exit...")  // aka <raise> MyCoolException
    }
```

- [panic](https://stepik.org/lesson/264473/step/5?unit=245397) - throw error
- [defer](https://stepik.org/lesson/264473/step/7?unit=245397) - LIFO

> Всякий раз, когда функция panic вступает в действие, она выполняет все функции defer, связанные с текущим потоком. Функции отложенного вызова defer могут применяться для освобождения ресурсов, использующихся в функции. Эти функции defer выполняются непосредственно перед завершением текущей функции.
