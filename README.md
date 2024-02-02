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
go run . word -path=./text
```

output 👇
```
Lorem {{c1::ipsum}} dolor {{c1::sit}} amet, {{c1::consectetur}} adipiscing {{c1::elit}}, sed {{c1::do}} eiusmod {{c1::tempor}} incididunt {{c1::ut}} labore {{c1::et}} dolore {{c1::magna}} aliqua. {{c1::In}} hac {{c1::habitasse}} platea {{c1::dictumst}} vestibulum {{c1::rhoncus}} est. {{c1::Quam}} nulla {{c1::porttitor}} massa {{c1::id}} neque {{c1::aliquam}}. Sed {{c1::blandit}} libero {{c1::volutpat}} sed {{c1::cras}}. Pretium {{c1::fusce}} id {{c1::velit}} ut {{c1::tortor}} pretium {{c1::viverra}}. In {{c1::vitae}} turpis {{c1::massa}} sed {{c1::elementum}} tempus {{c1::egestas}} sed. {{c1::Placerat}} vestibulum {{c1::lectus}} mauris {{c1::ultrices}} eros {{c1::in}}. Vel {{c1::quam}} elementum {{c1::pulvinar}} etiam {{c1::non}} quam {{c1::lacus}} suspendisse {{c1::faucibus}}. Et {{c1::pharetra}} pharetra {{c1::massa}} massa {{c1::ultricies}}. Neque {{c1::viverra}} justo {{c1::nec}} ultrices {{c1::dui}} sapien. {{c1::Amet}} commodo {{c1::nulla}} facilisi {{c1::nullam}} vehicula {{c1::ipsum}}. Ullamcorper {{c1::malesuada}} proin {{c1::libero}} nunc {{c1::consequat}} interdum {{c1::varius}} sit.
```