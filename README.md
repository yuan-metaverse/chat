# chat
char serve

```text
├── cmd                     // 启动服务入口
│   └── main.go
├── internal
│   ├── handler             // WebSocket 连接处理，消息广播等
│   │   └── websocket.go
│   ├── models              // 数据模型（如房间、消息、用户等）
│   │   ├── room.go
│   │   ├── message.go
│   │   └── user.go
│   ├── storage             // 数据存储（如MongoDB操作）
│   │   └── mongo.go
│   ├── service             // 核心业务逻辑（如房间管理、消息发送）
│   │   ├── room_service.go
│   │   ├── message_service.go
│   │   └── user_service.go
│   ├── utils               // 公共工具类（如时间处理、ID生成等）
│   └── config              // 配置文件和初始化
│       └── config.go
├── scripts                 // 启动和部署相关脚本
└── go.mod                  // Go模块定义


chat-service/
├── cmd/
│   └── server/
│       └── main.go               # 入口文件，启动服务
├── internal/
│   ├── config/                   # 配置相关
│   │   └── config.go             # 配置管理
│   ├── handler/                  # HTTP / WebSocket 请求处理
│   │   └── websocket.go          # WebSocket 处理逻辑
│   ├── message/                  # 消息处理和数据库交互
│   │   └── message.go            # 消息模型与存储
│   ├── user/                     # 用户相关功能
│   │   └── user.go               # 用户管理
│   ├── group/                    # 群聊相关功能
│   │   └── group.go              # 群组管理
│   ├── storage/                  # 存储相关操作
│   │   └── mongo.go              # MongoDB 连接和操作
│   └── service/                  # 核心业务逻辑
│       └── chat_service.go       # 聊天业务处理
├── pkg/                          # 公共库或者第三方集成
│   └── websocket/                # WebSocket 实现封装
│       └── wsclient.go           # WebSocket 客户端封装
├── scripts/                      # 启动脚本或工具
│   └── run.sh                    # 启动脚本（可选）
├── api/                          # API 定义（可选，适用于与前端或其他服务协作）
│   └── chat.proto                # 如果使用 gRPC 或 Protobuf 定义接口
├── .env                          # 环境变量配置（如数据库连接等）
├── go.mod                        # Go module 配置文件
├── go.sum                        # Go checksum 文件
└── README.md                     # 项目说明文档
```