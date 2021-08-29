#!/bin/bash
go build ./src/service
sudo systemctl start service
sudo systemctl enable service