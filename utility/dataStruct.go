package utility

import (
    // "labix.org/v2/mgo/bson"
)

//Error is one part of out message that is returned to client, it consists of a error code and a error message
//when error code is zero, it means everything is going well.
type Error struct {
	ErrorCode int
	ErrorMessage string
}

//out message is consist of 
//1. the error message, 
//2. which command it is corresponding
//3. user code of person who start the command
//4. the result of the command
type OutMessage struct {
	Error Error
	Command string
    UserCode string
	Result string
}

//in message consist of
//1. the user's command
//2. parameter of this command
type InMessage struct {
	Command string
	Parameter string
}

//user login struct consist of
//1. user name
//2. use password
type UserLogin struct{
    UserName string
    Password string
}

//user struct is corresponding to the mysql user table
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

//project code struct
type ProjectCode struct {
    ProjectCode string
}

type MissionCode struct {
    MissionCode string
}

type GraphCode struct {
    GraphCode string
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

type DailyCode struct {
    DailyCode string
}

type ChartCode struct {
    ChartCode string
}

//project struct is corresponding to the mysql project table
//project contains many missions
//mission can contain other missions
//if mission contain other missions, it is called a campaign and its iscampaign
//should be set to 1
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

//mission struct is corresponding to the mysql mission table
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

//graph struct is corresponding to the mysql graph table
//every campaign contains many missions which is formed out of
//many missions, these messions make up a directed acyclic graph
//this struct only contains the node position
//the dependency relationship is stored in dependency table
type Graph struct {
    GraphCode string
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

//dependency struct is corresponding to the mysql dependency table
//every depency is a vector which contains a start mission code
//and a end mission code
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

//target struct is corresponding to mysql target table
//every mission may have one or many target
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

//daily struct is corresponding to mysql daily table
//every mission may have one or many daily
//daily is used to review the rate of progress of mission
type Daily struct {
    DailyCode string
    MissionCode string
    ProjectCode string
    VersionTag string
    StoragePosition string
    Picture string
    InsertDatetime string
    UpdateDatetime string
}

//chart struct is corresponding to mongodb chart table
//a chart maybe a picture or a string
//chart is from someuser to someuser
type Chart struct {
    Id string `json:"id"    bson:"_id"`
    IsPicture int
    Message string
    From string
    SendTime string
    To string
    ReceivedTime string
    IsReceived int
    Deleted int
    DeletedTime string
}

//post struct is corresponding to mongodb post table
type Post struct {
    Id string `json:"id"    bson:"_id"`
    MissionCode string
    PostType int
    Code string
    IsPicture int
    Message string
    ReplyTo string
    UserCode string
    PostTime string
    Deleted int
    DeletedTime string
}