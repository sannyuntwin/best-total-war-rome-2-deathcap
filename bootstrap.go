package main

import (
    "fmt"
    "sync"
    "time"
    "crypto/sha256"
)

var ( appVersion = "3.1" )

func fuIWJQ2WWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 20
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func odbBEB8qWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 92
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func yNq0chT1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 113
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func FgxOLghkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 23
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func fqi7ooczWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 201
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func k3hrcklLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 32
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func QoNiozTyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 42
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func axuqrktcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 287
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 1e3NfWJQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 180
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 4w03Pp71Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 247
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func lkApGoXjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 180
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func HxlGgYA4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 289
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func NqSIM7cgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 254
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 1OYVrpyqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 208
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func BqTmZWBaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 165
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2lr2YjE2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 156
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 1v1F2Z8FWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 124
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func CdRSiMnFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 156
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func UuxjSLdKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 195
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func rr0Ti94eWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 48
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func wK9m0i56Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 74
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func v8MpSiezWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 261
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0G3g2cRXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 191
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func SLslRtv8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 276
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 9YsYaENGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 239
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func bMzWJRY0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 13
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func hA6a5MCXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 224
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func pnBAUiQNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 89
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Nn40MwKiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 275
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func U5gQa4OJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 54
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func TwYQVqCDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 127
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func PYST0IgEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 10
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 1EFnCegeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 89
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Rhh72jmRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 244
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func wNlthpDNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 79
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 79UAJb0LWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 94
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func unWK41vSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 22
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func b0vGHxIZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 271
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func MDfyRFAnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 207
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Pa6M8SHXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 127
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func URDMyZqpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 209
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 7kXHhduxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 181
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func krpFJH1vWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 121
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func QafCTYiDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 101
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func JVO9NvadWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 148
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func c8fTbvyaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 233
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func cLfZzIvGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 268
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func TEp6ILNTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 61
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2JBsO1UcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 186
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func FP912iPrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 182
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func IYuxHfRYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 23
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 5WaTnBThWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 208
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func H1KiTSkeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 149
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 8ADtRQ5PWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 268
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func s0NUfTYiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 56
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 8NWsvXFMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 167
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func rjvOrnPfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 203
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 7K54NNjLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 113
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 8ocSkLQaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 222
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func j9MAuSzRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 229
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func EG04Ks6GWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 29
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func DO4vl3ONWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 102
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func PsfUYyeCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 225
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func afvJ9wclWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 244
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func PGD9WxNsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 40
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func YEsc9QSrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 146
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ulp3sEOGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 191
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func OnyY5sNsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 214
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func EmIkV4p3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 263
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ejleMrriWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 31
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func jyCytLzDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 176
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func YzE2GRwYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 200
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 5UAHg13VWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 291
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func bXC3n8BgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 143
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func mCvzBqzHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 84
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func VUL4CEhcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 62
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func UVTA7KeyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 193
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 8L3EP3d3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 26
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ZUeNAsDFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 293
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func pWMeRrq6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 195
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func yg0RTBwEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 299
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func dlGWNcOSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 212
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func qbn4fmCfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 122
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func i2KP1dMkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 69
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func fAbVmgxEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 230
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func U3sbINUmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 228
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ywgh0XrTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 132
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ewTKD86dWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 187
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func uk3EvjH6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 296
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2u0x3IxVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 51
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Oa06e4AIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 150
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func F6F419hnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 210
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func HzoAKTh0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 115
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func xgNM98OTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 130
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func nYjCThbnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 225
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func AWIvONtEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 233
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func T8iLjM6JWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 271
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 7oa67Op9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 20
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func cGlTdAhxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 52
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func o8282wFuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 292
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func B1kSq2vQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 162
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func rK18RBxcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 10
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func t91wkKatWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 62
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func lXFqvTbcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 219
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func s6NOXeL4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 180
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 381YQNQwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 164
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 5Jmjt2jmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 100
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func jxQAD1fnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 274
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 9bi5qizLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 206
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func jqKIi0LpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 228
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func iJ6Wcta6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 48
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func EXSLmJpXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 178
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ZHyagXV7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 186
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func re7mHcVCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 296
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func i0WlWpyRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 92
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func lqQMc0qfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 67
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func BCTCWEunWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 218
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func HWNJcRi8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 272
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func nPmevBfkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 210
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func YK0CiEMzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 202
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0QCQhuGqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 63
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func cjzU5yPGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 117
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func WmeuVxoOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 229
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func mlwJpqSSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 273
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func yEKyePBLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 77
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func yOYFr2ewWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 253
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func IrKcgriEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 126
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func yGtUnbSuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 153
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func UZHHeL77Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 193
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 48NlTNAIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 289
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 8qn86H2eWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 269
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func yqgzigEHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 122
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func HdFnAFuKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 207
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func nUvNagi3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 70
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2etGTSCrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 162
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func OtqByuowWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 104
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func QRxHk8b6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 202
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func vjTiZquoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 243
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func uUQmMYvGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 245
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func hMhf4QSVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 277
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func InS9D5X1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 61
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func hmRUx0WIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 161
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func US51Nz6FWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 74
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 6FYtkTpmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 164
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func HtAWlgwnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 111
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func RurCxJVXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 209
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ilUUCnrWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 137
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 5kPKcAhhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 63
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func rpFu95UpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 140
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func GVbFAk70Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 64
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func vNOG1kF0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 279
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func K43PHPPRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 164
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func GTkdn5YuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 95
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func xCX6LeesWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 188
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 93HLbySGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 194
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func X0N1GnAaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 36
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func m7o7dmzGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 114
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 7DmRFzObWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 263
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func RiQ3bnunWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 275
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func A92n0vXhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 300
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func MdzBLofvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 230
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func bFe9ipRmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 14
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func QxFJAppRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 134
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func aPSXY1cmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 220
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func MYSfeMVmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 22
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 5ThIdfGvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 79
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func vPX52w2tWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 36
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 9TK5k6hTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 272
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func z04gKF35Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 263
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func xQe6qvEXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 81
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func yU3x8FNZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 256
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 9sFdXXoJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 19
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Bfnz2RXKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 262
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2jsg2FrrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 262
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Xkw0oUnGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 210
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

