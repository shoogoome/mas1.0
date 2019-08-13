#!/bin/bash

cd ..
WORK_DIR=`pwd`

# 初始化文件路径
init_path() {
    mkdir -p ${WORK_DIR}'/redis'
    mkdir -p ${WORK_DIR}'/mongo'
    for i in `seq 0 5`; do
        mkdir -p ${WORK_DIR}'/server/root'
        mkdir -p ${WORK_DIR}'/server/tmp'
        mkdir -p ${WORK_DIR}'/server/logs'
    done
}

init_path