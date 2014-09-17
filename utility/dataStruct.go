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

type ChartMessage struct {
    Message string
    To string
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
    MissionCode string
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


type Dependency struct {
    ProjectCode string
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

//This is used for add project
type AddProjectStruct struct {
    UserName string
    ProjectName string
    ProjectDetail string
    PlanBeginDatetime string
    PlanEndDatetime string
    RealBeginDatetime string
    RealEndDatetime string
    PersonInCharge string
    Status string
    Picture string
}

//when modify projects, replace the old one with a new instance without change the project code
type ModifyProjectStruct struct {
    UserName string
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
}

//the userName and projectCode are the only things you need to delete a project
type DeleteProjectStruct struct {
    UserName string
    ProjectCode string
}