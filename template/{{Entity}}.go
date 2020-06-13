package {{.Package}}
import "github.com/google/uuid"
type {{.Entity}} struct {
	ID   uuid.UUID
}
type {{.Entity}}Service interface {
	List() ([]{{.Entity}}, error)
	Create(person {{.Entity}}) (uuid.UUID, error)
	Remove(uuid.UUID, error)
	Get(uuid.UUID) ({{.Entity}}, error)
}
