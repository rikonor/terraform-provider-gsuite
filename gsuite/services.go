package gsuite

import admin "google.golang.org/api/admin/directory/v1"

type GroupsService interface {
	Insert(grp *admin.Group) (*admin.Group, error)
	Get(grpID string) (*admin.Group, error)
	Delete(grpID string) error
}

type StubGroupsService struct {
	InsertFunc func(grp *admin.Group) (*admin.Group, error)
	GetFunc    func(grpID string) (*admin.Group, error)
	DeleteFunc func(grpID string) error
}

func (grpSvc *StubGroupsService) Insert(grp *admin.Group) (*admin.Group, error) {
	return grpSvc.InsertFunc(grp)
}

func (grpSvc *StubGroupsService) Get(grpID string) (*admin.Group, error) {
	return grpSvc.GetFunc(grpID)
}

func (grpSvc *StubGroupsService) Delete(grpID string) error {
	return grpSvc.DeleteFunc(grpID)
}

func WrapGroupsService(grpSvc *admin.GroupsService) GroupsService {
	return &StubGroupsService{
		InsertFunc: func(grp *admin.Group) (*admin.Group, error) {
			return grpSvc.Insert(grp).Do()
		},
		GetFunc: func(grpID string) (*admin.Group, error) {
			return grpSvc.Get(grpID).Do()
		},
		DeleteFunc: func(grpID string) error {
			return grpSvc.Delete(grpID).Do()
		},
	}
}

type UsersService interface {
	Insert(usr *admin.User) (*admin.User, error)
	Get(usrID string) (*admin.User, error)
	Delete(userID string) error
}

type StubUsersService struct {
	InsertFunc func(usr *admin.User) (*admin.User, error)
	GetFunc    func(usrID string) (*admin.User, error)
	DeleteFunc func(userID string) error
}

func (us *StubUsersService) Insert(usr *admin.User) (*admin.User, error) {
	return us.InsertFunc(usr)
}

func (us *StubUsersService) Get(usrID string) (*admin.User, error) {
	return us.GetFunc(usrID)
}

func (us *StubUsersService) Delete(usrID string) error {
	return us.DeleteFunc(usrID)
}

func WrapUsersService(usrSvc *admin.UsersService) UsersService {
	return &StubUsersService{
		InsertFunc: func(usr *admin.User) (*admin.User, error) {
			return usrSvc.Insert(usr).Do()
		},
		GetFunc: func(usrID string) (*admin.User, error) {
			return usrSvc.Get(usrID).Do()
		},
		DeleteFunc: func(userID string) error {
			return usrSvc.Delete(userID).Do()
		},
	}
}

type MembersService interface {
	Insert(groupKey string, member *admin.Member) (*admin.Member, error)
	Get(groupKey string, memberKey string) (*admin.Member, error)
	Delete(groupKey string, memberKey string) error
}

type StubMembersService struct {
	InsertFunc func(groupKey string, member *admin.Member) (*admin.Member, error)
	GetFunc    func(groupKey string, memberKey string) (*admin.Member, error)
	DeleteFunc func(groupKey string, memberKey string) error
}

func (ms *StubMembersService) Insert(groupKey string, member *admin.Member) (*admin.Member, error) {
	return ms.InsertFunc(groupKey, member)
}

func (ms *StubMembersService) Get(groupKey string, memberKey string) (*admin.Member, error) {
	return ms.GetFunc(groupKey, memberKey)
}

func (ms *StubMembersService) Delete(groupKey string, memberKey string) error {
	return ms.DeleteFunc(groupKey, memberKey)
}

func WrapMembersService(membersSvc *admin.MembersService) MembersService {
	return &StubMembersService{
		InsertFunc: func(groupKey string, member *admin.Member) (*admin.Member, error) {
			return membersSvc.Insert(groupKey, member).Do()
		},
		GetFunc: func(groupKey string, memberKey string) (*admin.Member, error) {
			return membersSvc.Get(groupKey, memberKey).Do()
		},
		DeleteFunc: func(groupKey string, memberKey string) error {
			return membersSvc.Delete(groupKey, memberKey).Do()
		},
	}
}
