# shared memory for go

因有需要通过golang访问其他语言创建的共享内存，所以封装了这个库来直接操作共享内存

## Api

### 打开共享内存
```
func Open(key int) (Segment, error)
```

### 创建共享内存

```
func Create(size int, flags int, mode int) (Segment, error)
```

### 分配内存地址

```
segment.Attach() (uintptr, error)
```

### 分离内存地址

```
segment.Detach() error
```

## Tips

