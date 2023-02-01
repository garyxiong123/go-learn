# 异常-错误 处理

## 1:错误和异常如何区分？

- 错误【error】

    - 意料之中
    - 场景
        - 不合法的输入
        - 请求超时http/DB
    - 定义
        - 错误指的是可能出现问题的地方出现了问题，比如打开一个文件时失败，这种情况在人们的
        - 错误是业务过程的一部分，而异常不是
    - 
    - 原则

        - 1：失败的原因只有一个时，不使用error
        - 2：没有失败时，不使用error
        - 3：error应放在返回值类型列表的最后
        - 4：错误值统一定
        - 5：错误逐层传递时，层层都加日志
        - 6：错误处理使用defer 【finnaly】
            - 资源回收

        - 7：当尝试几次可以避免失败时，不要立即返回错误

            - 有时第二次或第三次尝试时会成功。
            - 在重试时，
                - 间间隔+重试的次数!=>无限重试

        - 8：当上层函数不关心错误时，建议不返回error

            - 关心了也没用=》不关心

                - 对于一些资源清理相关的函数（destroy/delete/clear），如果子函数出错，打印日志即可，而无需将错误进一步反馈到上层函数，
                - 因为一般情况下，上层函数是不关心执行结果的，或者即使关心也无能为力，于是我们建议将相关函数设计为不返回error

        - 9:当发生错误时，不忽略有用的返回值

            - 当函数返回non-nil的error时，其他的返回值是未定义的(undefined)，这些未定义的返回值应该被忽略。然而，有少部分函数在发生错误时，仍然会返回一些有用的返回值。比如，当读取文件发生错误时，Read函数会返回可以读取的字节数以及错误信息。对于这种情况，应该将读取到的字符串和错误信息一起打印出来

    - 错误=》异常

        - 重试URL多次

    - 
    - 

- 异常【panic】

    - 错误《=异常

        - 异常转错误，比如panic触发的异常被recover恢复后，将返回值中error类型的变量进行赋值，以便上层函数继续走错误处理流程。
        - 异常=》定义系统错误=》告警

    - 原则

        - 1:速错
        - 2：在程序部署后，应恢复异常避免程序终止

            - panic了，没有recover，就会异常退出。所以，一旦Golang程序部署后
            - recover响应该异常：

                - 1:打印堆栈的异常调用信息和关键的业务信息，
                - 2:将异常转换为错误，

                    - 以便调用者让程序恢复到健康状态并继续安全运行

        - 3：对于不应该出现的分支，使用异常处理

            -

    - 不应该出现问题的地方出现了问题，

        - 比如引用了空指针

    - panic

        - 如果遇到引用空指针、下标越界或显式调用panic函数等情况，则先触发panic函数的执行，然后调用延迟函数。调用者继续传递panic，因此该过程一直在调用栈中重复发生：函数停止执行，调用延迟执行函数等。如果一路在延迟函数中没有recover函数的调用，则会到达该携程的起点，该携程结束，然后终止其他所有携程，包括主携程（类似于C语言中的主线程，该携程ID为1）

    - 意料之外/不应该
    - 场景

        - 异常处理的作用域（场景）

            - 空指针引用
            - 下标越界
            - 除数为0
            - 不应该出现的分支，比如default
            - 输入不应该引起函数错误

    - 异常终止程序

        - 启动

    - 捕获异常/不能终止

        - 影响面太大

## 2:系统异常 HttpCode != 200

### 400

- StatusBadRequest                   = 400

    - 400 - GET /api/v1/block?by=height\u0026value=174 - 127.0.0.1:33248 - Go-http-client/1.1","duration":"0.4ms","level":"info","span":"31e0d4872c9166c8","trace":"604e4ded5f9df3e4798ec3255cef8772"}

- StatusUnauthorized                 = 401
- StatusForbidden                    = 403
- StatusNotFound                     = 404
- StatusMethodNotAllowed             = 405
- StatusRequestTimeout               = 408
- StatusLengthRequired               = 411
- StatusPreconditionFailed           = 412
- StatusRequestEntityTooLarge        = 413
- StatusTooManyRequests              = 429

    - 限流---
    -

### 500

- StatusInternalServerError           = 500
- StatusNotImplemented                = 501
- StatusBadGateway                    = 502

    -

- StatusServiceUnavailable            = 503

    - 超时

- StatusGatewayTimeout                = 504
- StatusHTTPVersionNotSupported       = 505
- StatusVariantAlsoNegotiates         = 506
- StatusInsufficientStorage           = 507
- StatusLoopDetected                  = 508
- StatusNotExtended                   = 510
- StatusNetworkAuthenticationRequired = 511

## 3:业务错误 HttpCode = 200

### 通用异常【需要跳转页面】

- 登陆过期

### 特点错误【前端Alert】

### 

