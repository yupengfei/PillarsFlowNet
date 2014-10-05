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
    IsPicture int//0 false, 1 true
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

type ProjectCode struct {
    ProjectCode string
}

type MissionCode struct {
    MissionCode string
}

type CampaignCode struct {
    CampaignCode string
}

type DependencyCode struct {
    DependencyCode string
}

type TargetCode struct {
    TargetCode string
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
    Status int
    Picture string
    InsertDatetime string
    UpdateDatetime string
}

type Mission struct {
    //MissionId string
    MissionCode string
    MissionName string
    ProjectCode string
    ProductType string
    IsCampaign int
    MissionType string
    MissionDetail string
    PlanBeginDatetime string
    PlanEndDatetime string
    RealBeginDatetime string
    RealEndDatetime string
    PersonIncharge string
    Status int
    Picture string
    InsertDatetime string
    UpdateDatetime string
}

type Campaign struct {
    CampaignCode string
    ProjectCode string
    NodeCode string
    Width int
    Height int
    XCoordinate int
    YCoordinate int
    InsertDatetime string
    UpdateDatetime string
}


type Dependency struct {
    DependencyCode string
    CampaignCode string
    ProjectCode string
    StartMissionCode string
    EndMissionCode string
    DependencyType int
    InsertDatetime string
    UpdateDatetime string
}

type Target struct {
    TargetCode string
    MissionCode string
    ProjectCode string
    VersionTag string
    StoragePosition string
    Picture string
    InsertDatetime string
    UpdateDatetime string
}

// 

type ChartMgo struct {
    ChartCode string
    Message string
    From string
    SendTime string
    To string
    ReceivedTime string
    Receipt int
    IsRead int
    Deleted int
    DeletedTime string
}