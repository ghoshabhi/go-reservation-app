package list

// TODO: Try to extend/override this interface for concrete types like Room, RoomType, etc.

// Helper defines the interface to perform operations on a list
type Helper interface {
	Filter(f func(interface{}) bool) interface{}
	Map(f func(string) string) []string
}
