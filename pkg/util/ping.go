package util

import (
    "bytes"
    "context"
    "fmt"
    "os/exec"
    "sync"
    "time"
)

const DefaultPingTimeout = 20 * time.Second

func Ping(ipList ...string) map[string]error {
    return PingWithTimeout(DefaultPingTimeout, ipList...)
}

func PingSingleIpWithTimeout(timeout time.Duration, ip string) error {
    return PingWithTimeout(timeout, ip)[ip]
}

func PingWithTimeout(timeout time.Duration, ipList ...string) map[string]error {
    if len(ipList) == 0 {
        return nil
    }

    ret := make(map[string]error)
    ch := make(chan map[string]error, len(ipList))
    wg := &sync.WaitGroup{}

    for _, ip := range ipList {
        wg.Add(1)
        go func(ip string) {
            mp := make(map[string]error)
            mp[ip] = fping(ip, timeout)
            //mp[ip] = ping(ip, 1, 1)
            //mp[ip] = pingWithCtx(ip, timeout)

            ch <- mp
            wg.Done()
        }(ip)
    }
    wg.Wait()

    for i := 0; i < len(ipList); i++ {
        for ip, err := range <- ch {
            ret[ip] = err
        }
    }

    return ret
}

// TODO ctx控制的超时会不会使cmd先退出，进而产生ping的孤儿进程？
func pingWithCtx(ip string, timeout time.Duration) error {
    ctx, cancel := context.WithTimeout(context.Background(), timeout)
    defer cancel()

    cmd := exec.CommandContext(ctx, "/bin/ping", "-c1", "-t1", ip)
    var stdout bytes.Buffer
    var stderr bytes.Buffer

    cmd.Stdout = &stdout
    cmd.Stderr = &stderr
    // Run()方法等待执行结果；Start()方法调用Wait()方法等待；Run()内部就是Start+Wait
    if err := cmd.Run(); err != nil {
        return err
    }

    return nil
}

func ping(ip string, count, interval int) error {
    if interval <= 0 { interval = 1 }
    cmd := exec.Command("/bin/ping", fmt.Sprintf("-c%d", count), fmt.Sprintf("-t%d", interval), ip)
    if err := cmd.Run(); err != nil {
        return err
    }
    return nil
}

// 超时时间单位是毫秒
func fping(ip string, timeout time.Duration) error {
    cmd := exec.Command("/usr/sbin/fping", "-A", fmt.Sprintf("-t%d", timeout), ip)
    if err := cmd.Run(); err != nil {
        return err
    }
    return nil
}
