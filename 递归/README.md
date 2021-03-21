# 递归

```golang
func recursion(level , p1,p2) result{
    if  level > MAX_LEVEL{
        print_result
        return
    }

    process_data(level,data)

    recursion(level,p1,p2)

    reverse_state(level)
}
```