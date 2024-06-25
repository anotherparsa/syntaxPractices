#!/bin/bash

# Author: AnotherParsa
# Date Created: 2024/06/25
# Last Modified: 2024/06/25

# Description:
# Creates Backup

# Usage:
# backup_script
tar -cvf ~/Desktop/backup" $(date +%d-%m-%Y_%H-%M-%S)".tar ~/Desktop/backup/* 2>/dev/null
exit 0
