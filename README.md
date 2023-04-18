# ERNIE-API

文心一言 API

## 可用性

- Chat文心千帆接口已调试通，需要企业应用密钥调用
- 文心Ernie大模型接口由于未获取正式访问密钥，未经过验证测试

## 依赖
```go
import "github.com/google/go-querystring/query"
```

## 使用

```go
import "github.com/zjy282/ernie-api"
```
1. 获取Token
    
```go
//Chat文心千帆
ctx := context.Background()
req := &OAuthTokenRequest{
    ClientID:     "client_id",
    ClientSecret: "client_secret",
}
response, err := CreateBCEOAuthToken(ctx, req)
//文心Ernie大模型
ctx := context.Background()
req := &OAuthTokenRequest{
    ClientID:     "client_id",
    ClientSecret: "client_secret",
}
response, err := CreateOAuthToken(ctx, req)
```

2. 调用Chat文心千帆接口

```go
client := NewClientWithConfig(DefaultBCEConfig("AccessToken"))
ctx := context.Background()
req := &ChatRequest{
    User: "test",
    Messages: []ChatRequestMessage{
        {Role: ChatRoleUser, Content: "介绍一下你自己"},
    },
}
response, err := client.CreateChat(ctx, req)
```

3. 调用文心一言自定义接口

```go
client := NewClient("AccessToken")
ctx := context.Background()
req := &V3CustomizeRequest{
    Async:            1,
    Text:             "标题：芍药香氛的沐浴乳\\n文案：",
    MinDecLen:        32,
    SeqLen:           512,
    TopP:             0.9,
    TaskPrompt:       TaskPromptAdText,
    PenaltyScore:     1.2,
    IsUnidirectional: 0,
    TypeId:           1,
}

response, err := client.CreateV3Customize(ctx, req)
```

4. 调用文心一言结果返回接口

```go
client := NewClient("")
ctx := context.Background()
req := &V3CustomizeResultRequest{
    TaskId: 1,
}
response, err := client.GetV3CustomizeResult(ctx, req)
```

5. 调用文心一言图片生成接口

```go
client := NewClient("AccessToken")
ctx := context.Background()
req := &Txt2ImgRequest{
    Text:       "睡莲",
    Style:      StyleOilPainting,
    Resolution: ResolutionSquareChart,
    Num:        1,
    Image:      *multipart.FileHeader,
}
response, err := client.CreateTxt2Img(ctx, req)
```

6. 调用文心一言图片生成结果返回接口

```go
client := NewClient("")
ctx := context.Background()
req := &Txt2ImgResultRequest{
    TaskId: 1,
}
response, err := client.GetTxt2ImgResult(ctx, req)
```

## 注意

- 有问题请在issue中发起讨论
- 此仓库仅供学习
