# 编译
go mod tidy
CGO_ENABLED=1 go build pimonitor.go
# 启用服务
sudo cp pimonitor.service /etc/systemd/system 
sudo systemctl enable pimonitor.service
sudo systemctl start pimonitor.service