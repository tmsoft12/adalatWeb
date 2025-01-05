#!/bin/bash

echo "Web deploy .........."
sleep 2
echo "web stop ............"
sudo systemctl stop adalatWeb.service
echo "remove old main file."
sleep 2
sudo rm /usr/local/bin/web/main
echo "build go file ......."
sleep 2
go build main.go
echo "mv new go build file."
sleep 5
sudo mv main /usr/local/bin/web/
echo "systemctl enable ...."
sleep 2
sudo systemctl enable adalatWeb.service
echo "systemctl start ...."
sleep 2
sudo systemctl start adalatWeb.service
sleep 2 
sudo systemctl status adalatWeb.service
