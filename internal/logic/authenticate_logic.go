package logic

import (
    "context"
    "github.com/zeromicro/go-zero/core/logc"
    "github.com/zeromicro/go-zero/core/logx"
    "google.golang.org/grpc/status"
    "mooon-auth-example/internal/svc"
    "mooon-auth-example/pb/mooon_auth"
    "strconv"

    moooncrypto "github.com/eyjian/gomooon/crypto"
    mooonutils "github.com/eyjian/gomooon/utils"
)

const (
    NoSessionId      = 2024020301 // cookie 中没有会话 ID（鉴权需要）
    SessionNotExists = 2024020302 // 会话不存在
)

type AuthenticateLogic struct {
    ctx    context.Context
    svcCtx *svc.ServiceContext
    logx.Logger
}

// 鉴权数据
type authData struct {
    uid  uint32 // 用户 ID
    role string // 用户角色
}

var authDataTable map[string]*authData // Key 为会话 ID

// 初始化鉴权数据
func init() {
    authDataTable = make(map[string]*authData)

    authDataTable = map[string]*authData{
        "1234567890": &authData{
            role: "super",
            uid:  2024020101,
        },
        "1234567891": &authData{
            role: "admin",
            uid:  2024020102,
        },
        "1234567892": &authData{
            role: "ordinary",
            uid:  2024020103,
        },
    }
}

func NewAuthenticateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AuthenticateLogic {
    return &AuthenticateLogic{
        ctx:    ctx,
        svcCtx: svcCtx,
        Logger: logx.WithContext(ctx),
    }
}

func (l *AuthenticateLogic) Authenticate(in *mooon_auth.AuthReq) (*mooon_auth.AuthResp, error) {
    // todo: add your logic here and delete this line
    out := &mooon_auth.AuthResp{}

    httpCookie, ok := in.HttpCookies["sessionid"]
    if !ok {
        logc.Error(l.ctx, "no sessionid in cookie")
        return nil, status.Error(NoSessionId, "no sessionid in cookie")
    }

    sessionId := httpCookie.Value
    authData, ok := authDataTable[sessionId]
    if !ok {
        logc.Error(l.ctx, "session not exists")
        return nil, status.Error(SessionNotExists, "session not exists")
    }

    out.HttpHeaders = make(map[string]string)
    // 在调用端的名必须以“Grpc-Metadata-”打头，而被调端必须以“gateway-”打头，这是 go-zero 的 gateway/internal/headerprocessor.go 写死的规则
    out.HttpHeaders["Grpc-Metadata-uid"] = strconv.FormatUint(uint64(authData.uid), 10)
    out.HttpHeaders["role"] = authData.role

    cookie := &mooon_auth.Cookie{
        Name:  "token",
        Value: getToken(),
    }
    out.HttpCookies = append(out.HttpCookies, cookie)

    return out, nil
}

func getToken() string {
    nonceStr := mooonutils.GetNonceStr(64)
    return moooncrypto.Md5Sum(nonceStr, false)
}
