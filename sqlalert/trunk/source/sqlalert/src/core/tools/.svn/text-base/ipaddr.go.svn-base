// File: ipaddr.go
//
// This file implements the misc tool functions for IP address.
//
// Copyright (C) 2017 YUN Li Lai, Nanjiing, Inc. All Rights Reserved.
// Written by ZHANG Li Dan <lidan.zhang@clearclouds-global.com>.
package tools

import (
	"errors"
	"fmt"
	"net"
	"strings"
)

// IPToUint32() - Convert IP string to uint32 number.
//
// @addr: String of IP address.
func IPToUint32(addr interface{}) (uint32, error) {
	var ip net.IP

	switch addr.(type) {
	case net.IP:
		ip = addr.(net.IP).To4()
		if ip == nil {
			msg := fmt.Sprintf("'%s' not IPv4 string", addr)
			return 0, errors.New(msg)
		}

	case string:
		ip = net.ParseIP(addr.(string))
		if ip == nil {
			msg := fmt.Sprintf("'%s' not IPv4 string", addr)
			return 0, errors.New(msg)
		}

		if ip = ip.To4(); ip == nil {
			msg := fmt.Sprintf("'%s' not IPv4 string", addr)
			return 0, errors.New(msg)
		}

	default:
		msg := fmt.Sprintf("'%v' not STR or net.IP type", addr)
		return 0, errors.New(msg)
	}

	ipValue := uint32(ip[0])
	ipValue = (ipValue << 8) + uint32(ip[1])
	ipValue = (ipValue << 8) + uint32(ip[2])
	ipValue = (ipValue << 8) + uint32(ip[3])

	return ipValue, nil
}

// IPInList() - Check whether IP address in the the given list.
//
// @addr: String of IP address.
// @list: List of IP string.
func IPInList(addr string, list []string) (bool, error) {
	ipNum, ipErr := IPToUint32(addr)

	for _, item := range list {
		if addr == item {
			return true, nil
		} else if ipErr == nil {
			if _, ipnet, err := net.ParseCIDR(item); err == nil {
				if itemNum, err := IPToUint32(ipnet.IP); err == nil {
					ones, bits := ipnet.Mask.Size()
					mask := ^((uint32(1) << uint(bits-ones)) - 1)
					if (ipNum & mask) == (itemNum & mask) {
						return true, nil
					}
				}
			} else if strList := strings.Fields(item); len(strList) == 3 && strings.ToLower(strList[1]) == "to" {
				matchLeft, matchRight := false, false
				if strList[0] == "*" {
					matchLeft = true
				} else if num, err := IPToUint32(strList[0]); err == nil {
					if ipNum >= num {
						matchLeft = true
					}
				}

				if strList[2] == "*" {
					matchRight = true
				} else if num, err := IPToUint32(strList[2]); err == nil {
					if ipNum <= num {
						matchRight = true
					}
				}

				if matchLeft && matchRight {
					return true, nil
				}
			}
		}
	}

	return false, nil
}
