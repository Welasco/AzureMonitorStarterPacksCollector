#!/bin/bash

INSTALL_DIR=/opt/azuremonitorstarterpackscollector

#Step 1) Check if root--------------------------------------
if [[ $EUID -ne 0 ]]; then
   echo "Please execute the installation script as root."
   exit 1
fi
#-----------------------------------------------------------

echo "Uninstalling AzureMonitorStarterPacksCollector"

uninstall (){
     
        echo "AzureMonitorStarterPacksCollector Uninstall"
        rm -rf $INSTALL_DIR
        echo "AzureMonitorStarterPacksCollector Uninstalled"  
}

disablesetupsystemd (){

    sudo systemctl disable azuremonitorstarterpackscollector
    sudo systemctl stop azuremonitorstarterpackscollector
    sudo rm /lib/systemd/system/azuremonitorstarterpackscollector.service
    sudo systemctl daemon-reload
}

disablesetupsystemd
uninstall