#!/bin/bash

INSTALL_DIR=/opt/azuremonitorstarterpackscollector

#Step 1) Check if root--------------------------------------
if [[ $EUID -ne 0 ]]; then
   echo "Please execute the installation script as root."
   exit 1
fi
#-----------------------------------------------------------

echo "Installing AzureMonitorStarterPacksCollector"

install (){
 
    echo "AzureMonitorStarterPacksCollector Install"
    mkdir -p $INSTALL_DIR
    cd $INSTALL_DIR
    wget https://raw.githubusercontent.com/Welasco/AzureMonitorStarterPacksCollector/main/AzureMonitorStarterPacksCollector
    wget https://raw.githubusercontent.com/Welasco/AzureMonitorStarterPacksCollector/main/config_collector.ini
    chmod +x AzureMonitorStarterPacksCollector


    echo "AzureMonitorStarterPacksCollector Installed"

  
}

setupsystemd (){
    echo "Setting AzureMonitorStarterPacksCollector as a systemd service"
    wget https://raw.githubusercontent.com/Welasco/AzureMonitorStarterPacksCollector/main/azuremonitorstarterpackscollector.service
    mv azuremonitorstarterpackscollector.service /lib/systemd/system
    if [ $? -eq 0 ]; then
        echo "Copied azuremonitorstarterpackscollector.service success"
    else
        echo "Fail to copy azuremonitorstarterpackscollector.service to /lib/systemd/system. Exiting instalation. Error Code:" $?
        exit 1
    fi

    sudo systemctl daemon-reload
    sudo systemctl enable azuremonitorstarterpackscollector
    if [ $? -eq 0 ]; then
        echo "Enabling azuremonitorstarterpackscollector.service success"
    else
        echo "Fail to enable azuremonitorstarterpackscollector.service. Exiting instalation. Error Code:" $?
        exit 1
    fi

    echo "Starting azuremonitorstarterpackscollector service"
    sudo systemctl stop azuremonitorstarterpackscollector
    sudo systemctl start azuremonitorstarterpackscollector
    if [ $? -eq 0 ]; then
        echo "Service azuremonitorstarterpackscollector.service started"
    else
        echo "Fail to start azuremonitorstarterpackscollector.service. Exiting instalation. Error Code:" $?
        exit 1
    fi
}

install
setupsystemd