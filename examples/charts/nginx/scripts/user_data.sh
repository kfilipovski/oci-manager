#!/bin/bash

apt-get update
apt-get install -y nginx-light
iptables -F