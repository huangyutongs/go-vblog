package user

import "context"

// 面向对象
// user.Service, 设计你这个模块提供的接口
// 接口定义，一定要考虑兼容性，接口的参数不能变
type Service interface {
	// 用户创建
	// CreateUser(username, password, role string, lable map[string]string)
	// 设计CreateUserRequest, 可以扩展对象，而不影响接口的定义
	// 1. 这个接口支持取消吗？要支持取消应该怎么办?
	// 2. 这个接口支持Trace，TraceID怎么传递？
	// 中间件参数，取消/Trace/... 怎么产生怎么传递
	CreateUser(context.Context, *CreateUserRequest) (*User, error)
	// 查询用户列表，对象列表 [{}]
	QueryUser(context.Context, *QueryUserRequest) (*UserSet, error)
	// 查询用户详情，通过Id查询，
	DescribeUser(context.Context, *DescribeUserRequest) (*User, error)
	// 用户信息修改，修改用户名、密码、角色和标签，
	PatchUser(context.Context, *PatchUserRequest) (*User, error)
	// 用户删除，通过用户名删除
	DeleteUser(context.Context, *DeleteUserRequest) (*User, error)
}

// 用户创建的参数
type CreateUserRequest struct {
	Username string
	Password string
	Role     string
	Lable    map[string]string
}

// 查询用户列表
type QueryUserRequest struct {
	// 分页大小，一个多少个
	PageSize int
	// 当前页，查询哪一页的数据
	PageNumber int
	// 根据用户name查找用户
	Username string
}

type UserSet struct {
	// 总共有多少个
	Totle int64
	// 当前查询的数据清单
	Iteams []*User
}

type DescribeUserRequest struct {
	UserId int
}

type DeleteUserRequest struct {
	Username string
}

type PatchUserRequest struct {
	Username string
	Password string
	Role     string
	Lable    map[string]string
}
