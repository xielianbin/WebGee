package internal

// 使用前缀树实现动态路由
type INode interface {
	MatchChild(part string) INode      //查找第一个匹配成功的节点
	MatchChildren(part string) []INode //查找所有匹配成功的节点
	Insert(pattern string, parts []string, height int)
	Search(parts []string, height int) INode
	GetPattern() string
	SetPattern(pattern string)
	GetPart() string
	SetPart(part string)
	GetChildren() []INode
	SetChildren(children []INode)
	GetIsWild() bool
	SetIsWild(isWild bool)
}
