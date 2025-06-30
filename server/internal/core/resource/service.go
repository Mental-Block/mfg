package resource

/*
	High level overview of ResourceService should not be directly imported.
	Copy interface and use dependancy injection over direct import.
*/
type IResourceService interface {

}

type ResourceService struct {
	resourceStore IResourceStore
	resourceBucket IResourceBlobStore
}

func NewResourceService(resourceStore IResourceStore) *ResourceService {
	return &ResourceService{
		resourceStore: resourceStore,
	}
}

