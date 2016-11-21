/*
Copyright (c) 2016, UPMC Enterprises
All rights reserved.
Redistribution and use in source and binary forms, with or without
modification, are permitted provided that the following conditions are met:
    * Redistributions of source code must retain the above copyright
      notice, this list of conditions and the following disclaimer.
    * Redistributions in binary form must reproduce the above copyright
      notice, this list of conditions and the following disclaimer in the
      documentation and/or other materials provided with the distribution.
    * Neither the name UPMC Enterprises nor the
      names of its contributors may be used to endorse or promote products
      derived from this software without specific prior written permission.
THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS "AS IS" AND
ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT LIMITED TO, THE IMPLIED
WARRANTIES OF MERCHANTABILITY AND FITNESS FOR A PARTICULAR PURPOSE ARE
DISCLAIMED. IN NO EVENT SHALL UPMC ENTERPRISES BE LIABLE FOR ANY
DIRECT, INDIRECT, INCIDENTAL, SPECIAL, EXEMPLARY, OR CONSEQUENTIAL DAMAGES
(INCLUDING, BUT NOT LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR SERVICES;
LOSS OF USE, DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER CAUSED AND
ON ANY THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY, OR TORT
(INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE OF THIS
*/

package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/Sirupsen/logrus"
)

var (
	appVersion = "0.0.1"

	printVersion bool
	baseImage    string
)

func init() {
	flag.BoolVar(&printVersion, "version", false, "Show version and quit")
	flag.StringVar(&baseImage, "baseImage", "upmcenterprises/docker-elasticsearch-kubernetes:2.4.1.1", "Base image to use when spinning up the elasticsearch components.")
	flag.Parse()
}

func main() {
	if printVersion {
		fmt.Println("elasticsearch-operator", appVersion)
		os.Exit(0)
	}

	logrus.Info("elasticsearch operator starting up!")

	// Print params configured
	logrus.Info("Using Variables:")
	logrus.Infof("   baseImage: %s", baseImage)
}