package user

// 存放需要入口的数据结构(P0)

// 用户创建成功后返回一个User对象
// CreateAt 为啥没用time.Time，int64(TimeStramp), 统一标准化，避免时区对程序产生影响
// 在需要对时间进行展示的时候，由前端根据具体展示哪个时区的时间

type User struct {
	// 用户Id
	Id int
	// 创建时间，时间戳 10位，秒
	CreateAt int64
	// 更新时间，时间戳 10位，秒
	UpdateAt int64
	// 用户参数
	*CreateUserRequest
}
