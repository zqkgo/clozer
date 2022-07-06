# Clozer

a tool to cloze code, text, etc.

```
go run . code -path=./code
```

output 👇
```
{{c1::func (wq *loopQueue) len() int {}}
        {{c2::if wq.size == 0 {}}
                {{c3::return 0}}
        }

        {{c4::if wq.head == wq.tail {}}
                {{c5::if wq.isFull {}}
                        {{c6::return wq.size}}
                }
                {{c7::return 0}}
        }

        {{c8::if wq.tail > wq.head {}}
                {{c9::return wq.tail - wq.head}}
        }

        {{c10::return wq.size - wq.head + wq.tail}}
}
```

```
go run . text -path=./text
```