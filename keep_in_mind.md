When to use locks over channels and goroutines?
https://quii.gitbook.io/learn-go-with-tests/go-fundamentals/sync#when-to-use-locks-over-channels-and-goroutines
Use channels when passing ownership of data
Use mutexes for managing state

在 Go 中，`channels` 和 `mutexes` 都可以用于同步并发操作，但它们适用于不同的场景：

### 使用 `Mutex` 的场景：
- **保护共享资源**：当多个 goroutine 需要读写共享的内存数据时，使用 `sync.Mutex` 可以确保同时只有一个 goroutine 能够访问该资源。
- **简单的锁定机制**：`Mutex` 提供了一种简单、低开销的方式来锁定和解锁资源，适合在需要强一致性或简单同步的场景下使用。

### 使用 `Channels` 的场景：
- **消息传递**：`Channels` 更适合用于 goroutine 之间的通信，通过传递数据来协调工作，避免直接共享内存。
- **处理并发任务**：使用 `Channels` 可以让一个 goroutine 发送数据，另一个 goroutine 接收并处理数据，适合用于任务队列、工作池等场景。
- **避免锁**：`Channels` 通过传递数据而非锁定资源来同步，通常可以让代码更容易理解且避免死锁。

### 选择指南：
- 如果你只需要简单地保护共享资源，`Mutex` 是更直接的选择。
- 如果你的程序需要通过 goroutine 之间传递数据来同步，`Channels` 会更合适。

总之，`Mutex` 适合用于简单的锁定和解锁操作，而 `Channels` 适合用于 goroutine 之间的复杂通信和协调。