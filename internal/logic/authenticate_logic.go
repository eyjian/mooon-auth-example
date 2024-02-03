package logic

import (
    "context"
    "github.com/zeromicro/go-zero/core/logc"
    "google.golang.org/grpc/metadata"

    "mooon-auth-example/internal/svc"
    "mooon-auth-example/pb/mooon_auth"

    "github.com/zeromicro/go-zero/core/logx"
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
        "mooon": &authData{
            role: "super",
            uid:  2024020101,
        },
        "zhangsan": &authData{
            role: "admin",
            uid:  2024020102,
        },
        "wangwu": &authData{
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
    md, ok := metadata.FromIncomingContext(l.ctx)
    if !ok {
        logc.Error(l.ctx, "no any")
    } else {
        logc.Info(l.ctx, md)
    }
    return out, nil
}
