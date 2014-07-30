package utility

type UserLogin struct{
	UserName string
	Password string
}
type Error struct {
	ErrorCode int
	ErrorMessage string
}
type OutMessage struct {
	Error Error
	Command string
	Result string
}

type InMessage struct {
	Command string
	Parameter string
}


type LoginInMessage struct {
	Auth string
	AuthMessage string
}

type User struct {
    UserCode string
    UserName string
    Password string
    Group string
    DisplayName string
    Position string
    Picture string
    Email string
    Phone string
    InsertDatetime string
    UpdateDatetime string
}


type Project struct {
    ProjectCode string
    ProjectName string
    ProjectDetail string
    PlanBeginDatetime string
    PlanEndDatetime string
    RealBeginDatetime string
    RealEndDatetime string
    PersonInCharge string
    Status string
    Picture string
    InsertDatetime string
    UpdateDatetime string
}

type Mission struct {
    MissionId string
    MissionName string
    ProjectCode string
    ProductType string
    MissionType string
    MissionDetail string
    PlanBeginDatetime string
    PlanEndDatetime string
    RealBeginDatetime string
    RealEndDatetime string
    PersonIncharge string
    Status string
    Picture string
    InsertDatetime string
    UpdateDatetime string
}


type WorkFLow struct {
    StartMissionCode string
    EndMissionCode string
    DependencyType string
    InsertDatetime string
    UpdateDatetime string
}

type Target struct {
    TargetCode string
    MissionCode string
    VersionTag string
    StoragePosition string
    Picture string
    InsertDatetime string
    UpdateDatetime string
}