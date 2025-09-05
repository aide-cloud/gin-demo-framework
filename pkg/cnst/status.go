package cnst

type Status int

const (
	StatusPending Status = iota
	StatusInProgress
	StatusCompleted
	StatusCancelled
)

func (s Status) String() string {
	switch s {
	case StatusPending:
		return "未开始"
	case StatusInProgress:
		return "进行中"
	case StatusCompleted:
		return "已完成"
	case StatusCancelled:
		return "已取消"
	default:
		return "未知"
	}
}
