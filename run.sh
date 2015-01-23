echo 'PillarsFlowNet项目启动脚本'
export C=$(cd `dirname $0`; pwd)
cd $C
go run socketServer.go
