package main

import (
    "fmt"
    "sync"
    "time"
    "crypto/sha256"
)

var ( appVersion = "5.5" )

func CKJhT4vlWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 184
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func mugRyF5VWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 114
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func g0tZd6m4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 125
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func yOzEl9nMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 234
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Cma7ytBsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 149
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func upf65P0aWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 190
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func tJ8lq6TYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 179
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 7DC0tnlNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 108
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func XuJr1gNDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 79
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0jxgpHNOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 70
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 9obyNgGMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 168
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func xNnEgKmkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 39
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func gnmHMnS3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 152
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 28SUbc6SWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 223
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func TfU39CGFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 135
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func iMAuycGIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 251
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func cZEzICYzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 22
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func TIDDQwEPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 21
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ZItCgEJmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 148
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func oZGe3ygdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 154
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func nu7rYrd3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 254
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func drOcv9QAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 96
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2wVhgcaCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 300
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func f4vjTJtEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 274
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func mJFRPjVUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 103
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Qv9CdRwOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 184
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func tlmnIpulWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 263
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func FiGOB07GWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 250
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func GJvRBmYgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 150
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func fWT3za9sWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 244
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Z0c5yJR9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 171
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func xwZWE5dcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 186
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func x80oVneBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 265
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func YPgMSo0qWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 16
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ocIgO7OaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 272
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func TuYCeTSQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 231
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 7uvuZMM1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 187
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Hrv40DyGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 256
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Mqt6M2eCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 273
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 1eA1T83yWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 99
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func KPKUxnUvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 240
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func eBDjgp3DWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 295
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 8hSd1SA6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 157
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 4WVct7HIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 77
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func mVi2Ma6fWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 212
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func I3xdfTkAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 142
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func MNMhK9lNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 157
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func QVjrURwqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 274
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func F8MhR7SXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 47
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func nKZutiCbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 259
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func gCRDkBg8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 200
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func fr9orNUGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 243
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func EOPLitADWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 234
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func tTuSPrvCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 290
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 4EdrYstwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 13
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func YQNEqYJrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 45
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func V3tHzIN2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 116
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func sdPaq2sYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 182
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func QVVckqs2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 271
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 66jAjIIPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 289
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func iXaQeTImWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 94
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func z3h7C45hWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 20
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func cFzlTDZPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 290
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func L25Z9fp0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 282
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func p71PJlZLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 185
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 4kPd6pdGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 279
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func iA5PesbjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 264
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func QUOrfFoQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 228
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func GvbwYp0IWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 141
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func APC9PzyLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 211
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func C4lY6XltWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 135
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 5eYqlbhlWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 167
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func WESJ0jXpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 293
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func DxIq5oEVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 238
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func VH5L1TuSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 151
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func DgJcEsvQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 54
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func f7stNuPiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 262
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func wV5A7Iy6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 240
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3nRokzXtWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 50
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func XzGhxydWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 204
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 8e4D8FMLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 123
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Xdk20Fz6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 32
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func JUQoUBMAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 282
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func xlK1snk9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 243
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func jYP3TzI2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 139
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func wjKJQHUjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 69
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func CiLKMkaPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 300
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func CIUT47wVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 192
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func tiOkovFZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 50
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func qETOHZpHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 117
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func XURVK7saWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 125
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0vjifTweWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 218
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Y3WsNEJ0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 290
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func nKK4sJfoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 96
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func EBIlDmU8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 61
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func NEGQPMCmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 33
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 4t7UKPwkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 117
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ETK2FGdlWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 190
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 35oeIqLiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 300
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func YEC58UH3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 254
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 8eAUJnF1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 95
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 7bz2efQ8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 213
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func KPQmz0kqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 186
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func FD3EDyVlWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 135
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func u7myucgQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 163
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func qKnYKjJoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 175
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func oZoMTCqQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 257
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func FgEq2loJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 206
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func XQx3FgltWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 267
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func lNJCLE8wWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 28
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func uBtmgpzZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 171
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func hW8dTcegWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 40
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ef68IWXdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 178
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func GcLgBSyUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 230
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 6sSoNtlcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 179
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func J2ZbhHaPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 28
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2PwUj50cWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 102
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func rEr1DcEPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 257
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func qZAoi0pqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 264
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func l04MmXAGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 219
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func KNNDL3GrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 189
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 5zqYOQ08Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 289
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func q2SLNhKMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 289
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func UZZ1vQnCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 39
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func MW8NVim5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 259
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func jb67ZMGBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 234
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func wg41zLsqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 113
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func S7KrVBOmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 40
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

